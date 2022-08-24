package services

import (
	"arvan-credit-service/types/interfaces/repositories"
	"arvan-credit-service/types/structs"
	. "arvan-credit-service/utils"
	"encoding/json"
	"github.com/jackc/pgconn"
)

type SpendCodeService struct {
	repo     repositories.ICreditCodeRepository
	clientId uint
}

func (c *SpendCodeService) SetRepo(repository repositories.ICreditCodeRepository) {
	c.repo = repository
}

func (c *SpendCodeService) SetClientId(clientId uint) {
	c.clientId = clientId
}

func (c *SpendCodeService) Spend(code string, referenceId uint) (uint, uint, *structs.CustomError) {
	res := c.repo.FindByText(code)

	if res.Error != nil {
		return 0, 0, CustomError(structs.Categories.Internal, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return 0, 0, CustomError(structs.Categories.BusinessLogic, "unable to find credit code [%s]", code)
	}

	if res.Model.CurrentUsedCount >= res.Model.MaxUsableCount {
		return 0, 0, CustomError(structs.Categories.BusinessLogic, "Credit code limitation reached")
	}

	logId, err := c.repo.RunSpendCodeDBFunction(code, referenceId, c.clientId)

	return res.Model.Amount, logId, c.processResponseError(err)
}

func (c *SpendCodeService) processResponseError(err error) *structs.CustomError {
	if err != nil {

		pgErr := err.(*pgconn.PgError)

		var errorDetail map[string]interface{}
		if err = json.Unmarshal([]byte(pgErr.Detail), &errorDetail); err == nil {
			switch errorDetail["code"].(float64) {
			case 1:
				return CustomError(structs.Categories.BusinessLogic, "Unable to find credit code")
			case 2:
				return CustomError(structs.Categories.BusinessLogic, "Credit code limitation reached")
			case 3:
				return CustomError(structs.Categories.BusinessLogic, "already utilized for this reference id")
			case 6:
				return CustomError(structs.Categories.UnAuthorized, "unknown client service")
			case 4, 5:
				fallthrough
			default:
				return CustomError(structs.Categories.Internal, "internal server error")

			}
		}
	}

	return nil
}

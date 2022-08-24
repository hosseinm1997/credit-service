package repositories

import (
	"arvan-credit-service/db/models"
	"arvan-credit-service/types/structs"
	. "arvan-credit-service/utils"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreditCodeRepository struct{}

func (repo *CreditCodeRepository) FindByText(text string) structs.RepositoryResult[models.Code] {
	model := &models.Code{Text: text}
	response := DB().Take(model)
	return RepoResult[models.Code](model, response)
}

func (repo CreditCodeRepository) RunSpendCodeDBFunction(code string, referenceId uint, clientId uint) (uint, error) {
	var usageLogId uint
	response := DB().Raw("select utilize_credit_code_v1_1($1, $2, $3)", code, referenceId, clientId).Scan(&usageLogId)
	return usageLogId, response.Error

}

func (repo *CreditCodeRepository) InsertNewUsageLog(code string, referenceId uint) (uint, uint, error) {
	var usageLogId uint

	model := &models.Code{
		Text: code,
	}

	err := DB().Transaction(func(tx *gorm.DB) error {

		res := tx.Take(model)

		if res.Error != nil || res.RowsAffected == 0 {
			return fmt.Errorf("undefined credit code")
		}

		if model.CurrentUsedCount >= model.MaxUsableCount {
			return fmt.Errorf("Credit code limitation reached")
		}

		model.CurrentUsedCount++
		response := tx.Save(model)

		if response.Error != nil {
			return response.Error
		}

		codeUsageLog := &models.CodeUsageLog{
			CodeID:      model.ID,
			ReferenceID: referenceId,
		}

		response = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(codeUsageLog)

		//err := tx.Model(model).
		//	Association("CodeUsageLogs").
		//	Append(&models.CodeUsageLog{ReferenceID: referenceId})

		if response.Error != nil {
			return response.Error
		}

		usageLogId = codeUsageLog.ID

		return nil
	})

	if err != nil {
		return 0, 0, err
	}

	return model.Amount, usageLogId, nil
}

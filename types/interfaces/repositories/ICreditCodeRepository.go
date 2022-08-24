package repositories

import (
	"arvan-credit-service/db/models"
	"arvan-credit-service/types/structs"
)

type ICreditCodeRepository interface {
	FindByText(text string) structs.RepositoryResult[models.Code]
	RunSpendCodeDBFunction(code string, referenceId uint, clientId uint) (uint, error)
	InsertNewUsageLog(code string, referenceId uint) (uint, uint, error)
}

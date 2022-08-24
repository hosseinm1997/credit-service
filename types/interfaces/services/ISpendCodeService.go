package services

import (
	"arvan-credit-service/types/interfaces/repositories"
	"arvan-credit-service/types/structs"
)

type ISpendCodeService interface {
	SetRepo(repository repositories.ICreditCodeRepository)
	SetClientId(clientId uint)
	Spend(code string, referenceId uint) (uint, uint, *structs.CustomError)
}

package providers

import (
	"arvan-credit-service/db/repositories"
	"arvan-credit-service/infrastructures"
	repositoryInterfaces "arvan-credit-service/types/interfaces/repositories"
)

func RepositoryProviders() {
	infrastructures.Register[repositoryInterfaces.ICreditCodeRepository](
		func(params ...any) repositoryInterfaces.ICreditCodeRepository {
			return &repositories.CreditCodeRepository{}
		},
	)
}

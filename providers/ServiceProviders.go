package providers

import (
	"arvan-credit-service/infrastructures"
	"arvan-credit-service/services"
	"arvan-credit-service/types/interfaces/repositories"
	serviceInterfaces "arvan-credit-service/types/interfaces/services"
)

func ServiceProvider() {

	infrastructures.Register[serviceInterfaces.IClientSecretChecker](
		func(params ...any) serviceInterfaces.IClientSecretChecker {
			return &services.ClientSecretChecker{}
		},
	)

	infrastructures.Register[serviceInterfaces.ISpendCodeService](
		func(params ...any) serviceInterfaces.ISpendCodeService {
			s := &services.SpendCodeService{}
			s.SetRepo(infrastructures.Resolve[repositories.ICreditCodeRepository]())
			s.SetClientId(params[0].(uint))
			return s
		},
	)
}

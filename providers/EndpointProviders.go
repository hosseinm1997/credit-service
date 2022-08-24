package providers

import (
	"arvan-credit-service/http/endpoints"
	"arvan-credit-service/infrastructures"
	endpointInterfaces "arvan-credit-service/types/interfaces/endpoints"
)

func EndpointProviders() {
	infrastructures.Register[endpointInterfaces.ISpendCodeEndpoint](
		func(params ...any) endpointInterfaces.ISpendCodeEndpoint {
			return endpoints.SpendCode{}
		},
	)
}

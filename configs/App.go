package configs

import (
	"arvan-credit-service/providers"
	"arvan-credit-service/types/structs"
	"github.com/spf13/viper"
)

func App() map[string]any {
	return structs.Config{
		"environment": viper.GetString("APP_ENV"),
		"providers": []func(){
			providers.SystemProvides,
			providers.RepositoryProviders,
			providers.ServiceProvider,
			providers.EndpointProviders,
		},
	}
}

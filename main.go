package main

import (
	"arvan-credit-service/configs"
	"arvan-credit-service/http/routing"
	"arvan-credit-service/infrastructures"
	"arvan-credit-service/infrastructures/interfaces"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	infrastructures.KernelBuilder().Build(configs.App())
	routingSystem := infrastructures.Resolve[interfaces.IChiRouter]()

	// todo: use logger service for following error
	http.ListenAndServe(
		getAddress(),
		routingSystem.InitRouter(routing.Routes),
	)

}

func getAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
}

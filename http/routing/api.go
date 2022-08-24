package routing

import (
	"arvan-credit-service/infrastructures"
	"arvan-credit-service/types/interfaces/endpoints"
	"github.com/go-chi/chi"
)

func Routes(r *chi.Mux) {
	r.Post("/credit/code/{code}/{referenceId}", infrastructures.Resolve[endpoints.ISpendCodeEndpoint]().Spend)
	r.Post("/credit/code/{code}/inquiry", infrastructures.Resolve[endpoints.ISpendCodeEndpoint]().Inquiry)
}

package endpoints

import "net/http"

type ISpendCodeEndpoint interface {
	Spend(res http.ResponseWriter, req *http.Request)
	Inquiry(res http.ResponseWriter, req *http.Request)
}

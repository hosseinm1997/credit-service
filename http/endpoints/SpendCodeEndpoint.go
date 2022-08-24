package endpoints

import (
	"arvan-credit-service/infrastructures"
	"arvan-credit-service/types/interfaces/repositories"
	"arvan-credit-service/types/interfaces/services"
	"arvan-credit-service/types/structs"
	. "arvan-credit-service/utils"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type SpendCode struct{}

func (r SpendCode) Spend(res http.ResponseWriter, req *http.Request) {

	// todo: make a struct for response
	clientId, code, referenceId, err := r.prepareParameters(req)

	if err != nil {
		Respond(req).WithError(err)
		return
	}

	// todo: mobile validation
	service := infrastructures.Resolve[services.ISpendCodeService](clientId)
	amount, usageLogId, err := service.Spend(code, referenceId)

	if err != nil {
		Respond(req).WithError(err)
		return
	}

	Respond(req).WithOkResult(map[string]any{"amount": amount, "log_id": usageLogId})
}

func (r SpendCode) prepareParameters(req *http.Request) (uint, string, uint, *structs.CustomError) {
	clientId, err := r.getClientId(req)
	if err != nil {
		return 0, "", 0, err
	}

	code := chi.URLParam(req, "code")
	referenceId, convErr := strconv.ParseUint(chi.URLParam(req, "referenceId"), 10, 64)

	if convErr != nil {
		return 0, "", 0, CustomError(structs.Categories.BusinessLogic, "invalid reference id [%s]", chi.URLParam(req, "referenceId"))
	}

	return clientId, code, uint(referenceId), nil
}

func (r SpendCode) getClientId(req *http.Request) (uint, *structs.CustomError) {
	token := req.Header.Get("X-Client-Token")

	if token == "" {
		return 0, CustomError(structs.Categories.UnAuthorized, "no secret provided")
	}

	clientId, ok := infrastructures.Resolve[services.IClientSecretChecker]().GetSecret(token)

	if !ok {
		return 0, CustomError(structs.Categories.UnAuthorized, "invalid secret provided")
	}

	return clientId, nil
}

func (r SpendCode) Inquiry(res http.ResponseWriter, req *http.Request) {

	_, err := r.getClientId(req)
	if err != nil {
		Respond(req).WithError(err)
		return
	}

	code := chi.URLParam(req, "code")
	creditCode := infrastructures.Resolve[repositories.ICreditCodeRepository]().FindByText(code)

	if creditCode.RowsAffected == 0 {
		Respond(req).WithBusinessLogicExceptionResult(fmt.Errorf("No such credit code"))
		return
	}

	if creditCode.Model.CurrentUsedCount >= creditCode.Model.MaxUsableCount {
		Respond(req).WithBusinessLogicExceptionResult(fmt.Errorf("Credit code limitation reached"))
		return
	}

	Respond(req).WithOkResult(map[string]any{
		"status": "Credit code is usable",
	})
}

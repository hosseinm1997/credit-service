package utils

import (
	"arvan-credit-service/infrastructures"
	infrastructureInterfaces "arvan-credit-service/infrastructures/interfaces"
	"arvan-credit-service/types/structs"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

func DB() *gorm.DB {
	return infrastructures.DB()
}

func CustomError(category structs.ErrorCategory, err string, bindings ...any) *structs.CustomError {
	return &structs.CustomError{Err: fmt.Sprintf(err, bindings...), Category: category}
}

func Respond(req *http.Request) infrastructureInterfaces.IResponseFormatter {
	return infrastructures.Resolve[infrastructureInterfaces.IResponseFormatter](req)
}

func RepoResult[T any](data any, response *gorm.DB) structs.RepositoryResult[T] {
	rr := structs.RepositoryResult[T]{}
	rr.Set(data, response)
	return rr
}

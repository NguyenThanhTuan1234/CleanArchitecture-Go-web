package driver

import (
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewSignupUsecase(handlerRepo)
	err := use.CreateSignupPage(w, r)
	if err != nil {
		panic(err)
	}
}

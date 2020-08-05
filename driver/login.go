package driver

import (
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewLoginUsecase(handlerRepo)
	err := use.CreateLoginPage(w, r)
	if err != nil {
		panic(err)
	}
}

package driver

import (
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewUserUsecase(handlerRepo)
	err := use.CreateUserPage(w, r)
	if err != nil {
		panic(err)
	}
}

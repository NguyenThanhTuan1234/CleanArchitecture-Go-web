package driver

import (
	"CleanArchitecture-Go-web/adapter/controller"
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	file := &controller.FileControllerPreset{
		Value: "nf",
	}
	sessionRepo := gateway.NewSessionClient()
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewUserUsecase(file, sessionRepo, handlerRepo)
	err := use.CreateUserPage(w, r)
	if err != nil {
		panic(err)
	}
}

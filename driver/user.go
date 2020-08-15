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
	fileRepo := gateway.NewFileRepository()
	paramIn := &controller.ParamControllerPreset{
		Param: "2ee4d15b6e7b6898dcda631a426ccc2234f28cc1.JPG",
	}
	cookieRepo := gateway.NewCookieRepository()
	sessionRepo := gateway.NewSessionClient()
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewUserUsecase(file, paramIn, fileRepo, cookieRepo, sessionRepo, handlerRepo)
	err := use.CreateUserPage(w, r)
	if err != nil {
		panic(err)
	}
}

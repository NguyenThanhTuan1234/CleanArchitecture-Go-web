package driver

import (
	"CleanArchitecture-Go-web/adapter/controller"
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/config"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	file := &controller.FileControllerPreset{
		Value: "nf",
	}
	fileRepo := gateway.NewFileRepository()
	paramIn := &controller.ParamControllerPreset{
		Param: "2ee4d15b6e7b6898dcda631a426ccc2234f28cc1.JPG",
	}
	cookieRepo := gateway.NewCookieRepository()
	sessionRepo := gateway.NewSessionClient()
	rdbclient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbclient)
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewUserUsecase(file, paramIn, fileRepo, cookieRepo, sessionRepo, postgresRepo, handlerRepo)
	err1 := use.CreateUserPage(w, r)
	if err1 != nil {
		panic(err1)
	}
}

package driver

import (
	"CleanArchitecture-Go-web/adapter/controller"
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/config"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()
	formIn1 := &controller.FormControllerPreset{
		Value: "username",
	}
	formIn2 := &controller.FormControllerPreset{
		Value: "password",
	}
	formIn3 := &controller.FormControllerPreset{
		Value: "firstname",
	}
	formIn4 := &controller.FormControllerPreset{
		Value: "lastname",
	}
	formIn5 := &controller.FormControllerPreset{
		Value: "role",
	}
	rdbclient, err := gateway.NewRDBClient(conf.GetPostgres())
	if err != nil {
		panic(err)
	}
	postgresRepo := gateway.NewPostgresRepository(rdbclient)
	bcryptRepo := gateway.NewBcrypt()
	sessionRepo := gateway.NewSessionClient()
	sessionname := &controller.ParamControllerPreset{
		Param: "session",
	}
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewSignupUsecase(formIn1, formIn2, formIn3, formIn4, formIn5, postgresRepo, bcryptRepo, sessionname, sessionRepo, handlerRepo)
	err1 := use.CreateSignupPage(w, r)
	if err1 != nil {
		panic(err)
	}
}

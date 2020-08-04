package driver

import (
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/usecase"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	handlerRepo := gateway.NewTemplateRepository()
	use := usecase.NewIndexUsecase(handlerRepo)
	err := use.CreateIndexPage(w, r)
	if err != nil {
		panic(err)
	}
}

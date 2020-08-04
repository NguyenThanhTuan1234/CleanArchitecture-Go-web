package main

import (
	"CleanArchitecture-Go-web/adapter/gateway"
	"CleanArchitecture-Go-web/config"
	"CleanArchitecture-Go-web/usecase"
)

func main() {
	// conf := config.GetConfig()
	x, err := gateway.NewRDBClient(config.GetConfig().GetPostgres())
	if err != nil {
		panic(err)
	}
	repo := gateway.NewPostgresRepository(x)
	use := usecase.NewUserUsecase(repo)
	err1 := use.CreateUser()
	if err1 != nil {
		panic(err)
	}
}

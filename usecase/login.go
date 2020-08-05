package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type LoginUsecase interface {
	service.LoginService
}

type loginUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewLoginUsecase(
	handlerRepo repository.HandlerRepository,
) LoginUsecase {
	return &loginUsecase{
		handlerRepo: handlerRepo,
	}
}

package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type UserUsecase interface {
	service.UserService
}

type userUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewUserUsecase(
	handlerRepo repository.HandlerRepository,
) UserUsecase {
	return &userUsecase{
		handlerRepo: handlerRepo,
	}
}

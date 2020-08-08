package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type UserUsecase interface {
	service.UserService
}

type userUsecase struct {
	fileIn      FileInput
	sessionRepo repository.SessionRepository
	handlerRepo repository.HandlerRepository
}

func NewUserUsecase(
	fileIn FileInput,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) UserUsecase {
	return &userUsecase{
		fileIn:      fileIn,
		sessionRepo: sessionRepo,
		handlerRepo: handlerRepo,
	}
}

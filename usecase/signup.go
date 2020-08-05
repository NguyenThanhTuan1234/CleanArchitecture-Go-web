package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type SignupUsecase interface {
	service.SignupService
}

type signupUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewSignupUsecase(
	handlerRepo repository.HandlerRepository,
) SignupUsecase {
	return &signupUsecase{
		handlerRepo: handlerRepo,
	}
}

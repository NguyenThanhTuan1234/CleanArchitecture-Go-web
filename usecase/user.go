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
	paramIn     ParamInput
	fileRepo    repository.FileRepository
	cookieRepo  repository.CookieRepository
	sessionRepo repository.SessionRepository
	handlerRepo repository.HandlerRepository
}

func NewUserUsecase(
	fileIn FileInput,
	paramIn ParamInput,
	fileRepo repository.FileRepository,
	cookieRepo repository.CookieRepository,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) UserUsecase {
	return &userUsecase{
		fileIn:      fileIn,
		paramIn:     paramIn,
		fileRepo:    fileRepo,
		cookieRepo:  cookieRepo,
		sessionRepo: sessionRepo,
		handlerRepo: handlerRepo,
	}
}

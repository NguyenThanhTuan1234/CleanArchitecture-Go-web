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
	fileRepo    repository.FileRepository
	cookieRepo  repository.CookieRepository
	sessionRepo repository.SessionRepository
	handlerRepo repository.HandlerRepository
}

func NewUserUsecase(
	fileIn FileInput,
	fileRepo repository.FileRepository,
	cookieRepo repository.CookieRepository,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) UserUsecase {
	return &userUsecase{
		fileIn:      fileIn,
		fileRepo:    fileRepo,
		cookieRepo:  cookieRepo,
		sessionRepo: sessionRepo,
		handlerRepo: handlerRepo,
	}
}

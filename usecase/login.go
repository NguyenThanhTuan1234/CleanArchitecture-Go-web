package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type LoginUsecase interface {
	service.LoginService
}

type loginUsecase struct {
	formIn1      FormInput
	formIn2      FormInput
	postgresRepo repository.PostgresRepository
	bcryptRepo   repository.BcryptRepository
	param        ParamInput
	sessionRepo  repository.SessionRepository
	handlerRepo  repository.HandlerRepository
}

func NewLoginUsecase(
	formIn1 FormInput,
	formIn2 FormInput,
	postgresRepo repository.PostgresRepository,
	bcryptRepo repository.BcryptRepository,
	param ParamInput,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) LoginUsecase {
	return &loginUsecase{
		formIn1:      formIn1,
		formIn2:      formIn2,
		postgresRepo: postgresRepo,
		bcryptRepo:   bcryptRepo,
		param:        param,
		sessionRepo:  sessionRepo,
		handlerRepo:  handlerRepo,
	}
}

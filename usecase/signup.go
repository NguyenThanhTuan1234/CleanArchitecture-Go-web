package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type SignupUsecase interface {
	service.SignupService
}

type signupUsecase struct {
	formIn1      FormInput
	formIn2      FormInput
	formIn3      FormInput
	formIn4      FormInput
	formIn5      FormInput
	postgresRepo repository.PostgresRepository
	bcryptRepo   repository.BcryptRepository
	param        ParamInput
	sessionRepo  repository.SessionRepository
	handlerRepo  repository.HandlerRepository
}

func NewSignupUsecase(
	formIn1 FormInput,
	formIn2 FormInput,
	formIn3 FormInput,
	formIn4 FormInput,
	formIn5 FormInput,
	postgresRepo repository.PostgresRepository,
	bcryptRepo repository.BcryptRepository,
	param ParamInput,
	sessionRepo repository.SessionRepository,
	handlerRepo repository.HandlerRepository,
) SignupUsecase {
	return &signupUsecase{
		formIn1:      formIn1,
		formIn2:      formIn2,
		formIn3:      formIn3,
		formIn4:      formIn4,
		formIn5:      formIn5,
		postgresRepo: postgresRepo,
		bcryptRepo:   bcryptRepo,
		param:        param,
		sessionRepo:  sessionRepo,
		handlerRepo:  handlerRepo,
	}
}

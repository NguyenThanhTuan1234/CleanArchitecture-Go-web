package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type UserUsecase interface {
	service.UserService
}

type userUsecase struct {
	postgresRepo repository.PostgresRepository
}

func NewUserUsecase(postgresRepo repository.PostgresRepository) UserUsecase {
	return &userUsecase{
		postgresRepo: postgresRepo,
	}
}

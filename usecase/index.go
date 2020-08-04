package usecase

import (
	"CleanArchitecture-Go-web/entity/repository"
	"CleanArchitecture-Go-web/entity/service"
)

type IndexUsecase interface {
	service.IndexService
}

type indexUsecase struct {
	handlerRepo repository.HandlerRepository
}

func NewIndexUsecase(
	handlerRepo repository.HandlerRepository,
) IndexUsecase {
	return &indexUsecase{
		handlerRepo: handlerRepo,
	}
}

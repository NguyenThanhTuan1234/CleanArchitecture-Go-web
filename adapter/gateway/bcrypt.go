package gateway

import "CleanArchitecture-Go-web/entity/repository"

type bcyptClient struct{}

type BcryptClient interface {
	repository.BcryptRepository
}

func NewBcrypt() BcryptClient {
	return &bcyptClient{}
}

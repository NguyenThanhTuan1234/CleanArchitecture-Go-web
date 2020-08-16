package gateway

import "CleanArchitecture-Go-web/entity/repository"

type fileRepository struct{}

type FileRepository interface {
	repository.FileRepository
}

func NewFileRepository() FileRepository {
	return &fileRepository{}
}

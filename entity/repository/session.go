package repository

import (
	"CleanArchitecture-Go-web/entity"
	"mime/multipart"
	"net/http"
)

type SessionRepository interface {
	CreateSession(http.ResponseWriter, *http.Request, string, int) error
	CheckSessionIfExist(http.ResponseWriter, *http.Request, string) bool
	DeleteSession(http.ResponseWriter, *http.Request, string) error
	GetSessionInfo(http.ResponseWriter, *http.Request) (*entity.Session, error)
	UploadImage(multipart.File, *multipart.FileHeader) (string, error)
}

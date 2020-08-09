package repository

import (
	"CleanArchitecture-Go-web/entity"
	"net/http"
)

type SessionRepository interface {
	CreateSession(http.ResponseWriter, *http.Request, string, int) error
	CheckSessionIfExist(http.ResponseWriter, *http.Request, string) bool
	DeleteSession(http.ResponseWriter, *http.Request, string) error
	GetSessionInfo(http.ResponseWriter, *http.Request) (*entity.Session, error)
}

package repository

import (
	"net/http"
)

type HandlerRepository interface {
	Index(http.ResponseWriter, *http.Request) error
	User(http.ResponseWriter, *http.Request) error
	Login(http.ResponseWriter, *http.Request) error
}

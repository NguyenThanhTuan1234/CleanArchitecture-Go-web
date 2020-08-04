package service

import "net/http"

type UserService interface {
	CreateUserPage(http.ResponseWriter, *http.Request) error
}

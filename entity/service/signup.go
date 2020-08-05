package service

import "net/http"

type SignupService interface {
	CreateSignupPage(http.ResponseWriter, *http.Request) error
}

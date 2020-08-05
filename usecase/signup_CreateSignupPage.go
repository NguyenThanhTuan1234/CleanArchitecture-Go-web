package usecase

import (
	"net/http"
)

func (i *signupUsecase) CreateSignupPage(w http.ResponseWriter, r *http.Request) error {
	err := i.handlerRepo.Signup(w, r)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
	"net/http"
)

func (i *loginUsecase) CreateLoginPage(w http.ResponseWriter, r *http.Request) error {
	err := i.handlerRepo.Login(w, r)
	if err != nil {
		return err
	}
	return nil
}

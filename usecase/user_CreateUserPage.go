package usecase

import (
	"net/http"
)

func (i *userUsecase) CreateUserPage(w http.ResponseWriter, r *http.Request) error {
	err := i.handlerRepo.User(w, r)
	if err != nil {
		return err
	}
	return nil
}

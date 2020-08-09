package usecase

import (
	"net/http"
	"strings"
)

func (u *userUsecase) CreateUserPage(w http.ResponseWriter, r *http.Request) error {
	var xs []string
	c := u.cookieRepo.GetCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := u.fileIn.GetFile(w, r)
		if err != nil {
			return err
		}
		fname, err1 := u.sessionRepo.UploadImage(mf, fh)
		if err1 != nil {
			return err1
		}
		c = u.cookieRepo.AppendValue(w, c, fname)
		xs = strings.Split(c.Value, "|")
	}
	if xs == nil {
		err2 := u.handlerRepo.User(w, r, "2ee4d15b6e7b6898dcda631a426ccc2234f28cc1.JPG")
		if err2 != nil {
			return err2
		}
	} else {
		err2 := u.handlerRepo.User(w, r, xs[len(xs)-1])
		if err2 != nil {
			return err2
		}
	}
	return nil
}

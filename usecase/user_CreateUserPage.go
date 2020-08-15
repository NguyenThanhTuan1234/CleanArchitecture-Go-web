package usecase

import (
	"fmt"
	"net/http"
	"strings"
)

func (u *userUsecase) CreateUserPage(w http.ResponseWriter, r *http.Request) error {
	var xs []string
	var filename string
	var err3 error
	var c = u.cookieRepo.GetCookie(w, r)
	fmt.Println(c.Value)
	if r.Method == http.MethodPost {
		mf, fh, err := u.fileIn.GetFile(w, r)
		if err != nil {
			return err
		}
		fname, err1 := u.fileRepo.UploadImage(mf, fh)
		if err1 != nil {
			return err1
		}
		xs = strings.Split(c.Value, "|")
		// fmt.Println(xs)
		session, err := u.sessionRepo.GetSessionInfo(w, r)
		c = u.cookieRepo.AppendValue(w, c, fname, session.Id)
		xs = strings.Split(c.Value, "|")
		// fmt.Println(xs)
		filename, err3 = u.fileRepo.MapUidToImage(session.Id, xs[len(xs)-1])
		if err3 != nil {
			return err3
		}
	}
	xs = strings.Split(c.Value, "|")
	fmt.Println(xs)
	session, _ := u.sessionRepo.GetSessionInfo(w, r)
	filename, err3 = u.fileRepo.MapUidToImage(session.Id, xs[len(xs)-1])
	if err3 != nil {
		return err3
	}
	if len(xs) < 2 {
		err2 := u.handlerRepo.User(w, r, "2ee4d15b6e7b6898dcda631a426ccc2234f28cc1.JPG")
		if err2 != nil {
			return err2
		}
	} else {
		err2 := u.handlerRepo.User(w, r, filename)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

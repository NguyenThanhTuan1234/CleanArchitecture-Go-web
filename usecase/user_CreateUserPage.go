package usecase

import (
	"net/http"
	"strings"
)

func (u *userUsecase) CreateUserPage(w http.ResponseWriter, r *http.Request) error {
	var xs []string
	var filename string
	var err3 error
	var c = u.cookieRepo.GetCookie(w, r)
	if r.Method == http.MethodPost {
		mf, fh, err := u.fileIn.GetFile(w, r)
		if err != nil {
			return err
		}
		fname, err1 := u.fileRepo.UploadImage(mf, fh)
		if err1 != nil {
			return err1
		}
		session, err := u.sessionRepo.GetSessionInfo(w, r)
		c = u.cookieRepo.AppendValue(w, c, fname, session.Id)
		xs = strings.Split(c.Value, "|")
		filename, err3 = u.fileRepo.MapUidToImage(session.Id, xs[len(xs)-1])
		if err3 != nil {
			return err3
		}
	}
	xs = strings.Split(c.Value, "|")
	session, _ := u.sessionRepo.GetSessionInfo(w, r)
	filename, err3 = u.fileRepo.MapUidToImage(session.Id, xs[len(xs)-1])
	if err3 != nil {
		return err3
	}
	defaultImg, err := u.paramIn.ReadParam()
	if err != nil {
		return err
	}
	if len(xs) < 2 {
		err2 := u.handlerRepo.User(w, r, defaultImg)
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

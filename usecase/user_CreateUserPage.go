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
		err4 := u.postgresRepo.SaveImage(filename, session.Id)
		if err != nil {
			return err4
		}
	}
	xs = strings.Split(c.Value, "|")
	fmt.Println(xs)
	session, _ := u.sessionRepo.GetSessionInfo(w, r)
	// filename, err3 = u.fileRepo.MapUidToImage(session.Id, xs[len(xs)-1])
	// if err3 != nil {
	// 	return err3
	// }
	img, err := u.postgresRepo.GetImage(session.Id)
	fmt.Println(img)
	fmt.Println(session.Id)
	if err != nil {
		return err
	}
	defaultImg, err := u.paramIn.ReadParam()
	if err != nil {
		return err
	}
	if img.Name == "" {
		err2 := u.handlerRepo.User(w, r, defaultImg)
		if err2 != nil {
			return err2
		}
	} else {
		err2 := u.handlerRepo.User(w, r, img.GetName())
		if err2 != nil {
			return err2
		}
	}
	return nil
}

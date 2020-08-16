package usecase

import (
	"CleanArchitecture-Go-web/entity"
	"fmt"
	"net/http"
)

func (u *loginUsecase) CreateLoginPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err := u.param.ReadParam()
	if err != nil {
		return err
	}
	if u.sessionRepo.CheckSessionIfExist(w, r, sessionname) == true {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
		return nil
	}
	var user *entity.User
	var err5 error
	if r.Method == http.MethodPost {
		username, err := u.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		password, err := u.formIn2.GetFormValue(w, r)
		if err != nil {
			return err
		}
		user, err5 = u.postgresRepo.GetUser(username)
		if err5 != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return nil
		}
		err1 := u.bcryptRepo.ComparePassword(user.GetPassWord(), password)
		if err1 != nil {
			fmt.Println("test")
			http.Error(w, "Username or password do not match", http.StatusForbidden)
			return nil
		}
		sessionname, err := u.param.ReadParam()
		if err != nil {
			return err
		}
		err3 := u.sessionRepo.CreateSession(w, r, sessionname, user.GetId())
		if err3 != nil {
			return err3
		}
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
	if user == nil {
		err4 := u.handlerRepo.Login(w, r)
		if err4 != nil {
			return err4
		}
	}
	return nil
}

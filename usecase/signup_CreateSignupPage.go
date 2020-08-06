package usecase

import (
	"net/http"
)

func (u *signupUsecase) CreateSignupPage(w http.ResponseWriter, r *http.Request) error {
	sessionname, err := u.param.ReadParam()
	if err != nil {
		return err
	}
	if u.sessionRepo.CheckSessionIfExist(w, r, sessionname) == true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
	if r.Method == http.MethodPost {
		username, err := u.formIn1.GetFormValue(w, r)
		if err != nil {
			return err
		}
		password, err := u.formIn2.GetFormValue(w, r)
		if err != nil {
			return err
		}
		first, err := u.formIn3.GetFormValue(w, r)
		if err != nil {
			return err
		}
		last, err := u.formIn4.GetFormValue(w, r)
		if err != nil {
			return err
		}
		role, err := u.formIn5.GetFormValue(w, r)
		if err != nil {
			return err
		}
		user, err := u.postgresRepo.GetUser(username)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return nil
		}
		if user.GetUserName() != "" {
			http.Error(w, "Username already taken", http.StatusForbidden)
		}
		sessionname, err := u.param.ReadParam()
		if err != nil {
			return err
		}
		err2 := u.sessionRepo.CreateSession(w, r, sessionname, user.GetId())
		if err != nil {
			return err2
		}
		hashPassword, err := u.bcryptRepo.GenerateHashPassword(password)
		if err != nil {
			return err
		}
		err3 := u.postgresRepo.CreateUser(username, hashPassword, first, last, role)
		if err3 != nil {
			return err3
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	err4 := u.handlerRepo.Signup(w, r)
	if err4 != nil {
		return err4
	}
	return nil
}

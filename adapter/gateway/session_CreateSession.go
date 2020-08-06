package gateway

import (
	"CleanArchitecture-Go-web/entity"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

const sessionLength int = 600

var session *entity.Session
var dbSessions = map[string]*entity.Session{}

func (c *sessionRepository) CreateSession(w http.ResponseWriter, r *http.Request, name string, id int) error {

	sID, err := uuid.NewV4()
	fmt.Println(sID.String())
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   name,
		Value:  sID.String(),
		MaxAge: sessionLength,
	}
	http.SetCookie(w, cookie)
	dbSessions[cookie.Value] = &entity.Session{id, time.Now()}
	fmt.Println(dbSessions[cookie.Value].Id)
	return nil
}

package gateway

import (
	"CleanArchitecture-Go-web/entity"
	"net/http"
)

func (c *sessionRepository) GetSessionInfo(w http.ResponseWriter, r *http.Request) (*entity.Session, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}
	return dbSessions[cookie.Value], nil
}

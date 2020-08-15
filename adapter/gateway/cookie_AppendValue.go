package gateway

import (
	"CleanArchitecture-Go-web/entity"
	"net/http"
	"strings"
	"time"
)

func (r *cookieRepository) AppendValue(w http.ResponseWriter, c *http.Cookie, fname string, id int) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	dbSessions[c.Value] = &entity.Session{id, time.Now()}
	return c
}

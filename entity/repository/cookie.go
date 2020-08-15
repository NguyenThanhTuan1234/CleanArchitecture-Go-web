package repository

import "net/http"

type CookieRepository interface {
	GetCookie(http.ResponseWriter, *http.Request) *http.Cookie
	AppendValue(http.ResponseWriter, *http.Cookie, string, int) *http.Cookie
}

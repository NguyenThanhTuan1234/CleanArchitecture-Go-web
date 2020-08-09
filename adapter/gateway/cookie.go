package gateway

import "CleanArchitecture-Go-web/entity/repository"

type cookieRepository struct{}

type CookieRepository interface {
	repository.CookieRepository
}

func NewCookieRepository() CookieRepository {
	return &cookieRepository{}
}

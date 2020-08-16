package repository

import "CleanArchitecture-Go-web/entity"

//　PostgresRepository　データベースとやりとる先
type PostgresRepository interface {
	GetUser(string) (*entity.User, error)
	CreateUser(string, string, string, string, string) error
	GetImage(int) (*entity.Image, error)
	SaveImage(string, int) error
	UpdateImage(string, int) error
}

package repository

import "CleanArchitecture-Go-web/entity"

//　PostgresRepository　データベースとやりとる先
type PostgresRepository interface {
	GetUser(string) (*entity.User, error)
}

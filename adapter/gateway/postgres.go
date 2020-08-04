package gateway

import (
	"CleanArchitecture-Go-web/entity/repository"
	"fmt"
)

// PostgresConfig PostgreSQL コンフィグ
type PostgresConfig interface {
	GetUser() string
	GetPassword() string
	GetDatabaseName() string
	GetSSLMode() string
}

// NewPostgresSpecific　PostgreSQLの接続情報を作成する
func NewPostgresSpecific(c PostgresConfig) RDBSpecific {
	return &rdbSpecific{
		driver: "postgres",
		source: fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", c.GetUser(), c.GetPassword(), c.GetDatabaseName(), c.GetSSLMode()),
	}
}

// PostgresRepository　PostgreSQLを用いるレポジトリ
type PostgresRepository interface {
	repository.PostgresRepository
}

type postgresRepository struct {
	client *rdbClient
}

// NewPostgresRepository　PostgreSQLを用いるレポジトリを作成する
func NewPostgresRepository(client *rdbClient) PostgresRepository {
	return &postgresRepository{
		client: client,
	}
}

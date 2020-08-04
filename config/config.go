package config

import "CleanArchitecture-Go-web/adapter/gateway"

// Config　コンフィグ
type Config interface {
	GetPostgres() gateway.PostgresConfig
}

var conf Config

// GetConfig　コンフィグを返す
func GetConfig() Config {
	return conf
}

type config struct {
	Postgres *postgresConfig `josn:"postgres"`
}

func (c *config) GetPostgres() gateway.PostgresConfig {
	return c.Postgres
}

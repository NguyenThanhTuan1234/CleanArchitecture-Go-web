package config

type postgresConfig struct {
	User         string `json: "user"`
	Password     string `json: "password"`
	DatabaseName string `json: "databasename"`
	SSLMode      string `json: "sslmode"`
}

func (p *postgresConfig) GetUser() string {
	return p.User
}

func (p *postgresConfig) GetPassword() string {
	return p.Password
}

func (p *postgresConfig) GetDatabaseName() string {
	return p.DatabaseName
}

func (p *postgresConfig) GetSSLMode() string {
	return p.SSLMode
}

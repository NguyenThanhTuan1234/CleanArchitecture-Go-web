package gateway

import (
	"database/sql"
	"reflect"
	"testing"
)

type postgresConfigMock struct {
	user         string
	password     string
	databaseName string
	sslMode      string
}

func (c *postgresConfigMock) GetUser() string         { return c.user }
func (c *postgresConfigMock) GetPassword() string     { return c.password }
func (c *postgresConfigMock) GetDatabaseName() string { return c.databaseName }
func (c *postgresConfigMock) GetSSLMode() string      { return c.sslMode }

type rdbClientMock struct {
	db *sql.DB
}

func TestNewPostgresSpecific(t *testing.T) {
	type args struct {
		c PostgresConfig
	}
	tests := []struct {
		name string
		args args
		want RDBSpecific
	}{
		{
			name: "test",
			args: args{
				c: &postgresConfigMock{
					user:         "user",
					password:     "password",
					databaseName: "databaseName",
					sslMode:      "sslMode",
				},
			},
			want: &rdbSpecific{
				driver: "postgres",
				source: "user=user password=password dbname=databaseName sslmode=sslMode",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostgresSpecific(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresSpecific() = %v, want %v", got, tt.want)
			}
		})
	}
}

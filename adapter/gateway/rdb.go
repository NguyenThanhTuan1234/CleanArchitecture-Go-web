package gateway

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// RDBSpecific　RDB接続情報
type RDBSpecific interface {
	GetDriver() string
	GetSource() string
}

type rdbSpecific struct {
	driver string
	source string
}

func (s *rdbSpecific) GetDriver() string { return s.driver }
func (s *rdbSpecific) GetSource() string { return s.source }

type rdbClient struct {
	db *sql.DB
}

func NewRDBClient(config PostgresConfig) (*rdbClient, error) {
	specific := NewPostgresSpecific(config)
	db, err := sql.Open(specific.GetDriver(), specific.GetSource())
	if err != nil {
		return nil, err
	}
	// fmt.Println("You connected to your database")
	return &rdbClient{
		db: db,
	}, nil
}

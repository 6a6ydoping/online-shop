package pg

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const usersTable = "users" // TODO: replace to config

type Postgres struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	DB       sqlx.DB
}

func New(opts ...Option) (*Postgres, error) {
	db := new(Postgres)

	for _, opt := range opts {
		opt(db)
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", db.host, db.username, db.password, db.port, db.dbName)

	database, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := database.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{
		host:     db.host,
		username: db.username,
		password: db.password,
		port:     db.port,
		dbName:   db.dbName,
		DB:       *database,
	}, nil
}

func (p *Postgres) Close() {
	if p.DB.Close() != nil {
		p.DB.Close()
	}
}

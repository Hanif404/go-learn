package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Handler interface {
	Rebind(string) string
	Exec(string, ...interface{}) (sql.Result, error)
	Get(interface{}, string, ...interface{}) error
}

func New(db *sqlx.DB) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db Handler
}

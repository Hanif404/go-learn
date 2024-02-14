package db

import (
	_ "embed"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Initialize() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:dev12345@tcp(localhost:3306)/level3?multiStatements=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

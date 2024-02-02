package db_client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DBClient *sql.DB

func InitialDBConn() {
	db, err := sql.Open("mysql", "sql6681335:vrdC25nRxr@tcp(sql6.freesqldatabase.com:3306)/sql6681335")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DBClient = db
}

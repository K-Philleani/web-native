package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

const DSN = "root:123456@tcp(124.70.71.78:3306)/webApi"

func init() {
	Db, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
}

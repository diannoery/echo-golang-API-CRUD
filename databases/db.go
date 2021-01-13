package databases

import (
	"database/sql"
	"echo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connect := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connect)
	if err != nil {
		panic("connec eroorr")
	}

	err = db.Ping()
	if err != nil {
		panic("DNS invalid")
	}
}

func CreateCon() *sql.DB {

	return db
}

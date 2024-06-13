package app

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup() {

	db_type := os.Getenv("DB_TYPE")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")

	d, err := sql.Open(db_type, db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_database)

	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *sql.DB {

	return db

}

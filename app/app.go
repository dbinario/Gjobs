package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
func Setup() error {

	db_type := os.Getenv("DB_TYPE")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_DATABASE")

	db, err := sql.Open(db_type, db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_database)

	if err != nil {

		return err

	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		return err
	}

	return nil

}*/

func Run() error {

	//en caso de que no existan las llaves
	if os.Getenv("PASETO_PRIVATE_KEY") == "" && os.Getenv("PASETO_PUBLIC_KEY") == "" {

		GenerarLlaves()

	}

	//iniciamos la conexion a la base de datos

	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	// Verifica que todas las variables de entorno necesarias están presentes
	if dbType == "" || dbUser == "" || dbHost == "" || dbPort == "" || dbDatabase == "" {
		return fmt.Errorf("Database configuration missing")
	}

	// Establece la conexión a la base de datos
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	db, err = sql.Open(dbType, dsn)
	if err != nil {
		return err
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		return err
	}

	return nil

}

func GetDB() *sql.DB {

	return db

}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

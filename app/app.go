package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Run() error {

	//en caso de que no existan las llaves
	if os.Getenv("PASETO_PRIVATE_KEY") == "" && os.Getenv("PASETO_PUBLIC_KEY") == "" {

		err := GenerarLlaves()
		if err != nil {
			return fmt.Errorf("failed to generate keys: %w", err)
		}

	}

	//iniciamos la conexion a la base de datos

	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	missingVars := []string{}

	if dbType == "" {
		missingVars = append(missingVars, "DB_TYPE")
	}
	/*if dbPass == "" {
		missingVars = append(missingVars, "DB_PASS")
	}*/
	if dbUser == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if dbHost == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if dbPort == "" {
		missingVars = append(missingVars, "DB_PORT")
	}
	if dbDatabase == "" {
		missingVars = append(missingVars, "DB_DATABASE")
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing environment variables: %v", missingVars)
	}

	// Establece la conexión a la base de datos
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	db, err = sql.Open(dbType, dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
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

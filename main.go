package main

import (
	"gjobs-back/app"
	"gjobs-back/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configura la conexión a la base de datos
	if err := app.Run(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer app.CloseDB() // Asegúrate de cerrar la conexión al finalizar

	routes.Run()

}

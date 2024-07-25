package routes

import (
	"gjobs-back/controllers"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	getRoutes()

	// Obtiene el puerto desde la variable de entorno o usa el puerto por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto
	}

	// Inicia el servidor y maneja posibles errores
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

func getRoutes() {

	router.POST("/registrar", controllers.PostRegistrar)
	router.POST("/autenticar", controllers.PostAutenticar)

}

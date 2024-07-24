package routes

import (
	"gjobs-back/controllers"
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

	router.Run(":8080")
}

func getRoutes() {

	router.POST("/registrar", controllers.PostRegistrar)
	router.POST("/autenticar", controllers.PostAutenticar)

}

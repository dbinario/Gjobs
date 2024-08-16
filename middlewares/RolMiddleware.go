package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RolMiddleware(rol string) gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.GetString("rol") == rol {

			// Si cumple, continuar con la ejecución del siguiente manejador
			c.Next()
		} else {
			// Si no cumple, responder con un error
			c.JSON(403, gin.H{
				"error": "No cumples con el rol para ejecutar esta funcion",
			})

			// Detener la ejecución del resto de los manejadores
			c.Abort()
		}
	}

}

package middlewares

import (
	"net/http"
	"os"
	"strings"

	"aidanwoods.dev/go-paseto"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Buscamos si existe el Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		publicKeyHex := os.Getenv("PASETO_PUBLIC_KEY")
		publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid public key format"})
			c.Abort()
			return
		}

		parser := paseto.NewParser()
		tokendec, err := parser.ParseV4Public(publicKey, tokenString, nil)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		data := tokendec.Claims()

		// Validamos que los claims necesarios estén presentes
		idUsuario, idOk := data["id_usuario"]
		rol, rolOk := data["rol"]
		idEmpresa, idEmpresaOk := data["id_empresa"]
		if !idOk || !rolOk || !idEmpresaOk {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Asignamos los datos del token al contexto de Gin
		c.Set("id_usuario", idUsuario)
		c.Set("rol", rol)
		c.Set("id_empresa", idEmpresa)
		c.Next()
	}
}

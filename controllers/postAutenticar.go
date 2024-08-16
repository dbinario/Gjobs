package controllers

import (
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostAutenticar(c *gin.Context) {
	db := app.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var autenticar interfaces.Autenticar
	if err := c.ShouldBindJSON(&autenticar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int
	if err := db.QueryRow("CALL sp_buscarEmail(?)", autenticar.Email).Scan(&count); err != nil || count == 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "Email no existe"})
		return
	}

	var hash string
	if err := db.QueryRow("CALL sp_autenticarUsuario(?)", autenticar.Email).Scan(&hash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(autenticar.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error en el usuario o contraseña"})
		return
	}

	var id, nombreCompleto, rol, id_empresa string
	if err := db.QueryRow("CALL sp_autenticacionCorrecta(?)", autenticar.Email).Scan(&id, &nombreCompleto, &rol, &id_empresa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
		return
	}

	token := app.GenerateTokenPaseto(id, rol, id_empresa)
	c.JSON(http.StatusOK, gin.H{
		"token":          token,
		"id":             id,
		"nombreCompleto": nombreCompleto,
	})
}

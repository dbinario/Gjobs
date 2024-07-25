package controllers

import (
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAutenticar(c *gin.Context) {

	db := app.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var autenticar interfaces.Autenticar

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&autenticar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int

	db.QueryRow("CALL sp_buscarEmail(?)", autenticar.Email).Scan(&count)

	if count == 0 { //si email no existe

		c.JSON(http.StatusConflict, gin.H{
			"message": "Email no existe",
		})

		return

	} else { //si email si existe

	}

}

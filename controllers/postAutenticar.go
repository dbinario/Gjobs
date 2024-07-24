package controllers

import (
	"database/sql"
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAutenticar(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	//cerramos la conexion despues de usarla
	defer db.Close()

	var autenticar interfaces.Autenticar

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&autenticar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

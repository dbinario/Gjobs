package controllers

import (
	"database/sql"
	"gjobs-back/app"

	"github.com/gin-gonic/gin"
)

func PostAutenticar(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	//cerramos la conexion despues de usarla
	defer db.Close()

}

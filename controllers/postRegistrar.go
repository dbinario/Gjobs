package controllers

import (
	"database/sql"
	"fmt"
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostRegistrar(c *gin.Context) {

	var db *sql.DB

	//se obtiene la conexion de la base de datos

	app.Setup()
	db = app.GetDB()
	//cerramos la conexion despues de usarla
	defer db.Close()

	var registrar interfaces.Registrar

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&registrar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int

	db.QueryRow("CALL sp_buscarEmail(?)", registrar.Email).Scan(&count)

	if count > 0 { // ya existe el email

		c.JSON(http.StatusConflict, gin.H{
			"message": "Email ya existe",
		})

	} else {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrar.Password), 10)

		if err != nil {
			fmt.Println("Error al hashear la contraseña:", err)
			return
		}

		// Insertar datos en la tabla
		stmt, err := db.Prepare("CALL sp_registrarUsuario(?,?,?,?,?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
			return
		}
		defer stmt.Close()

		// Ejecutar la consulta con los parámetros
		_, err = stmt.Exec(registrar.Name, registrar.LastName, registrar.Email, hashedPassword, registrar.Rol)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el usuario"})
			return
		}

		db.QueryRow("CALL sp_buscarEmail(?)", registrar.Email).Scan(&count)

		if count != 1 {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el usuario"})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario Registrado de manera exitosa",
		})

	}

}

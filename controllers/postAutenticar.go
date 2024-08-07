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

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&autenticar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int

	db.QueryRow("CALL sp_buscarEmail(?)", autenticar.Email).Scan(&count)

	if count == 1 { //si email no existe

		var hash string

		db.QueryRow("CALL sp_autenticarUsuario(?)", autenticar.Email).Scan(&hash)

		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(autenticar.Password))

		if err == nil { //pasword correcto

			var id, nombreCompleto, rol string

			db.QueryRow("CALL sp_autenticacionCorrecta(?)", autenticar.Email).Scan(&id, &nombreCompleto, &rol)

			//fmt.Println(id, nombreCompleto, rol)

			token := app.GenerateTokenPaseto(id, rol)

			c.JSON(http.StatusOK, gin.H{
				"token":          token,
				"id":             id,
				"nombreCompleto": nombreCompleto,
			})
			return

		} else {

			//retornamos el error
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "Error en el usuario o contraseña"},
			)
			return

		}

	} else { //si email si existe

		c.JSON(http.StatusConflict, gin.H{
			"message": "Email no existe",
		})

		return

	}

}

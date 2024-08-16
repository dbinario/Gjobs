package empresa

import (
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostActualizarDatosEmpresa(c *gin.Context) {

	db := app.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var actualizarDatosEmpresa interfaces.ActualizarDatosEmpresa

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&actualizarDatosEmpresa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.QueryRow("CALL sp_actualizarDatosEmpresas(?,?)", c.GetString("id_empresa"), actualizarDatosEmpresa.NombreEmpresa).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar los datos de la empresa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Datos de la empresa actualizados de manera correcta",
	})

}

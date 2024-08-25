package controllers

import (
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVacante(c *gin.Context) {

	db := app.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	idVacante := c.Param("id")

	//c.JSON(200, gin.H{"id": IdVacante})

	var vacante interfaces.Vacante

	//este es el sp_recuperarVacante

	db.QueryRow("CALL sp_recuperarVacante(?)", idVacante).Scan(&vacante.IdVacante, &vacante.IdEmpresa, &vacante.NombreEmpresa, &vacante.NombreVacante, &vacante.TipoContratacion, &vacante.Horario, &vacante.Modalidad, &vacante.Municipio, &vacante.Estado, &vacante.RangoMin, &vacante.RangoMax, &vacante.Descripcion)

	c.JSON(http.StatusOK, gin.H{
		"data": vacante,
	})

}

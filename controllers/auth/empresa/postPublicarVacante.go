package empresa

import (
	"gjobs-back/app"
	"gjobs-back/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPublicarVacante(c *gin.Context) {

	db := app.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	var vacante interfaces.Vacante

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&vacante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insertar datos en la tabla
	stmt, err := db.Prepare("CALL sp_publicarVacante(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al preparar la consulta"})
		return
	}
	defer stmt.Close()

	// Ejecutar la consulta con los parámetros
	_, err = stmt.Exec(c.GetString("id_empresa"), vacante.NombreVacante, vacante.OcultarEmpresa, vacante.Contratacion, vacante.Horario, vacante.Modalidad, vacante.Municipio, vacante.Estado, vacante.RangoMin, vacante.RangoMax, vacante.OcultarRango, vacante.Descripcion)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la vacante"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La Vacante se ha publicado de manera exitosa",
	})

}

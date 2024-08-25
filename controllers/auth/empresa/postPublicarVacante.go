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

	var publicarVacante interfaces.PublicarVacante

	//validamos que vengan con los parametos que requerimos
	if err := c.ShouldBindJSON(&publicarVacante); err != nil {
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
	_, err = stmt.Exec(c.GetString("id_empresa"), publicarVacante.NombreVacante, publicarVacante.OcultarEmpresa, publicarVacante.Contratacion, publicarVacante.Horario, publicarVacante.Modalidad, publicarVacante.Municipio, publicarVacante.Estado, publicarVacante.RangoMin, publicarVacante.RangoMax, publicarVacante.OcultarRango, publicarVacante.Descripcion)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la vacante"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "La Vacante se ha publicado de manera exitosa",
	})

}

package handlers

import (
	"net/http"
	"test/aveonline-api/database"
	"test/aveonline-api/models"

	"github.com/gin-gonic/gin"
)

func Medicamentos(c *gin.Context) {
	medicamentos := []models.Medicamento{}
	

	if err := database.DB.Find(&medicamentos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, medicamentos)
}

func CreateMedicamento(c *gin.Context) {
	var medicamento models.Medicamento

	if err := c.ShouldBindJSON(&medicamento); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := database.DB.Create(&medicamento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, medicamento)
}
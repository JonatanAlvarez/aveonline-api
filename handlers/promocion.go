package handlers

import (
	"net/http"
	"test/aveonline-api/database"
	"test/aveonline-api/models"

	"github.com/gin-gonic/gin"
)

func Promociones(c *gin.Context) {
	promociones := []models.Promocion{}
	

	if err := database.DB.Find(&promociones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, promociones)
}

func CreatePromocion(c *gin.Context) {
	var promocion models.Promocion

	if err := c.ShouldBindJSON(&promocion); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if promocion.Porcentaje > models.MAX_PORCENTAJE  {
		c.JSON(http.StatusBadRequest, "El porcentaje aplicado exece el limite permitido.")
		return
	}

	promociones := []models.Promocion{}

	if err := database.DB.Find(&promociones).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	if promocion.Fecha_inicio.Unix() > promocion.Fecha_fin.Unix() {
		c.JSON(http.StatusBadRequest, "La fecha de inicio no puede ser mayor a la de fin.")
		return
	}

	if promocion.CheckDateCollision(promociones) {
		c.JSON(http.StatusBadRequest, "Ya existe una promocion en la fecha dada.")
		return
	}

	if err := database.DB.Create(&promocion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, promocion)
}
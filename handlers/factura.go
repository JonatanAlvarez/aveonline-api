package handlers

import (
	"net/http"
	"test/aveonline-api/database"
	"test/aveonline-api/models"

	"github.com/gin-gonic/gin"
)

func Facturas(c *gin.Context) {
	facturas := []models.Factura{}
	

	if err := database.DB.Find(&facturas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, facturas)
}

func CreateFactura(c *gin.Context) {
	var factura models.Factura

	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	factura.Pago_total = 0

	for _, item := range factura.Medicamentos {
		factura.Pago_total += item.Precio
	}

	if factura.Promocion.Porcentaje > 0 {
		factura.Pago_total -= (factura.Pago_total * factura.Promocion.Porcentaje)
	}

	if err := database.DB.Create(&factura).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
    return
	}

	c.JSON(http.StatusOK, factura)
}

func SimularFactura(c *gin.Context) {
	var factura models.Factura

	if err := c.ShouldBindJSON(&factura); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	factura.Pago_total = 0

	for _, item := range factura.Medicamentos {
		factura.Pago_total += item.Precio
	}

	if factura.Promocion.Porcentaje > 0 {
		factura.Pago_total -= (factura.Pago_total * factura.Promocion.Porcentaje)
	}

	c.JSON(http.StatusOK, factura.Pago_total)
}
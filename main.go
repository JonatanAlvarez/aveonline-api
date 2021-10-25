package main

import (
	_ "test/aveonline-api/database"
	"test/aveonline-api/handlers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	v1 := router.Group("/v1")

	v1.GET("/medicamento", handlers.Medicamentos)
	v1.POST("/medicamento", handlers.CreateMedicamento)

	v1.GET("/promocion", handlers.Promociones)
	v1.POST("/promocion", handlers.CreatePromocion)

	v1.GET("/factura", handlers.Facturas)
	v1.POST("/factura", handlers.CreateFactura)
	v1.POST("/factura/simular", handlers.SimularFactura)

	router.Run("localhost:9092")
}
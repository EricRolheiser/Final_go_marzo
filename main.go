package main

import (
	"examen/handler"
	"examen/middlewares"
	"examen/service"

	"github.com/gin-gonic/gin"
)

var (
	router          *gin.Engine
	analisisHandler *handler.AnalisisHandler
)

func main() {
	router = gin.Default()
	dependencies()
	mappingRoutes()
	router.Run(":8080")
}

func mappingRoutes() {
	router.Use(middlewares.ValidateAuthHeader())
	groupAnalisis := router.Group("/analisis")
	{
		groupAnalisis.POST("/promedio", analisisHandler.CalcularPromedio)
		groupAnalisis.POST("/volatilidad", analisisHandler.CalcularVolatilidad)
		groupAnalisis.GET("/proyeccion", analisisHandler.ProyeccionValor)
	}
}

func dependencies() {
	analisisService := service.NewAnalisisService()
	analisisHandler = handler.NewAnalisisHandler(analisisService)
}

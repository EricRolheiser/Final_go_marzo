package handler

import (
	"examen/model"
	"examen/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnalisisHandler struct {
	analisisService service.AnalisisInterface
}

func NewAnalisisHandler(analisisService service.AnalisisInterface) *AnalisisHandler {
	return &AnalisisHandler{analisisService: analisisService}
}

func (h *AnalisisHandler) CalcularPromedio(c *gin.Context) {
	var accion model.Accion
	err := c.ShouldBindJSON(&accion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, resultado := h.analisisService.CalcularPromedio(&accion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"promedioCierre": resultado})
}

func (h *AnalisisHandler) CalcularVolatilidad(c *gin.Context) {
	var accion model.Accion
	err := c.ShouldBindJSON(&accion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, resultado := h.analisisService.CalcularVolatilidad(&accion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"volatilidad": resultado})
}

func (h *AnalisisHandler) ProyeccionValor(c *gin.Context) {
	precioActualStr := c.Query("precioActual")
	tasaStr := c.Query("tasaCrecimientoAnual")
	aniosStr := c.Query("anios")
	if precioActualStr == "" || tasaStr == "" || aniosStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan parámetros en la consulta"})
		return
	}
	precioActual, err := strconv.ParseFloat(precioActualStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "precioActual inválido"})
		return
	}
	tasa, err := strconv.ParseFloat(tasaStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tasaCrecimientoAnual inválido"})
		return
	}
	anios, err := strconv.Atoi(aniosStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "anios inválido"})
		return
	}
	err, proyeccion := h.analisisService.ProyeccionValorFuturo(precioActual, tasa, anios)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"proyeccionValor": proyeccion})
}

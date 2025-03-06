package service

import (
	"errors"
	"examen/model"
	"math"
)

type AnalisisInterface interface {
	CalcularPromedio(accion *model.Accion) (error, float64)
	CalcularVolatilidad(accion *model.Accion) (error, float64)
	ProyeccionValorFuturo(precioActual, tasa float64, anios int) (error, []model.ProyeccionItem)
}

type AnalisisService struct{}

func NewAnalisisService() AnalisisInterface {
	return &AnalisisService{}
}

func (s *AnalisisService) CalcularPromedio(accion *model.Accion) (error, float64) {
	if accion == nil || len(accion.PreciosCierre) == 0 {
		return errors.New("los datos de precios de cierre son nulos"), 0
	}
	var suma float64
	for _, precio := range accion.PreciosCierre {
		if precio < 0 {
			return errors.New("los precios no pueden ser negativos"), 0
		}
		suma += precio
	}
	promedio := suma / float64(len(accion.PreciosCierre))
	promedio = math.Round(promedio*100) / 100
	return nil, promedio
}

func (s *AnalisisService) CalcularVolatilidad(accion *model.Accion) (error, float64) {
	if accion == nil || len(accion.PreciosCierre) == 0 {
		return errors.New("los datos de precios de cierre son nulos"), 0
	}
	var suma float64
	for _, precio := range accion.PreciosCierre {
		if precio < 0 {
			return errors.New("los precios no pueden ser negativos"), 0
		}
		suma += precio
	}
	promedio := suma / float64(len(accion.PreciosCierre))
	var sumaCuadrados float64
	for _, precio := range accion.PreciosCierre {
		diferencia := precio - promedio
		sumaCuadrados += diferencia * diferencia
	}
	volatilidad := math.Sqrt(sumaCuadrados / float64(len(accion.PreciosCierre)))
	volatilidad = math.Round(volatilidad*100) / 100
	return nil, volatilidad
}

func (s *AnalisisService) ProyeccionValorFuturo(precioActual, tasa float64, anios int) (error, []model.ProyeccionItem) {
	if precioActual < 0 {
		return errors.New("el precio actual no puede ser negativo"), nil
	}
	if tasa < 0 || tasa > 100 {
		return errors.New("la tasa de crecimiento debe estar entre 0 y 100"), nil
	}
	if anios <= 0 {
		return errors.New("la cantidad de años debe ser mayor a 0"), nil
	}
	var proyeccion []model.ProyeccionItem
	valor := precioActual
	for i := 1; i <= anios; i++ {
		if i == 1 {
			valor = valor*(1+tasa/100) - 1
		} else {
			valor = valor * (1 + tasa/100)
		}
		valorRedondeado := math.Round(valor*100) / 100
		proyeccion = append(proyeccion, model.ProyeccionItem{
			Anio:  i,
			Valor: valorRedondeado,
		})
	}
	return nil, proyeccion
}

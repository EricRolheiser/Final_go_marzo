package model

type Accion struct {
	PreciosCierre []float64 `json:"preciosCierre"`
}

type PromedioResponse struct {
	PromedioCierre float64 `json:"promedioCierre"`
}

type VolatilidadResponse struct {
	Volatilidad float64 `json:"volatilidad"`
}

type ProyeccionItem struct {
	Anio  int     `json:"anio"`
	Valor float64 `json:"valor"`
}

type ProyeccionResponse struct {
	ProyeccionValor []ProyeccionItem `json:"proyeccionValor"`
}

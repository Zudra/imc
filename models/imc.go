package models

type IMCEntrada struct {
	Peso   float64 `json:"peso"`
	Altura float64 `json:"altura"`
}

type IMCResultado struct {
	IMC           float64 `json:"imc"`
	Classificacao string  `json:"classificacao"`
}

package services

import (
	"errors"
	"imc/models"
)

func CalcularIMC(e models.IMCEntrada) (models.IMCResultado, error) {
	if e.Altura <= 0  {
		return models.IMCResultado{}, errors.New("altura inválida")
	}

	if e.Peso <= 0 {
		return models.IMCResultado{}, errors.New("peso inválido")
	}

	imc := e.Peso / (e.Altura * e.Altura)
	classificacao := classificarIMC(imc)

	return models.IMCResultado{
		IMC: round(imc),
		Classificacao: classificacao,
	}, nil
}

func classificarIMC(imc float64) string {
	switch {
	case imc < 17.0:
		return "Muito abaixo do peso"
	case imc >= 17 && imc <= 18.49:
		return "Abaixo do peso"
	case imc >= 18.50 && imc <= 24.99:
		return "Peso normal"
	case imc >= 25 && imc <= 29.99:
		return "Acima do peso"
	case imc >= 30 && imc <= 34.99:
		return "Obesidade I"
	case imc >= 35 && imc <= 39.99:
		return "Obesidade II (Severa)"
	default:
		return "Obesidade III (Mórbida)"
	}
}

func round(f float64) float64 {
	return float64(int(f*100)) / 100
}
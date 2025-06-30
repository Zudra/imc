package tests

import (
	"testing"

	"imc/models"
	"imc/services"
)

func TestCalcularIMC(t *testing.T) {
	casos := []struct {
		entrada       models.IMCEntrada
		imcEsperado   float64
		classEsperada string
		esperaErro    bool
	}{
		{models.IMCEntrada{Peso: 70, Altura: 1.75}, 22.85, "Peso normal", false},
		{models.IMCEntrada{Peso: 70, Altura: 1.80}, 21.60, "Peso normal", false},
		{models.IMCEntrada{Peso: 95, Altura: 1.75}, 31.02, "Obesidade I", false},
		{models.IMCEntrada{Peso: 45, Altura: 1.60}, 17.57, "Abaixo do peso", false},
		{models.IMCEntrada{Peso: 70, Altura: 0}, 0, "", true},
		{models.IMCEntrada{Peso: 0, Altura: 0}, 0, "", true},
	}

	for _, c := range casos {
		resultado, err := services.CalcularIMC(c.entrada)

		if c.esperaErro && err == nil {
			t.Errorf("esperava erro, mas não houve para entrada: %+v", c.entrada)
		}

		if !c.esperaErro && err != nil {
			t.Errorf("erro inesperado para entrada %+v: %v", c.entrada, err)
		}

		if !c.esperaErro {
			if resultado.IMC != c.imcEsperado {
				t.Errorf("imc esperado: %.2f, obtido: %.2f", c.imcEsperado, resultado.IMC)
			}

			if resultado.Classificacao != c.classEsperada {
				t.Errorf("Classificação esperada: %s, obtida: %s", c.classEsperada, resultado.Classificacao)
			}
		}
	}
}

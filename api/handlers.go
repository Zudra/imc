package api

import (
	"encoding/json"
	"net/http"

	"imc/models"
	"imc/services"
)

func CalcularIMCHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var entrada models.IMCEntrada
	err := json.NewDecoder(r.Body).Decode(&entrada)
	if err != nil {
		http.Error(w, "Erro ao ler entrada: "+err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := services.CalcularIMC(entrada)
	if err != nil {
		http.Error(w, "Erro no cálculo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}
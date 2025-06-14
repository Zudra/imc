package main

import (
	"log"
	"net/http"

	"imc/api"
)

func main() {
	http.HandleFunc("/imc", api.CalcularIMCHandler)

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
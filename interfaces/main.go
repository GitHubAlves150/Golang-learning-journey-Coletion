package main

import (
	"log"
	"net/http"
)

func main() {
	// Registrar as rotas
	http.HandleFunc("/rotas", getTasks)   // POST /rotas, GET /rotas
	http.HandleFunc("/rotas/", getTasksID) // GET /rotas/1, PUT /rotas/1

	log.Println("🚀 Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
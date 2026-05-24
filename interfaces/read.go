package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// getAllTasks lista todas as tarefas
func getAllTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(banco); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("📤 GET /rotas - %d tarefas", len(banco))
}

// getTaskByID busca uma tarefa pelo ID
func getTaskByID(w http.ResponseWriter, id int) {
	for _, task := range banco {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			log.Printf("📤 GET /rotas/%d - Encontrada: %s", id, task.Title)
			return
		}
	}
	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
	log.Printf("❌ GET /rotas/%d - Não encontrada", id)
}

// getTasks lida com as requisições em "/rotas" (sem ID)
func getTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		created(w, r)
	case http.MethodGet:
		if r.URL.Path == "/rotas" {
			getAllTasks(w)
		} else {
			http.Error(w, "Rota não encontrada", http.StatusNotFound)
		}
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// getTasksID lida com as requisições em "/rotas/" (com ID)
// Aqui é o "Router" que extrai o ID e chama as funções corretas (GET, PUT, DELETE)
func getTasksID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da URL
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) != 2 || parts[0] != "rotas" {
		http.Error(w, "URL inválida", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Roteia baseado no método HTTP
	switch r.Method {
	case http.MethodGet:
		getTaskByID(w, id)
	case http.MethodPut:
		Update(w, r, id) // Chama o Update passando o ID!
	case http.MethodDelete: // Placeholder para o futuro
		Delete(w, r, id)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
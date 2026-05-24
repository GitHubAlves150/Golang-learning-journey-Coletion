package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Update atualiza uma tarefa existente
// O ID já vem extraído do Router (getTasksID)
func Update(w http.ResponseWriter, r *http.Request, id int) {
	log.Printf("🔄 PUT /rotas/%d - Atualizando...", id)

	// 1. Procura a tarefa pelo ID
	index := -1
	for i, task := range banco {
		if task.ID == id {
			index = i
			break
		}
	}

	// 2. Se não encontrou, erro 404
	if index == -1 {
		http.Error(w, "ID não encontrado", http.StatusNotFound)
		return
	}

	// 3. Lê o JSON do corpo da requisição
	var taskAtualizada EstruturaAPI
	err := json.NewDecoder(r.Body).Decode(&taskAtualizada)
	if err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 4. Preserva o ID original
	taskAtualizada.ID = id

	// 5. Substitui a tarefa no slice
	banco[index] = taskAtualizada

	// 6. Retorna a tarefa atualizada (status 200 OK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banco[index])

	log.Printf("✅ PUT /rotas/%d - Atualizada: %s", id, taskAtualizada.Title)
}
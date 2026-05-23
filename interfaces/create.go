package main

import (
	"encoding/json"
	"net/http"
)




func created(w http.ResponseWriter, r *http.Request) {
	//1. Mecanismo de proteção que irá aceitar somente metodo POST
	if r.Method != http.MethodPost {
		http.Error(w, "Use apenas POST", http.StatusMethodNotAllowed)
	}

	//2. Criar uma task vazia
	var taskvazia EstruturaAPI

	//3. Ler o JSON que vem do corpo da requisição do clieete
	retorno := json.NewDecoder(r.Body).Decode(&taskvazia)
	if retorno != nil {
		http.Error(w, "Not permission"+retorno.Error(), http.StatusBadRequest)
		return
	}

	//4. Gerar o ID automatico de cada created
	taskvazia.ID = len(banco) + 1

	//5. Adiciona no banco a requisição
	banco = append(banco, taskvazia)

	//6. Responde para o cliente o status da requisição
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskvazia)

}

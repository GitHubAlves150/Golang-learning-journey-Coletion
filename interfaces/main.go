package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ==============================================================
// Servidor que ENVIA JSON(GET)
// ==============================================================

// Pessoa é a estrutura de dados que vamos enviar
type Summary struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var summaries []Summary

func main() {

	// crias as rotas - listagem de summary
	http.HandleFunc("/summaries", summariesHandler)
	http.HandleFunc("/summaries/", summaryHandler)

	log.Println("Server is Runiing")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func summariesHandler(w http.ResponseWriter, r *http.Request) {
	//ResponseWriter - serve para alimentar com dados o solicitante
	//Request - Qunado o solicitante envia algo do front
	//Como neste código não vamos pegar nada do nosso solicitante, vou deixar com um underline

	switch r.Method {
	case http.MethodGet:
		getSummaries(w)
		break;
	case http.MethodPost:
		creatSummary(w, r)
		break;
	}


}

func getSummaries(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(summaries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func creatSummary(w http.ResponseWriter, r *http.Request) {
	var summary Summary

	if err := json.NewDecoder(r.Body).Decode(&summary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	summary.ID = len(summaries) + 1
	summaries = append(summaries, summary)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(summary)

}

// =========================================================
func summaryHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) != 2 || parts[0] != "summaries" {
		http.Error(w, "invalid url", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "invalid Sumary ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getSumary(w, id)
		break;
	case http.MethodPut:
		updateSummary(w, r, id)
		break;
	case http.MethodDelete:
		deleteSumary(w, id)
		break;
	}

}

func getSumary(w http.ResponseWriter, id int) {
	for _, s := range summaries {
		if s.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
			return
		}

	}
	http.Error(w, "Summary no found", http.StatusNotFound)

}

// ============
func updateSummary(w http.ResponseWriter, r *http.Request, id int) {
	var summary Summary

	if err := json.NewDecoder(r.Body).Decode(&summary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	for i, s := range summaries {
		if s.ID == id {
			summaries[i].Description = summary.Description
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(summaries[i])
			return
		}
	}

	http.Error(w, "invalid Sumary ID", http.StatusBadRequest)

}

// ================================================================
func deleteSumary(w http.ResponseWriter, id int) {

	index := -1

	for i, s := range summaries {
		if s.ID == id {
			index = i
		}
	}
	if index == -1 {
		http.Error(w, "Sumary not found", http.StatusNotFound)
		
		return
	}

	summaries = append(summaries[:index], summaries[index+1:]...)
	w.WriteHeader(http.StatusNoContent)

}

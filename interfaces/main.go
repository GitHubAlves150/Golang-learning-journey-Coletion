package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Summary struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

var summaries []Summary

func summariesHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("list summary")
	//retornar um tipo de json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(summaries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	//Rotas
	//C.R.U.D
	http.HandleFunc("/summary", summariesHandler)
	log.Println("Server is Running")
	log.Fatal(http.ListenAndServe(":8080", nil)	)


	fmt.Println("...FIM...")
}

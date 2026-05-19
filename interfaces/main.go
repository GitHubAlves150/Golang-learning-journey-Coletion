package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Summary struct {
	ID          int    `json:"id"` //Quando converter para JSON, o campo ID vira "id"
	Description string `json:"description"`//Quando converter para JSON, o campo Description vira "description"
}

var summaries []Summary

func summariesHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("list summary")
	
	//Define o cabeçalho: o que estamos enviando é um Json
	w.Header().Set("Content-Type", "application/json")

	//Converte a slice summaries para JSON e envia
	if err := json.NewEncoder(w).Encode(summaries); err != nil {//json como resposta
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	//Rotas
	//C.R.U.D
	http.HandleFunc("/summary", summariesHandler)//cria uma rota
	log.Println("Server is Running")
	log.Fatal(http.ListenAndServe(":8080", nil)	)//cria um servidor local


	fmt.Println("...FIM...")
}

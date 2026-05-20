package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
//==============================================================
//Servidor que ENVIA JSON(GET)
//==============================================================

// Pessoa é a estrutura de dados que vamos enviar
type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}


func main() {

	//Cria uma estrutura do tipo Pessoa
	people := Pessoa{
		Nome:  "Lucas Lorenço Alves",
		Idade: 38,
	}

	//Cria um Handler(manipulador) que retorna essa estrutura pessoa em formato JSON
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		//Diz que a resposta será um JSON
		w.Header().Set("Contente-Type", "application/json")

		//Converte a estrutura para formato JSON e envia
		json.NewEncoder(w).Encode(people)

		fmt.Println("✅ Enviei a estrutura")
	})

	fmt.Println("Servidor rodando")
	fmt.Println("Acesse: http:/localhost:8080/people")

	//Abre servidor
	http.ListenAndServe(":8080", nil)


}

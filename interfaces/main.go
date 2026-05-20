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

	//Cria um Handler(manipulador) que retorna essa estrutura pessoa em formato JSON
	http.HandleFunc("/receber", func(w http.ResponseWriter, r *http.Request) {

		//Só Aceita methodo POST
		if r.Method != http.MethodPost {
			http.Error(w, "Use POST", http.StatusMethodNotAllowed)
			fmt.Println("❌ Método inválido:", r.Method)
			return
		}

		//Cria uma estrutura vazia do tipo pessoa
		var people Pessoa

		//Ler o JSON que veio do corpo da requisição
		err:= json.NewDecoder(r.Body).Decode(&people)
		if err != nil{
			http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
			return
		}

		//Mostra no console o que recebeu
		fmt.Printf("📥 Recebi: %s, %d anos\n", people.Nome, people.Idade)
		
		//Responde se deu certo
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Recebido com sucesso!"))

	})

	fmt.Println("Servidor rodando")
	fmt.Println("Acesse: http:/localhost:8080/people")

	//Abre servidor
	http.ListenAndServe(":8080", nil)

}

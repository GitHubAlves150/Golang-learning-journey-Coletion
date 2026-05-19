package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* 
	{
	  "userId": 1,
	  "id": 1,
	  "title": "delectus aut autem",
	  "completed": false
	}
*/
type Post struct {
	UserId   int    `json:"userId"` //As flags tem que serem iguais ao corpo do json da API externa
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"completed"`
}

func main() {

	fmt.Println("Fazedno requisição para a API json placeholder")
	resposta, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Erro ao fazer requisição: ", err)
		return
	}

	defer resposta.Body.Close() //No final do programa a Defer se encarrega de fechar todas as conexões da API
	//cria uma instancia de Post
	var post Post
	//verificação de erro
	err = json.NewDecoder(resposta.Body).Decode(&post)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta: ", err)
		return
	}

	fmt.Println(post)

	fmt.Println("...FIM...")
}

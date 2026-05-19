package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println("requisição foi recebida")
	var p Person

	//Corpo da requisição
	err :=json.NewDecoder(r.Body).Decode(&p)
    if err != nil{
		http.Error(w, "Erro ao processar Json", http.StatusBadRequest)
		return
	}
	
	fmt.Println(p)
	fmt.Fprintf(w, "Recebido: %+v\n", p)

}


func main() {
	fmt.Println("Escutando na porta ")
	http.HandleFunc("/submit", handler)
	
	http.ListenAndServe(":8080", nil)

	fmt.Println("...FIM...")
}



















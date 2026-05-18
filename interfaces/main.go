package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"nome"`
 	Email string `json:"email,omitempty"`
	Senha string `json:"senha"`
	Ativo bool `json:"ativo"`
	Idade int `json:"idade"`
}



func main() {

	u:= User{
		Name: "Lucas",
		Email: "",
		Senha: "12345",
		Ativo: true,
		Idade: 32,
	}

	jsonDate,_ := json.Marshal(u)


	fmt.Println(string(jsonDate))

	fmt.Println("...FIM...")
}

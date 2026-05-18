package main

import (
	"encoding/json"
	"fmt"
)


//======================================================
//Decodificação Json -> GO
//======================================================


type User struct {
	Name  string `json:"Name"`
 	Email string `json:"email,omitempty"`
	Senha string `json:"senha"`
	Ativo bool `json:"ativo"`
	Idade int `json:"idade"`
}



func main() {

	//Json recebido de alguma API, Arquivo, etc
	u:=`{
		"Name": "Lucas",
		"Email": "",
		"Senha": "12345",
		"Ativo": true,
		"Idade": 32
	}`

	
    //Decodificar
	var usuario User
	
	err := json.Unmarshal( []byte(u), &usuario    )
	
	if err!= nil {
		fmt.Println("Erro::::", err)
		return
	}

	fmt.Printf("%s %s\n %s\n %t\n %d\n",usuario.Name, usuario.Email, usuario.Senha, usuario.Ativo, usuario.Idade)
	fmt.Println("...FIM...")
}

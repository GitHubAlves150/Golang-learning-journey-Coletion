package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position string   `json:"position"`
	Skills   []string `json:"skils"`
	Adress   `json:"address"`
}

type Adress struct {
	City    string
	Country string
}

// validação
func validateEmployee(e Employee) error {
	if e.Name == "" {
		return fmt.Errorf("O nome do funcionário é obrigatório")

	}
	if e.Position == "" {
		return fmt.Errorf("falha")
	}
	return  nil
}

func main() {

    //1. Crio meu json
    jsonData := `{
                    "id": 1,
                    "name": "", 
                    "position": "engenheiro de Software",
                    "skills": ["GO", "Python"],
                    "address": {
                        "city": "Osasco",
                        "country": "Brasil"
                    }
                }`

    var employee Employee

    err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println("Falha ao decodificar o json", err)
		return
	}
	err=validateEmployee(employee)
	if err!=nil {
		fmt.Println("",err)
		return 
	}

	fmt.Println(employee)
	fmt.Println("...FIM...")
}

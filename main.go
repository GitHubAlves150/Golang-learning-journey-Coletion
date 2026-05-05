/*
*Sintaxe básica
 */

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

//Criar comportamento para structs
//Um método referenciando uma struct
func (p Person) Greet_() string{
	return  "hello, my name is " + p.Name;
}


func main() {

	p:= Person{
		Name: "Lucas",
		Age: 38,
		Email: "senoratec.io@gmail.com",
	}

	lucas_greet:= p.Greet_()

	fmt.Print("..", lucas_greet)


	fmt.Println("\nfim...")
}

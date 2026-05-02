/*
*Sintaxe básica
 */

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

func main() {

	//laço for simples
	for i := 0; i < 10; i++ {
		fmt.Println(">> ", i)
	}
	//range
	numbers := []int{10, 20, 33, 43}

	for indices, _ := range numbers {
		fmt.Println("indice: ", indices)
		//fmt.Println("valorres: ", valores)

	}

	capitais := map[string]string{
		"Br" : "Brasilia", //chave - valor
		"Fr" : "Paris",
		"Jp" : "Tóquio",
	}

	for pais, capital:= range capitais{
		fmt.Println("Pais: ", pais)
		fmt.Println("Capital :", capital)
	}



	fmt.Println("ola...")
}




/*
*Sintaxe básica
 */

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

func main() {

	//O que é slice?
	//arrays mais flexiveis

	var numbers = [6]int{2, 54, 66} //com tamanho é array e sem tamanho é slice
	var numeros = []int{12, 44, 65} //slice - simples e mágico (uuui)
	fmt.Println(numbers)
	fmt.Println(numeros)

	days := []string{"Segunda", "Terça", "Quarta"}
	fmt.Println(days)
	days = append(days, "Quinta")

	fmt.Println(days)
	numeros = append(numeros, 2)
	fmt.Println(numeros)

	vetorNumeros := []int{2, 3, 12, 21, 342, 234}

	fmt.Println(vetorNumeros[:2])
	vetorNumeros = append(vetorNumeros[:2], vetorNumeros[3:]...)
	fmt.Println(vetorNumeros)

	matriz := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{7, 8, 9},
	}

	fmt.Println(matriz)

	numbers2 := make([]int, 5) //cria um slice de 5 posições de forma dinâmica -comprimento  e capacidade. Ou seja make([]int, comprimento, capacidade)

	fmt.Println(cap(numbers2))
	fmt.Println(len(numbers2))
	fmt.Println("ola...")
}

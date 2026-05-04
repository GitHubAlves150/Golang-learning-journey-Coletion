/*
*Sintaxe básica
 */

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

func main() {

	var numbers [5]int

	fmt.Println(numbers)

	numbers[0] = 10
	fmt.Println(numbers)
	//inicialização durante a declaração
	day := [3]string{"\nsegunda\n", "terça\n", "quarta\n"}
	fmt.Println(day)

	//inferencia de valores
	notas := [...]int{33, 22, 43, 21, 434, 23213}
	fmt.Println(notas)

	fmt.Println("FIM...")
}

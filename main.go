package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

func oddGenerator() func() int {
	i := -1

	return func() int {
		i += 2
		return i
	}
}


//funcao soma recebe varios numeros inteiros
	func soma( numeros ... int)int{
		total := 0;//começa com zero

		for _, conteudo:=range numeros{
			total += conteudo;
		}
		return total;
	}


	
func main() {

	nextOddgenerator := oddGenerator()

	fmt.Println(nextOddgenerator())
	fmt.Println(nextOddgenerator())
	fmt.Println(nextOddgenerator())

	fmt.Println("Funções variadicas.........")

	fmt.Println(soma(1, 2) );   //3
	fmt.Println(soma(1, 2, 3)); //6
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7)) //28

	fmt.Println("....Fim....")
}

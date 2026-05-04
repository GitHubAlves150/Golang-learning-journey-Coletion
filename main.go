/*
*Sintaxe básica
 */

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/modulosGo/util"
	"github.com/Rhymond/go-money"
)

func main() {

	fmt.Println("", util.Soma(4, 10))

	pound := money.NewFromFloat(10.0, money.BRL);

fmt.Println( pound.Display());

	fmt.Println("FIM...")
}

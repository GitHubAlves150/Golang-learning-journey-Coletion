/*
*Você foi contratado para implementar um sistema bancário que
 gerencia contas correntes. O sistema deve usar métodos
 (com receivers por valor e por ponteiro) para manipular as contas.

 =======================================================================
 Objetivos do Exercício

    ✅ Praticar métodos com receiver por valor (func (c Conta))

    ✅ Praticar métodos com receiver por ponteiro (func (c *Conta))

    ✅ Entender quando usar valor vs ponteiro

    ✅ Implementar validações e tratamento de erros

    ✅ Trabalhar com múltiplos tipos e interações entre objetos
*/

package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/conta"
)

func main() {

	minha_conta := conta.NovaConta(2344, "Lucas")
	
	fmt.Println("\nfim...")
}

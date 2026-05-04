package main

import "fmt"

//Função principal que orquestra todas as análises e exibe um relatório formatado.
func analizarTexto(texto string) {
	fmt.Println(contarPalavras(texto) )
	fmt.Println(frequenciaPalavras(texto) )
	fmt.Println(palavrasMaisLongas(texto, 3) )
	fmt.Println(palavrasUnicas(texto) )
}
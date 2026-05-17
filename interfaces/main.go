package main

import (
	"fmt"
	"os"
)



func main() {

	//1. Lê o arquvio inteiro de uma vez
	dados, err := os.ReadFile("log.txt")

	if err != nil{
		fmt.Errorf("Erro:", "não foi possivel ler o arquivo")
		return
	}
	//2. Converte para string e exibe
	fmt.Println(string(dados))

	

	fmt.Println("...FIM...")
}

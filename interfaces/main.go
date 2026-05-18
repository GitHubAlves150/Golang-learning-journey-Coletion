package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {

	//1. Abrir o arquivo
	file, err := os.Open("log.txt")

	if err !=nil{
		fmt.Errorf("Erro: ", "Erro ao abrir o arquivo")
		return
	}

	defer file.Close() // Essencial!! Fecha o arquivo automaticamente no final

	//2. Criar scanner para ler linha por linha do arquivo
	scanner := bufio.NewScanner(file)

	//3. ler cada linha
	for scanner.Scan(){
		linha := scanner.Text()
		fmt.Println(linha)		
	}

	//4. Verifica se houve erro na leitura
	if err := scanner.Err(); err !=nil{
		fmt.Println("Erro ao ler arquivo", err)
	}

	fmt.Println("...FIM...")
}

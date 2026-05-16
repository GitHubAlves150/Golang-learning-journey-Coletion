package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Erro!Não é possível dividri por zero")
	}
	return a / b, nil
}

func openFile(name string) error {
	if name == "" {
		return fmt.Errorf("Erro ao abrir o arquivo %w", errors.New("O nome do arquivos está vazio"))
	}
	return nil

}

func main() {

	result, erro := divide(10, 10)
	if erro != nil {
		fmt.Println("Erro! ", erro)
		return
	}
	fmt.Println("Resultado", result)
	fmt.Println("Erro", erro)

	erro = openFile("")
	if erro!=nil{
		fmt.Println("não foi possível abrir o arquivo")
		return
	}

	fmt.Println("Arquivo abrto com sucesso")
	fmt.Println("....FIM....")
}

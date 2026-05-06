package main

import (
	"bancario/internal/banco"

	"fmt"
)

func main() {

	//minhaConta := conta.Conta{}

	//===========================================
	//criar um novo banco
	//===========================================
	meuBanco := banco.NovoBanco("Banco do Brasil")
	fmt.Println("Bnaco: ", meuBanco.GetNome())

	//===========================================
	//Abrir uma conta (método do banco)
	//===========================================
	conta1 := meuBanco.AbrirConta("Lucas Alves")
	fmt.Println("conta1", conta1)

	conta2 := meuBanco.AbrirConta("Maria Azevedo")
	fmt.Println("conta2", conta2)

	conta3 := meuBanco.AbrirConta("Constâncio Branito")
	fmt.Println("conta3", conta3)

	//Lucas Abre uma segunda conta
	conta4 := meuBanco.AbrirConta("Lucas Alves")
	fmt.Println("conta4", conta4)

	// ========================================
	// 3. REALIZAR OPERAÇÕES
	// ========================================
	fmt.Println("Deposito.....................")
	err :=conta1.Deposito(12.90)
	if err == nil {
		fmt.Println("Depositado......")
	}else{
		fmt.Println("Não foi possivel depositar")
	}

	fmt.Println("saldo: ",conta1)



	//===========================================
	//Exibir contas
	//===========================================

}

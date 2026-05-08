//=================================================================================================
//Com apenas interface
/*
|
|
|
|
|
*/
package main

import (
	"fmt"
)

// =================================================================================================
// 1. A interface (O contrato)
// =================================================================================================
type Pagamento interface {
	Pagar(valor float64) string
	GetNome() string
}

//=================================================================================================
//2. Cartão (implementa o pagamento)
//=================================================================================================

type Cartao struct {
	NomeTitular string
	Limite      float64
}

func (c Cartao) Pagar(valor float64) string {
	if valor > c.Limite {
		return "Cartão recusado: limite insuficiente"
	}
	return fmt.Sprintf("Pagamento de R$ %.2f aprovado no cartão de %s", valor, c.NomeTitular)
}

func (c Cartao) GetNome() string {
	return "Cartão"
}

//=================================================================================================
//3. Pix (implementa o pagamento)
//=================================================================================================

type Pix struct {
	Chave string
	Saldo float64
}

func (p Pix) Pagar(valor float64) string {
	if valor > p.Saldo {
		return "Saldo insuficiente:"
	}
	return fmt.Sprintf("Pagamento de R$ %.2f aprovado na chave de %s", valor, p.Chave)
}

func (p Pix) GetNome() string {
	return "Pix"
}

//=================================================================================================
//4. Boleto (implementa o pagamento)
//=================================================================================================

type Boleto struct {
	codigo string
}

func (b Boleto) Pagar(valor float64) string {

	return fmt.Sprintln("Boleto gerado com sucesso. Numero:", b.codigo[:10])
}

func (p Boleto) GetNome() string {
	return "Boleto"
}

// =================================================================================================
// 5. Função que aceita todos os tipos de pagamento(Polimorfismo)
// =================================================================================================
func FinalizarPagamento(pagamento Pagamento, valor float64) {
	fmt.Println("Processando pagamento com ", pagamento.GetNome())
	resultado := pagamento.Pagar(valor)
	fmt.Println(resultado)
}

func main() {

	//criando diferentes formas de pagamento(cosntrutor)
	cartao := Cartao{NomeTitular: "Lucas", Limite: 10000.00}
	pix := Pix{Chave: "048988450724", Saldo: 30000.0}
	boleto := Boleto{codigo: "AHSTFR654389FDSJFSDFKDMM998880-9"}

	//Polimorfismo em ação
	//A mesma função FinalizarPagamento funciona para TODOS!

	fmt.Println("----------------\n\n")
	FinalizarPagamento(cartao, 500.0)

	fmt.Println("----------------\n\n")
	FinalizarPagamento(pix, 500.0)
	
	fmt.Println("----------------\n\n")
    FinalizarPagamento(boleto, 200)

	fmt.Println("....FIM.....")
}

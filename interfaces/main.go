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
	if valor < c.Limite {
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
	if valor < p.Saldo {
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

func main() {

	cartao := Cartao{NomeTitular: "Lucas", Limite: 2000000.0}
	pix := Pix{Chave: "048988450724", Saldo: 3000.0}
	boleto := Boleto{codigo: "AHSTFR654389FDSJFSDFKDMM998880-9"}

	a := Pagamento.Pagar(cartao, 3000000)
	b := Pagamento.Pagar(pix, 40000)
	c := Pagamento.Pagar(boleto, 30)

	fmt.Println("..", a)
	fmt.Println("..", b)
	fmt.Println("..", c)

	fmt.Println("....FIM.....")
}

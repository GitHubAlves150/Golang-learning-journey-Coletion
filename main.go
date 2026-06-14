package main

import (
	"fmt"
	"strings"
)

// Interface para veiculo
type Veiculo interface {
	Mover() string
	GetNome() string
}

// implementação 1
type Carro struct {
	Modelo string
    Marca string
}

func (c Carro) Mover() string {
	return "Carro acelerando"
}

func (c Carro) GetNome() string {
	return c.Modelo
}

// Implementação 2
type Moto struct {
	Modelo string
}

func (m Moto) Mover() string {
	return "🏍️ Moto empinando"
}
func (m Moto) GetNome() string {
	return m.Modelo
}

//=============================================
//FUNCOES GENERICAS (Aceitam qualquer veículo)
//=============================================

// Função que recebe QUALQUER veículo
func IniciaCorrida(veiculos []Veiculo) {
	fmt.Println("Corrida Iniciada")
	for _, valores := range veiculos {
		fmt.Println("..", valores.GetNome(), valores.Mover())
	}

}

// Função que retorna QUALQUER veículo
func EscolherVeiculo(tipo string) Veiculo { //devolve uma inteface
	if tipo == "Carro" {
		return Carro{Modelo: "Wolksvager"}
	}
	return Moto{Modelo: "Biz"}
}

func main() {

	//Exemplo 1: Slice de veículo (Aceita carro e moto juntos)
	vehicle := []Veiculo{
		Carro{Modelo: "BYD", Marca: "Ainda nao sei"},
		Moto{Modelo: "Yamaha"},
		Carro{Modelo: "Fusca"},
	}
	
    //f := vehicle[0].(Carro)	
    //fmt.Println("Modelo", vehicle[0].Mover()) // Caso queira ver como se os membros de dados de cada indice

    fmt.Println(strings.Repeat("=", 10) )

    IniciaCorrida(vehicle)


    fmt.Println(strings.Repeat("=", 10) )

    v1:= EscolherVeiculo("Carro");
    v2:= EscolherVeiculo("Moto");

    fmt.Println("Escolhi: ", v1.GetNome(), v1.Mover());
    fmt.Println("Escolhi: ", v2.GetNome(), v2.Mover());


}

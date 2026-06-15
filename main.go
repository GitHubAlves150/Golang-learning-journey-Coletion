package main

import (
	"fmt"
	"strings"
)

//============================================
//INTERFACE INDIVIDUAIS
//============================================

type Localizavel interface {
	ObterCoordenadas() (lat, lon float64)
}

type Velocidade interface {
	ObterVelocidade() float64
}

type Direcao interface {
	ObterDirecao() float64
}

// ============================================
// COMPOSIÇÃO: Interface que COMBINA outras
// ============================================
// Rastreador completo tem TODOS os métodos das interfaces acima
type RastreadorCompleto interface {
	Localizavel //Tem ObterCoordenadas
	Velocidade  //Tem ObterVelocidade
	Direcao     //Tem ObterDireção

	//Dá para adicionar mpetodos próprios
	GetID() string
}

// ============================================
// IMPLEMENTAÇÃO
// ============================================
// --Tipo comcreto
type RastreadorReal struct {
	ID         string
	UltimaLat  float64
	UltimaLon  float64
	Velocidade float64
	Direcao    float64
}

// Implementação
func (r RastreadorReal) ObterCoordenadas() (float64, float64) {
	return r.UltimaLat, r.UltimaLon
}

func (v RastreadorReal) ObterVelocidade() float64 {
	return v.Velocidade
}

func (d RastreadorReal) ObterDirecao() float64 {
	return d.Direcao
}

func (r RastreadorReal) GetID() string {
	return r.ID
}

// ============================================
// Funções que usam as interfaces
// ============================================

// Função que só precisa de localização
func MostrarLocal(locallizavel Localizavel) {
	lat, lon := locallizavel.ObterCoordenadas()
	fmt.Printf("\nLocalização:%.4f, %.4f ", lat, lon)
}

// Função que precisa de tudo (rastreador completo)
func MostrarStatusCompleto(rastrear RastreadorCompleto) {
	fmt.Println("\n=====STATUS COMPLETO==========")
	fmt.Printf("\nID: %s", rastrear.GetID())

	lat, lon := rastrear.ObterCoordenadas()
	fmt.Printf("\nPosição [%.4f - %.4f]", lat, lon)

	fmt.Printf("\nvelocidade: %.1f", rastrear.ObterVelocidade())
	fmt.Printf("\nDireção: %.1f", rastrear.ObterDirecao())
	fmt.Printf(strings.Repeat("=", 50))
}

func main() {
	rastreador := RastreadorReal{
	ID:           "ES",
	UltimaLat:   -32.2222,
	UltimaLon:   -45.5555,
	Velocidade:  34.2,
	Direcao:     84.3,
	}

	//RastreadorReal implementa Localizavel (porque tem ObterCoordenadas)
	MostrarLocal(rastreador)

	//RastreadorReal implementa RastreadorCompleto (tem todos os métodos)
	MostrarStatusCompleto(rastreador)


	// ========================================
    // DEMONSTRAÇÃO DE FLEXIBILIDADE
    // ========================================

	fmt.Printf("\n--Flexibilidade: Funções aceitam diferentes interfaces --")

	//Função que aceita Localizavel (pode ser qualquer coisa com coordenadas)
	coordenadas := []Localizavel{
        rastreador,
         // struct anônima
    }

	for _, valores:= range coordenadas{
		//fmt.Printf("|", valores)
		MostrarLocal(valores)
	}

}

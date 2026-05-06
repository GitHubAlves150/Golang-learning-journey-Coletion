package main

import "fmt"

// 1. A interface (contrato)
type Campeao interface {
	Atacar() string
	Defender() string
	GetNome() string
}

// Guerreiro (Implementa campeao)
type Guerreiro struct {
	Nome string
}

func (g Guerreiro) Atacar() string {
	return g.Nome + " golpeia com a espada!\n"
}

func (g Guerreiro) Defender() string {
	return g.Nome + " ergue o escudo \n"

}

func (g Guerreiro) GetNome() string {
	return g.Nome
}

// Mago (Implementa campeao)
type Mago struct {
	Nome string
}

func (m Mago) Atacar() string {
	return m.Nome + " tira a varinha!\n"
}

func (m Mago) Defender() string {
	return m.Nome + " lança feitiço\n"

}

func (m Mago) GetNome() string {
	return m.Nome
}

// Arqueiro (Implementa campeao)
type Arqueiro struct {
	Nome string
}

func (a Arqueiro) Atacar() string {
	return a.Nome + " lança uma flecha!\n"
}

func (a Arqueiro) Defender() string {
	return a.Nome + " desvia do golpe \n"

}

func (a Arqueiro) GetNome() string {
	return a.Nome
}

// Funçao que aceita qualquer campeao
func Batalha(campeao Campeao) {
	fmt.Println(campeao.GetNome() + "Entrou na batalha")
	fmt.Println(campeao.Atacar() )
	fmt.Println(campeao.Defender() )

}

func main() {

	ranger := Guerreiro{Nome: "Lucas"}
    Batalha(ranger)
}

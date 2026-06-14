package main    

import(
    "fmt"
)

//Mesma interface de branch Topics/interface_basica
type Animal interface{
    EmitirSom() string
}

//Implementação 1
type Cachorro struct{
    Nome string
}

func (c Cachorro) EmitirSom() string{
    return "Au-Au" + c.Nome
}

//Implementação 2
type Gato struct{
    Nome string
}

func (g Gato)EmitirSom()string{
    return "Miau" + g.Nome
}

//Implementação 3
type Vaca struct{
    Nome string
}

func (v Vaca)EmitirSom()string{
    return "Muuuu" + v.Nome
}


//Função que aceita qualquer animal
func FazerBarulho(animal Animal){
    fmt.Println(animal.EmitirSom())
}



func main(){

    cachorro := Cachorro{Nome:"Dog 1"}
    gato := Gato{Nome: "Cat 1"}
    vaca:= Vaca{Nome: "Cow 1"}

    //Todas funcionam por que implementam Animal
    FazerBarulho(cachorro);
    FazerBarulho(gato);
    FazerBarulho(vaca);

}
package main    

import(
    "fmt"
)



//Passo 1: Interface básica - Defiinir a interface (o contrato)
type Animal interface{
    EmitirSom() string
}

//Passo 2: Criar a struct (tipo concreto)
type Cachorro struct{
    Nome string
} 

//passo 3: Implementar o método da interface
func (c Cachorro) EmitirSom() string{
    return "Au-Au"
}

func main(){

    var a Animal  //Variavél do tipo interface
    a = Cachorro{Nome: "Guapeca"}
    fmt.Println("..", a.EmitirSom())
}
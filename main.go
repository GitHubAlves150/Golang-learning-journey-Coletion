package main

import (
	"context"
	"fmt"
	"time"
)



func ExemploBasico(){
 ctx,cancel := context.WithCancel( context.Background())

 //Goroutine que escuta o contexto
 go func ()  {
	select{
	case <-ctx.Done():
		fmt.Println("Goroutine recebeu sinal de cancelamento")

	}
 }()
 time.Sleep( 1*time.Second)
 cancel()//Envia sinal de cancelamento
 time.Sleep( 100 *time.Millisecond)
}

func main() {

	ExemploBasico()
	
}
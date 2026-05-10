package main

import (
	"fmt"
	"time"
)

// ============================================
// channels
// ============================================

func main() {

	ch := make(chan string) //declaração do channel

	go func() {
		fmt.Println("Escrevendo no channel....")
		ch <- "Menssagem 1"  //para escrever no channel usa a seta para a esquerda <-
		fmt.Println("Menssagem enviada")
	}() //abro e fecho parenteses para chamar a função

	time.Sleep(2*time.Second)

	go func ()  {
		msg:= <-ch
		fmt.Println("Recebido: ", msg)
	}()

	time.Sleep(2*time.Second)


	fmt.Println("....FIM.....")
}

package main

import (
	"fmt"
	"time"
)

// ============================================
// channels
// ============================================

func main() {

	ch := make(chan string, 3) //declaração do channel

	go func()  {
		fmt.Println("escrevendo no channel com buffer")
		ch <- "menssagem 1"
		fmt.Println("Menssagem 1 enviada")
		ch <- "menssagem 2"
		fmt.Println("Menssagem 2 enviada")
		ch <- "menssagem 3"
		fmt.Println("Menssagem 3 enviada")
		ch <- "menssagem 4"
		fmt.Println("Menssagem 4 enviada")
	}()

	time.Sleep(2 *time.Second)

	go func ()  {
		for i:=0; i<4;i++{
			msg:= <-ch
			fmt.Println("Recebido: ", msg)
			time.Sleep(1*time.Second)
		}

	}()


	time.Sleep(10*time.Second)


	fmt.Println("....FIM.....")
}

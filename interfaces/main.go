package main

import (
	"fmt"
)

// ============================================
// channels
// ============================================

func Sendmessage(ch chan<- string) {
	ch <- "Menssagem 1"
	ch <- "Menssagem 2"
	ch <- "Menssagem 3"

}

func main() {

	ch := make(chan string,2)

	go Sendmessage(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("....FIM.")
}

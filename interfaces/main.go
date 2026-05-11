package main

import (
	"fmt"
)

// ============================================
// channels
// ============================================

func Readmessage(ch <-chan string) {
	msg1:= <-ch
	msg2:= <-ch
	fmt.Println("menssagem recebida: \n", msg1," e ", msg2)
}

func main() {

	ch:=make(chan string, 2)


	ch <-"Menssagem 1"
	ch <-"Menssagem 2"

	Readmessage(ch)

	fmt.Println("....FIM.")
}

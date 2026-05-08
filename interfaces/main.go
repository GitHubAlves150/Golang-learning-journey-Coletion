package main

import (
	"fmt"
	"time"
)


func FazerCafe(){
	fmt.Println("Fazendo café")
	time.Sleep(3*time.Second)
	fmt.Println("Café pronto")
}

func FritarOvos(){
	fmt.Println("Fritando Ovos")
	time.Sleep(3*time.Second)
	fmt.Println("ovos prontos")
}

func FazerTorrada(){
	fmt.Println("FazerTorrada")
	time.Sleep(3*time.Second)
	fmt.Println("Torrada Pronta")
}


func main() {

	start := time.Now()
	go FazerCafe()
	go FritarOvos()
	go FazerTorrada()

	time.Sleep(3* time.Second)
fmt.Println( time.Since(start))
	fmt.Println("....FIM.....")
}

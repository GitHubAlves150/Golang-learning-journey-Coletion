package main

import (
	"fmt"
	"sync"
	"time"
)

//=============================================
//código base com problemas
//=============================================

var contador int = 0

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	//cada goroutine vai incrementar 1000 vezes
	for i := 0; i < 1000; i++ {
		contador++ //Problema aqui
	}
}

func main() {
	var wg sync.WaitGroup

	numGoroutine := 100

	inicio := time.Now()

	for i := 0; i < numGoroutine; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}

	wg.Wait()

	fmt.Println("Esperado ", 100*1000) //100.000
	fmt.Println("Resultado ", contador)
	fmt.Println("Diferença ", (100*1000)-contador)
	fmt.Println("Tempo ", time.Since(inicio))

	fmt.Println("....FIM.")
}

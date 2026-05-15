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
var mu sync.Mutex

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()
	//cada goroutine vai incrementar 1000 vezes
	for i := 0; i < 1000; i++ {
		mu.Lock()// Comente esta linha para ver o caso de Race Condition - Lock gerencia o acesso a variavel sendo um acesso de cada vez (pega a chave abre a porta e tranca)
		contador++ //Problema aqui
		mu.Unlock()// Comente esta linha para ver o caso de Race Condition - Aqui o Unlock desbloqueia o acesso (libera a porta e devolve a chave)
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

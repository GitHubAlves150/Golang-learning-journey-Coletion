package main

import (
	"fmt"
	"sync"
)

// ============================================
// Mutex
// ============================================

var (//Declarnado mais de uma variavel 
	T int=10000;
	counter int = 0 //recurso compartilhado
	mutex   sync.Mutex
)

func IncrementCount(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock() //Pega a tarefa e segura, é como na analogia do banheiro. Aqui ele está fazendo "Pegar a chave e trancar a porta" (trava)
	counter++
	mutex.Unlock() //desbloqueia e segua adiante - seguinda a analogia do banheiro "Devolve a chave (destrava)"
}

func main() {

	var wg sync.WaitGroup

	wg.Add(T)

	fmt.Println("......", counter)

	for i := 0; i < T; i++ {

		go IncrementCount(&wg)

		//fmt.Println("..", counter)

	}

	wg.Wait()
	fmt.Println(".>>", counter)

	fmt.Println("....FIM.")
}

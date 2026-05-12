package main

import (
	"fmt"
	"sync"
)

// ============================================
// Mutex
// ============================================

var T int = 100000
var counter int = 0 //recurso compartilhado

func IncrementCount(wg *sync.WaitGroup) {
	defer wg.Done()
	counter++
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

package main

import (
	"fmt"
	"time"
	"sync"
)

func ProcessarVeiculo(id int, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("\nProcessando veiculo: %d", id)
	time.Sleep(1*time.Second)//simula trabalho
	fmt.Printf("\nveiculo %d Processado\n", id)

}



func main() {

	fmt.Println("====PROCESSANDO VEICULO COM WAITGROUP======")

	var wg sync.WaitGroup

	//Processar 5 vecículo em paralelo
	for i:=1; i<5; i++{
		wg.Add(1) //DIz! mais uma gouroutine
		go ProcessarVeiculo(i, &wg) //Passa o ponteiro do wg
	}

	fmt.Println("\nAguardanddo todos os veículos serem processados\n")
	wg.Wait()// Bloqueia até todas gouroutines chamaram Done()

	fmt.Println("\nTodos os gouroutines foram processados\n")


}

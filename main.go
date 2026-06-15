package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simula um veiculo enviando dados de GPS
func EnviandoVeiculos(veiculoID int, wg *sync.WaitGroup) {
	defer wg.Done()

	//Simula 3 leituras de GPS
	for leitura := 0; leitura <= 3; leitura++ {
		lat := -23.5 + rand.Float64()
		lon := -45.3 + rand.Float64()
		velocidade := rand.Float64() * 100
		fmt.Printf("\nVelocidade do veiculo %d [Leitura %d]: lat: %.4f, lon: %.4f, velocidade: %.1f km/h", veiculoID, leitura, lat, lon, velocidade)
		time.Sleep(100 * time.Millisecond) //Simula intervalo entre as leituras
	}
	fmt.Printf("\nveiculo %d finalizou suas leituras.\n", veiculoID)

}

func main() {

	fmt.Println("=== RASTREADOR COM MÚLTIPLOS VEÍCULOS ===\n")
	var wg sync.WaitGroup

	numeroDeVeiculos := 3

	//Inicia Varios veiculos em paralelos
	for i := 0; i <= numeroDeVeiculos; i++ {
		wg.Add(1) //diz quantas goo deve ser chamadas. Como está num FOR, entçao serṕa chamada uma por vez
		go EnviandoVeiculos(i, &wg)
		//time.Sleep(1 * time.Second) //Simula intervalo entre as leituras

	}

	fmt.Printf("\nTodos os vecículos estçao enviando dados simultaneamente\n")
	wg.Wait()
	fmt.Println("\n🎉 Todos os veículos finalizaram o envio!\n")

}

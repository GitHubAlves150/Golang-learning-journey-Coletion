package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Cenário

Você tem 20 veículos que precisam ser processados. Cada veículo tem um ID (de 1 a 20) e um tempo de processamento variável (entre 100ms e 1000ms).
O que o sistema deve fazer

    Processar os veículos em PARALELO (usando goroutines)

    Manter um contador global de quantos veículos já foram processados (usando mutex)

    Exibir quando cada veículo começar e terminar

    Ao final, exibir o tempo total de processamento e o contador final
*/

type Processador struct {
	contador int        //total de veículos  procesados
	mu       sync.Mutex //Protege o contador
}

func (p *Processador) ProcessadorVeiculo(veiculoID, tempoMS int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("\nVeiculo %d iniciou | TempoMS %d\n", veiculoID, tempoMS)
	time.Sleep(time.Duration(tempoMS) * time.Millisecond)
	
	p.mu.Lock()
	p.contador += 1
    p.mu.Unlock()

	fmt.Printf("\nVeiculo %d TERMINOU\n", veiculoID)

}

func main() {

	inicio := time.Now()

	processador := &Processador{
		contador: 0,
	}

	var wg sync.WaitGroup

	// Criar 20 veículos com tempos aleatórios
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		tempoMs := rand.Intn(900) + 100 // 100 a 1000ms

		go processador.ProcessadorVeiculo(i, tempoMs, &wg)
	}

	wg.Wait()
	fmt.Println("Demorou- ", time.Since(inicio) )
	fmt.Printf("Total de veiculos processados %d\n", processador.contador)

}

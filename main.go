// Estudos de Goroutine
// Lucas Lorenço Alves
// Ano 09/2026

/*
	Recaptular os estudos de Goroutine.
	[*]Nível 1 : Gouroutine básica
	[*]Nível 2 : Waitgroup (esperar gorourine)
    [*]Nível 3 : Channel unbuffered
	[ ]Nível 4 : Channel buffered
	[ ]Nível 5 : Armadilha do loop
	[ ]Nível 6 : Mutex (proteger estado)
	[ ]Nível 7 : Workerpool
*/

// Criar uma função que calcula o consumo de combustível
// Criar uma função que calcula a velocidade do veiculo

// Motrar na tela o valor do combustível
// Mostrar o toal de veículo
// Mostrar o número de veículo que esta dentro do limite
// Mostrar o número de veículo que está acima do limite

package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID         string
	Velocidade float64
	KmRodados  float64
}

type Resultado struct {
	ID     string
	Status string
	Custo  float64
}

// Crinado a função worker
func worker(id int, jobs <-chan Job, resultado chan<- Resultado, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs { // lê até o channel terminar
		status := "OK"

		if j.Velocidade > 80 {
			status = "ACIMA"
		}
		custo := j.KmRodados * 0.10 * 6.50
		resultado <- Resultado{
			ID:     j.ID,
			Status: status,
			Custo:  custo,
		}
	}
}

func main() {

	veiculos := []Job{
		{"CAR-000", 14.90, 60.4},
		{"CAR-001", 44.90, 10.4},
		{"CAR-002", 24.90, 69.4},
		{"CAR-003", 84.90, 10.1},
		{"CAR-004", 64.90, 50.1},
	}

	jobs := make(chan Job, 10)
	results := make(chan Resultado, 10)
	var wg sync.WaitGroup

	// 1. lança os workers
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results, &wg)
	}

	// 2. envia os jobs numa goroutine separada
	go func() {
		for _, v := range veiculos {
			jobs <- Job{
				ID:         v.ID,
				Velocidade: v.Velocidade,
				KmRodados:  v.KmRodados,
			}
		}
		close(jobs) // avisa workers que acabou
	}()

	// 3. fecha results quando todos workers terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. coleta resultados
	for r := range results {
		fmt.Printf("%s | %s | Custo: R$ %.2f\n", r.ID, r.Status, r.Custo)
	}
}

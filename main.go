package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type DadosVeiculo struct {
	Lat       float64
	Lon       float64
	VeiculoID int
}

// O primeiro argumento AGORA é o Context!
func SimulaVeiculo(ctx context.Context, id int, ch chan<- DadosVeiculo, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// 1. SEU DESAFIO: Criar o select aqui dentro.
		// Caso 1: Se o ctx.Done() disparar, imprime que o veículo parou e dá um 'return' para encerrar.
		// Caso 2: Se o time.After(200 * time.Millisecond) disparar, gera os dados e envia para o canal 'ch'.

		/* ESCREVA O SELECT AQUI */
		select {
		case <-ctx.Done():
			fmt.Println("EROR: ", ctx.Err())
			return

		case <-time.After(8 * time.Second):
			fmt.Println("time de 200 milisegundos")
			date := DadosVeiculo{
				Lat: -45.44,
				Lon: -33.22,
				VeiculoID: id,
			}
			ch <-date
			fmt.Println("Dado enviado..", id)
		}
	}
}

func ProcessaLeitura(ch <-chan DadosVeiculo) {
	for leitura := range ch {
		fmt.Printf("🚗 [Processador] Veículo %d coletado -> Lat: %.4f | Lon: %.4f\n",
			leitura.VeiculoID, leitura.Lat, leitura.Lon)
	}
	fmt.Println("🏁 Processador de dados encerrado com segurança.")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan DadosVeiculo, 10)

	// 2. Criamos um contexto com TIMEOUT de 1 Segundo.
	// Isso significa que todo o sistema vai rodar por exatamente 1s e parar sozinho!
	ctxPai := context.Background()
	ctx, cancel := context.WithTimeout(ctxPai, 6*time.Second)
	defer cancel()

	fmt.Println("🚀 Inicializando a frota de veículos em background...")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// Passamos o contexto filho para cada veículo
		go SimulaVeiculo(ctx, i, ch, &wg)
	}

	// O processador fica ouvindo o canal
	go ProcessaLeitura(ch)

	// Aguarda todos os veículos pararem (quando o timeout de 1s acontecer)
	wg.Wait()

	// Fecha o canal para liberar o range do processador
	close(ch)

	// Tempo pequeno para o print do processador aparecer antes do main morrer
	time.Sleep(50 * time.Millisecond)
	fmt.Println("🎉 Sistema finalizado com sucesso!")
}

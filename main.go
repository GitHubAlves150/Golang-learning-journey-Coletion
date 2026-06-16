// exercicio1_frota.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DadosVeiculo struct {
	Lat        float64
	Lon        float64
	Velocidade float64
	VeiculoID  int
}










// Simula veículos e envia dados GPS
func SimulaVeiculo(veiculoID int, ch chan<- DadosVeiculo, wg *sync.WaitGroup) {
	defer wg.Done()

	Leitura := make([]DadosVeiculo, 10)

	for i := 0; i < 10; i++ {
		Leitura[i].Lat = -23.4 + rand.Float64()
		Leitura[i].Lon = -13.4 + rand.Float64()
		Leitura[i].Velocidade = 10 + rand.Float64()
		Leitura[i].VeiculoID = veiculoID

		ch <- Leitura[i] //Envia para o canal
		time.Sleep(800 * time.Millisecond)

	}

	fmt.Printf("\n✅Veiculo %d finalizou envio", veiculoID)
}











// Processa dados
func ProcessaLeitura(ch <-chan DadosVeiculo ,wg *sync.WaitGroup) {
	defer wg.Done()

	for leitura := range ch { // range recebe até o canl fechar
		fmt.Printf("\n📡Veiculo %d | lat: %.4f - lon: %.4f - velo: %.4f", leitura.VeiculoID, leitura.Lat, leitura.Lon, leitura.Velocidade)
	}

	fmt.Println("Processador finalizou")

}










func main() {

	var wg sync.WaitGroup
	ch := make(chan DadosVeiculo, 15) //cria um buffer de 10 canais

	//cria 10 veiculos e faz eles enviarem dados parelamente
	fmt.Println("Inicializa os veículos")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go SimulaVeiculo(i, ch, &wg)
	}

	//inicia o processador em background
	go ProcessaLeitura(ch, &wg)

	

	//time.Sleep(100 * time.Millisecond)
	//Aguarda todos os veículos terminarem
	wg.Wait()

	//fechar o canal
	close(ch)

	fmt.Println("Sisitema Finalizado")

}

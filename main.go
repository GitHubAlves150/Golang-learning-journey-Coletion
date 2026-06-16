// 1_select_basico.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Telemetria struct {
	VeiculoID int
	Lon       float64
	Lat       float64
}

func main() {

	//crio canal
	canal1 := make(chan Telemetria)

	//Simula veiculos enviando dados
	go func() {
		//laço for
		for i := 0; i < 10; i++ {
			leitura := Telemetria{
				VeiculoID: i,
				Lon:       -23.5 + rand.Float64(),
				Lat:       -43.55 + rand.Float64(),
			}
			canal1 <- leitura //envia uma estrutura por vez
			time.Sleep(500 * time.Millisecond)
		}
		close(canal1)
	}()

	//processa o timeout
	
	for{  //è igual ao while(1)
		select {

		case leitura, ok := <-canal1:
			if !ok {
				fmt.Println("\nCanal fechado")
				return
			}
			fmt.Printf("\nveiculo %d: Lat: %.4f - Lon: %.4f", leitura.VeiculoID, leitura.Lat, leitura.Lon)

		case <-time.After(2 * time.Second):
			fmt.Println("nenhum dado recebido po 2 segundos")
			return
		}
	}


}

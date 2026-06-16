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

type Clima struct {
	VentoNorte int
	VentoSul   int
}

func main() {
	canal1 := make(chan Telemetria)
	canal2 := make(chan Clima)

	// Produtor 1: Telemetria
	go func() {
		for i := 0; i < 5; i++ {
			leitura := Telemetria{
				VeiculoID: i,
				Lat:       -34.4 + rand.Float64(),
				Lon:       -33.2 + rand.Float64(),
			}
			canal1 <- leitura
			fmt.Printf("📡 Telemetria %d enviada\n", i)
			time.Sleep(3 * time.Second)
		}
		fmt.Println("✅ Produtor 1 finalizou")
		close(canal1)
	}()

	// Produtor 2: Clima
	go func() {
		for i := 0; i < 5; i++ {
			leituraClima := Clima{
				VentoNorte: 23 + i,
				VentoSul:   22 + i,
			}
			canal2 <- leituraClima
			fmt.Printf("🌤️  Clima %d enviado\n", i)
			time.Sleep(3 * time.Second)
		}
		fmt.Println("✅ Produtor 2 finalizou")
		close(canal2)
	}()

	// Processador central
	canal1Fechado := false
	canal2Fechado := false

	for {
		// Se ambos os canais estiverem fechados, sai do loop
		if canal1Fechado && canal2Fechado {
			fmt.Println("\n🎉 Todos os dados processados!")
			break
		}

		select {
		case msg, ok := <-canal1:
			if !ok {
				if !canal1Fechado {
					fmt.Println("📭 Canal 1 fechado")
					canal1Fechado = true
				}
				continue
			}
			fmt.Printf("📍 [Canal 1] Veículo %d: Lat=%.4f, Lon=%.4f\n",
				msg.VeiculoID, msg.Lat, msg.Lon)

		case msg, ok := <-canal2:
			if !ok {
				if !canal2Fechado {
					fmt.Println("📭 Canal 2 fechado")
					canal2Fechado = true
				}
				continue
			}
			fmt.Printf("🌤️  [Canal 2] Vento Sul=%d, Vento Norte=%d\n",
				msg.VentoSul, msg.VentoNorte)

		case <-time.After(4 * time.Second): // ← 4 segundos (maior que 3)
			fmt.Println("⏰ Timeout: nenhum dado recebido em 4 segundos")
			return
		}
	}
}
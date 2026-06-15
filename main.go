package main

import (
	"fmt"
	"strings"
)

//==============================================
//INTERFACE DO GPS
//==============================================

// Qualquer fonte de GPS deve ter este método
type FonteGPS interface {
	ObterPosicao() (lat, long float64)
}

//==============================================
//IMPLEMENTAÇÕES DA INTERFACES
//==============================================

// 1. GPS real (hardware L76K)
type GPSL76K struct {
	PortaString string
}

func (g GPSL76K) ObterPosicao() (float64, float64) {
	fmt.Println("Lendo GPS real na porta: ", g.PortaString)
	// simula leitura do hardware
	return -23.5543, -65.3433
}

// 2. GPS Simulado (para testes)
type GPSMock struct {
	Lat, Long float64
}

func (g GPSMock) ObterPosicao() (float64, float64) {
	fmt.Println("Usando GPS simulado para testes")
	return g.Lat, g.Long
}

//==============================================
//Funçoes que usam a interface
//==============================================

// Esta função funciona com QUALQUER FonteGPS
func ColetarLocalizacao(gps FonteGPS, veiculoID string) {
	fmt.Println("Coletando posicao do veiculo: ", veiculoID)

	lat, lon := gps.ObterPosicao()

	fmt.Println("Posicao: ", lat, " / ", lon)
	fmt.Println("Locailização coletada com sucesso")

}

func main() {

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("RASTREADOR COM INTERFACE")

	// Cenário 1: Hardware real
	gpsReal := GPSL76K{PortaString: "/dev/ttyACM0"}
	ColetarLocalizacao(gpsReal, "ESP32-001")

	fmt.Println(strings.Repeat("=", 50))

	gpsMockup := GPSMock{Lat: -55.6544, Long: -43.2233}
	ColetarLocalizacao(gpsMockup, "ESP32-0002")

    fmt.Println("\nMesma função comportamento doferente\n")

    ColetarLocalizacao(gpsReal, "Carro-ABC")
    ColetarLocalizacao(gpsMockup, "Carro-xyz")


}

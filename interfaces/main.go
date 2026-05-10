package main

import (
	"fmt"
	"time"
)

// ============================================
// TRÊS MÉTODOS (cada um demora 1 segundo)
// ============================================

func ValidarPedido(pedidoID int) string {
	time.Sleep(1 * time.Second) //simula trabalho
	return fmt.Sprintln("Pedido :",pedidoID ,"Validação concluída", pedidoID)
}

func ProcessarPedido(pedidoID int) string {
	time.Sleep(1 * time.Second) //simula trabalho
	return fmt.Sprintln("Pedido :",pedidoID, "Processamento concluído", pedidoID)
}

func NotificarCliente(pedidoID int) string {
	time.Sleep(1 * time.Second) //simula trabalho
	return fmt.Sprintln("Pedido ",pedidoID, ": Notificacão enviada")
}

// ============================================
// EXECUÇÃO SEQUENCIAL (sem goroutines)
// ============================================

func main() {

	inicio := time.Now()
	pedidoID := 123
	fmt.Println("Iniciando processamento sequencial....\n")

	//Executa um depois do outro
	go ValidarPedido(pedidoID)

	go ProcessarPedido(pedidoID)
	
	go NotificarCliente(pedidoID)

	fmt.Println("Tempo total: ", time.Since(inicio))
	fmt.Println("....FIM.....")
}

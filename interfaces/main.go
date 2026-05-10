package main

import (
	"fmt"
	"sync" //adicionado para uso do waitgroup
	"time"
)

// ============================================
// TRÊS MÉTODOS (adaptados para usar WaitGroup)
// ============================================

func ValidarPedido(pedidoID int, wg *sync.WaitGroup) {
	defer wg.Done()             //Avisa que terminou
	time.Sleep(1 * time.Second) //simula trabalho
    fmt.Println("Pedido :", pedidoID, "Validação concluída", pedidoID)
}

func ProcessarPedido(pedidoID int, wg *sync.WaitGroup){
	defer wg.Done()             //Avisa que terminou
	time.Sleep(1 * time.Second) //simula trabalho
	fmt.Println("Pedido :", pedidoID, "Processamento concluído", pedidoID)
}

func NotificarCliente(pedidoID int, wg *sync.WaitGroup) {
	defer wg.Done() //Avisa que terminou
	time.Sleep(1 * time.Second) //simula trabalho
	fmt.Println("Pedido ", pedidoID, ": Notificacão enviada")
}

// ============================================
// EXECUÇÃO COM GOROUTINES + WAITGROUP
// ============================================


func main() {

	inicio := time.Now()
	pedidoID := 123
	fmt.Println("Iniciando processamento COM goroutine+waitgroup....\n")

	var wg sync.WaitGroup

	//Adicioan 3 tarefas ao contador
	wg.Add(3)


	//Executa um depois do outro
	go ValidarPedido(pedidoID, &wg)
	

	go ProcessarPedido(pedidoID, &wg)

	go NotificarCliente(pedidoID, &wg)


	//Espera todos terminarem
	wg.Wait()

	fmt.Println("Tempo total: ", time.Since(inicio))
	fmt.Println("....FIM.....")
}

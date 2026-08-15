package main

import (
	"context"
	"fmt"
	"time"
)

// O SOLDADO (nossa goroutine)
// Repare que o Context é sempre o primeiro parâmetro da função
func missãoDoSoldado(ctx context.Context) {
	fmt.Println("Soldado: Entrei na floresta e comecei a procurar por suprimentos")

	// Críamos um canal falso para simular o tempo que leva para achar os suprimentos
	suprimentosAchados := time.After(3 * time.Second)

	// O soldado fica ouvindo o rádio através do SELECT, enquanto o trabalho.
	select {
	case <-suprimentosAchados:
		// Se os 3 segundos passarem antes do rádio apitar, a missão foi um sucesso
		fmt.Printf("\nMissão com sucesso\n")
	case <-ctx.Done():
		// O rádio apitou ! o canal do contexto fechou porque o tempo acabou lá na base
		fmt.Println("Soldado! Recebi ordens pelo rádio, para abortar a missão e voltar pra base")
		fmt.Printf("Soldado (Motivo gravado no relatório): %v\n", ctx.Err())
	}
}

// O COMANDANTE (A função Main)
func main() {

	// O comandante pega um contexto base "vazio"
	contextBase := context.Background()

	// O Comandante define o limite : "Só temos 2 segundos"
	// Ele ganha o "contextoComComPrazo" (o rádio) e a função abortaMissão (o botão de emergência)
	contextoComPrazo, abortaMissao:=context.WithTimeout(contextBase, 5 * time.Second)

	// O defer garante que o botão de cancelar seja limpo da memória assim que a main acabar
	defer abortaMissao()

	// O Comandante envia o Soldado para a missão e entrega o rádio para ele
	go missãoDoSoldado(contextoComPrazo)


	// A Main (comandante precisa esperar um pouco na base para ver o soldado responder)
	time.Sleep(4 * time.Second)
	fmt.Println("Comandante: Operação encerrada")
}

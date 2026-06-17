package main

import (
	"context"
	"fmt"
	"time"
)

func InspecionaOperacao(ctxFinal context.Context) {

	//1. Método: Value(key) -> Recupera dados do "crachá" do contexto
	if operador, ok := ctxFinal.Value("operador_id").(string); ok {
		fmt.Printf("\n👤 Operador responsável: %s\n", operador)
	} else {
		fmt.Printf("\n👤 Nenhum operador identificado no conexto.")
	}

	//2. Método: Deadline()-> Devolve QUANDO o contexto vai expirar (no formato time.Time)
	//O 'ok' devolve true se o contexto tiver um tempo limite definido, ou false se for eterno.
	if horarioLimite, ok := ctxFinal.Deadline(); ok {
		tempoRestante := time.Until(horarioLimite)
		fmt.Printf("\n⏱️ Horário limite: %s (Resta exatamente: %v)\n", horarioLimite.Format("15:04:10"), tempoRestante)

	} else {
		fmt.Printf("\n⏱️ Este horário é externo, não tem horário limite\n")
	}

	fmt.Printf("\n⏳ Iniciando processamento pesado....\n")

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("\n ✅ Processamento concluído com sucesso\n")
	case <-ctxFinal.Done():
		//3. Método: Err() -> Só devolve algo  DEPOIS que o Done() fecha.
		//Ele diz a razão do cancelamento: ou "context deadline exceeded" (timeout)
		//ou "context canceled" (cancelamento manual)
		fmt.Printf("\n❌ERROR: \n", ctxFinal.Err())
	}

}

func main() {

	ctx := context.Background()
	operador := "operador_id"
	driver := "Lucas_Dev_2026"

	//injetamos um dado usando value
	ctxComvalor := context.WithValue(ctx, operador, driver)

	//Criamos um timeout curto de 200ms apartir do contexto que já tinha o valor
	ctxFinal, cancel := context.WithTimeout(ctxComvalor, 200*time.Millisecond)
	defer cancel()

	//Executa a inspeção
	InspecionaOperacao(ctxFinal)

}

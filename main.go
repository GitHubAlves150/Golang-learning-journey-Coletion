package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//w!j9Q%yqdfZL:n?  Stefanni Melo

// Funcao que processa o pagamento no gateway (Pargar.me)
func ProcessarPagamento(ctx context.Context, valorCompra float64) {
	fmt.Printf("\n💳 [GATEWAY] Iniciando validação do cartão..\n")

	//1. Método Value() - Recupera o ID do cliente que veio no "crachá" do contexto
	if clienteID, ok := ctx.Value("usuario_id").(int); ok {
		fmt.Printf("\n👤 [GATEWAY] Identificando CPF/Dados do cliente ID: %d..\n", clienteID)
	}
	//2. Método Deadline() -  verificamos quanto tempo temos antes da nossa API estourar
	if limite, ok := ctx.Deadline(); ok {
		tempoRestante := time.Until(limite)
		fmt.Printf("\n⏱️[GATEWAY] Janela de tempo segura; a requisição expira em %v\n", tempoRestante)

		//Decisão inteligente baseada no deadline
		if tempoRestante < 50*time.Millisecond {
			fmt.Println("⚠️ [Gateway] Tempo restante insuficiente para transação bancária segura. Abortando antes de cobrar!")
			return
		}
	}

	//Simula a chamada de rede para o banco (demora 300ms)
	chBanco := make(chan string, 1)
	go func() {
		time.Sleep(300 * time.Millisecond)
		chBanco <- "PAGAMENTO_APROVADO_TOKEN_9988"
	}()

	//O select monitora o watchdog
	select {
	case resposta := <- chBanco:
		fmt.Printf("✅ [Gateway] Sucesso! Cartão cobrado no valor de R$ %.2f. Código: %s\n", valorCompra, resposta)
	case <-ctx.Done():
		//3. Método Err() - O Done() disparou. Vamos dar o veredito do motivo usando Err()
	motivo:=ctx.Err()
		fmt.Printf("\n🚨 [Gateway] OPERAÇÃO ABORTADA PELO SISTEMA!\n")

		if errors.Is(motivo, context.DeadlineExceeded) {
			fmt.Println("❌ Motivo: O servidor do banco demorou mais que o limite permitido(Timeout)")
		} else if errors.Is(motivo, context.Canceled) {
			fmt.Println("❌ Motivo: O cliente cancelou a operação ou fechou a aba do navegador.")
		}

		fmt.Println("🛡️ [Gateway] Estorno de segurança garantido. Nenhuma cobrança foi feita.")

	}

}
func main() {

	//Criamos o contexto base da requisição HTTP
	ctxPAI := context.Background()

	//2. injetamos o ID do cliente logado na sessão (withValue)
	ctxUsuario := context.WithValue(ctxPAI, "usuario_id", 4042)

	//CASO 1: o gateway responde a tempo (Timeout de 500ms, banco demora 300ms) obs! Sempre liberar os recursos
	fmt.Println("---SIMULAÇÃO 1: FLUXO PERFEITO---")
	ctxSucesso, cancel1 := context.WithTimeout(ctxUsuario, 500*time.Millisecond)
	ProcessarPagamento(ctxSucesso, 150.90)
	cancel1() //Sempre liberar recursos

	//CASO 2: O banco fica lente e estoura o limite (Timeout de 100ms, banco demora 300ms)
	fmt.Println("---SIMULAÇÃO 2: BANCO LENTO----")
	ctxTimeout, cancel2 := context.WithTimeout(ctxUsuario, 100*time.Millisecond)
	ProcessarPagamento(ctxTimeout, 89.90)
	cancel2()

	//CASO 3: O cliente desiste  e clica em "Cancelar" manualmente
	fmt.Println("--SIMULAÇÃO 3: CLIENTE DESISITE (CANCELAMENTO MANUAL) ---")
	ctxCancelamento, cancelManual := context.WithCancel(ctxUsuario)

	//Simula o cliente clicando em fechar a página após 50 milisegundos
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("💻 [Navegador] Usuário fechou a aba do e-comerce!")
		cancelManual()
	}()

	ProcessarPagamento(ctxCancelamento, 450.00)

}

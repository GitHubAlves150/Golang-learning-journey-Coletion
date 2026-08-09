package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"calculadora-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Conecta no servidor que está rodando na porta 50052
	conexao, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Erro ao se conectar ao servidor: %v", err)
	}
	defer conexao.Close()

	// 2. Cria o cliente gRPC baseado no nosso contrato
	cliente := pb.NewCalculadoraServiceClient(conexao)

	// 3. Criamos o contexto com timeout de segurança
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Nossos números de teste
	numero1 := int32(105)
	numero2 := int32(300)

	fmt.Printf("📞 Enviando requisição para somar %d + %d via gRPC...\n", numero1, numero2)

	// 4. Executamos a chamada como se fosse uma função comum do Go!
	resposta, err := cliente.Somar(ctx, &pb.SomaRequest{
		Num1: numero1,
		Num2: numero2,
	})
	if err != nil {
		log.Fatalf("Erro ao chamar a função Somar: %v", err)
	}

	// 5. Exibe o resultado
	fmt.Printf("🎉 O servidor respondeu! O resultado é: %d\n", resposta.Resultado)
}

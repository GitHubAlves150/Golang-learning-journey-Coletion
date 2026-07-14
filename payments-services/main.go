package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"payments-services/pb" // Importa a pasta pb que o protoc gerou

	"google.golang.org/grpc"
)

// 1. Criamos a struct que representar o servidor.
// Ela PRECISA herdar a estrutura padrão gerada pelo gRPC
type Server struct {
	pb.UnimplementedPaymentServiceServer
}

// 2. Implementamos a função que descrevemos lá no arquivo .proto
// Perceba como o gRPC nos obriga a usar o 'context.Context' nativamente
func (s *Server) ProcessarPagamento(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("💳 Recebida tentativa de pagamento para o Pedido ID: %s no valor de: R$ %.2f", req.PedidoId, req.Valor)

	// Aqui entraria a integração com a API do Stripe, Asaas, Pagar.me, etc.
	// Para o nosso lab, vamos apenas simular que se o valor for maior que 1000, o pagamento é recusado
	status := "Aprovado"
	if req.Valor > 1000 {
		status = "Recusado (Sem limite)"
	}

	//Retornamos a struct exata que o .proto espera
	return &pb.PaymentResponse{
		Status:      status,
		TransacaoId: "tx_6672859652", // ID fictício da transação
	}, nil

}

func main() {
	port := ":50051"

	// Abrimos uma porta TCP para escutar as conexões de rede
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Falha ao abrir a porta %s: %v", port, err)
	}

	// 4. Criamos uma instância do servidor gRPC puro
	grpcServer := grpc.NewServer()

	// 5. Registramos o nosso Server dentro do ecossistema gRPC
	pb.RegisterPaymentServiceServer(grpcServer, &Server{})

	fmt.Printf("🚀 Microserviço de Pagamentos gRPC rodando na porta %s...\n", port)

	// 6. Ligamos o servidor para aceitar requisições externas
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Falha ao rodar o servidor gRPC: %v", err)
	}

}

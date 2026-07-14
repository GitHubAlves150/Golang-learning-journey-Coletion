package main

import (
	"calculadora-grpc/pb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Criamos nossa struct do servidor
type ServidorCalculadora struct {
	pb.UnimplementedCalculadoraServiceServer
}

// Implementamos a função de soma definida no .proto
func (s *ServidorCalculadora) Somar(ctx context.Context, req *pb.SomaRequest) (*pb.SomaResponse, error) {
	log.Printf("📥 Servidor recebeu os números: %d e %d:", req.Num1, req.Num2)

	soma := req.Num1 + req.Num2

	// Retornamos a resposta esperada pelo gRPC
	return &pb.SomaResponse{
		Resultado: soma,
	}, nil
}

func main() {
	// 1. Abrimos a porta de rede 50052 para executar conexões TCP
	lister, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Erro ao abrir a porta : %v", err)
	}

	// Criamos o servidor gRPC
	grpcServer := grpc.NewServer()

	// 3. registramos nossa lógica de calculadora dentro do servidor gRPC
	pb.RegisterCalculadoraServiceServer(grpcServer, &ServidorCalculadora{})

	fmt.Println("🚀 Servidor rodando na porta :50052....")

	// 4. Começamos a escutar
	if err := grpcServer.Serve(lister); err != nil {
		log.Fatalln("erro ao rodar o gRPC: %v", err)
	}

}

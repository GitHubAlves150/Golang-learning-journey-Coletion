package main
import (
	"context"
	"fmt"
	"log"
	"time"

	"orders-services/pb" // Importa o pb local de pedidos

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



func main(){
	// 1. Conectar no microserviço de Pagamentos (Porta 50051)
	// Usamos 'Withtransportcredencial(insecure...)' por que estamos em ambiente local sem SSL/TLS
	conn , err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Não foi possível conectar ao serviço de pagamentos: %v", err)
	}

	defer conn.Close()

	// 2. Cria o cliente grpc baseado no contrato gerado
	client := pb.NewPaymentServiceClient(conn)

	// 3. Criamos um Context com Timeout de 3 segundos (Resiliencia de microserviços!)
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second) 
	defer cancel()

	// ------Sumulação de disparo de pedido-------------
	pedidoID := "req_998877"
	valorPedido := 250.00 // Teste para mudar para 1500 depois ver recusar

    fmt.Printf("🛒 Criando pedido %s no valor de R$ %.2f...\n", pedidoID, valorPedido)
	fmt.Println("📞 Chamando microserviço de pagamentos via gRPC...")

	// 4. Faz a chamada remota (RPC) como se fosse uma função local do seu código
	resposta, err := client.ProcessarPagamento(ctx, &pb.PaymentRequest{
		PedidoId: pedidoID,
		Valor: valorPedido,
	})
	if err != nil {
		log.Fatalf("Erro ao processar pagamento: %v", err)
	}


	// 5. Exibe o resultado retornado pelo outro microserviço
	fmt.Println("\n--- 🧾 RETORNO DO MICROSERVIÇO DE PAGAMENTOS ---")
	fmt.Printf("Status do Pagamento: %s\n", resposta.Status)
	fmt.Printf("ID da Transação:     %s\n", resposta.TransacaoId)


}

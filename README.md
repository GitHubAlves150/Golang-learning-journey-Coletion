# Exemplo de POST em Golang com Gin

Este material mostra a evolução de uma API em Go usando o framework Gin, em três etapas:

1. Um endpoint simples de teste.
2. Um POST síncrono recebendo JSON.
3. Um POST assíncrono usando goroutine e canal para processamento em background.

O objetivo é entender como o fluxo da requisição pode sair de algo básico até uma arquitetura mais parecida com um caso real de fila de processamento.

---

## 1) Etapa inicial: endpoint de teste

```go
package main

import (
	"github.com/gin-gonic/gin"
)

// DadosVeiculo representa o que vai chegar no corpo da requisição POST.
type DadosVeiculo struct {
	VeiculoId  string  `json:"veiculo_id" binding:"required"`
	Lat        float64 `json:"lat" binding:"required"`
	Lon        float64 `json:"lon" binding:"required"`
	Velocidade float64 `json:"velocidade"`
}

func main() {
	// Cria o roteador padrão do Gin.
	// Ele já vem com logs e tratamento básico de erros.
	router := gin.Default()

	// Rota GET simples para testar se o servidor está funcionando.
	router.GET("/ping", func(c *gin.Context) {
		// O c aqui é o contexto do Gin, usado para responder a requisição.
		c.JSON(200, gin.H{"mensagem": "Pong"})
	})

	// Sobe o servidor na porta 8080.
	router.Run(":8080")
}
```

### O que essa etapa ensina

Essa versão serve para validar que o servidor está no ar e que o Gin está configurado corretamente.  
O `router.GET("/ping")` responde com JSON e ajuda a testar rapidamente a API.  
Aqui ainda não existe POST, nem leitura de corpo da requisição, nem validação de campos.

---

## 2) Etapa POST: recebendo JSON e validando dados

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DadosVeiculo representa o que vai chegar no corpo da requisição POST.
type DadosVeiculo struct {
	VeiculoId  string  `json:"veiculo_id" binding:"required"`
	Lat        float64 `json:"lat" binding:"required"`
	Lon        float64 `json:"lon" binding:"required"`
	Velocidade float64 `json:"velocidade"`
}

func main() {
	router := gin.Default()

	// Rota POST para receber a telemetria do veículo.
	router.POST("/telemetria", func(c *gin.Context) {
		var novoDado DadosVeiculo

		// Lê o JSON recebido e preenche a struct.
		if err := c.ShouldBindJSON(&novoDado); err != nil {
			// Se o JSON estiver inválido ou faltar campo obrigatório, retorna 400.
			c.JSON(http.StatusBadRequest, gin.H{
				"erro":    "JSON inválido ou campos obrigatórios ausentes",
				"detalhe": err.Error(),
			})
			return
		}

		// Aqui o JSON já foi validado e convertido para struct.
		// Neste ponto você poderia salvar no banco ou enviar para outro serviço.
		println("📡 Posição recebida do veículo:", novoDado.VeiculoId)
		println("🌍 Coordenadas:", novoDado.Lat, novoDado.Lon)

		// Resposta de sucesso.
		c.JSON(http.StatusCreated, gin.H{
			"status":  "Telemetria recebida com sucesso",
			"veiculo": novoDado.VeiculoId,
		})
	})

	router.Run(":8080")
}
```

### O que essa etapa ensina

Agora a API recebe dados reais via `POST`.  
O método `ShouldBindJSON` faz o parsing do corpo JSON diretamente para a struct `DadosVeiculo`.  
Os campos com `binding:"required"` são obrigatórios, então a requisição falha se `veiculo_id`, `lat` ou `lon` não vierem preenchidos [web:49][web:51].

### Exemplo de JSON de envio

```json
{
  "veiculo_id": "CAR-123",
  "lat": -27.5954,
  "lon": -48.5480,
  "velocidade": 82.5
}
```

### Resposta esperada

```json
{
  "status": "Telemetria recebida com sucesso",
  "veiculo": "CAR-123"
}
```

---

## 3) Etapa assíncrona: POST com goroutine e canal

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DadosVeiculo representa o que vai chegar no corpo da requisição POST.
type DadosVeiculo struct {
	VeiculoId  string  `json:"veiculo_id" binding:"required"`
	Lat        float64 `json:"lat" binding:"required"`
	Lon        float64 `json:"lon" binding:"required"`
	Velocidade float64 `json:"velocidade"`
}

// processadorDeFila consome dados do canal e processa em background.
func processadorDeFila(canalTelemetria <-chan DadosVeiculo) {
	fmt.Println("👷 Worker: pronto e aguardando dados no canal...")

	// O range fica lendo mensagens do canal até ele ser fechado.
	for telemetria := range canalTelemetria {
		fmt.Printf("⚙️ [PROCESSADOR BACKGROUND] Tratando telemetria do veículo %s\n", telemetria.VeiculoId)
		fmt.Printf("Coordenadas: Lat: %.4f - Lon: %.4f | Velocidade: %.2f\n",
			telemetria.Lat, telemetria.Lon, telemetria.Velocidade)

		// Aqui poderia entrar gravação em banco, fila, log ou integração externa.
	}
}

func main() {
	// Opcional: reduz logs de debug.
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Canal que intermedeia a API e o processador.
	canalTelemetria := make(chan DadosVeiculo, 100)

	// Inicia workers em background.
	for i := 0; i < 3; i++ {
		go processadorDeFila(canalTelemetria)
	}

	// Rota POST que recebe os dados e envia para o canal.
	router.POST("/telemetria-post", func(c *gin.Context) {
		var novoDado DadosVeiculo

		if err := c.ShouldBindJSON(&novoDado); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
			return
		}

		// Em vez de processar aqui, enviamos para a fila interna.
		canalTelemetria <- novoDado

		// Responde imediatamente ao cliente.
		c.JSON(http.StatusCreated, gin.H{
			"status": "telemetria aceita e enviada para a fila de processamento",
		})
	})

	fmt.Println("🚀 Servidor HTTP rodando na porta :8080")
	router.Run(":8080")
}
```

### O que essa etapa ensina

Aqui a API responde rápido, mas o processamento pesado acontece fora da requisição HTTP.  
A rota só valida o JSON e manda o dado para o canal.  
As goroutines `processadorDeFila` ficam ouvindo esse canal e processam as mensagens em paralelo [web:50][web:54][web:58].

### Por que isso é útil

Esse padrão é muito comum quando você quer:

- responder rápido ao cliente;
- evitar deixar a requisição presa em tarefa demorada;
- distribuir trabalho entre vários workers;
- separar a camada HTTP da camada de processamento.

## Diferença entre as 3 etapas

| Etapa | O que faz | Vantagem | Limitação |
|---|---|---|---|
| GET `/ping` | Testa o servidor | Simples e útil para debug | Não recebe dados |
| POST síncrono | Recebe JSON e processa na hora | Fácil de entender | Pode deixar a resposta lenta |
| POST assíncrono | Recebe JSON e joga para canal | Resposta rápida e escalável | Exige cuidado com concorrência e fechamento do canal |

---

## Pontos importantes do código

### `binding:"required"`
Esse tag informa ao Gin que o campo não pode ficar vazio na validação.  
Se o JSON vier incompleto, `ShouldBindJSON` retorna erro [web:51][web:49].

### `ShouldBindJSON`
É a forma recomendada quando você já sabe que o corpo vem em JSON.  
Ele tenta converter o corpo para a struct e validar os campos ao mesmo tempo [web:51].

### Goroutine
A goroutine permite que o processamento aconteça em paralelo ao fluxo principal da API.  
Isso evita travar a resposta HTTP enquanto o trabalho acontece em segundo plano [web:50][web:54].

### Canal
O canal é a ponte entre a rota HTTP e o worker.  
Ele transporta `DadosVeiculo` com segurança entre goroutines [web:50].

---

## Exemplo de teste com curl

```bash
curl -X POST http://localhost:8080/telemetria-post \
  -H "Content-Type: application/json" \
  -d '{
    "veiculo_id": "CAR-123",
    "lat": -27.5954,
    "lon": -48.5480,
    "velocidade": 82.5
  }'
```

---

## Conclusão

Essas três etapas mostram a evolução natural de uma API em Go com Gin: primeiro um teste simples, depois um POST validado e, por fim, uma versão mais próxima de produção com processamento assíncrono.  
A combinação de `POST + goroutine + channel` ajuda a construir sistemas mais rápidos e organizados, principalmente quando existe trabalho pesado após o recebimento da requisição [web:50][web:54][web:56].
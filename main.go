package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DadosVeiculo representa o que vai chegar no corpa da requisição POST
type DadosVeiculo struct {
	VeiculoId  string  `json:"veiculo_id" binding:"required"` //biding:required torna o campo obrigatório
	Lat        float64 `json:"lat" binding:"required"`
	Lon        float64 `json:"lon" binding:"required"`
	Velocidade float64 `json:"velocidade"`
}

func processadorDeFila(canalTelemetria <-chan DadosVeiculo) {
	fmt.Println("👷 Worker: Pronto e aguardando dados no canal...")

	//O loop 'range' fica ouvindo o canal perpetuamente
	for Telemetria := range canalTelemetria {
		fmt.Printf("⚙️[PROCESSADOR BACKGROUD] TRATANDO TELEMETRIA DO %s", Telemetria.VeiculoId)
		fmt.Printf("Coordenadas: Lat: %.4f - Lon: %.4f Velecidade: %.4f", Telemetria.Lat, Telemetria.Lon, Telemetria.Velocidade)
		/*--código aqui para gravação no banco de dados--*/
	}

}

func main() {

	//Desliga os logs de debug para limpar o terminal (opcional
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//2. Criamos o canal que vai intermediar a API e o processador
	canalTeletria := make(chan DadosVeiculo, 100) //buffer de 100 para arguentar picos de entrada

	//3. Inicializamos o processador em background (Goroutine)
	for i := 0; i < 3; i++ {
		go processadorDeFila(canalTeletria)

	}

	//Rota POST
	router.POST("/telemetriaPost", func(ctx *gin.Context) {
		//crio uma variavel do tipo struct do JSON
		var novoDado DadosVeiculo

		if err := ctx.ShouldBindJSON(&novoDado); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"erro": err.Error(),
			})
			return
		}

		//4. EM VEZ PROCESSAR AQUI, JOGAMOS NO CANAL
		//operação ultra rápida: joga no cano e libera o http
		canalTeletria <- novoDado

		//Responde imediatamente para o veiculo
		ctx.JSON(http.StatusCreated, gin.H{
			"Status": "telemetria aceita e evia a fila de processamento",
		})

	})

	close(canalTeletria)
	fmt.Println("🚀 Servidor HTTP rodando na porta :8080")

	router.Run(":8080")
}

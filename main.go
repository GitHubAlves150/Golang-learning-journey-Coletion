package main

import(
	"github.com/gin-gonic/gin"
)

func main(){

	//1. Cria o roteador padrão do gin (Já vem com logs e tratamento de crach embutidos)
	router := gin.Default()

	//2. Define a nossa primeira rota HTTP do tipo GET
	router.GET("/ping", func (c* gin.Context)  {
		//O 'c' aqui dentro é o Context do Gin (Que encapsula o context.Context do GO que você aprendeu)
		c.JSON(200, gin.H{ "menssagem": "Pong"})		
	})

	//3. Roda o servidor na porta 8080 (http://localhost:8080)
	router.Run(":8080")
}
package main

import (
	"fmt"      // Formatação de strings (ex: fmt.Sprintf) fmt → Para formatar a string do endereço (":8080")
	"log"      // Logs no console (ex: log.Println, log.Fatal) log → Para mostrar mensagens no terminal
	"net/http" // Constantes HTTP (StatusBadRequest, StatusOK, etc) net/http → Para usar códigos HTTP como 200, 400, 404, etc

	"github.com/labstack/echo/v4" // Framework web Echo echo/v4 → Para criar o servidor web
	"github.com/spf13/viper"      // Gerenciador de configurações viper → Para ler configurações (porta do servidor)
)

// ============================================================
// 🏗️ PARTE 2: MODEL (Estrutura de Dados)
// ============================================================

type Product struct { //type Product struct → Cria um novo tipo chamado Product
	ID          int    `json:"id"` //ID int → Campo para o número identificador `json:"id"` → Tag: quando converter para JSON, o campo ID vira "id"
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"img"`
}

// "Banco de dados" em memória
var products []Product
var nextID = 1 //var nextID = 1 → Próximo ID disponível (começa em 1)

// ============================================================
// ⚙️ PARTE 3: CONFIGURAÇÃO (Viper)
// ============================================================
//Define a estrutura das configurações:
//  type Config struct → Tipo para guardar configurações
// Server struct {...} → Configurações do servidor
// Port int → Porta onde o servidor vai rodar
// `mapstructure:"port"` → Tag para o Viper mapear o YAML
// `mapstructure:"server"` → Tag para mapear a seção server do YAML

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

// Declaração da função → Retorna ponteiro para Config e possivelmente um erro.
func loadConfig() (*Config, error) {

	//Cria uma nova instância do Viper → Cada chamada tem seu próprio Viper.
	v := viper.New()

	// 1. Onde procurar as configurações?
	v.SetConfigName("config")   // Procura por arquivo config.yaml ou config.yml
	v.SetConfigType("yaml")     // Define o formato como YAML
	v.AddConfigPath(".")        //Procura na pasta atual
	v.AddConfigPath("./config") //Procura na subpasta config/

	v.SetDefault("server.port", 8080) //Define valor padrão → Se não encontrar em nenhum lugar, usa porta 8080.

	//Pela ordem de precedência
	v.AutomaticEnv()                 //AutomaticEnv() → Ativa leitura automática de env vars
	v.BindEnv("server.port", "PORT") //BindEnv("server.port", "PORT") → A chave server.port pode vir da env PORT

	//Tenta ler o arquivo de configuração → Se não encontrar, só avisa e usa valores padrão.
	if err := v.ReadInConfig(); err != nil {
		log.Println("⚠️ Arquivo de config não encontrado, usando padrões")
	}

	//Ler e guardar na struct config
	var config Config

	//Unmarshal(&config) → Preenche a struct com os valores do Viper
	if err := v.Unmarshal(&config); err != nil {
		return nil, err

	}
	//Devolver as configurações
	return &config, nil
}

// ============================================================
// 🎮 PARTE 4: HANDLERS (Lógica das Rotas)
// ============================================================
// CreateProdutc - POST /products
// Declaração do handler → Recebe c echo.Context e retorna error.
func createProducts(c echo.Context) error {
	//1. Cria um produto vazio
	var newProduct Product

	//2. Extrair os dados do JSON que veio na requisição e coloca em newproduct
	//c.Bind(&newProduct) → Converte o JSON recebido para a struct
	if err := c.Bind(&newProduct); err != nil { //c.Bind tira o pedido vindo da requisição e coloca em newProduct
		//Devolve um status bad request 400 se caso for nil
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Dados inválidos: " + err.Error(),
		})
	}
	log.Println(">>>", newProduct.Title) //Isso prova que o c.Bind extraiu os dados da requisição
	//3.validar dados obrigatórios - um deles é o nome que não pode vir vazio
	if newProduct.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "O campo nome é obrigatório",
		})
	}

	//4.Gerar ID Automático
	newProduct.ID = nextID
	nextID++

	//5. Adicionar ao banco de dados (Memória local para este exemplo)
	products = append(products, newProduct)

	//6. Retornar o produto criado (status 201 Created)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "produto criado com sucesso",
		"product": newProduct,
	})

}

func listProducts(c echo.Context) error {
	//c.JSON() → Converte o map para JSON e envia ao cliente
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products":      products,
		"count":         len(products),
		"Qualquercoisa": "opaaa",
	})
}

// ============================================================
// 🔧 PARTE 5: SETUP DO SERVIDOR
// ============================================================

func setupServer() *echo.Echo {
	//Cria uma nova instância do servidor Echo.
	e := echo.New()

	//Middleware de log
	//Adiciona um middleware (função que roda antes de cada requisição):
	//e.Use() → Registra um middleware
	e.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc { //A função recebe next (próximo handler)
			return func(c echo.Context) error { //Retorna uma função que loga o método e a URL
				log.Println("...", c.Request().Method, c.Request().URL.Path)
				return next(c)
			}
		})

	//Rotas
	//e.POST("/products", createProducts) → Quando receber POST em /products, chama createProducts
	e.POST("/products", createProducts)
	//e.GET("/products", listProducts) → Quando receber GET em /products, chama listProducts
	e.GET("/products", listProducts)

	//Retorna o servidor configurado.
	return e

}

// ============================================================
// 🚀 PARTE 6: FUNÇÃO PRINCIPAL (main)
// ============================================================

func main() {

	//Carregar as configurações
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Erro ao carregar as configurações: ", err)

	}

	//Cria o servidor
	server := setupServer()

	//Iniciar
	addr := fmt.Sprintf(":%d", config.Server.Port)

	log.Printf("Servidor rodando em http://localhost:%s", addr)
	log.Printf("Endpoints:")
	log.Printf("  POST /products       -Criar produto")
	log.Printf("  GET  /products       -Listar produto")

	if err := server.Start(addr); err != nil {
		log.Fatal("Servidor falhou ", err)
	}

}

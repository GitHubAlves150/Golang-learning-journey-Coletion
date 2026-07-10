package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// usuario representa uma tabela no banco de dados
type User struct {
	ID       uint   `gorm:"primarykey; autoIncrement" json:"id"`
	Nome     string `gorm:"type:varchar" json:"nome"`
	Idade    uint   `gorm:"type:varchar(20)" json:"idade"`
	Telefone string `gorm:"type:varchar(30)" json:"telefone"`
}

// Global para simplificar neste orimeiro passao (vamos remover isso quando aplicarmos SOLID)
var db *gorm.DB

func main() {

	//1. Conexão com obanco de dados
	//Altere as credenciais abaixo de acordo com o seu ambiente local
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
		return
	}

	//Automigrate -  Cria a tabela automaticamente se ela não existir
	//db.AutoMigrate(&User{})

	//2. Configuração do roteador 'Chi'
	r := chi.NewRouter()
	r.Use(middleware.Logger) //middleware de log para ver as requisições no terminal
	r.Use(middleware.Recoverer)

	//3. definição das rotas - Neste caso GET
	r.Get("/usuary", GetusuarioHandler)
	r.Get("/usuary/{id}", GetUsuarioByIDHandler)

	//4. inicialização de Servidor HTTP
	log.Printf("Servidor rodando no porta :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao abrir o servidor: %v", err)
	}

}

// =======================================
// Preciso das funçoes Handler
// =======================================
// funcao handler para listar todos os usuarios (GET /usuarios)
func GetusuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuarios []User //crio um slice que cada indice é uma struct

	//busca todos os registros usando o GORM
	if result := db.Find(&usuarios); result.Error != nil {
		http.Error(w, "Erro ao selecionar os indices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(usuarios)
}

// funcação handler para buscar um usuario especifico por ID (GET /usuario/{id})
func GetUsuarioByIDHandler(w http.ResponseWriter, r *http.Request) {
	//Extrai o parametro id da URL
	id := chi.URLParam(r, "id")

	//Simula a tabela no banco
	var usuario []User

	//busca o primeiro registro que coincidir com o ID
	if result := db.First(&usuario, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Usuario nao encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

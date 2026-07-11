package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid" // Precisamos desse pacote para lidar com UUID no Go
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Usuario mapeia exatamente a sua tabela do banco de dados
type Usuario struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Nome     string    `gorm:"type:varchar(30)" json:"nome"`
	Email    string    `gorm:"type:varchar(30)" json:"email"`
	Sexo     string    `gorm:"type:varchar(6)" json:"sexo"`
	CriadoEm time.Time `gorm:"column:criado_em;default:CURRENT_TIMESTAMP" json:"criado_em"`
}

// iniciamos ao GORM o nome exato da sua tabela
func (Usuario) TableName() string {
	return "usuarios"
}

var db *gorm.DB

func main() {

	//1. Conexão com o banco de dados
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("erro ao conectar o banco: %s", err)
	}

	//2. Configuração do chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//Apenas Rota POST
	r.Post("/usuario", CriarUserHandler)

	//3. Iniciar o servidor
	log.Println("Servidor rodando na porta :8080 (Apenas POST)...")
	http.ListenAndServe(":8080", r)
}

// handler simples que recebe o JSON do cliente e salva no banco
func CriarUserHandler(w http.ResponseWriter, r *http.Request) {
	var NovoUsuario Usuario

	//Passo A: Pega o json enviado no corpo (Body) da requisição e joga para a struct espelho
	err := json.NewDecoder(r.Body).Decode(&NovoUsuario)
	if err != nil {
		log.Fatalln("erro ao pegar a requisição do corpo: %s", err)
		return
	}

	//Passo B: Salva no banco de dados usando GORM
	//O postgres var gerar um UUID e o CriadoEm automaticamente por causa dos DEAFAULTS do seu banco
	result := db.Create(&NovoUsuario)
	if result.Error != nil {
		http.Error(w, "Erro: ", http.StatusInternalServerError)
		return
	}

	//Passo C: Devolve o usuário criado de volta para quem chamou (já com o UUID e data gerados)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //status 201
}

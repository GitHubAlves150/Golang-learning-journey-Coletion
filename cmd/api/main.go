package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Precisamos desse pacote para lidar com UUID no Go

// Nossa struct idêntica ao banco de dados
type Usuario struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Nome     string    `gorm:"type:varchar(30)" json:"nome"`
	Email    string    `gorm:"type:varchar(30)" json:"email"`
	Sexo     string    `gorm:"type:varchar(6)" json:"sexo"`
	CriadoEm time.Time `gorm:"column:criado_em;default:CURRENT_TIMESTAMP" json:"criado_em"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

var db *gorm.DB

func main() {

	// 1. Conexão com o banco
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Put("/usuarios/{id}", AtualizarUsuarioHandler)

	log.Println("Servidor rodando na porta :8080 (Apenas PUT)...")
	http.ListenAndServe(":8080", r)
}

func AtualizarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	//Passo A: Pegar o UUID enviado na URL
	idParam := chi.URLParam(r, "id")

	//Validar se o ID enviado na URL é realmente um UUID válido
	userID, err := uuid.Parse(idParam)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	//Passo B: Verifica se o usuário existe no banco de dados antes de atualizar
	var usuarioExisitente Usuario
	if err := db.First(&usuarioExisitente, "id=?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Usuario nao encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	//Passo C: Decodificar os novos dados vindos do json do Body
	var dadosAtualizados Usuario
	if err := json.NewDecoder(r.Body).Decode(&dadosAtualizados); err != nil {
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	//Passo D: Atualizar os campos da struct que pegamos do banco
	usuarioExisitente.Nome = dadosAtualizados.Nome
	usuarioExisitente.Email = dadosAtualizados.Email
	usuarioExisitente.Sexo = dadosAtualizados.Sexo

	//Passo E: Salvar as alteração de volta no banco de dados usando o .Save() do GORM
	//O .Save() faz um UPDATE baseado na primaryKey (ID) da struct
	if err := db.Save(&usuarioExisitente).Error; err != nil {
		http.Error(w, "Erro ao atualizar no banco de dados", http.StatusInternalServerError)
		return
	}

	//Passo F: Retornar o usuário já atualizado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //status 200

}

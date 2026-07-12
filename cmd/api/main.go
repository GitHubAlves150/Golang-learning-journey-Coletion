package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	// 🔥 Nova rota DELETE cadastrada
	r.Delete("/usuarios/{id}", DeletarUsuarioHandler)

	log.Println("Servidor rodando na porta :8080 (Apenas PUT)...")
	http.ListenAndServe(":8080", r)
}

func DeletarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// Passo A: Extrai o ID da URL
	idParam := chi.URLParam(r, "id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "ID em formato inválido", http.StatusBadRequest)
		return
	}

	// Passo B: Verifica se o usuário existe no banco de dados
	var usuario Usuario
	if err := db.First(&usuario, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound) // 404
			return
		}
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	// Passo C: Deleta fisicamente do banco de dados (Hard Delete)
	// Passamos a struct populada com o ID para o GORM saber quem deletar
	if err := db.Delete(&usuario).Error; err != nil {
		http.Error(w, "Erro ao deletar do banco", http.StatusInternalServerError)
		return
	}

	// Passo D: Retorna Status 204 (No Content) - Padrão REST para deleção bem-sucedida que não retorna corpo
	w.WriteHeader(http.StatusNoContent)

}

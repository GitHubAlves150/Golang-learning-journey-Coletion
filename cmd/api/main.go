package main

import (
	"log"
	"net/http"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/handler"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// 1. Conexão com o banco
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	// Inicialização das camadas (SOLID - Injeção de Dependência)
	userRepo := repository.NewUserRepositoryDB(db)
	userhandler := handler.NewUserHandler(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Put("/usuarios/{id}", userhandler.AtualizaUsuario)

	log.Println("Servidor rodando com arquitetura limpa na porta :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

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

	// Montagem do grafo de dependências (SOLID)
	userRepo := repository.NewUserRepositoryDB(db)
	userHandler := handler.NewUserHandler(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Única rota registrada no sistema inteiro
	r.Delete("/usuarios/{id}", userHandler.DeletarUsuario)

	log.Println("Servidor rodando na porta :8080 (Apenas PUT)...")
	http.ListenAndServe(":8080", r)
}

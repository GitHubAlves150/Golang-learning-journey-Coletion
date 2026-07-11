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
	// 1. Banco de Dados
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar no banco: %v", err)
	}

	repository.RodaScriptSQL(db)
	//db.AutoMigrate(&entity.Usuario{})

	if err != nil {
		log.Fatalf("Erro ao executar os comandos do arquivo.sql: %v", err)
	}
	// 2. Injeção de Dependências (O coração do SOLID)
	userRepo := repository.NewUserRepositoryDB(db)
	userHandler := handler.NewUserHandler(userRepo)

	// 3. Roteador Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rotas apontando para os métodos do Handler
	r.Get("/usuarios", userHandler.GetUsuarios)
	r.Get("/usuarios/{id}", userHandler.GetUsuarioByID)

	log.Println("Servidor limpo e performático rodando na porta :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/handler"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware" // Precisamos desse pacote para lidar com UUID no Go
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	//1. Conexão com o banco de dados
	dsn := "host=localhost user=lucas_user password=123456 dbname=postgress_for_crud port=5432 sslmode=disable"

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("erro ao conectar o banco: %s", err)
	}

	//Injeção de dependência
	userRepo := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepo)

	//2. Configuração do chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//Apenas Rota POST
	r.Post("/usuario", userHandler.CriarUsuario)

	//3. Iniciar o servidor
	log.Println("Servidor rodando na porta :8080 (Apenas POST)...")
	http.ListenAndServe(":8080", r)
}

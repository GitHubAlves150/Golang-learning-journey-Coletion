package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"
)

type UserHandler struct{
	repo repository.UserRepository //Depende do contrato, não do banco direto
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo:repo}
}

func (h *UserHandler)CriarUsuario(w http.ResponseWriter, r *http.Request){
	var NovoUsuario entity.Usuario

	if err := json.NewDecoder(r.Body).Decode(&NovoUsuario); err != nil{
		http.Error(w, "json Inválido", http.StatusInternalServerError)
		return
	}

	if err:=h.repo.Create(&NovoUsuario); err!= nil{
		http.Error(w, "erro ao salvar no banco", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //status 201
	
}
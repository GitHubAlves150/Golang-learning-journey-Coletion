package handler

import (
	"encoding/json"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	repo repository.UserRepository // Depende da interface, não da implementação!
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarios)
}

func (h *UserHandler) GetUsuarioByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	usuario, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}

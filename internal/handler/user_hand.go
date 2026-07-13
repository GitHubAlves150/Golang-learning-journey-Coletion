package handler

import (
	"net/http"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	repo repository.UserRepository // Inversão de dependência
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// 1. Extrai o parâmetro id da URL do Chi
	idParam := chi.URLParam(r, "id")

	// 2. Valida se a string enviada é de fato um UUID válido
	userID, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "ID em formato inválido (deve ser um UUID válido)", http.StatusBadRequest)
		return
	}

	// 3. Aciona o repositório para verificar se o usuário existe no banco
	usuario, err := h.repo.FindByID(userID)
	if err != nil {
		// Se o GORM não achar, retornamos 404 Not Found imediatamente
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// 4. Manda o repositório executar a exclusão física
	if err := h.repo.Delete(usuario); err != nil {
		http.Error(w, "Erro ao deletar registro no banco", http.StatusInternalServerError)
		return
	}

	// 5. Devolve o Status 204 No Content (padrão REST para deleções de sucesso sem corpo)
	w.WriteHeader(http.StatusNoContent)
}
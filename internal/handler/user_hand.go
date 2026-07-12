package handler

import(
	"encoding/json"
	"net/http"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct{
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository)*UserHandler{
	return &UserHandler{repo: repo}
}


func (h *UserHandler)AtualizaUsuario(w http.ResponseWriter, r *http.Request){
	//1. Extrai e valida o UUID da URL
	idParam := chi.URLParam(r, "id")
	userId, err:=uuid.Parse(idParam)
	if err != nil{
		http.Error(w, "ID em formato inválido (deve ser em UUID)", http.StatusBadRequest)
		return
	}
	//2.Verifica se o usuário de formato existe no banco
	usuarioExistente, err := h.repo.FindByID(userId)  
	if err != nil{
		http.Error(w, "Usuário não encontrado", http.StatusBadRequest)
		return
	}

	//3. Decodifica os dados que o cliente que atualizar
	var dadosNovos entity.Usuario
	if err:= json.NewDecoder(r.Body).Decode(&dadosNovos); err!=nil{
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	//4. Executa avalidação que criamos lá na Entity
	if err:= dadosNovos.Validar(); err != nil{
		http.Error(w, "Dados não validados", http.StatusBadRequest)
		return
	}


	//5. Substitue dados antigospelos dados novos (preservando o ID do usuário e data-hora)
	usuarioExistente.Nome =dadosNovos.Nome
	usuarioExistente.Email=dadosNovos.Email
	usuarioExistente.Sexo=dadosNovos.Sexo

	//6. Envia para o repositório salvar no postgres
	if err:=h.repo.UpDate(usuarioExistente);err!=nil{
		http.Error(w, "Dados não validados", http.StatusBadRequest)
		return
	}

	//7. Retorna o usuário atualizado com Status 200 OK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
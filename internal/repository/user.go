package repository

import (
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository dita o contrato estrito para esta operação
type UserRepository interface {
	FindByID(id uuid.UUID) (*entity.Usuario, error)
	Delete(usuario *entity.Usuario) error
}

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(database *gorm.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: database}
}

func (r *UserRepositoryDB) FindByID(id uuid.UUID) (*entity.Usuario, error) {
	var usuario entity.Usuario
	// Busca o usuário pelo UUID enviado
	if err := r.db.First(&usuario, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UserRepositoryDB) Delete(usuario *entity.Usuario) error {
	// Remove o registro fisicamente do banco de dados (Hard Delete)
	return r.db.Delete(usuario).Error
}
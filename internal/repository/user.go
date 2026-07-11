package repository

import (
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"
	"gorm.io/gorm"
)

// UserRepository dita o contrato (Interface)
type UserRepository interface {
	Create(usuario *entity.Usuario)error
}

//UserRepositoryDB implementa a interface usando  GORM/POstgres (Trabalhador)
type UserRepositoryDB struct{
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepositoryDB{
	return &UserRepositoryDB{db:database}
}


func (r *UserRepositoryDB)Create(usuario *entity.Usuario) error{
	return r.db.Create(usuario).Error
}



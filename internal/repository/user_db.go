package repository

import (
	"log"
	"os"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"

	"gorm.io/gorm"
)

// UserRepository define o contrato (Inversão de Dependência)
type UserRepository interface {
	FindAll() ([]entity.Usuario, error)
	FindByID(id string) (*entity.Usuario, error)
}

// UserRepositoryDB implementa a interface usando GORM
type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(database *gorm.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: database}
}

func (r *UserRepositoryDB) FindAll() ([]entity.Usuario, error) {
	var usuarios []entity.Usuario
	if err := r.db.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *UserRepositoryDB) FindByID(id string) (*entity.Usuario, error) {
	var usuario entity.Usuario
	if err := r.db.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func RodaScriptSQL(db *gorm.DB) {
	//conteudo, err := os.ReadFile("github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/Created/000001_aluno.sql")
	conteudo, err := os.ReadFile("internal/CreateDB/000001_aluno.sql")

	if err != nil {
		log.Fatal("Erro ao briar o arquivo sql")
	}

	err = db.Exec(string(conteudo)).Error
	if err != nil {
		log.Fatal("não foi possivel  rodar db.Exec()")
	}

	log.Println("banco de dados rodando com sucesso...")

}

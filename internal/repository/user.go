/*
A camada do banco de dados não faz validações de negócio, ela apenas executa o que mandamos. Ela precisa 
de dois métodos: buscar o registro atual (FindByID) e salvar as alterações (Update).
*/
package repository


import(
	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uuid.UUID)(*entity.Usuario, error)
	UpDate(usuario *entity.Usuario)error
}

type UserRepositoryDB struct{
	db *gorm.DB
}

func NewUserRepositoryDB(database *gorm.DB) *UserRepositoryDB{
	return &UserRepositoryDB{db:database}
}

//FindByID busca no banco o usuario antes de altera-lo.
func (r *UserRepositoryDB)FindByID(id uuid.UUID)(*entity.Usuario, error){
	var usuario entity.Usuario

	if err:= r.db.First(&usuario, "id=?", id).Error; err != nil{
		return nil, err
	}
	return &usuario, nil
}

//Update salva as alterações usando .Save do GORM
func (r *UserRepositoryDB)UpDate(usuario *entity.Usuario)error{
	return r.db.Save(usuario).Error
}





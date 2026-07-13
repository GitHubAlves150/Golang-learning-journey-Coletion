/*
A entidade apenas diz ao Go e ao GORM qual é a estrutura e o nome da tabela do dado que o sistema vai manipular.
*/


package entity

import (
	"time"
	"github.com/google/uuid"
)

type Usuario struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	Sexo     string    `json:"sexo"`
	CriadoEm time.Time `gorm:"column:criado_em" json:"criado_em"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
package entity

import (
	"time"

	"github.com/google/uuid"
)

//A nossa Entidade pura. Ela apenas define como o usuário é estruturado no nosso sistema e no banco.

type Usuario struct {
	ID       uuid.UUID `gorm:"typeuuid;primaryKey; default:gen_ramdom_uuid()" json:"id"`
	Nome     string    `gorm:"type:varchar(20)" json:"nome"`
	Email    string    `gorm:"type:varchar(30)" json:"email"`
	Sexo     string    `gorm:"type:varchar(6)" json:"sexo"`
	CriadoEm time.Time `gorm:"column:criado_em;deafault:CURRENT_TIMESTAMP" json:"criado_em"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

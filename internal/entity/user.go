package entity

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Nome     string       `gorm:"type:varchar(20)" json:"nome"`
	Email    string       `gorm:"type:varchar(30)" json:"email"`
	Sexo     string       `gorm:"type:varchar(8)" json:"sexo"`
	Criadoem time.Time    `gorm:"type:timestamp with time zone" json:"criado_em"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

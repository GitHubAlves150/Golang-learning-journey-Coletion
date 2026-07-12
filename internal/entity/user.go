/*
A melhor prática do SOLID (S - Responsabilidade Única) dita que a própria entidade deve saber se ela é válida ou não. Então, colocamos um método Validar() diretamente na nossa Struct.
*/

package entity

import (
	"errors"
	"time"
)

type Usuario struct {
	ID       string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Nome     string    `gorm:"type:varchar(50)" json:"nome"`
	Email    string    `gorm:"type:varchar(40)" json:"email"`
	Sexo     string    `gorm:"type:varchar(10)" json:"sexo"`
	CriadoEm time.Time `gorm:"column:criado_em;default:CURRENT_TIMESTAMP" json:"criado_em"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

// Validar garante que nehum dados obrigatorio vá em branco para o banco
func (u *Usuario) Validar() error {
	if u.Nome == "" || len(u.Nome) < 3 {
		return errors.New("O nome é obrigatório e deve ter pelo menos 3 caracteres")
	}
	if u.Email == "" {
		return errors.New("O Email é obrigatório")
	}
	if u.Sexo != "Masc" && u.Sexo != "Fem" {
		return errors.New("O sexo da pessoa é obrigatório")
	}
	return nil
}

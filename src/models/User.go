package models

import (
	"errors"
	"strings"
	"time"
)

// Estrutura de usuários da aplicação
type User struct {
	ID        uint      `json:"id,omitempty`
	Name      string    `json:"name,omitempty`
	CPF       string    `json:"cpf,omitempty`
	Email     string    `json:"email,omitempty`
	Password  string    `json: password,omitempty`
	Address   string    `json:"address,omitempty`
	Phone     string    `json:"phone,omitempty`
	CreatedAt time.Time `json:"createdAt,omitempty`
	EditedAt  time.Time `json:"editedAt,omitempty`
	DeletedAt time.Time `json:"deletedAt,omitempty`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.formatting()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O campo nome deve ser preenchido")
	}
	if user.CPF == "" {
		return errors.New("O campo CPF deve ser preenchido")
	}
	if user.Email == "" {
		return errors.New("O campo Email deve ser preenchido")
	}
	if user.Password == "" {
		return errors.New("O campo Senha deve ser preenchido")
	}
	if user.Address == "" {
		return errors.New("O campo Endereço deve ser preenchido")
	}
	if user.Phone == "" {
		return errors.New("O campo Telefone deve ser preenchido")
	}
	return nil
}

func (user *User) formatting() {
	user.Name = strings.TrimSpace(user.Name)
	user.CPF = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Name)
	user.Address = strings.TrimSpace(user.Name)
	user.Phone = strings.TrimSpace(user.Name)
}

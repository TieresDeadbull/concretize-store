package models

import "time"

//Estrutura de usuários da aplicação
type User struct {
	ID        uint      `json:"id,omitempty`
	Name      string    `json:"name,omitempty`
	CPF       string    `json:"cpf,omitempty`
	Email     string    `json:"email,omitempty`
	Password  string    `json: password, omitempty`
	Address   uint      `json:"address,omitempty`
	Phone     string    `json:"phone,omitempty`
	CreatedAt time.Time `json:"createdAt,omitempty`
	EditedAt  time.Time `json:"editedAt,omitempty`
	DeletedAt time.Time `json:"deletedAt,omitempty`
}

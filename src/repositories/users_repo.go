package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// Cria novo repositório de usuários
func NewUsersRepo(db *sql.DB) *users {
	return &users{db}
}

// Função que manipula o banco para criação de usuário
func (u users) CreateUser(user models.User) (uint64, error) {
	return 0, nil
}

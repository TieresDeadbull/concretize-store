package repositories

import (
	"api/src/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

// Cria novo repositório de usuários
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

// Função que manipula o banco para criação de usuário
func (userRepo UsersRepo) CreateUser(user models.User) (uint64, error) {
	statement, err := userRepo.db.Prepare(
		"insert into users(name, cpf, email,password, address, phone) values (?,?,?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		user.Name,
		user.CPF,
		user.Email,
		user.Password,
		user.Address,
		user.Phone,
	)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil

}

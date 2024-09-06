package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
		"insert into users (name, cpf, email, password, address, phone) values (?, ?, ?, ?, ?, ?)",
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

// Função que manipula o banco para listagem de usuários
func (userRepo UsersRepo) ListUsers() error {
	lines, err := userRepo.db.Query(
		"select * from users",
	)
	if err != nil {
		return err
	}

	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.CPF,
			&user.Email,
			&user.Password,
			&user.Address,
			&user.Phone,
			&user.CreatedAt,
			&user.EditedAt,
			&user.DeletedAt,
		); err != nil {
			return err
		}
		users = append(users, user)
	}

	fmt.Printf("%v", users)

	return nil
}

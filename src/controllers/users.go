package controllers

import (
	"api/src/database.go"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Errors(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Errors(w, http.StatusBadRequest, err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		responses.Errors(w, http.StatusInternalServerError, err)
	}

	userRepository := repositories.NewUsersRepo(db)
	user.ID, err = userRepository.CreateUser(user)
	if err != nil {
		responses.Errors(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusCreated, user)

}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUsersRepo(db)
	err = userRepository.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("Listando todos os usuários"))
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário por ID"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando Usuário"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}

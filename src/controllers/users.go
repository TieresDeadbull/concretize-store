package controllers

import (
	"api/src/database.go"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUsersRepo(db)
	userRepository.CreateUser(user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
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

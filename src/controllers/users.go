package controllers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
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

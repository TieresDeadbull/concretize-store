package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Retorna resposta em JSON para a requisição
func Errors(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})
}

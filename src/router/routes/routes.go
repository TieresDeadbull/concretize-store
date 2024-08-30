package routes

import "net/http"

//Estrutura de cada rota da aplicação
type Route struct {
	URI          string
	Method       string
	Function     func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

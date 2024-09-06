package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura de cada rota da aplicação
type Route struct {
	URI          string
	Method       string
	Function     func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

// Adiciona as rotas ao router
func ConfigRoutes(m *mux.Router) *mux.Router {
	routes := healthcheckRoute
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		m.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return m
}

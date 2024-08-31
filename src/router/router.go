package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Retorna o router com rotas configuradas
func GenerateRouter() *mux.Router {
	return routes.ConfigRoutes(mux.NewRouter())
}

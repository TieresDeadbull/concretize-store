package router

import "github.com/gorilla/mux"

//Retorna o router com rotas configuradas
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}

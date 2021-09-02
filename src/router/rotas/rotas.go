package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota=route(all routes of api).
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar=(put all routes into the router)
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuários

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}

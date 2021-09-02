package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CriarUsuario=(insert user into database)
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

//(BuscarUsuarios=(search  all users in database)
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search users!"))
}

//BuscarUsuario=(search user in database)
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search user!"))
}

//AtualizarUsuario=(update data user in database)
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user!"))
}

//DeletarUsuario=(delete data user)
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user!"))
}

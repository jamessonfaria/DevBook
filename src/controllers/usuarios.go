package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// CriarUsuario cria um usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Error(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuario especifico no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
w	.Write([]byte("Buscando usuario!"))
}

// AtualizarUsuario atualiza usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario!"))
}

// DeletarUsuario deleta usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario!"))
}
package controllers

import "net/http"

// CriarUsuario cria um usuario no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario!"))
}

// BuscarUsuarios busca todos os usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios!"))
}

// BuscarUsuario busca um usuario especifico no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario!"))
}

// AtualizarUsuario atualiza usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario!"))
}

// DeletarUsuario deleta usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario!"))
}
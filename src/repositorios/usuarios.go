package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioUsuarios cria um repositorio de usuarios
func NovoRepositorioUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar inseri usuarios no banco de dados
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
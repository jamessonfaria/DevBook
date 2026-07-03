package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario modelo de dados do usuario
type Usuario struct {
	ID       uint64 	`json:"id,omitempty"`
	Nome     string 	`json:"nome,omitempty"`
	Nick     string 	`json:"nick,omitempty"`
	Email    string 	`json:"email,omitempty"`
	Senha    string 	`json:"senha,omitempty"`
	CriadoEm time.Time 	`json:"criadoEm,omitempty"`
}

// Preparar validar e formatar o usuario
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("Nome e obrigatorio e nao pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("Nick e obrigatorio e nao pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("Email e obrigatorio e nao pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("Senha e obrigatoria e nao pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
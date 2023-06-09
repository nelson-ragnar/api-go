package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEM time.Time `json:"CriadoEM,omitempty"`
}

func (Usuario *Usuario) Preparar() error {
	if erro := Usuario.validar(); erro != nil {
		return erro
	}
	Usuario.formatar()
	return nil

}

func (Usuario *Usuario) validar() error {
	if Usuario.Nome == "" {
		return errors.New("O Nome é obrigatorio e não pode estar em branco")
	}
	if Usuario.Nick == "" {
		return errors.New("O Nick é obrigatorio e não pode estar em branco")
	}
	if Usuario.Email == "" {
		return errors.New("O E-mail é obrigatorio e não pode estar em branco")
	}
	if Usuario.Senha == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (Usuario *Usuario) formatar() {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)
}

package model

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

// Estrutura Clientes
type Clientes struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome" validate:"required,min=3,max=40"`
	CPF      string `json:"cpf" validate:"required,len=11"`
	Email    string `json:"email" validate:"required,email"`
	Telefone string `json:"telefone" validate:"required,len=11"`
}

// validador
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// função de validação
func (clientes *Clientes) Validate() error {
	return validate.Struct(clientes)
}

// Metado para cadastrar clientes.
func (clientes *Clientes) Create(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO clientes(nome, cpf, email, telefone) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(clientes.Nome, clientes.CPF, clientes.Email, clientes.Telefone)
	return err
}

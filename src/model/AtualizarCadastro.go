package model

import (
	"database/sql"
)

// Estrutura ClientesUpdates
type ClientesUpdates struct {
	ID       int     `json:"id"`
	Nome     *string `json:"nome" validate:",min=3,max=40"`
	CPF      *string `json:"cpf" validate:",len=11"`
	Email    *string `json:"email" validate:",email"`
	Telefone *string `json:"telefone" validate:",len=11"`
}

// Método de atualizar dados dos clientes
func (clientesUpdates *ClientesUpdates) UpdateClientes(db *sql.DB) error {
	// variavel para obter o dado existente
	var existing Clientes
	err := db.QueryRow("SELECT id, nome, cpf, email, telefone FROM clientes WHERE id=?", clientesUpdates.ID).Scan(
		&existing.ID, &existing.Nome, &existing.CPF, &existing.Email, &existing.Telefone)
	if err != nil {
		return err
	}

	// Verificação dos campos nulo
	if clientesUpdates.Nome != nil {
		existing.Nome = *clientesUpdates.Nome
	}
	if clientesUpdates.CPF != nil {
		existing.CPF = *clientesUpdates.CPF
	}
	if clientesUpdates.Email != nil {
		existing.Email = *clientesUpdates.Email
	}
	if clientesUpdates.Telefone != nil {
		existing.Telefone = *clientesUpdates.Telefone
	}

	stmt, err := db.Prepare("UPDATE clientes SET nome=?, cpf=?, email=?, telefone=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(existing.Nome, existing.CPF, existing.Email, existing.Telefone, existing.ID)
	return err
}

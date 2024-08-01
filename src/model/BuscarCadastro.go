package model

import "database/sql"


// Metado para listar todos os clientes cadastrado.
func GetAllClientes(db *sql.DB) ([]*Clientes, error) {
	rows, err := db.Query("SELECT id, nome, cpf, email, telefone FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []*Clientes
	for rows.Next() {
		var cliente Clientes
		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.CPF, &cliente.Email, &cliente.Telefone); err != nil {
			return nil, err
		}
		clientes = append(clientes, &cliente)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return clientes, nil
}

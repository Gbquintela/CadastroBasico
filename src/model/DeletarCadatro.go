package model

import "database/sql"

func (clientes *Clientes) Delete(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM clientes WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(clientes.ID)
	return err
}

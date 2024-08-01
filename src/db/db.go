package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Altere os detalhes da string de conexão para o seu banco de dados
	DB, err = sql.Open("mysql", "root:Mafioso55@@tcp(127.0.0.1:3306)/cadastrosimples")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"CadastroBasico/src/controller"
	"CadastroBasico/src/db"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	router := mux.NewRouter()
	log.Println("Executando ...")

	router.HandleFunc("/cadastro", controller.CreateCadastro).Methods("POST")
	router.HandleFunc("/clientes", controller.GetCadastros).Methods("GET")
	router.HandleFunc("/atualizar", controller.UpdateCadastro).Methods("PUT")
	router.HandleFunc("/deletar", controller.DeleteCadastro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

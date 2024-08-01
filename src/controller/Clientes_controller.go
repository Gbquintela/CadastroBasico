package controller

import (
	"CadastroBasico/src/db"
	"CadastroBasico/src/model"
	"CadastroBasico/src/view"
	"encoding/json"
	"net/http"
)

type RespostaMensagem struct {
	Mensagem string `json:"mensagem"`
}

// Função responsável por criar um Cadastro
func CreateCadastro(w http.ResponseWriter, r *http.Request) {
	var newCadastro model.Clientes
	err := json.NewDecoder(r.Body).Decode(&newCadastro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := newCadastro.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = newCadastro.Create(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resposta := RespostaMensagem{Mensagem: "Cliente cadastrado com sucesso"}
	view.JSONResponse(w, http.StatusOK, newCadastro, resposta)
}

// Função responsável por obter todos os Cadastros
func GetCadastros(w http.ResponseWriter, r *http.Request) {
	Cadastros, err := model.GetAllClientes(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resposta := RespostaMensagem{Mensagem: "Lista de Clientes"}
	view.JSONResponse(w, http.StatusOK, Cadastros, resposta)
}

// Função responsável por atualizar um Cadastro
func UpdateCadastro(w http.ResponseWriter, r *http.Request) {
	var updatedCadastro model.ClientesUpdates
	err := json.NewDecoder(r.Body).Decode(&updatedCadastro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updatedCadastro.UpdateClientes(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resposta := RespostaMensagem{Mensagem: "Cliente deletado com sucesso"}
	view.JSONResponse(w, http.StatusOK, updatedCadastro, resposta)
}

// Função responsável por deletar um Cadastro

func DeleteCadastro(w http.ResponseWriter, r *http.Request) {
	var CadastroToDelete model.Clientes
	err := json.NewDecoder(r.Body).Decode(&CadastroToDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CadastroToDelete.Delete(db.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resposta := RespostaMensagem{Mensagem: "Cliente deletado com sucesso"}
	view.JSONResponse(w, http.StatusOK, DeleteCadastro, resposta)
}

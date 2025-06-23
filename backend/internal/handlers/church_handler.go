package handlers

import (
	tabledatagateways "church-for-all/internal/data/table_data_gateways"
	"church-for-all/internal/models"
	"encoding/json"
	"net/http"
)

// CreateChurch é uma função que cria uma igreja
func CreateChurch(w http.ResponseWriter, r *http.Request) {
	// Responsável por definir o tipo de conteúdo da resposta
	w.Header().Set("Content-Type", "application/json")

	// Responsável por decodificar o corpo da requisição e armazenar em uma variável
	var church models.Church
	err := json.NewDecoder(r.Body).Decode(&church)

	// Responsável por verificar se ocorreu algum erro ao decodificar o corpo da requisição
	if err != nil {
		// Responsável por retornar um erro 400 caso o corpo da requisição seja inválido
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Responsável por criar uma instância da estrutura ChurchTableDataGateway
	churchTableDataGateway := tabledatagateways.NewChurchTableDataGateway()
	// Responsável por criar uma igreja no banco de dados
	err = churchTableDataGateway.CreateChurch(church)

	// Responsável por verificar se ocorreu algum erro ao criar uma igreja no banco de dados
	if err != nil {
		// Responsável por retornar um erro 500 caso ocorra um erro ao criar uma igreja no banco de dados
		http.Error(w, "Failed to create church", http.StatusInternalServerError)
		return
	}
	// Responsável por retornar a igreja cadastrada
	json.NewEncoder(w).Encode(church)
}

// GetChurchs é uma função que busca todas as igrejas
func GetChurchs(w http.ResponseWriter, r *http.Request) {
	// Responsável por definir o tipo de conteúdo da resposta
	w.Header().Set("Content-Type", "application/json")

	// Responsável por buscar todas as igrejas no banco de dados
	churchTableDataGateway := tabledatagateways.NewChurchTableDataGateway()
	// Responsável por buscar todas as igrejas no banco de dados
	churches, err := churchTableDataGateway.GetChurchs()

	// Responsável por verificar se ocorreu algum erro ao buscar as igrejas no banco de dados
	if err != nil {
		// Responsável por retornar um erro 500 caso ocorra um erro ao buscar as igrejas no banco de dados
		http.Error(w, "Failed to get churches", http.StatusInternalServerError)
		return
	}

	// Responsável por retornar a lista de igrejas
	json.NewEncoder(w).Encode(churches)
}

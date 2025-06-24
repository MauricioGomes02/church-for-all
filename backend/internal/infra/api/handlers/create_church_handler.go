package handlers

import (
	usecases "church-for-all/internal/application/use_cases"
	"church-for-all/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateChurchHandler é responsável por encapsular o tratamento de uma requisição de criação de uma igreja
type CreateChurchHandler struct {
	createChurchUseCase *usecases.CreateChurchUseCase
}

// NewCreateChurchHandler é um método responsável por criar uma nova instância de CreateChurchHandler
func NewCreateChurchHandler(c *usecases.CreateChurchUseCase) *CreateChurchHandler {
	return &CreateChurchHandler{createChurchUseCase: c}
}

// Create é um método responsável por tratar uma requisição de criação de uma igreja
func (c *CreateChurchHandler) Create(ginContext *gin.Context) {
	var church entities.Church
	err := ginContext.ShouldBindJSON(&church)

	// Responsável por verificar se ocorreu algum erro ao decodificar o corpo da requisição
	if err != nil {
		// Responsável por retornar um erro 400 caso o corpo da requisição seja inválido
		ginContext.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	err = c.createChurchUseCase.Create(&church)

	// Responsável por verificar se ocorreu algum erro ao criar uma igreja no banco de dados
	if err != nil {
		// Responsável por retornar um erro 500 caso ocorra um erro ao criar uma igreja no banco de dados
		ginContext.JSON(http.StatusInternalServerError, "Failed to create church")
		return
	}

	// Responsável por retornar a igreja cadastrada
	ginContext.JSON(http.StatusOK, church)
}

package handlers

import (
	usecases "church-for-all/internal/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetChurchsHandler é responsável por encapsular o tratamento de uma requisição de obter igrejas
type GetChurchsHandler struct {
	getChurchsUseCase *usecases.GetChurchsUseCase
}

// NewGetChurchsHandler é um método responsável por criar uma nova instância de GetChurchsHandler
func NewGetChurchsHandler(g *usecases.GetChurchsUseCase) *GetChurchsHandler {
	return &GetChurchsHandler{getChurchsUseCase: g}
}

// Get é um método responsável por tratar a requisição de obter igrejas
func (g *GetChurchsHandler) Get(ginContext *gin.Context) {
	// Responsável por obter todas as igrejas
	churches, err := g.getChurchsUseCase.Get()

	// Responsável por verificar se ocorreu algum erro ao obter as igrejas
	if err != nil {
		// Responsável por retornar um erro 500 caso ocorra um erro ao obter as igrejas
		ginContext.JSON(http.StatusInternalServerError, "Failed to get churches")
		return
	}

	// Responsável por retornar a lista de igrejas
	ginContext.JSON(http.StatusOK, churches)
}

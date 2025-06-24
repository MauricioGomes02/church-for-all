package usecases

import (
	"church-for-all/internal/domain/entities"
	"church-for-all/internal/domain/interfaces/infra/data"
)

// GetChurchsUseCase é responável por encapsular o caso de uso de obter igrejas
type GetChurchsUseCase struct {
	getChurchsDataManagement data.GetChurchsDataManagement
}

// NewGetChurchsUseCase é um método responsável por criar uma nova instância de GetChurchsUseCase
func NewGetChurchsUseCase(g data.GetChurchsDataManagement) *GetChurchsUseCase {
	return &GetChurchsUseCase{getChurchsDataManagement: g}
}

// Get é um método responsável por orquestrar o caso de uso de obter igrejas
func (g *GetChurchsUseCase) Get() ([]entities.Church, error) {
	return g.getChurchsDataManagement.Get()
}

package usecases

import (
	"church-for-all/internal/domain/entities"
	"church-for-all/internal/domain/interfaces/infra/data"
)

// CreateChurchUseCase é responsável por encapsular toda a lógica do caso de uso de criação de uma igreja
type CreateChurchUseCase struct {
	createChurchDataManagement data.CreateChurchDataManagement
}

// NewCreateChurchUseCase é um método responsável por criar uma nova instância de CreateChurchUseCase
func NewCreateChurchUseCase(c data.CreateChurchDataManagement) *CreateChurchUseCase {
	return &CreateChurchUseCase{createChurchDataManagement: c}
}

// Create é um método responsável por orquestrar o caso de uso de criação de uma igreja
func (c CreateChurchUseCase) Create(church *entities.Church) error {
	err := c.createChurchDataManagement.Create(church)
	return err
}

package data

import "church-for-all/internal/domain/entities"

// GetChurchsDataManagement é responsável por encapsular toda a lógica de gerenciamento de dados para a obtenção de igrejas
type GetChurchsDataManagement interface {
	Get() ([]entities.Church, error)
}

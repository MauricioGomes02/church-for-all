package data

import "church-for-all/internal/domain/entities"

type CreateChurchDataManagement interface {
	Create(church *entities.Church) error
}

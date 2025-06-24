package usecases

import (
	"church-for-all/internal/domain/entities"
	"church-for-all/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Dado_um_conjunto_de_igrejas_Quando_obter_Deve_listar_todas(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	createChurchDataManagement := mocks.NewMockGetChurchsDataManagement(control)
	sut := NewGetChurchsUseCase(createChurchDataManagement)

	churchs := []entities.Church{
		{
			ID:      1,
			Name:    "Igreja Central",
			Address: "Rua das Flores, 123",
			Phones:  []string{"11999999999", "11888888888"},
			Email:   "contato@igreja.com",
		},
	}
	createChurchDataManagement.EXPECT().Get().Return(churchs, nil)

	storedChurchs, err := sut.Get()

	assert.NotEmpty(t, storedChurchs)
	assert.Nil(t, err)
}

package usecases

import (
	"church-for-all/internal/domain/entities"
	"church-for-all/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Dado_uma_igreja_Quando_criar_Nao_deve_lancar_erro(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	createChurchDataManagement := mocks.NewMockCreateChurchDataManagement(control)
	sut := NewCreateChurchUseCase(createChurchDataManagement)

	createChurchDataManagement.EXPECT().Create(gomock.Any()).Return(nil)

	church := &entities.Church{
		ID:      1,
		Name:    "Igreja Central",
		Address: "Rua das Flores, 123",
		Phones:  []string{"11999999999", "11888888888"},
		Email:   "contato@igreja.com",
	}

	err := sut.Create(church)

	assert.Nil(t, err)
}

package mocks

import (
	entities "church-for-all/internal/domain/entities"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCreateChurchDataManagement is a mock of CreateChurchDataManagement interface.
type MockCreateChurchDataManagement struct {
	ctrl     *gomock.Controller
	recorder *MockCreateChurchDataManagementMockRecorder
	isgomock struct{}
}

// MockCreateChurchDataManagementMockRecorder is the mock recorder for MockCreateChurchDataManagement.
type MockCreateChurchDataManagementMockRecorder struct {
	mock *MockCreateChurchDataManagement
}

// NewMockCreateChurchDataManagement creates a new mock instance.
func NewMockCreateChurchDataManagement(ctrl *gomock.Controller) *MockCreateChurchDataManagement {
	mock := &MockCreateChurchDataManagement{ctrl: ctrl}
	mock.recorder = &MockCreateChurchDataManagementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateChurchDataManagement) EXPECT() *MockCreateChurchDataManagementMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreateChurchDataManagement) Create(church *entities.Church) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", church)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCreateChurchDataManagementMockRecorder) Create(church any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateChurchDataManagement)(nil).Create), church)
}

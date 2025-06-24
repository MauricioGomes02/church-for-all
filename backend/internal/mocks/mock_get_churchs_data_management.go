package mocks

import (
	entities "church-for-all/internal/domain/entities"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGetChurchsDataManagement is a mock of GetChurchsDataManagement interface.
type MockGetChurchsDataManagement struct {
	ctrl     *gomock.Controller
	recorder *MockGetChurchsDataManagementMockRecorder
	isgomock struct{}
}

// MockGetChurchsDataManagementMockRecorder is the mock recorder for MockGetChurchsDataManagement.
type MockGetChurchsDataManagementMockRecorder struct {
	mock *MockGetChurchsDataManagement
}

// NewMockGetChurchsDataManagement creates a new mock instance.
func NewMockGetChurchsDataManagement(ctrl *gomock.Controller) *MockGetChurchsDataManagement {
	mock := &MockGetChurchsDataManagement{ctrl: ctrl}
	mock.recorder = &MockGetChurchsDataManagementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetChurchsDataManagement) EXPECT() *MockGetChurchsDataManagementMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGetChurchsDataManagement) Get() ([]entities.Church, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]entities.Church)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGetChurchsDataManagementMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGetChurchsDataManagement)(nil).Get))
}

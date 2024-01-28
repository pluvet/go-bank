// Package mock_repositories is a generated GoMock package.
package repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepositoryInterface is a mock of IUserRepositoryInterface interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryInterfaceMockRecorder is the mock recorder for MockIUserRepositoryInterface
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepositoryInterface creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockIUserRepository) CreateUser(arg0, arg1, arg2 string) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockIUserRepositoryMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateUser), arg0, arg1, arg2)
}

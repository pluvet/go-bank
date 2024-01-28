package repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/pluvet/go-bank/app/models"
)

// MockIAccountRepository is a mock of IAccountRepository interface
type MockIAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRepositoryMockRecorder
}

// MockIAccountRepositoryMockRecorder is the mock recorder for MockIAccountRepository
type MockIAccountRepositoryMockRecorder struct {
	mock *MockIAccountRepository
}

// NewMockIAccountRepository creates a new mock instance
func NewMockIAccountRepository(ctrl *gomock.Controller) *MockIAccountRepository {
	mock := &MockIAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAccountRepository) EXPECT() *MockIAccountRepositoryMockRecorder {
	return m.recorder
}

// FindAccount mocks base method
func (m *MockIAccountRepository) FindAccount(arg0 int) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAccount", arg0)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAccount indicates an expected call of FindAccount
func (mr *MockIAccountRepositoryMockRecorder) FindAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAccount", reflect.TypeOf((*MockIAccountRepository)(nil).FindAccount), arg0)
}

// UpdateAccount mocks base method
func (m *MockIAccountRepository) UpdateAccount(arg0 *models.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockIAccountRepositoryMockRecorder) UpdateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockIAccountRepository)(nil).UpdateAccount), arg0)
}

// CreateAccount mocks base method
func (m *MockIAccountRepository) CreateAccount(arg0 int) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockIAccountRepositoryMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIAccountRepository)(nil).CreateAccount), arg0)
}

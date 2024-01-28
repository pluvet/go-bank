package services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAccountService is a mock of IAccountService interface
type MockIAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountServiceMockRecorder
}

// MockIAccountServiceMockRecorder is the mock recorder for MockIAccountService
type MockIAccountServiceMockRecorder struct {
	mock *MockIAccountService
}

// NewMockIAccountService creates a new mock instance
func NewMockIAccountService(ctrl *gomock.Controller) *MockIAccountService {
	mock := &MockIAccountService{ctrl: ctrl}
	mock.recorder = &MockIAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAccountService) EXPECT() *MockIAccountServiceMockRecorder {
	return m.recorder
}

// AccountDeposit mocks base method
func (m *MockIAccountService) AccountDeposit(arg0 int, arg1 float32) (*float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountDeposit", arg0, arg1)
	ret0, _ := ret[0].(*float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountDeposit indicates an expected call of AccountDeposit
func (mr *MockIAccountServiceMockRecorder) AccountDeposit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountDeposit", reflect.TypeOf((*MockIAccountService)(nil).AccountDeposit), arg0, arg1)
}

// AccountWithdraw mocks base method
func (m *MockIAccountService) AccountWithdraw(arg0 int, arg1 float32) (*float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountWithdraw", arg0, arg1)
	ret0, _ := ret[0].(*float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountWithdraw indicates an expected call of AccountWithdraw
func (mr *MockIAccountServiceMockRecorder) AccountWithdraw(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountWithdraw", reflect.TypeOf((*MockIAccountService)(nil).AccountWithdraw), arg0, arg1)
}

// CreateAccount mocks base method
func (m *MockIAccountService) CreateAccount(arg0 int) (*int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockIAccountServiceMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockIAccountService)(nil).CreateAccount), arg0)
}

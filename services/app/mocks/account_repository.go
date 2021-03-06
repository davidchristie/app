// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/davidchristie/app/services/app/repositories (interfaces: AccountRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/davidchristie/app/services/app/entities"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockAccountRepository) FindByID(arg0 context.Context, arg1 uuid.UUID) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockAccountRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAccountRepository)(nil).FindByID), arg0, arg1)
}

// FindByProvider mocks base method.
func (m *MockAccountRepository) FindByProvider(arg0 context.Context, arg1, arg2, arg3 string) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByProvider", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByProvider indicates an expected call of FindByProvider.
func (mr *MockAccountRepositoryMockRecorder) FindByProvider(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByProvider", reflect.TypeOf((*MockAccountRepository)(nil).FindByProvider), arg0, arg1, arg2, arg3)
}

// Insert mocks base method.
func (m *MockAccountRepository) Insert(arg0 context.Context, arg1 *entities.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAccountRepositoryMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAccountRepository)(nil).Insert), arg0, arg1)
}

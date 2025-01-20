// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/dduuddeekk/go-restaurant-app/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of Repository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CheckRegistered mocks base method.
func (m *MockUserRepository) CheckRegistered(ctx context.Context, username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRegistered", ctx, username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckRegistered indicates an expected call of CheckRegistered.
func (mr *MockUserRepositoryMockRecorder) CheckRegistered(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRegistered", reflect.TypeOf((*MockUserRepository)(nil).CheckRegistered), ctx, username)
}

// CheckSession mocks base method.
func (m *MockUserRepository) CheckSession(ctx context.Context, data model.UserSession) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSession", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockUserRepositoryMockRecorder) CheckSession(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockUserRepository)(nil).CheckSession), ctx, data)
}

// CreateUserSession mocks base method.
func (m *MockUserRepository) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserSession", ctx, userID)
	ret0, _ := ret[0].(model.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserSession indicates an expected call of CreateUserSession.
func (mr *MockUserRepositoryMockRecorder) CreateUserSession(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserSession", reflect.TypeOf((*MockUserRepository)(nil).CreateUserSession), ctx, userID)
}

// GenerateUserHash mocks base method.
func (m *MockUserRepository) GenerateUserHash(ctx context.Context, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUserHash", ctx, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateUserHash indicates an expected call of GenerateUserHash.
func (mr *MockUserRepositoryMockRecorder) GenerateUserHash(ctx, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUserHash", reflect.TypeOf((*MockUserRepository)(nil).GenerateUserHash), ctx, password)
}

// GetUserData mocks base method.
func (m *MockUserRepository) GetUserData(ctx context.Context, username string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserData", ctx, username)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockUserRepositoryMockRecorder) GetUserData(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockUserRepository)(nil).GetUserData), ctx, username)
}

// RegisterUser mocks base method.
func (m *MockUserRepository) RegisterUser(ctx context.Context, userData model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, userData)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserRepositoryMockRecorder) RegisterUser(ctx, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserRepository)(nil).RegisterUser), ctx, userData)
}

// VerifyLogin mocks base method.
func (m *MockUserRepository) VerifyLogin(ctx context.Context, username, password string, userData model.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyLogin", ctx, username, password, userData)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyLogin indicates an expected call of VerifyLogin.
func (mr *MockUserRepositoryMockRecorder) VerifyLogin(ctx, username, password, userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyLogin", reflect.TypeOf((*MockUserRepository)(nil).VerifyLogin), ctx, username, password, userData)
}

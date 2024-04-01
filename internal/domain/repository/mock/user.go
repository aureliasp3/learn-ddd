// Code generated by MockGen. DO NOT EDIT.
// Source: ./user.go
//
// Generated by this command:
//
//	mockgen -package=mock -source=./user.go -destination=./mock/user.go
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "learn-ddd/internal/domain/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
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

// Fetch mocks base method.
func (m *MockUserRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockUserRepositoryMockRecorder) Fetch(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockUserRepository)(nil).Fetch), ctx)
}

// Select mocks base method.
func (m *MockUserRepository) Select(ctx context.Context, id int64) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockUserRepositoryMockRecorder) Select(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockUserRepository)(nil).Select), ctx, id)
}

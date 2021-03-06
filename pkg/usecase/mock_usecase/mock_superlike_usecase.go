// Code generated by MockGen. DO NOT EDIT.
// Source: superlike_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

// MockSuperlikeUseCase is a mock of SuperlikeUseCase interface.
type MockSuperlikeUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSuperlikeUseCaseMockRecorder
}

// MockSuperlikeUseCaseMockRecorder is the mock recorder for MockSuperlikeUseCase.
type MockSuperlikeUseCaseMockRecorder struct {
	mock *MockSuperlikeUseCase
}

// NewMockSuperlikeUseCase creates a new mock instance.
func NewMockSuperlikeUseCase(ctrl *gomock.Controller) *MockSuperlikeUseCase {
	mock := &MockSuperlikeUseCase{ctrl: ctrl}
	mock.recorder = &MockSuperlikeUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuperlikeUseCase) EXPECT() *MockSuperlikeUseCaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSuperlikeUseCase) Delete(ctx context.Context, entity *model.Superlike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSuperlikeUseCaseMockRecorder) Delete(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSuperlikeUseCase)(nil).Delete), ctx, entity)
}

// Insert mocks base method.
func (m *MockSuperlikeUseCase) Insert(ctx context.Context, entity *model.Superlike, UID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity, UID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSuperlikeUseCaseMockRecorder) Insert(ctx, entity, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSuperlikeUseCase)(nil).Insert), ctx, entity, UID)
}

// Select mocks base method.
func (m *MockSuperlikeUseCase) Select(ctx context.Context, UID string) ([]*model.Superlike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx, UID)
	ret0, _ := ret[0].([]*model.Superlike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockSuperlikeUseCaseMockRecorder) Select(ctx, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockSuperlikeUseCase)(nil).Select), ctx, UID)
}

// Update mocks base method.
func (m *MockSuperlikeUseCase) Update(ctx context.Context, entity *model.Superlike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSuperlikeUseCaseMockRecorder) Update(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSuperlikeUseCase)(nil).Update), ctx, entity)
}

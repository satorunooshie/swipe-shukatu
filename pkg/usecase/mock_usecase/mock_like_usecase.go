// Code generated by MockGen. DO NOT EDIT.
// Source: like_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

// MockLikeUseCase is a mock of LikeUseCase interface.
type MockLikeUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockLikeUseCaseMockRecorder
}

// MockLikeUseCaseMockRecorder is the mock recorder for MockLikeUseCase.
type MockLikeUseCaseMockRecorder struct {
	mock *MockLikeUseCase
}

// NewMockLikeUseCase creates a new mock instance.
func NewMockLikeUseCase(ctrl *gomock.Controller) *MockLikeUseCase {
	mock := &MockLikeUseCase{ctrl: ctrl}
	mock.recorder = &MockLikeUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLikeUseCase) EXPECT() *MockLikeUseCaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockLikeUseCase) Delete(ctx context.Context, entity *model.Like) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLikeUseCaseMockRecorder) Delete(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLikeUseCase)(nil).Delete), ctx, entity)
}

// Insert mocks base method.
func (m *MockLikeUseCase) Insert(ctx context.Context, entity *model.Like, UID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity, UID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockLikeUseCaseMockRecorder) Insert(ctx, entity, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLikeUseCase)(nil).Insert), ctx, entity, UID)
}

// Select mocks base method.
func (m *MockLikeUseCase) Select(ctx context.Context, UID string) ([]*model.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx, UID)
	ret0, _ := ret[0].([]*model.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockLikeUseCaseMockRecorder) Select(ctx, UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockLikeUseCase)(nil).Select), ctx, UID)
}

// Update mocks base method.
func (m *MockLikeUseCase) Update(ctx context.Context, entity *model.Like) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLikeUseCaseMockRecorder) Update(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLikeUseCase)(nil).Update), ctx, entity)
}

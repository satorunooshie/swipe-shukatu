// Code generated by MockGen. DO NOT EDIT.
// Source: superlike_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

// MockSuperlikeRepository is a mock of SuperlikeRepository interface.
type MockSuperlikeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSuperlikeRepositoryMockRecorder
}

// MockSuperlikeRepositoryMockRecorder is the mock recorder for MockSuperlikeRepository.
type MockSuperlikeRepositoryMockRecorder struct {
	mock *MockSuperlikeRepository
}

// NewMockSuperlikeRepository creates a new mock instance.
func NewMockSuperlikeRepository(ctrl *gomock.Controller) *MockSuperlikeRepository {
	mock := &MockSuperlikeRepository{ctrl: ctrl}
	mock.recorder = &MockSuperlikeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuperlikeRepository) EXPECT() *MockSuperlikeRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockSuperlikeRepository) Delete(ctx context.Context, entity *model.SuperLike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSuperlikeRepositoryMockRecorder) Delete(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSuperlikeRepository)(nil).Delete), ctx, entity)
}

// Insert mocks base method.
func (m *MockSuperlikeRepository) Insert(ctx context.Context, entity *model.SuperLike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockSuperlikeRepositoryMockRecorder) Insert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSuperlikeRepository)(nil).Insert), ctx, entity)
}

// Select mocks base method.
func (m *MockSuperlikeRepository) Select(ctx context.Context) ([]*model.SuperLike, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx)
	ret0, _ := ret[0].([]*model.SuperLike)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockSuperlikeRepositoryMockRecorder) Select(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockSuperlikeRepository)(nil).Select), ctx)
}

// Update mocks base method.
func (m *MockSuperlikeRepository) Update(ctx context.Context, entity *model.SuperLike) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockSuperlikeRepositoryMockRecorder) Update(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSuperlikeRepository)(nil).Update), ctx, entity)
}

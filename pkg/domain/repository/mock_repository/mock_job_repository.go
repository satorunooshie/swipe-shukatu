// Code generated by MockGen. DO NOT EDIT.
// Source: job_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

// MockJobRepository is a mock of JobRepository interface.
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository.
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance.
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockJobRepository) Delete(ctx context.Context, entity *model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockJobRepositoryMockRecorder) Delete(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockJobRepository)(nil).Delete), ctx, entity)
}

// Insert mocks base method.
func (m *MockJobRepository) Insert(ctx context.Context, entity *model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockJobRepositoryMockRecorder) Insert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockJobRepository)(nil).Insert), ctx, entity)
}

// Select mocks base method.
func (m *MockJobRepository) Select(ctx context.Context) ([]*model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx)
	ret0, _ := ret[0].([]*model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockJobRepositoryMockRecorder) Select(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockJobRepository)(nil).Select), ctx)
}

// Update mocks base method.
func (m *MockJobRepository) Update(ctx context.Context, entity *model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockJobRepositoryMockRecorder) Update(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockJobRepository)(nil).Update), ctx, entity)
}

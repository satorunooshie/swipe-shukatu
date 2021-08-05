// Code generated by MockGen. DO NOT EDIT.
// Source: message_usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/satorunooshie/swipe-shukatu/pkg/domain/model"
)

// MockMessageUseCase is a mock of MessageUseCase interface.
type MockMessageUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockMessageUseCaseMockRecorder
}

// MockMessageUseCaseMockRecorder is the mock recorder for MockMessageUseCase.
type MockMessageUseCaseMockRecorder struct {
	mock *MockMessageUseCase
}

// NewMockMessageUseCase creates a new mock instance.
func NewMockMessageUseCase(ctrl *gomock.Controller) *MockMessageUseCase {
	mock := &MockMessageUseCase{ctrl: ctrl}
	mock.recorder = &MockMessageUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageUseCase) EXPECT() *MockMessageUseCaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMessageUseCase) Delete(ctx context.Context, entity *model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageUseCaseMockRecorder) Delete(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageUseCase)(nil).Delete), ctx, entity)
}

// Insert mocks base method.
func (m *MockMessageUseCase) Insert(ctx context.Context, entity *model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockMessageUseCaseMockRecorder) Insert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMessageUseCase)(nil).Insert), ctx, entity)
}

// Select mocks base method.
func (m *MockMessageUseCase) Select(ctx context.Context) ([]*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockMessageUseCaseMockRecorder) Select(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockMessageUseCase)(nil).Select), ctx)
}

// Update mocks base method.
func (m *MockMessageUseCase) Update(ctx context.Context, entity *model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMessageUseCaseMockRecorder) Update(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMessageUseCase)(nil).Update), ctx, entity)
}

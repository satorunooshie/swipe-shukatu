// Code generated by MockGen. DO NOT EDIT.
// Source: superlike_handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSuperlikeHandler is a mock of SuperlikeHandler interface.
type MockSuperlikeHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSuperlikeHandlerMockRecorder
}

// MockSuperlikeHandlerMockRecorder is the mock recorder for MockSuperlikeHandler.
type MockSuperlikeHandlerMockRecorder struct {
	mock *MockSuperlikeHandler
}

// NewMockSuperlikeHandler creates a new mock instance.
func NewMockSuperlikeHandler(ctrl *gomock.Controller) *MockSuperlikeHandler {
	mock := &MockSuperlikeHandler{ctrl: ctrl}
	mock.recorder = &MockSuperlikeHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSuperlikeHandler) EXPECT() *MockSuperlikeHandlerMockRecorder {
	return m.recorder
}

// HandleDelete mocks base method.
func (m *MockSuperlikeHandler) HandleDelete() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDelete")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleDelete indicates an expected call of HandleDelete.
func (mr *MockSuperlikeHandlerMockRecorder) HandleDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDelete", reflect.TypeOf((*MockSuperlikeHandler)(nil).HandleDelete))
}

// HandleInsert mocks base method.
func (m *MockSuperlikeHandler) HandleInsert() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInsert")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleInsert indicates an expected call of HandleInsert.
func (mr *MockSuperlikeHandlerMockRecorder) HandleInsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInsert", reflect.TypeOf((*MockSuperlikeHandler)(nil).HandleInsert))
}

// HandleSelect mocks base method.
func (m *MockSuperlikeHandler) HandleSelect() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSelect")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleSelect indicates an expected call of HandleSelect.
func (mr *MockSuperlikeHandlerMockRecorder) HandleSelect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSelect", reflect.TypeOf((*MockSuperlikeHandler)(nil).HandleSelect))
}

// HandleUpdate mocks base method.
func (m *MockSuperlikeHandler) HandleUpdate() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockSuperlikeHandlerMockRecorder) HandleUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockSuperlikeHandler)(nil).HandleUpdate))
}

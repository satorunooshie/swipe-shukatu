// Code generated by MockGen. DO NOT EDIT.
// Source: nope_handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNopeHandler is a mock of NopeHandler interface.
type MockNopeHandler struct {
	ctrl     *gomock.Controller
	recorder *MockNopeHandlerMockRecorder
}

// MockNopeHandlerMockRecorder is the mock recorder for MockNopeHandler.
type MockNopeHandlerMockRecorder struct {
	mock *MockNopeHandler
}

// NewMockNopeHandler creates a new mock instance.
func NewMockNopeHandler(ctrl *gomock.Controller) *MockNopeHandler {
	mock := &MockNopeHandler{ctrl: ctrl}
	mock.recorder = &MockNopeHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNopeHandler) EXPECT() *MockNopeHandlerMockRecorder {
	return m.recorder
}

// HandleDelete mocks base method.
func (m *MockNopeHandler) HandleDelete() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDelete")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleDelete indicates an expected call of HandleDelete.
func (mr *MockNopeHandlerMockRecorder) HandleDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDelete", reflect.TypeOf((*MockNopeHandler)(nil).HandleDelete))
}

// HandleInsert mocks base method.
func (m *MockNopeHandler) HandleInsert() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInsert")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleInsert indicates an expected call of HandleInsert.
func (mr *MockNopeHandlerMockRecorder) HandleInsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInsert", reflect.TypeOf((*MockNopeHandler)(nil).HandleInsert))
}

// HandleSelect mocks base method.
func (m *MockNopeHandler) HandleSelect() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSelect")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleSelect indicates an expected call of HandleSelect.
func (mr *MockNopeHandlerMockRecorder) HandleSelect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSelect", reflect.TypeOf((*MockNopeHandler)(nil).HandleSelect))
}

// HandleUpdate mocks base method.
func (m *MockNopeHandler) HandleUpdate() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockNopeHandlerMockRecorder) HandleUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockNopeHandler)(nil).HandleUpdate))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: job_handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJobHandler is a mock of JobHandler interface.
type MockJobHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJobHandlerMockRecorder
}

// MockJobHandlerMockRecorder is the mock recorder for MockJobHandler.
type MockJobHandlerMockRecorder struct {
	mock *MockJobHandler
}

// NewMockJobHandler creates a new mock instance.
func NewMockJobHandler(ctrl *gomock.Controller) *MockJobHandler {
	mock := &MockJobHandler{ctrl: ctrl}
	mock.recorder = &MockJobHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobHandler) EXPECT() *MockJobHandlerMockRecorder {
	return m.recorder
}

// HandleDelete mocks base method.
func (m *MockJobHandler) HandleDelete() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDelete")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleDelete indicates an expected call of HandleDelete.
func (mr *MockJobHandlerMockRecorder) HandleDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDelete", reflect.TypeOf((*MockJobHandler)(nil).HandleDelete))
}

// HandleInsert mocks base method.
func (m *MockJobHandler) HandleInsert() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInsert")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleInsert indicates an expected call of HandleInsert.
func (mr *MockJobHandlerMockRecorder) HandleInsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInsert", reflect.TypeOf((*MockJobHandler)(nil).HandleInsert))
}

// HandleSelect mocks base method.
func (m *MockJobHandler) HandleSelect() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSelect")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleSelect indicates an expected call of HandleSelect.
func (mr *MockJobHandlerMockRecorder) HandleSelect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSelect", reflect.TypeOf((*MockJobHandler)(nil).HandleSelect))
}

// HandleUpdate mocks base method.
func (m *MockJobHandler) HandleUpdate() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockJobHandlerMockRecorder) HandleUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockJobHandler)(nil).HandleUpdate))
}

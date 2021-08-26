// Code generated by MockGen. DO NOT EDIT.
// Source: matchlist_handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMatchlistHandler is a mock of MatchlistHandler interface.
type MockMatchlistHandler struct {
	ctrl     *gomock.Controller
	recorder *MockMatchlistHandlerMockRecorder
}

// MockMatchlistHandlerMockRecorder is the mock recorder for MockMatchlistHandler.
type MockMatchlistHandlerMockRecorder struct {
	mock *MockMatchlistHandler
}

// NewMockMatchlistHandler creates a new mock instance.
func NewMockMatchlistHandler(ctrl *gomock.Controller) *MockMatchlistHandler {
	mock := &MockMatchlistHandler{ctrl: ctrl}
	mock.recorder = &MockMatchlistHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchlistHandler) EXPECT() *MockMatchlistHandlerMockRecorder {
	return m.recorder
}

// HandleDelete mocks base method.
func (m *MockMatchlistHandler) HandleDelete() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDelete")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleDelete indicates an expected call of HandleDelete.
func (mr *MockMatchlistHandlerMockRecorder) HandleDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDelete", reflect.TypeOf((*MockMatchlistHandler)(nil).HandleDelete))
}

// HandleInsert mocks base method.
func (m *MockMatchlistHandler) HandleInsert() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInsert")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleInsert indicates an expected call of HandleInsert.
func (mr *MockMatchlistHandlerMockRecorder) HandleInsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInsert", reflect.TypeOf((*MockMatchlistHandler)(nil).HandleInsert))
}

// HandleSelect mocks base method.
func (m *MockMatchlistHandler) HandleSelect() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleSelect")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleSelect indicates an expected call of HandleSelect.
func (mr *MockMatchlistHandlerMockRecorder) HandleSelect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleSelect", reflect.TypeOf((*MockMatchlistHandler)(nil).HandleSelect))
}

// HandleUpdate mocks base method.
func (m *MockMatchlistHandler) HandleUpdate() http.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdate")
	ret0, _ := ret[0].(http.HandlerFunc)
	return ret0
}

// HandleUpdate indicates an expected call of HandleUpdate.
func (mr *MockMatchlistHandlerMockRecorder) HandleUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdate", reflect.TypeOf((*MockMatchlistHandler)(nil).HandleUpdate))
}
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aimzeter/wuts/three (interfaces: CitizenStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	three "github.com/aimzeter/wuts/three"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCitizenStore is a mock of CitizenStore interface
type MockCitizenStore struct {
	ctrl     *gomock.Controller
	recorder *MockCitizenStoreMockRecorder
}

// MockCitizenStoreMockRecorder is the mock recorder for MockCitizenStore
type MockCitizenStoreMockRecorder struct {
	mock *MockCitizenStore
}

// NewMockCitizenStore creates a new mock instance
func NewMockCitizenStore(ctrl *gomock.Controller) *MockCitizenStore {
	mock := &MockCitizenStore{ctrl: ctrl}
	mock.recorder = &MockCitizenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCitizenStore) EXPECT() *MockCitizenStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockCitizenStore) Get(arg0 string) three.Citizen {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(three.Citizen)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockCitizenStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCitizenStore)(nil).Get), arg0)
}

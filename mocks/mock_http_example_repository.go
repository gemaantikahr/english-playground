// Code generated by MockGen. DO NOT EDIT.
// Source: app/repositories/http_example_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "english_playground/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpExampleInterface is a mock of HttpExampleInterface interface.
type MockHttpExampleInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHttpExampleInterfaceMockRecorder
}

// MockHttpExampleInterfaceMockRecorder is the mock recorder for MockHttpExampleInterface.
type MockHttpExampleInterfaceMockRecorder struct {
	mock *MockHttpExampleInterface
}

// NewMockHttpExampleInterface creates a new mock instance.
func NewMockHttpExampleInterface(ctrl *gomock.Controller) *MockHttpExampleInterface {
	mock := &MockHttpExampleInterface{ctrl: ctrl}
	mock.recorder = &MockHttpExampleInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpExampleInterface) EXPECT() *MockHttpExampleInterfaceMockRecorder {
	return m.recorder
}

// GetPost mocks base method.
func (m *MockHttpExampleInterface) GetPost() models.Posts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost")
	ret0, _ := ret[0].(models.Posts)
	return ret0
}

// GetPost indicates an expected call of GetPost.
func (mr *MockHttpExampleInterfaceMockRecorder) GetPost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockHttpExampleInterface)(nil).GetPost))
}

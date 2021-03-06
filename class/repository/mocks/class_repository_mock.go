// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/goweb4/class/repository (interfaces: ClassRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	class "github.com/goweb4/class"
	reflect "reflect"
)

// MockClassRepository is a mock of ClassRepository interface
type MockClassRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClassRepositoryMockRecorder
}

// MockClassRepositoryMockRecorder is the mock recorder for MockClassRepository
type MockClassRepositoryMockRecorder struct {
	mock *MockClassRepository
}

// NewMockClassRepository creates a new mock instance
func NewMockClassRepository(ctrl *gomock.Controller) *MockClassRepository {
	mock := &MockClassRepository{ctrl: ctrl}
	mock.recorder = &MockClassRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClassRepository) EXPECT() *MockClassRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClassRepository) Create(arg0 string, arg1 int) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockClassRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClassRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockClassRepository) Delete(arg0 int) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClassRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClassRepository)(nil).Delete), arg0)
}

// Update mocks base method
func (m *MockClassRepository) Update(arg0 class.Class) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockClassRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClassRepository)(nil).Update), arg0)
}

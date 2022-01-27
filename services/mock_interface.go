// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"
	models "user-curd/models"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeleteByUserId mocks base method.
func (m *MockService) DeleteByUserId(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUserId indicates an expected call of DeleteByUserId.
func (mr *MockServiceMockRecorder) DeleteByUserId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserId", reflect.TypeOf((*MockService)(nil).DeleteByUserId), id)
}

// InsertUserDetails mocks base method.
func (m *MockService) InsertUserDetails(arg0 models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserDetails", arg0)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserDetails indicates an expected call of InsertUserDetails.
func (mr *MockServiceMockRecorder) InsertUserDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserDetails", reflect.TypeOf((*MockService)(nil).InsertUserDetails), arg0)
}

// IsEmailValid mocks base method.
func (m *MockService) IsEmailValid(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmailValid", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmailValid indicates an expected call of IsEmailValid.
func (mr *MockServiceMockRecorder) IsEmailValid(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmailValid", reflect.TypeOf((*MockService)(nil).IsEmailValid), email)
}

// SearchAll mocks base method.
func (m *MockService) SearchAll() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAll")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAll indicates an expected call of SearchAll.
func (mr *MockServiceMockRecorder) SearchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAll", reflect.TypeOf((*MockService)(nil).SearchAll))
}

// SearchByUserId mocks base method.
func (m *MockService) SearchByUserId(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByUserId", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByUserId indicates an expected call of SearchByUserId.
func (mr *MockServiceMockRecorder) SearchByUserId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByUserId", reflect.TypeOf((*MockService)(nil).SearchByUserId), id)
}

// UpdateByUserId mocks base method.
func (m *MockService) UpdateByUserId(arg0 models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByUserId", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByUserId indicates an expected call of UpdateByUserId.
func (mr *MockServiceMockRecorder) UpdateByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByUserId", reflect.TypeOf((*MockService)(nil).UpdateByUserId), arg0)
}

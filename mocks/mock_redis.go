// Code generated by MockGen. DO NOT EDIT.
// Source: redis/redis_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GenerateUserId mocks base method
func (m *MockDataStore) GenerateUserId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUserId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GenerateUserId indicates an expected call of GenerateUserId
func (mr *MockDataStoreMockRecorder) GenerateUserId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUserId", reflect.TypeOf((*MockDataStore)(nil).GenerateUserId))
}

// GetRabbitQueueNames mocks base method
func (m *MockDataStore) GetRabbitQueueNames(clientIds []uint64) (map[string]bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRabbitQueueNames", clientIds)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetRabbitQueueNames indicates an expected call of GetRabbitQueueNames
func (mr *MockDataStoreMockRecorder) GetRabbitQueueNames(clientIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRabbitQueueNames", reflect.TypeOf((*MockDataStore)(nil).GetRabbitQueueNames), clientIds)
}

// CheckUserIn mocks base method
func (m *MockDataStore) CheckUserIn(clientId uint64, queueName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckUserIn", clientId, queueName)
}

// CheckUserIn indicates an expected call of CheckUserIn
func (mr *MockDataStoreMockRecorder) CheckUserIn(clientId, queueName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserIn", reflect.TypeOf((*MockDataStore)(nil).CheckUserIn), clientId, queueName)
}

// CheckUserOut mocks base method
func (m *MockDataStore) CheckUserOut(clientId uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CheckUserOut", clientId)
}

// CheckUserOut indicates an expected call of CheckUserOut
func (mr *MockDataStoreMockRecorder) CheckUserOut(clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserOut", reflect.TypeOf((*MockDataStore)(nil).CheckUserOut), clientId)
}

// GetAllConnectedUsers mocks base method
func (m *MockDataStore) GetAllConnectedUsers(caller uint64) []uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllConnectedUsers", caller)
	ret0, _ := ret[0].([]uint64)
	return ret0
}

// GetAllConnectedUsers indicates an expected call of GetAllConnectedUsers
func (mr *MockDataStoreMockRecorder) GetAllConnectedUsers(caller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllConnectedUsers", reflect.TypeOf((*MockDataStore)(nil).GetAllConnectedUsers), caller)
}
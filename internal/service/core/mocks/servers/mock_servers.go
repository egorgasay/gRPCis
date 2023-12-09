// Code generated by MockGen. DO NOT EDIT.
// Source: itisadb/internal/memory-balancer/core (interfaces: IServers)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	"itisadb/internal/service/servers"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIServers is a mock of IServers interface.
type MockIServers struct {
	ctrl     *gomock.Controller
	recorder *MockIServersMockRecorder
}

func (m *MockIServers) DelFromAll(ctx context.Context, key string) (atLeastOnce bool) {
	//TODO implement me
	panic("implement me")
}

// MockIServersMockRecorder is the mock recorder for MockIServers.
type MockIServersMockRecorder struct {
	mock *MockIServers
}

// NewMockIServers creates a new mock instance.
func NewMockIServers(ctrl *gomock.Controller) *MockIServers {
	mock := &MockIServers{ctrl: ctrl}
	mock.recorder = &MockIServersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServers) EXPECT() *MockIServersMockRecorder {
	return m.recorder
}

// AddServer mocks base method.
func (m *MockIServers) AddServer(arg0 string, arg1, arg2 uint64, arg3 int32) (int32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddServer indicates an expected call of AddServer.
func (mr *MockIServersMockRecorder) AddServer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServer", reflect.TypeOf((*MockIServers)(nil).AddServer), arg0, arg1, arg2, arg3)
}

// DeepSearch mocks base method.
func (m *MockIServers) DeepSearch(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepSearch", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeepSearch indicates an expected call of DeepSearch.
func (mr *MockIServersMockRecorder) DeepSearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepSearch", reflect.TypeOf((*MockIServers)(nil).DeepSearch), arg0, arg1)
}

// Disconnect mocks base method.
func (m *MockIServers) Disconnect(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect", arg0)
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockIServersMockRecorder) Disconnect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockIServers)(nil).Disconnect), arg0)
}

// Exists mocks base method.
func (m *MockIServers) Exists(arg0 int32) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockIServersMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIServers)(nil).Exists), arg0)
}

// GetServer mocks base method.
func (m *MockIServers) GetServer() (*servers.RemoteServer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServer")
	ret0, _ := ret[0].(*servers.RemoteServer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetServer indicates an expected call of GetServer.
func (mr *MockIServersMockRecorder) GetServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServer", reflect.TypeOf((*MockIServers)(nil).GetServer))
}

// GetServerByID mocks base method.
func (m *MockIServers) GetServerByID(arg0 int32) (*servers.RemoteServer, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerByID", arg0)
	ret0, _ := ret[0].(*servers.RemoteServer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetServerByID indicates an expected call of GetServerByID.
func (mr *MockIServersMockRecorder) GetServerByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByID", reflect.TypeOf((*MockIServers)(nil).GetServerByID), arg0)
}

// GetServers mocks base method.
func (m *MockIServers) GetServers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetServers indicates an expected call of GetServers.
func (mr *MockIServersMockRecorder) GetServers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServers", reflect.TypeOf((*MockIServers)(nil).GetServers))
}

// Len mocks base method.
func (m *MockIServers) Len() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int32)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockIServersMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockIServers)(nil).Len))
}

// SetToAll mocks base method.
func (m *MockIServers) SetToAll(arg0 context.Context, arg1, arg2 string, arg3 bool) []int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetToAll", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]int32)
	return ret0
}

// SetToAll indicates an expected call of SetToAll.
func (mr *MockIServersMockRecorder) SetToAll(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToAll", reflect.TypeOf((*MockIServers)(nil).SetToAll), arg0, arg1, arg2, arg3)
}

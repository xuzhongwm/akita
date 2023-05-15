// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/akita/v3/tracing (interfaces: NamedHookable)

package tracing

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sim "gitlab.com/akita/akita/v3/sim"
)

// MockNamedHookable is a mock of NamedHookable interface.
type MockNamedHookable struct {
	ctrl     *gomock.Controller
	recorder *MockNamedHookableMockRecorder
}

// MockNamedHookableMockRecorder is the mock recorder for MockNamedHookable.
type MockNamedHookableMockRecorder struct {
	mock *MockNamedHookable
}

// NewMockNamedHookable creates a new mock instance.
func NewMockNamedHookable(ctrl *gomock.Controller) *MockNamedHookable {
	mock := &MockNamedHookable{ctrl: ctrl}
	mock.recorder = &MockNamedHookableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedHookable) EXPECT() *MockNamedHookableMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockNamedHookable) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockNamedHookableMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockNamedHookable)(nil).AcceptHook), arg0)
}

// Hooks mocks base method.
func (m *MockNamedHookable) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockNamedHookableMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockNamedHookable)(nil).Hooks))
}

// InvokeHook mocks base method.
func (m *MockNamedHookable) InvokeHook(arg0 sim.HookCtx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InvokeHook", arg0)
}

// InvokeHook indicates an expected call of InvokeHook.
func (mr *MockNamedHookableMockRecorder) InvokeHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeHook", reflect.TypeOf((*MockNamedHookable)(nil).InvokeHook), arg0)
}

// Name mocks base method.
func (m *MockNamedHookable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockNamedHookableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockNamedHookable)(nil).Name))
}

// NumHooks mocks base method.
func (m *MockNamedHookable) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockNamedHookableMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockNamedHookable)(nil).NumHooks))
}

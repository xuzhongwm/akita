// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sarchlab/akita/v4/sim (interfaces: Port,Engine,BufferedSender)

package tlb

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sim "github.com/sarchlab/akita/v4/sim"
)

// MockPort is a mock of Port interface.
type MockPort struct {
	ctrl     *gomock.Controller
	recorder *MockPortMockRecorder
}

// MockPortMockRecorder is the mock recorder for MockPort.
type MockPortMockRecorder struct {
	mock *MockPort
}

// NewMockPort creates a new mock instance.
func NewMockPort(ctrl *gomock.Controller) *MockPort {
	mock := &MockPort{ctrl: ctrl}
	mock.recorder = &MockPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPort) EXPECT() *MockPortMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockPort) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockPortMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockPort)(nil).AcceptHook), arg0)
}

// CanSend mocks base method.
func (m *MockPort) CanSend() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockPortMockRecorder) CanSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockPort)(nil).CanSend))
}

// Component mocks base method.
func (m *MockPort) Component() sim.Component {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Component")
	ret0, _ := ret[0].(sim.Component)
	return ret0
}

// Component indicates an expected call of Component.
func (mr *MockPortMockRecorder) Component() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Component", reflect.TypeOf((*MockPort)(nil).Component))
}

// Deliver mocks base method.
func (m *MockPort) Deliver(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deliver", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Deliver indicates an expected call of Deliver.
func (mr *MockPortMockRecorder) Deliver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deliver", reflect.TypeOf((*MockPort)(nil).Deliver), arg0)
}

// Hooks mocks base method.
func (m *MockPort) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockPortMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockPort)(nil).Hooks))
}

// Name mocks base method.
func (m *MockPort) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPortMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPort)(nil).Name))
}

// NotifyAvailable mocks base method.
func (m *MockPort) NotifyAvailable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyAvailable")
}

// NotifyAvailable indicates an expected call of NotifyAvailable.
func (mr *MockPortMockRecorder) NotifyAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAvailable", reflect.TypeOf((*MockPort)(nil).NotifyAvailable))
}

// NumHooks mocks base method.
func (m *MockPort) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockPortMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockPort)(nil).NumHooks))
}

// PeekIncoming mocks base method.
func (m *MockPort) PeekIncoming() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekIncoming")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// PeekIncoming indicates an expected call of PeekIncoming.
func (mr *MockPortMockRecorder) PeekIncoming() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekIncoming", reflect.TypeOf((*MockPort)(nil).PeekIncoming))
}

// PeekOutgoing mocks base method.
func (m *MockPort) PeekOutgoing() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekOutgoing")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// PeekOutgoing indicates an expected call of PeekOutgoing.
func (mr *MockPortMockRecorder) PeekOutgoing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekOutgoing", reflect.TypeOf((*MockPort)(nil).PeekOutgoing))
}

// RetrieveIncoming mocks base method.
func (m *MockPort) RetrieveIncoming() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveIncoming")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// RetrieveIncoming indicates an expected call of RetrieveIncoming.
func (mr *MockPortMockRecorder) RetrieveIncoming() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveIncoming", reflect.TypeOf((*MockPort)(nil).RetrieveIncoming))
}

// RetrieveOutgoing mocks base method.
func (m *MockPort) RetrieveOutgoing() sim.Msg {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveOutgoing")
	ret0, _ := ret[0].(sim.Msg)
	return ret0
}

// RetrieveOutgoing indicates an expected call of RetrieveOutgoing.
func (mr *MockPortMockRecorder) RetrieveOutgoing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveOutgoing", reflect.TypeOf((*MockPort)(nil).RetrieveOutgoing))
}

// Send mocks base method.
func (m *MockPort) Send(arg0 sim.Msg) *sim.SendError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(*sim.SendError)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockPortMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockPort)(nil).Send), arg0)
}

// SetConnection mocks base method.
func (m *MockPort) SetConnection(arg0 sim.Connection) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConnection", arg0)
}

// SetConnection indicates an expected call of SetConnection.
func (mr *MockPortMockRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConnection", reflect.TypeOf((*MockPort)(nil).SetConnection), arg0)
}

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// AcceptHook mocks base method.
func (m *MockEngine) AcceptHook(arg0 sim.Hook) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptHook", arg0)
}

// AcceptHook indicates an expected call of AcceptHook.
func (mr *MockEngineMockRecorder) AcceptHook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptHook", reflect.TypeOf((*MockEngine)(nil).AcceptHook), arg0)
}

// Continue mocks base method.
func (m *MockEngine) Continue() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Continue")
}

// Continue indicates an expected call of Continue.
func (mr *MockEngineMockRecorder) Continue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Continue", reflect.TypeOf((*MockEngine)(nil).Continue))
}

// CurrentTime mocks base method.
func (m *MockEngine) CurrentTime() sim.VTimeInSec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentTime")
	ret0, _ := ret[0].(sim.VTimeInSec)
	return ret0
}

// CurrentTime indicates an expected call of CurrentTime.
func (mr *MockEngineMockRecorder) CurrentTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentTime", reflect.TypeOf((*MockEngine)(nil).CurrentTime))
}

// Hooks mocks base method.
func (m *MockEngine) Hooks() []sim.Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]sim.Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockEngineMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockEngine)(nil).Hooks))
}

// NumHooks mocks base method.
func (m *MockEngine) NumHooks() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumHooks")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumHooks indicates an expected call of NumHooks.
func (mr *MockEngineMockRecorder) NumHooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumHooks", reflect.TypeOf((*MockEngine)(nil).NumHooks))
}

// Pause mocks base method.
func (m *MockEngine) Pause() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Pause")
}

// Pause indicates an expected call of Pause.
func (mr *MockEngineMockRecorder) Pause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockEngine)(nil).Pause))
}

// Run mocks base method.
func (m *MockEngine) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockEngineMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEngine)(nil).Run))
}

// Schedule mocks base method.
func (m *MockEngine) Schedule(arg0 sim.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Schedule", arg0)
}

// Schedule indicates an expected call of Schedule.
func (mr *MockEngineMockRecorder) Schedule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schedule", reflect.TypeOf((*MockEngine)(nil).Schedule), arg0)
}

// MockBufferedSender is a mock of BufferedSender interface.
type MockBufferedSender struct {
	ctrl     *gomock.Controller
	recorder *MockBufferedSenderMockRecorder
}

// MockBufferedSenderMockRecorder is the mock recorder for MockBufferedSender.
type MockBufferedSenderMockRecorder struct {
	mock *MockBufferedSender
}

// NewMockBufferedSender creates a new mock instance.
func NewMockBufferedSender(ctrl *gomock.Controller) *MockBufferedSender {
	mock := &MockBufferedSender{ctrl: ctrl}
	mock.recorder = &MockBufferedSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBufferedSender) EXPECT() *MockBufferedSenderMockRecorder {
	return m.recorder
}

// CanSend mocks base method.
func (m *MockBufferedSender) CanSend(arg0 int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSend", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSend indicates an expected call of CanSend.
func (mr *MockBufferedSenderMockRecorder) CanSend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSend", reflect.TypeOf((*MockBufferedSender)(nil).CanSend), arg0)
}

// Clear mocks base method.
func (m *MockBufferedSender) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockBufferedSenderMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockBufferedSender)(nil).Clear))
}

// Send mocks base method.
func (m *MockBufferedSender) Send(arg0 sim.Msg) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send.
func (mr *MockBufferedSenderMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBufferedSender)(nil).Send), arg0)
}

// Tick mocks base method.
func (m *MockBufferedSender) Tick() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Tick indicates an expected call of Tick.
func (mr *MockBufferedSenderMockRecorder) Tick() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockBufferedSender)(nil).Tick))
}

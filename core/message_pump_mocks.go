// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/message_pump.go

// Package core is a generated GoMock package.
package core

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransporter is a mock of Transporter interface
type MockTransporter struct {
	ctrl     *gomock.Controller
	recorder *MockTransporterMockRecorder
}

// MockTransporterMockRecorder is the mock recorder for MockTransporter
type MockTransporterMockRecorder struct {
	mock *MockTransporter
}

// NewMockTransporter creates a new mock instance
func NewMockTransporter(ctrl *gomock.Controller) *MockTransporter {
	mock := &MockTransporter{ctrl: ctrl}
	mock.recorder = &MockTransporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransporter) EXPECT() *MockTransporterMockRecorder {
	return m.recorder
}

// Messages mocks base method
func (m *MockTransporter) Messages() <-chan []*Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Messages")
	ret0, _ := ret[0].(<-chan []*Message)
	return ret0
}

// Messages indicates an expected call of Messages
func (mr *MockTransporterMockRecorder) Messages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockTransporter)(nil).Messages))
}

// Delete mocks base method
func (m *MockTransporter) Delete(arg0 *Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTransporterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTransporter)(nil).Delete), arg0)
}

// Poison mocks base method
func (m *MockTransporter) Poison(arg0 *Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poison", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Poison indicates an expected call of Poison
func (mr *MockTransporterMockRecorder) Poison(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poison", reflect.TypeOf((*MockTransporter)(nil).Poison), arg0)
}

// MockDispatcher is a mock of Dispatcher interface
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockDispatcher) Dispatch(arg0 *Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockDispatcherMockRecorder) Dispatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockDispatcher)(nil).Dispatch), arg0)
}

// MockBrokerBinder is a mock of BrokerBinder interface
type MockBrokerBinder struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerBinderMockRecorder
}

// MockBrokerBinderMockRecorder is the mock recorder for MockBrokerBinder
type MockBrokerBinderMockRecorder struct {
	mock *MockBrokerBinder
}

// NewMockBrokerBinder creates a new mock instance
func NewMockBrokerBinder(ctrl *gomock.Controller) *MockBrokerBinder {
	mock := &MockBrokerBinder{ctrl: ctrl}
	mock.recorder = &MockBrokerBinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrokerBinder) EXPECT() *MockBrokerBinderMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockBrokerBinder) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockBrokerBinderMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBrokerBinder)(nil).Start), arg0)
}

// Awaiter mocks base method
func (m *MockBrokerBinder) Awaiter() *Awaiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Awaiter")
	ret0, _ := ret[0].(*Awaiter)
	return ret0
}

// Awaiter indicates an expected call of Awaiter
func (mr *MockBrokerBinderMockRecorder) Awaiter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Awaiter", reflect.TypeOf((*MockBrokerBinder)(nil).Awaiter))
}

// MockMessagePumpRunner is a mock of MessagePumpRunner interface
type MockMessagePumpRunner struct {
	ctrl     *gomock.Controller
	recorder *MockMessagePumpRunnerMockRecorder
}

// MockMessagePumpRunnerMockRecorder is the mock recorder for MockMessagePumpRunner
type MockMessagePumpRunnerMockRecorder struct {
	mock *MockMessagePumpRunner
}

// NewMockMessagePumpRunner creates a new mock instance
func NewMockMessagePumpRunner(ctrl *gomock.Controller) *MockMessagePumpRunner {
	mock := &MockMessagePumpRunner{ctrl: ctrl}
	mock.recorder = &MockMessagePumpRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessagePumpRunner) EXPECT() *MockMessagePumpRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockMessagePumpRunner) Run(arg0 context.Context, arg1 Transporter, arg2 Config) *Awaiter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Awaiter)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockMessagePumpRunnerMockRecorder) Run(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMessagePumpRunner)(nil).Run), arg0, arg1, arg2)
}

// MockBrokerBinderBuilder is a mock of BrokerBinderBuilder interface
type MockBrokerBinderBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerBinderBuilderMockRecorder
}

// MockBrokerBinderBuilderMockRecorder is the mock recorder for MockBrokerBinderBuilder
type MockBrokerBinderBuilderMockRecorder struct {
	mock *MockBrokerBinderBuilder
}

// NewMockBrokerBinderBuilder creates a new mock instance
func NewMockBrokerBinderBuilder(ctrl *gomock.Controller) *MockBrokerBinderBuilder {
	mock := &MockBrokerBinderBuilder{ctrl: ctrl}
	mock.recorder = &MockBrokerBinderBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrokerBinderBuilder) EXPECT() *MockBrokerBinderBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockBrokerBinderBuilder) Build(arg0 MessagePumpRunner, arg1 *Config, arg2 interface{}) (BrokerBinder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2)
	ret0, _ := ret[0].(BrokerBinder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockBrokerBinderBuilderMockRecorder) Build(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBrokerBinderBuilder)(nil).Build), arg0, arg1, arg2)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/http/metric (interfaces: EventRecorder)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	metric "github.com/influxdata/influxdb/v2/http/metric"
	reflect "reflect"
)

// MockEventRecorder is a mock of EventRecorder interface
type MockEventRecorder struct {
	ctrl     *gomock.Controller
	recorder *MockEventRecorderMockRecorder
}

// MockEventRecorderMockRecorder is the mock recorder for MockEventRecorder
type MockEventRecorderMockRecorder struct {
	mock *MockEventRecorder
}

// NewMockEventRecorder creates a new mock instance
func NewMockEventRecorder(ctrl *gomock.Controller) *MockEventRecorder {
	mock := &MockEventRecorder{ctrl: ctrl}
	mock.recorder = &MockEventRecorderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventRecorder) EXPECT() *MockEventRecorderMockRecorder {
	return m.recorder
}

// Record mocks base method
func (m *MockEventRecorder) Record(arg0 context.Context, arg1 metric.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Record", arg0, arg1)
}

// Record indicates an expected call of Record
func (mr *MockEventRecorderMockRecorder) Record(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockEventRecorder)(nil).Record), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: MetricsDatabaseCursor)

// Package mock_moira_alert is a generated GoMock package.
package mock_moira_alert

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetricsDatabaseCursor is a mock of MetricsDatabaseCursor interface.
type MockMetricsDatabaseCursor struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsDatabaseCursorMockRecorder
}

// MockMetricsDatabaseCursorMockRecorder is the mock recorder for MockMetricsDatabaseCursor.
type MockMetricsDatabaseCursorMockRecorder struct {
	mock *MockMetricsDatabaseCursor
}

// NewMockMetricsDatabaseCursor creates a new mock instance.
func NewMockMetricsDatabaseCursor(ctrl *gomock.Controller) *MockMetricsDatabaseCursor {
	mock := &MockMetricsDatabaseCursor{ctrl: ctrl}
	mock.recorder = &MockMetricsDatabaseCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsDatabaseCursor) EXPECT() *MockMetricsDatabaseCursorMockRecorder {
	return m.recorder
}

// Free mocks base method.
func (m *MockMetricsDatabaseCursor) Free() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Free")
	ret0, _ := ret[0].(error)
	return ret0
}

// Free indicates an expected call of Free.
func (mr *MockMetricsDatabaseCursorMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockMetricsDatabaseCursor)(nil).Free))
}

// Next mocks base method.
func (m *MockMetricsDatabaseCursor) Next() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockMetricsDatabaseCursorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockMetricsDatabaseCursor)(nil).Next))
}
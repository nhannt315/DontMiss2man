// Code generated by MockGen. DO NOT EDIT.
// Source: clock.go

// Package mock_datetime is a generated GoMock package.
package mock_datetime

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	datetime "github.com/nhannt315/real_estate_api/pkg/datetime"
)

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// NewTime mocks base method.
func (m *MockClock) NewTime(year int, month time.Month, day, hour, min, sec, nsec int) *datetime.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTime", year, month, day, hour, min, sec, nsec)
	ret0, _ := ret[0].(*datetime.Time)
	return ret0
}

// NewTime indicates an expected call of NewTime.
func (mr *MockClockMockRecorder) NewTime(year, month, day, hour, min, sec, nsec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTime", reflect.TypeOf((*MockClock)(nil).NewTime), year, month, day, hour, min, sec, nsec)
}

// Now mocks base method.
func (m *MockClock) Now() *datetime.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(*datetime.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
}

// ParseTime mocks base method.
func (m *MockClock) ParseTime(layout, value string) (*datetime.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTime", layout, value)
	ret0, _ := ret[0].(*datetime.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTime indicates an expected call of ParseTime.
func (mr *MockClockMockRecorder) ParseTime(layout, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTime", reflect.TypeOf((*MockClock)(nil).ParseTime), layout, value)
}

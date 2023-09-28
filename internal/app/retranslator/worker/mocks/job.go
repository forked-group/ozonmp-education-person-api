// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/app/retranslator/worker (interfaces: Job)

// Package mock_worker is a generated GoMock package.
package mock_worker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJob is a mock of Job interface.
type MockJob struct {
	ctrl     *gomock.Controller
	recorder *MockJobMockRecorder
}

// MockJobMockRecorder is the mock recorder for MockJob.
type MockJobMockRecorder struct {
	mock *MockJob
}

// NewMockJob creates a new mock instance.
func NewMockJob(ctrl *gomock.Controller) *MockJob {
	mock := &MockJob{ctrl: ctrl}
	mock.recorder = &MockJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJob) EXPECT() *MockJobMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockJob) Do() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do")
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockJobMockRecorder) Do() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockJob)(nil).Do))
}

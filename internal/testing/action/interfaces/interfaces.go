// Code generated by MockGen. DO NOT EDIT.
// Source: internal/automation/action/interfaces/interfaces.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"
	model "tcms/m/internal/connections/db/model"

	interfaces "github.com/BlenderistDev/automation/interfaces"
	gomock "github.com/golang/mock/gomock"
)

// MockActionWithModel is a mock of ActionWithModel interface.
type MockActionWithModel struct {
	ctrl     *gomock.Controller
	recorder *MockActionWithModelMockRecorder
}

// MockActionWithModelMockRecorder is the mock recorder for MockActionWithModel.
type MockActionWithModelMockRecorder struct {
	mock *MockActionWithModel
}

// NewMockActionWithModel creates a new mock instance.
func NewMockActionWithModel(ctrl *gomock.Controller) *MockActionWithModel {
	mock := &MockActionWithModel{ctrl: ctrl}
	mock.recorder = &MockActionWithModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionWithModel) EXPECT() *MockActionWithModelMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockActionWithModel) Execute(action model.Action, trigger interfaces.Trigger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", action, trigger)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockActionWithModelMockRecorder) Execute(action, trigger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockActionWithModel)(nil).Execute), action, trigger)
}
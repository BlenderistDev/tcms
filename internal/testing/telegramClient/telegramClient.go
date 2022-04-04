// Code generated by MockGen. DO NOT EDIT.
// Source: internal/telegramClient/telegramClient.go

// Package mock_telegramClient is a generated GoMock package.
package mock_telegramClient

import (
	reflect "reflect"
	telegram "tcms/pkg/telegram"

	gomock "github.com/golang/mock/gomock"
)

// MockTelegramClient is a mock of TelegramClient interface.
type MockTelegramClient struct {
	ctrl     *gomock.Controller
	recorder *MockTelegramClientMockRecorder
}

// MockTelegramClientMockRecorder is the mock recorder for MockTelegramClient.
type MockTelegramClientMockRecorder struct {
	mock *MockTelegramClient
}

// NewMockTelegramClient creates a new mock instance.
func NewMockTelegramClient(ctrl *gomock.Controller) *MockTelegramClient {
	mock := &MockTelegramClient{ctrl: ctrl}
	mock.recorder = &MockTelegramClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelegramClient) EXPECT() *MockTelegramClientMockRecorder {
	return m.recorder
}

// AuthSignIn mocks base method.
func (m *MockTelegramClient) AuthSignIn(code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthSignIn", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthSignIn indicates an expected call of AuthSignIn.
func (mr *MockTelegramClientMockRecorder) AuthSignIn(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthSignIn", reflect.TypeOf((*MockTelegramClient)(nil).AuthSignIn), code)
}

// Authorization mocks base method.
func (m *MockTelegramClient) Authorization(phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorization", phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authorization indicates an expected call of Authorization.
func (mr *MockTelegramClientMockRecorder) Authorization(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorization", reflect.TypeOf((*MockTelegramClient)(nil).Authorization), phone)
}

// Dialogs mocks base method.
func (m *MockTelegramClient) Dialogs() (*telegram.DialogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dialogs")
	ret0, _ := ret[0].(*telegram.DialogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dialogs indicates an expected call of Dialogs.
func (mr *MockTelegramClientMockRecorder) Dialogs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dialogs", reflect.TypeOf((*MockTelegramClient)(nil).Dialogs))
}

// GetCurrentUser mocks base method.
func (m *MockTelegramClient) GetCurrentUser() (*telegram.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser")
	ret0, _ := ret[0].(*telegram.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser.
func (mr *MockTelegramClientMockRecorder) GetCurrentUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockTelegramClient)(nil).GetCurrentUser))
}

// MuteChat mocks base method.
func (m *MockTelegramClient) MuteChat(id string, unMute bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MuteChat", id, unMute)
	ret0, _ := ret[0].(error)
	return ret0
}

// MuteChat indicates an expected call of MuteChat.
func (mr *MockTelegramClientMockRecorder) MuteChat(id, unMute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteChat", reflect.TypeOf((*MockTelegramClient)(nil).MuteChat), id, unMute)
}

// MuteUser mocks base method.
func (m *MockTelegramClient) MuteUser(id, accessHash string, unMute bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MuteUser", id, accessHash, unMute)
	ret0, _ := ret[0].(error)
	return ret0
}

// MuteUser indicates an expected call of MuteUser.
func (mr *MockTelegramClientMockRecorder) MuteUser(id, accessHash, unMute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MuteUser", reflect.TypeOf((*MockTelegramClient)(nil).MuteUser), id, accessHash, unMute)
}

// SendMessage mocks base method.
func (m *MockTelegramClient) SendMessage(peer, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", peer, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockTelegramClientMockRecorder) SendMessage(peer, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockTelegramClient)(nil).SendMessage), peer, message)
}

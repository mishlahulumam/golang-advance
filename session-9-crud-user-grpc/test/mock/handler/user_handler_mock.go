package mock_handler

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

type MockIUserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIUserHandlerMockRecorder
}

type MockIUserHandlerMockRecorder struct {
	mock *MockIUserHandler
}

func NewMockIUserHandler(ctrl *gomock.Controller) *MockIUserHandler {
	mock := &MockIUserHandler{ctrl: ctrl}
	mock.recorder = &MockIUserHandlerMockRecorder{mock}
	return mock
}

func (m *MockIUserHandler) EXPECT() *MockIUserHandlerMockRecorder {
	return m.recorder
}

func (m *MockIUserHandler) CreateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", c)
}

func (mr *MockIUserHandlerMockRecorder) CreateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserHandler)(nil).CreateUser), c)
}

func (m *MockIUserHandler) DeleteUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUser", c)
}

func (mr *MockIUserHandlerMockRecorder) DeleteUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserHandler)(nil).DeleteUser), c)
}

func (m *MockIUserHandler) GetAllUsers(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAllUsers", c)
}

func (mr *MockIUserHandlerMockRecorder) GetAllUsers(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserHandler)(nil).GetAllUsers), c)
}

func (m *MockIUserHandler) GetUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUser", c)
}

func (mr *MockIUserHandlerMockRecorder) GetUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserHandler)(nil).GetUser), c)
}

func (m *MockIUserHandler) UpdateUser(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateUser", c)
}

func (mr *MockIUserHandlerMockRecorder) UpdateUser(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserHandler)(nil).UpdateUser), c)
}

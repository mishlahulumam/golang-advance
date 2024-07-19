package mock_postgres_gorm

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

type MockGormDBIface struct {
	ctrl     *gomock.Controller
	recorder *MockGormDBIfaceMockRecorder
}

type MockGormDBIfaceMockRecorder struct {
	mock *MockGormDBIface
}

func NewMockGormDBIface(ctrl *gomock.Controller) *MockGormDBIface {
	mock := &MockGormDBIface{ctrl: ctrl}
	mock.recorder = &MockGormDBIfaceMockRecorder{mock}
	return mock
}

func (m *MockGormDBIface) EXPECT() *MockGormDBIfaceMockRecorder {
	return m.recorder
}

func (m *MockGormDBIface) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGormDBIface)(nil).Create), value)
}

func (m *MockGormDBIface) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{value}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) Delete(value interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGormDBIface)(nil).Delete), varargs...)
}

func (m *MockGormDBIface) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) Find(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGormDBIface)(nil).Find), varargs...)
}

func (m *MockGormDBIface) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockGormDBIface)(nil).First), varargs...)
}

func (m *MockGormDBIface) Save(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGormDBIface)(nil).Save), value)
}

func (m *MockGormDBIface) WithContext(ctx context.Context) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

func (mr *MockGormDBIfaceMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockGormDBIface)(nil).WithContext), ctx)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/harryosmar/hash-go (interfaces: JwtSign)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	jwt "github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
)

// MockJwtSign is a mock of JwtSign interface.
type MockJwtSign struct {
	ctrl     *gomock.Controller
	recorder *MockJwtSignMockRecorder
}

// MockJwtSignMockRecorder is the mock recorder for MockJwtSign.
type MockJwtSignMockRecorder struct {
	mock *MockJwtSign
}

// NewMockJwtSign creates a new mock instance.
func NewMockJwtSign(ctrl *gomock.Controller) *MockJwtSign {
	mock := &MockJwtSign{ctrl: ctrl}
	mock.recorder = &MockJwtSignMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtSign) EXPECT() *MockJwtSignMockRecorder {
	return m.recorder
}

// Sign mocks base method.
func (m *MockJwtSign) Sign(arg0 context.Context, arg1 *jwt.MapClaims) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockJwtSignMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockJwtSign)(nil).Sign), arg0, arg1)
}

// Validate mocks base method.
func (m *MockJwtSign) Validate(arg0 context.Context, arg1 string) (*jwt.MapClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(*jwt.MapClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockJwtSignMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockJwtSign)(nil).Validate), arg0, arg1)
}

// ValidateReturnBytes mocks base method.
func (m *MockJwtSign) ValidateReturnBytes(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateReturnBytes", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateReturnBytes indicates an expected call of ValidateReturnBytes.
func (mr *MockJwtSignMockRecorder) ValidateReturnBytes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateReturnBytes", reflect.TypeOf((*MockJwtSign)(nil).ValidateReturnBytes), arg0, arg1)
}

// ValidateReturnReader mocks base method.
func (m *MockJwtSign) ValidateReturnReader(arg0 context.Context, arg1 string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateReturnReader", arg0, arg1)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateReturnReader indicates an expected call of ValidateReturnReader.
func (mr *MockJwtSignMockRecorder) ValidateReturnReader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateReturnReader", reflect.TypeOf((*MockJwtSign)(nil).ValidateReturnReader), arg0, arg1)
}

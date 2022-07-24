// Code generated by MockGen. DO NOT EDIT.
// Source: ./iParser/parser.go

// Package mock_iParser is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser.
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// IsValid mocks base method.
func (m *MockParser) IsValid(expression string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid", expression)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockParserMockRecorder) IsValid(expression interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockParser)(nil).IsValid), expression)
}

// MaxAllowedValue mocks base method.
func (m *MockParser) MaxAllowedValue() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxAllowedValue")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxAllowedValue indicates an expected call of MaxAllowedValue.
func (mr *MockParserMockRecorder) MaxAllowedValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxAllowedValue", reflect.TypeOf((*MockParser)(nil).MaxAllowedValue))
}

// MinAllowedValue mocks base method.
func (m *MockParser) MinAllowedValue() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinAllowedValue")
	ret0, _ := ret[0].(int)
	return ret0
}

// MinAllowedValue indicates an expected call of MinAllowedValue.
func (mr *MockParserMockRecorder) MinAllowedValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinAllowedValue", reflect.TypeOf((*MockParser)(nil).MinAllowedValue))
}

// Parse mocks base method.
func (m *MockParser) Parse(expression string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", expression)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserMockRecorder) Parse(expression interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), expression)
}
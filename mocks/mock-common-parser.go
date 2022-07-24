// Code generated by MockGen. DO NOT EDIT.
// Source: ./common/parser/common-parser.go

// Package mock_parser is a generated GoMock package.
package mocks

import (
	"cron-expression-parser/parser-interface"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommonParser is a mock of CommonParser interface.
type MockCommonParser struct {
	ctrl     *gomock.Controller
	recorder *MockCommonParserMockRecorder
}

// MockCommonParserMockRecorder is the mock recorder for MockCommonParser.
type MockCommonParserMockRecorder struct {
	mock *MockCommonParser
}

// NewMockCommonParser creates a new mock instance.
func NewMockCommonParser(ctrl *gomock.Controller) *MockCommonParser {
	mock := &MockCommonParser{ctrl: ctrl}
	mock.recorder = &MockCommonParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonParser) EXPECT() *MockCommonParserMockRecorder {
	return m.recorder
}

// Parse mocks base method.
func (m *MockCommonParser) Parse(expression string, parser parser_interface.Parser) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", expression, parser)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockCommonParserMockRecorder) Parse(expression, parser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockCommonParser)(nil).Parse), expression, parser)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase.go
//
// Generated by this command:
//
//	mockgen -source ./usecase.go -destination ../../internal/mocks/usecase/usecase_mock.go -package mocks_usecase
//

// Package mocks_usecase is a generated GoMock package.
package mocks_usecase

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// BuildResponse mocks base method.
func (m *MockUsecase) BuildResponse(message, sender string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildResponse", message, sender)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildResponse indicates an expected call of BuildResponse.
func (mr *MockUsecaseMockRecorder) BuildResponse(message, sender any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildResponse", reflect.TypeOf((*MockUsecase)(nil).BuildResponse), message, sender)
}

// IsValid mocks base method.
func (m *MockUsecase) IsValid(message string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid", message)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockUsecaseMockRecorder) IsValid(message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockUsecase)(nil).IsValid), message)
}

// MockUCBuilder is a mock of UCBuilder interface.
type MockUCBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockUCBuilderMockRecorder
}

// MockUCBuilderMockRecorder is the mock recorder for MockUCBuilder.
type MockUCBuilderMockRecorder struct {
	mock *MockUCBuilder
}

// NewMockUCBuilder creates a new mock instance.
func NewMockUCBuilder(ctrl *gomock.Controller) *MockUCBuilder {
	mock := &MockUCBuilder{ctrl: ctrl}
	mock.recorder = &MockUCBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUCBuilder) EXPECT() *MockUCBuilderMockRecorder {
	return m.recorder
}

// BuildResponseContext mocks base method.
func (m *MockUCBuilder) BuildResponseContext(message, sender string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildResponseContext", message, sender)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildResponseContext indicates an expected call of BuildResponseContext.
func (mr *MockUCBuilderMockRecorder) BuildResponseContext(message, sender any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildResponseContext", reflect.TypeOf((*MockUCBuilder)(nil).BuildResponseContext), message, sender)
}

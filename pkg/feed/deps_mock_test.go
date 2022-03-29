// Code generated by MockGen. DO NOT EDIT.
// Source: deps.go

// Package feed is a generated GoMock package.
package feed

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mxpv/podsync/pkg/model"
)

// MockfeedProvider is a mock of feedProvider interface.
type MockfeedProvider struct {
	ctrl     *gomock.Controller
	recorder *MockfeedProviderMockRecorder
}

// MockfeedProviderMockRecorder is the mock recorder for MockfeedProvider.
type MockfeedProviderMockRecorder struct {
	mock *MockfeedProvider
}

// NewMockfeedProvider creates a new mock instance.
func NewMockfeedProvider(ctrl *gomock.Controller) *MockfeedProvider {
	mock := &MockfeedProvider{ctrl: ctrl}
	mock.recorder = &MockfeedProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfeedProvider) EXPECT() *MockfeedProviderMockRecorder {
	return m.recorder
}

// GetFeed mocks base method.
func (m *MockfeedProvider) GetFeed(ctx context.Context, feedID string) (*model.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeed", ctx, feedID)
	ret0, _ := ret[0].(*model.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed.
func (mr *MockfeedProviderMockRecorder) GetFeed(ctx, feedID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockfeedProvider)(nil).GetFeed), ctx, feedID)
}
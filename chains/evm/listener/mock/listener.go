// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/listener/event-handler.go

// Package mock_listener is a generated GoMock package.
package mock_listener

import (
	context "context"
	big "math/big"
	reflect "reflect"

	events "github.com/ChainSafe/sygma-core/chains/evm/calls/events"
	message "github.com/ChainSafe/sygma-core/relayer/message"
	types "github.com/ChainSafe/sygma-core/types"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockEventListener is a mock of EventListener interface.
type MockEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockEventListenerMockRecorder
}

// MockEventListenerMockRecorder is the mock recorder for MockEventListener.
type MockEventListenerMockRecorder struct {
	mock *MockEventListener
}

// NewMockEventListener creates a new mock instance.
func NewMockEventListener(ctrl *gomock.Controller) *MockEventListener {
	mock := &MockEventListener{ctrl: ctrl}
	mock.recorder = &MockEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventListener) EXPECT() *MockEventListenerMockRecorder {
	return m.recorder
}

// FetchDeposits mocks base method.
func (m *MockEventListener) FetchDeposits(ctx context.Context, address common.Address, startBlock, endBlock *big.Int) ([]*events.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDeposits", ctx, address, startBlock, endBlock)
	ret0, _ := ret[0].([]*events.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDeposits indicates an expected call of FetchDeposits.
func (mr *MockEventListenerMockRecorder) FetchDeposits(ctx, address, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDeposits", reflect.TypeOf((*MockEventListener)(nil).FetchDeposits), ctx, address, startBlock, endBlock)
}

// MockDepositHandler is a mock of DepositHandler interface.
type MockDepositHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDepositHandlerMockRecorder
}

// MockDepositHandlerMockRecorder is the mock recorder for MockDepositHandler.
type MockDepositHandlerMockRecorder struct {
	mock *MockDepositHandler
}

// NewMockDepositHandler creates a new mock instance.
func NewMockDepositHandler(ctrl *gomock.Controller) *MockDepositHandler {
	mock := &MockDepositHandler{ctrl: ctrl}
	mock.recorder = &MockDepositHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepositHandler) EXPECT() *MockDepositHandlerMockRecorder {
	return m.recorder
}

// HandleDeposit mocks base method.
func (m *MockDepositHandler) HandleDeposit(sourceID, destID uint8, nonce uint64, resourceID types.ResourceID, calldata, handlerResponse []byte) (*message.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeposit", sourceID, destID, nonce, resourceID, calldata, handlerResponse)
	ret0, _ := ret[0].(*message.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleDeposit indicates an expected call of HandleDeposit.
func (mr *MockDepositHandlerMockRecorder) HandleDeposit(sourceID, destID, nonce, resourceID, calldata, handlerResponse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeposit", reflect.TypeOf((*MockDepositHandler)(nil).HandleDeposit), sourceID, destID, nonce, resourceID, calldata, handlerResponse)
}

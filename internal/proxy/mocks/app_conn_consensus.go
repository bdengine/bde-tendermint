// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	abciclient "github.com/tendermint/tendermint/abci/client"

	mock "github.com/stretchr/testify/mock"

	types "github.com/tendermint/tendermint/abci/types"
)

// AppConnConsensus is an autogenerated mock type for the AppConnConsensus type
type AppConnConsensus struct {
	mock.Mock
}

// BeginBlock provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) BeginBlock(_a0 context.Context, _a1 types.RequestBeginBlock) (*types.ResponseBeginBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseBeginBlock
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestBeginBlock) *types.ResponseBeginBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseBeginBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestBeginBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: _a0
func (_m *AppConnConsensus) Commit(_a0 context.Context) (*types.ResponseCommit, error) {
	ret := _m.Called(_a0)

	var r0 *types.ResponseCommit
	if rf, ok := ret.Get(0).(func(context.Context) *types.ResponseCommit); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseCommit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeliverTx provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) DeliverTx(_a0 context.Context, _a1 types.RequestDeliverTx) (*types.ResponseDeliverTx, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseDeliverTx
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestDeliverTx) *types.ResponseDeliverTx); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseDeliverTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestDeliverTx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlock provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) EndBlock(_a0 context.Context, _a1 types.RequestEndBlock) (*types.ResponseEndBlock, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseEndBlock
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestEndBlock) *types.ResponseEndBlock); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseEndBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestEndBlock) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnConsensus) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitChain provides a mock function with given fields: _a0, _a1
func (_m *AppConnConsensus) InitChain(_a0 context.Context, _a1 types.RequestInitChain) (*types.ResponseInitChain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ResponseInitChain
	if rf, ok := ret.Get(0).(func(context.Context, types.RequestInitChain) *types.ResponseInitChain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ResponseInitChain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.RequestInitChain) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetResponseCallback provides a mock function with given fields: _a0
func (_m *AppConnConsensus) SetResponseCallback(_a0 abciclient.Callback) {
	_m.Called(_a0)
}

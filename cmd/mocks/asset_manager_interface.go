// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// AssetManagerInterface is an autogenerated mock type for the AssetManagerInterface type
type AssetManagerInterface struct {
	mock.Mock
}

// CreateCollection provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *AssetManagerInterface) CreateCollection(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint16, _a3 int8, _a4 uint32, _a5 []uint16, _a6 string) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint16, int8, uint32, []uint16, string) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint16, int8, uint32, []uint16, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7
func (_m *AssetManagerInterface) CreateJob(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint8, _a3 int8, _a4 uint8, _a5 string, _a6 string, _a7 string) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *AssetManagerInterface) GetActiveStatus(_a0 *ethclient.Client, _a1 *bind.CallOpts, _a2 uint16) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint16) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCollectionStatus provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *AssetManagerInterface) SetCollectionStatus(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 bool, _a3 uint16) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, bool, uint16) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, bool, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCollection provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *AssetManagerInterface) UpdateCollection(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint16, _a3 uint16, _a4 uint32, _a5 int8, _a6 []uint16) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint16, uint16, uint32, int8, []uint16) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint16, uint16, uint32, int8, []uint16) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJob provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7
func (_m *AssetManagerInterface) UpdateJob(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint16, _a3 uint8, _a4 int8, _a5 uint8, _a6 string, _a7 string) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint16, uint8, int8, uint8, string, string) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint16, uint8, int8, uint8, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
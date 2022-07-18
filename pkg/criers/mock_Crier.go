// Code generated by mockery v1.0.0. DO NOT EDIT.

package criers

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockCrier is an autogenerated mock type for the Crier type
type MockCrier struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockCrier) Close() {
	_m.Called()
}

// Send provides a mock function with given fields: ctx, data
func (_m *MockCrier) Send(ctx context.Context, data []byte) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

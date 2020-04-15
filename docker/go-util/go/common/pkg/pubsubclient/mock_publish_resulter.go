// Code generated by mockery v1.0.0. DO NOT EDIT.
package pubsubclient

import context "context"
import mock "github.com/stretchr/testify/mock"

// MockPublishResulter is an autogenerated mock type for the PublishResulter type
type MockPublishResulter struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx
func (_m *MockPublishResulter) Get(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
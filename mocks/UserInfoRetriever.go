// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserInfoRetriever is an autogenerated mock type for the UserInfoRetriever type
type UserInfoRetriever struct {
	mock.Mock
}

type UserInfoRetriever_Expecter struct {
	mock *mock.Mock
}

func (_m *UserInfoRetriever) EXPECT() *UserInfoRetriever_Expecter {
	return &UserInfoRetriever_Expecter{mock: &_m.Mock}
}

// Retrieve provides a mock function with given fields: ctx, citizenID
func (_m *UserInfoRetriever) Retrieve(ctx context.Context, citizenID string) map[string]string {
	ret := _m.Called(ctx, citizenID)

	if len(ret) == 0 {
		panic("no return value specified for Retrieve")
	}

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, citizenID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// UserInfoRetriever_Retrieve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Retrieve'
type UserInfoRetriever_Retrieve_Call struct {
	*mock.Call
}

// Retrieve is a helper method to define mock.On call
//   - ctx context.Context
//   - citizenID string
func (_e *UserInfoRetriever_Expecter) Retrieve(ctx interface{}, citizenID interface{}) *UserInfoRetriever_Retrieve_Call {
	return &UserInfoRetriever_Retrieve_Call{Call: _e.mock.On("Retrieve", ctx, citizenID)}
}

func (_c *UserInfoRetriever_Retrieve_Call) Run(run func(ctx context.Context, citizenID string)) *UserInfoRetriever_Retrieve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserInfoRetriever_Retrieve_Call) Return(_a0 map[string]string) *UserInfoRetriever_Retrieve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserInfoRetriever_Retrieve_Call) RunAndReturn(run func(context.Context, string) map[string]string) *UserInfoRetriever_Retrieve_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserInfoRetriever creates a new instance of UserInfoRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserInfoRetriever(t interface {
	mock.TestingT
	Cleanup(func())
},
) *UserInfoRetriever {
	mock := &UserInfoRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

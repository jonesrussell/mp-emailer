// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockErrorHandlerInterface is an autogenerated mock type for the ErrorHandlerInterface type
type MockErrorHandlerInterface struct {
	mock.Mock
}

type MockErrorHandlerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockErrorHandlerInterface) EXPECT() *MockErrorHandlerInterface_Expecter {
	return &MockErrorHandlerInterface_Expecter{mock: &_m.Mock}
}

// HandleHTTPError provides a mock function with given fields: c, err, message, statusCode
func (_m *MockErrorHandlerInterface) HandleHTTPError(c echo.Context, err error, message string, statusCode int) error {
	ret := _m.Called(c, err, message, statusCode)

	if len(ret) == 0 {
		panic("no return value specified for HandleHTTPError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, error, string, int) error); ok {
		r0 = rf(c, err, message, statusCode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockErrorHandlerInterface_HandleHTTPError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleHTTPError'
type MockErrorHandlerInterface_HandleHTTPError_Call struct {
	*mock.Call
}

// HandleHTTPError is a helper method to define mock.On call
//   - c echo.Context
//   - err error
//   - message string
//   - statusCode int
func (_e *MockErrorHandlerInterface_Expecter) HandleHTTPError(c interface{}, err interface{}, message interface{}, statusCode interface{}) *MockErrorHandlerInterface_HandleHTTPError_Call {
	return &MockErrorHandlerInterface_HandleHTTPError_Call{Call: _e.mock.On("HandleHTTPError", c, err, message, statusCode)}
}

func (_c *MockErrorHandlerInterface_HandleHTTPError_Call) Run(run func(c echo.Context, err error, message string, statusCode int)) *MockErrorHandlerInterface_HandleHTTPError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(error), args[2].(string), args[3].(int))
	})
	return _c
}

func (_c *MockErrorHandlerInterface_HandleHTTPError_Call) Return(_a0 error) *MockErrorHandlerInterface_HandleHTTPError_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockErrorHandlerInterface_HandleHTTPError_Call) RunAndReturn(run func(echo.Context, error, string, int) error) *MockErrorHandlerInterface_HandleHTTPError_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockErrorHandlerInterface creates a new instance of MockErrorHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockErrorHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockErrorHandlerInterface {
	mock := &MockErrorHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

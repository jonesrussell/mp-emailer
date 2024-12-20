// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockFlashHandlerInterface is an autogenerated mock type for the FlashHandlerInterface type
type MockFlashHandlerInterface struct {
	mock.Mock
}

type MockFlashHandlerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFlashHandlerInterface) EXPECT() *MockFlashHandlerInterface_Expecter {
	return &MockFlashHandlerInterface_Expecter{mock: &_m.Mock}
}

// ClearSession provides a mock function with given fields: c
func (_m *MockFlashHandlerInterface) ClearSession(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ClearSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFlashHandlerInterface_ClearSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearSession'
type MockFlashHandlerInterface_ClearSession_Call struct {
	*mock.Call
}

// ClearSession is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockFlashHandlerInterface_Expecter) ClearSession(c interface{}) *MockFlashHandlerInterface_ClearSession_Call {
	return &MockFlashHandlerInterface_ClearSession_Call{Call: _e.mock.On("ClearSession", c)}
}

func (_c *MockFlashHandlerInterface_ClearSession_Call) Run(run func(c echo.Context)) *MockFlashHandlerInterface_ClearSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockFlashHandlerInterface_ClearSession_Call) Return(_a0 error) *MockFlashHandlerInterface_ClearSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlashHandlerInterface_ClearSession_Call) RunAndReturn(run func(echo.Context) error) *MockFlashHandlerInterface_ClearSession_Call {
	_c.Call.Return(run)
	return _c
}

// SetFlashAndSaveSession provides a mock function with given fields: c, message
func (_m *MockFlashHandlerInterface) SetFlashAndSaveSession(c echo.Context, message string) error {
	ret := _m.Called(c, message)

	if len(ret) == 0 {
		panic("no return value specified for SetFlashAndSaveSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(c, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFlashHandlerInterface_SetFlashAndSaveSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFlashAndSaveSession'
type MockFlashHandlerInterface_SetFlashAndSaveSession_Call struct {
	*mock.Call
}

// SetFlashAndSaveSession is a helper method to define mock.On call
//   - c echo.Context
//   - message string
func (_e *MockFlashHandlerInterface_Expecter) SetFlashAndSaveSession(c interface{}, message interface{}) *MockFlashHandlerInterface_SetFlashAndSaveSession_Call {
	return &MockFlashHandlerInterface_SetFlashAndSaveSession_Call{Call: _e.mock.On("SetFlashAndSaveSession", c, message)}
}

func (_c *MockFlashHandlerInterface_SetFlashAndSaveSession_Call) Run(run func(c echo.Context, message string)) *MockFlashHandlerInterface_SetFlashAndSaveSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockFlashHandlerInterface_SetFlashAndSaveSession_Call) Return(_a0 error) *MockFlashHandlerInterface_SetFlashAndSaveSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFlashHandlerInterface_SetFlashAndSaveSession_Call) RunAndReturn(run func(echo.Context, string) error) *MockFlashHandlerInterface_SetFlashAndSaveSession_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFlashHandlerInterface creates a new instance of MockFlashHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFlashHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFlashHandlerInterface {
	mock := &MockFlashHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

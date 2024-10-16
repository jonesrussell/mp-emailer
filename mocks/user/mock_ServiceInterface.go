// Code generated by mockery v2.46.3. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockServiceInterface is an autogenerated mock type for the ServiceInterface type
type MockServiceInterface struct {
	mock.Mock
}

type MockServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServiceInterface) EXPECT() *MockServiceInterface_Expecter {
	return &MockServiceInterface_Expecter{mock: &_m.Mock}
}

// RegisterUser provides a mock function with given fields: username, email, password
func (_m *MockServiceInterface) RegisterUser(username string, email string, password string) error {
	ret := _m.Called(username, email, password)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(username, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockServiceInterface_RegisterUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterUser'
type MockServiceInterface_RegisterUser_Call struct {
	*mock.Call
}

// RegisterUser is a helper method to define mock.On call
//   - username string
//   - email string
//   - password string
func (_e *MockServiceInterface_Expecter) RegisterUser(username interface{}, email interface{}, password interface{}) *MockServiceInterface_RegisterUser_Call {
	return &MockServiceInterface_RegisterUser_Call{Call: _e.mock.On("RegisterUser", username, email, password)}
}

func (_c *MockServiceInterface_RegisterUser_Call) Run(run func(username string, email string, password string)) *MockServiceInterface_RegisterUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockServiceInterface_RegisterUser_Call) Return(_a0 error) *MockServiceInterface_RegisterUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockServiceInterface_RegisterUser_Call) RunAndReturn(run func(string, string, string) error) *MockServiceInterface_RegisterUser_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyUser provides a mock function with given fields: username, password
func (_m *MockServiceInterface) VerifyUser(username string, password string) (string, error) {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUser")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServiceInterface_VerifyUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyUser'
type MockServiceInterface_VerifyUser_Call struct {
	*mock.Call
}

// VerifyUser is a helper method to define mock.On call
//   - username string
//   - password string
func (_e *MockServiceInterface_Expecter) VerifyUser(username interface{}, password interface{}) *MockServiceInterface_VerifyUser_Call {
	return &MockServiceInterface_VerifyUser_Call{Call: _e.mock.On("VerifyUser", username, password)}
}

func (_c *MockServiceInterface_VerifyUser_Call) Run(run func(username string, password string)) *MockServiceInterface_VerifyUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockServiceInterface_VerifyUser_Call) Return(_a0 string, _a1 error) *MockServiceInterface_VerifyUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServiceInterface_VerifyUser_Call) RunAndReturn(run func(string, string) (string, error)) *MockServiceInterface_VerifyUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServiceInterface creates a new instance of MockServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServiceInterface {
	mock := &MockServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
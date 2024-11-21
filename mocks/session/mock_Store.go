// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	sessions "github.com/gorilla/sessions"

	time "time"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

type MockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStore) EXPECT() *MockStore_Expecter {
	return &MockStore_Expecter{mock: &_m.Mock}
}

// Cleanup provides a mock function with given fields: threshold
func (_m *MockStore) Cleanup(threshold time.Time) error {
	ret := _m.Called(threshold)

	if len(ret) == 0 {
		panic("no return value specified for Cleanup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(threshold)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_Cleanup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cleanup'
type MockStore_Cleanup_Call struct {
	*mock.Call
}

// Cleanup is a helper method to define mock.On call
//   - threshold time.Time
func (_e *MockStore_Expecter) Cleanup(threshold interface{}) *MockStore_Cleanup_Call {
	return &MockStore_Cleanup_Call{Call: _e.mock.On("Cleanup", threshold)}
}

func (_c *MockStore_Cleanup_Call) Run(run func(threshold time.Time)) *MockStore_Cleanup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Time))
	})
	return _c
}

func (_c *MockStore_Cleanup_Call) Return(_a0 error) *MockStore_Cleanup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_Cleanup_Call) RunAndReturn(run func(time.Time) error) *MockStore_Cleanup_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: r, name
func (_m *MockStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	ret := _m.Called(r, name)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *sessions.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request, string) (*sessions.Session, error)); ok {
		return rf(r, name)
	}
	if rf, ok := ret.Get(0).(func(*http.Request, string) *sessions.Session); ok {
		r0 = rf(r, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sessions.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockStore_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - r *http.Request
//   - name string
func (_e *MockStore_Expecter) Get(r interface{}, name interface{}) *MockStore_Get_Call {
	return &MockStore_Get_Call{Call: _e.mock.On("Get", r, name)}
}

func (_c *MockStore_Get_Call) Run(run func(r *http.Request, name string)) *MockStore_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request), args[1].(string))
	})
	return _c
}

func (_c *MockStore_Get_Call) Return(_a0 *sessions.Session, _a1 error) *MockStore_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_Get_Call) RunAndReturn(run func(*http.Request, string) (*sessions.Session, error)) *MockStore_Get_Call {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: r, name
func (_m *MockStore) New(r *http.Request, name string) (*sessions.Session, error) {
	ret := _m.Called(r, name)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 *sessions.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request, string) (*sessions.Session, error)); ok {
		return rf(r, name)
	}
	if rf, ok := ret.Get(0).(func(*http.Request, string) *sessions.Session); ok {
		r0 = rf(r, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sessions.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request, string) error); ok {
		r1 = rf(r, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type MockStore_New_Call struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - r *http.Request
//   - name string
func (_e *MockStore_Expecter) New(r interface{}, name interface{}) *MockStore_New_Call {
	return &MockStore_New_Call{Call: _e.mock.On("New", r, name)}
}

func (_c *MockStore_New_Call) Run(run func(r *http.Request, name string)) *MockStore_New_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request), args[1].(string))
	})
	return _c
}

func (_c *MockStore_New_Call) Return(_a0 *sessions.Session, _a1 error) *MockStore_New_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_New_Call) RunAndReturn(run func(*http.Request, string) (*sessions.Session, error)) *MockStore_New_Call {
	_c.Call.Return(run)
	return _c
}

// RegenerateID provides a mock function with given fields: r, w
func (_m *MockStore) RegenerateID(r *http.Request, w http.ResponseWriter) (string, error) {
	ret := _m.Called(r, w)

	if len(ret) == 0 {
		panic("no return value specified for RegenerateID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request, http.ResponseWriter) (string, error)); ok {
		return rf(r, w)
	}
	if rf, ok := ret.Get(0).(func(*http.Request, http.ResponseWriter) string); ok {
		r0 = rf(r, w)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*http.Request, http.ResponseWriter) error); ok {
		r1 = rf(r, w)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_RegenerateID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegenerateID'
type MockStore_RegenerateID_Call struct {
	*mock.Call
}

// RegenerateID is a helper method to define mock.On call
//   - r *http.Request
//   - w http.ResponseWriter
func (_e *MockStore_Expecter) RegenerateID(r interface{}, w interface{}) *MockStore_RegenerateID_Call {
	return &MockStore_RegenerateID_Call{Call: _e.mock.On("RegenerateID", r, w)}
}

func (_c *MockStore_RegenerateID_Call) Run(run func(r *http.Request, w http.ResponseWriter)) *MockStore_RegenerateID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request), args[1].(http.ResponseWriter))
	})
	return _c
}

func (_c *MockStore_RegenerateID_Call) Return(_a0 string, _a1 error) *MockStore_RegenerateID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_RegenerateID_Call) RunAndReturn(run func(*http.Request, http.ResponseWriter) (string, error)) *MockStore_RegenerateID_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: r, w, s
func (_m *MockStore) Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	ret := _m.Called(r, w, s)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*http.Request, http.ResponseWriter, *sessions.Session) error); ok {
		r0 = rf(r, w, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockStore_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - r *http.Request
//   - w http.ResponseWriter
//   - s *sessions.Session
func (_e *MockStore_Expecter) Save(r interface{}, w interface{}, s interface{}) *MockStore_Save_Call {
	return &MockStore_Save_Call{Call: _e.mock.On("Save", r, w, s)}
}

func (_c *MockStore_Save_Call) Run(run func(r *http.Request, w http.ResponseWriter, s *sessions.Session)) *MockStore_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request), args[1].(http.ResponseWriter), args[2].(*sessions.Session))
	})
	return _c
}

func (_c *MockStore_Save_Call) Return(_a0 error) *MockStore_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_Save_Call) RunAndReturn(run func(*http.Request, http.ResponseWriter, *sessions.Session) error) *MockStore_Save_Call {
	_c.Call.Return(run)
	return _c
}

// SetOptions provides a mock function with given fields: options
func (_m *MockStore) SetOptions(options *sessions.Options) {
	_m.Called(options)
}

// MockStore_SetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetOptions'
type MockStore_SetOptions_Call struct {
	*mock.Call
}

// SetOptions is a helper method to define mock.On call
//   - options *sessions.Options
func (_e *MockStore_Expecter) SetOptions(options interface{}) *MockStore_SetOptions_Call {
	return &MockStore_SetOptions_Call{Call: _e.mock.On("SetOptions", options)}
}

func (_c *MockStore_SetOptions_Call) Run(run func(options *sessions.Options)) *MockStore_SetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sessions.Options))
	})
	return _c
}

func (_c *MockStore_SetOptions_Call) Return() *MockStore_SetOptions_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStore_SetOptions_Call) RunAndReturn(run func(*sessions.Options)) *MockStore_SetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// SetSameSite provides a mock function with given fields: mode
func (_m *MockStore) SetSameSite(mode http.SameSite) {
	_m.Called(mode)
}

// MockStore_SetSameSite_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSameSite'
type MockStore_SetSameSite_Call struct {
	*mock.Call
}

// SetSameSite is a helper method to define mock.On call
//   - mode http.SameSite
func (_e *MockStore_Expecter) SetSameSite(mode interface{}) *MockStore_SetSameSite_Call {
	return &MockStore_SetSameSite_Call{Call: _e.mock.On("SetSameSite", mode)}
}

func (_c *MockStore_SetSameSite_Call) Run(run func(mode http.SameSite)) *MockStore_SetSameSite_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.SameSite))
	})
	return _c
}

func (_c *MockStore_SetSameSite_Call) Return() *MockStore_SetSameSite_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStore_SetSameSite_Call) RunAndReturn(run func(http.SameSite)) *MockStore_SetSameSite_Call {
	_c.Call.Return(run)
	return _c
}

// SetSecure provides a mock function with given fields: secure
func (_m *MockStore) SetSecure(secure bool) {
	_m.Called(secure)
}

// MockStore_SetSecure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSecure'
type MockStore_SetSecure_Call struct {
	*mock.Call
}

// SetSecure is a helper method to define mock.On call
//   - secure bool
func (_e *MockStore_Expecter) SetSecure(secure interface{}) *MockStore_SetSecure_Call {
	return &MockStore_SetSecure_Call{Call: _e.mock.On("SetSecure", secure)}
}

func (_c *MockStore_SetSecure_Call) Run(run func(secure bool)) *MockStore_SetSecure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockStore_SetSecure_Call) Return() *MockStore_SetSecure_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStore_SetSecure_Call) RunAndReturn(run func(bool)) *MockStore_SetSecure_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

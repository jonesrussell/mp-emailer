// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	campaign "github.com/jonesrussell/mp-emailer/campaign"
	mock "github.com/stretchr/testify/mock"
)

// MockRepresentativeLookupServiceInterface is an autogenerated mock type for the RepresentativeLookupServiceInterface type
type MockRepresentativeLookupServiceInterface struct {
	mock.Mock
}

type MockRepresentativeLookupServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepresentativeLookupServiceInterface) EXPECT() *MockRepresentativeLookupServiceInterface_Expecter {
	return &MockRepresentativeLookupServiceInterface_Expecter{mock: &_m.Mock}
}

// FetchRepresentatives provides a mock function with given fields: postalCode
func (_m *MockRepresentativeLookupServiceInterface) FetchRepresentatives(postalCode string) ([]campaign.Representative, error) {
	ret := _m.Called(postalCode)

	if len(ret) == 0 {
		panic("no return value specified for FetchRepresentatives")
	}

	var r0 []campaign.Representative
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]campaign.Representative, error)); ok {
		return rf(postalCode)
	}
	if rf, ok := ret.Get(0).(func(string) []campaign.Representative); ok {
		r0 = rf(postalCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]campaign.Representative)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postalCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRepresentatives'
type MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call struct {
	*mock.Call
}

// FetchRepresentatives is a helper method to define mock.On call
//   - postalCode string
func (_e *MockRepresentativeLookupServiceInterface_Expecter) FetchRepresentatives(postalCode interface{}) *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call {
	return &MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call{Call: _e.mock.On("FetchRepresentatives", postalCode)}
}

func (_c *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call) Run(run func(postalCode string)) *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call) Return(_a0 []campaign.Representative, _a1 error) *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call) RunAndReturn(run func(string) ([]campaign.Representative, error)) *MockRepresentativeLookupServiceInterface_FetchRepresentatives_Call {
	_c.Call.Return(run)
	return _c
}

// FilterRepresentatives provides a mock function with given fields: representatives, filters
func (_m *MockRepresentativeLookupServiceInterface) FilterRepresentatives(representatives []campaign.Representative, filters map[string]string) []campaign.Representative {
	ret := _m.Called(representatives, filters)

	if len(ret) == 0 {
		panic("no return value specified for FilterRepresentatives")
	}

	var r0 []campaign.Representative
	if rf, ok := ret.Get(0).(func([]campaign.Representative, map[string]string) []campaign.Representative); ok {
		r0 = rf(representatives, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]campaign.Representative)
		}
	}

	return r0
}

// MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterRepresentatives'
type MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call struct {
	*mock.Call
}

// FilterRepresentatives is a helper method to define mock.On call
//   - representatives []campaign.Representative
//   - filters map[string]string
func (_e *MockRepresentativeLookupServiceInterface_Expecter) FilterRepresentatives(representatives interface{}, filters interface{}) *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call {
	return &MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call{Call: _e.mock.On("FilterRepresentatives", representatives, filters)}
}

func (_c *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call) Run(run func(representatives []campaign.Representative, filters map[string]string)) *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]campaign.Representative), args[1].(map[string]string))
	})
	return _c
}

func (_c *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call) Return(_a0 []campaign.Representative) *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call) RunAndReturn(run func([]campaign.Representative, map[string]string) []campaign.Representative) *MockRepresentativeLookupServiceInterface_FilterRepresentatives_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepresentativeLookupServiceInterface creates a new instance of MockRepresentativeLookupServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepresentativeLookupServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepresentativeLookupServiceInterface {
	mock := &MockRepresentativeLookupServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

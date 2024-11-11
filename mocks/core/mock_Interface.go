// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	core "github.com/fullstackdev42/mp-emailer/database/core"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Association provides a mock function with given fields: column
func (_m *MockInterface) Association(column string) core.AssociationInterface {
	ret := _m.Called(column)

	if len(ret) == 0 {
		panic("no return value specified for Association")
	}

	var r0 core.AssociationInterface
	if rf, ok := ret.Get(0).(func(string) core.AssociationInterface); ok {
		r0 = rf(column)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.AssociationInterface)
		}
	}

	return r0
}

// MockInterface_Association_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Association'
type MockInterface_Association_Call struct {
	*mock.Call
}

// Association is a helper method to define mock.On call
//   - column string
func (_e *MockInterface_Expecter) Association(column interface{}) *MockInterface_Association_Call {
	return &MockInterface_Association_Call{Call: _e.mock.On("Association", column)}
}

func (_c *MockInterface_Association_Call) Run(run func(column string)) *MockInterface_Association_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_Association_Call) Return(_a0 core.AssociationInterface) *MockInterface_Association_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Association_Call) RunAndReturn(run func(string) core.AssociationInterface) *MockInterface_Association_Call {
	_c.Call.Return(run)
	return _c
}

// AutoMigrate provides a mock function with given fields: dst
func (_m *MockInterface) AutoMigrate(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AutoMigrate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_AutoMigrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutoMigrate'
type MockInterface_AutoMigrate_Call struct {
	*mock.Call
}

// AutoMigrate is a helper method to define mock.On call
//   - dst ...interface{}
func (_e *MockInterface_Expecter) AutoMigrate(dst ...interface{}) *MockInterface_AutoMigrate_Call {
	return &MockInterface_AutoMigrate_Call{Call: _e.mock.On("AutoMigrate",
		append([]interface{}{}, dst...)...)}
}

func (_c *MockInterface_AutoMigrate_Call) Run(run func(dst ...interface{})) *MockInterface_AutoMigrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_AutoMigrate_Call) Return(_a0 error) *MockInterface_AutoMigrate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_AutoMigrate_Call) RunAndReturn(run func(...interface{}) error) *MockInterface_AutoMigrate_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: value
func (_m *MockInterface) Create(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - value interface{}
func (_e *MockInterface_Expecter) Create(value interface{}) *MockInterface_Create_Call {
	return &MockInterface_Create_Call{Call: _e.mock.On("Create", value)}
}

func (_c *MockInterface_Create_Call) Run(run func(value interface{})) *MockInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockInterface_Create_Call) Return(_a0 error) *MockInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Create_Call) RunAndReturn(run func(interface{}) error) *MockInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DB provides a mock function with given fields:
func (_m *MockInterface) DB() *gorm.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DB")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// MockInterface_DB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DB'
type MockInterface_DB_Call struct {
	*mock.Call
}

// DB is a helper method to define mock.On call
func (_e *MockInterface_Expecter) DB() *MockInterface_DB_Call {
	return &MockInterface_DB_Call{Call: _e.mock.On("DB")}
}

func (_c *MockInterface_DB_Call) Run(run func()) *MockInterface_DB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_DB_Call) Return(_a0 *gorm.DB) *MockInterface_DB_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_DB_Call) RunAndReturn(run func() *gorm.DB) *MockInterface_DB_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: value
func (_m *MockInterface) Delete(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - value interface{}
func (_e *MockInterface_Expecter) Delete(value interface{}) *MockInterface_Delete_Call {
	return &MockInterface_Delete_Call{Call: _e.mock.On("Delete", value)}
}

func (_c *MockInterface_Delete_Call) Run(run func(value interface{})) *MockInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockInterface_Delete_Call) Return(_a0 error) *MockInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Delete_Call) RunAndReturn(run func(interface{}) error) *MockInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: query, args
func (_m *MockInterface) Exec(query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockInterface_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Exec(query interface{}, args ...interface{}) *MockInterface_Exec_Call {
	return &MockInterface_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Exec_Call) Run(run func(query string, args ...interface{})) *MockInterface_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Exec_Call) Return(_a0 error) *MockInterface_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Exec_Call) RunAndReturn(run func(string, ...interface{}) error) *MockInterface_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: model, query, args
func (_m *MockInterface) Exists(model interface{}, query string, args ...interface{}) (bool, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) (bool, error)); ok {
		return rf(model, query, args...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) bool); ok {
		r0 = rf(model, query, args...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, ...interface{}) error); ok {
		r1 = rf(model, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockInterface_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - model interface{}
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Exists(model interface{}, query interface{}, args ...interface{}) *MockInterface_Exists_Call {
	return &MockInterface_Exists_Call{Call: _e.mock.On("Exists",
		append([]interface{}{model, query}, args...)...)}
}

func (_c *MockInterface_Exists_Call) Run(run func(model interface{}, query string, args ...interface{})) *MockInterface_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Exists_Call) Return(_a0 bool, _a1 error) *MockInterface_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Exists_Call) RunAndReturn(run func(interface{}, string, ...interface{}) (bool, error)) *MockInterface_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: model, query, args
func (_m *MockInterface) FindOne(model interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(model, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockInterface_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type MockInterface_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - model interface{}
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) FindOne(model interface{}, query interface{}, args ...interface{}) *MockInterface_FindOne_Call {
	return &MockInterface_FindOne_Call{Call: _e.mock.On("FindOne",
		append([]interface{}{model, query}, args...)...)}
}

func (_c *MockInterface_FindOne_Call) Run(run func(model interface{}, query string, args ...interface{})) *MockInterface_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_FindOne_Call) Return(_a0 error) *MockInterface_FindOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_FindOne_Call) RunAndReturn(run func(interface{}, string, ...interface{}) error) *MockInterface_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// GetSQLDB provides a mock function with given fields:
func (_m *MockInterface) GetSQLDB() (*sql.DB, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSQLDB")
	}

	var r0 *sql.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*sql.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_GetSQLDB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSQLDB'
type MockInterface_GetSQLDB_Call struct {
	*mock.Call
}

// GetSQLDB is a helper method to define mock.On call
func (_e *MockInterface_Expecter) GetSQLDB() *MockInterface_GetSQLDB_Call {
	return &MockInterface_GetSQLDB_Call{Call: _e.mock.On("GetSQLDB")}
}

func (_c *MockInterface_GetSQLDB_Call) Run(run func()) *MockInterface_GetSQLDB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_GetSQLDB_Call) Return(_a0 *sql.DB, _a1 error) *MockInterface_GetSQLDB_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_GetSQLDB_Call) RunAndReturn(run func() (*sql.DB, error)) *MockInterface_GetSQLDB_Call {
	_c.Call.Return(run)
	return _c
}

// Group provides a mock function with given fields: query
func (_m *MockInterface) Group(query string) core.Interface {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Group")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(string) core.Interface); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Group_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Group'
type MockInterface_Group_Call struct {
	*mock.Call
}

// Group is a helper method to define mock.On call
//   - query string
func (_e *MockInterface_Expecter) Group(query interface{}) *MockInterface_Group_Call {
	return &MockInterface_Group_Call{Call: _e.mock.On("Group", query)}
}

func (_c *MockInterface_Group_Call) Run(run func(query string)) *MockInterface_Group_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_Group_Call) Return(_a0 core.Interface) *MockInterface_Group_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Group_Call) RunAndReturn(run func(string) core.Interface) *MockInterface_Group_Call {
	_c.Call.Return(run)
	return _c
}

// Having provides a mock function with given fields: query, args
func (_m *MockInterface) Having(query string, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Having")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(string, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Having_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Having'
type MockInterface_Having_Call struct {
	*mock.Call
}

// Having is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Having(query interface{}, args ...interface{}) *MockInterface_Having_Call {
	return &MockInterface_Having_Call{Call: _e.mock.On("Having",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Having_Call) Run(run func(query string, args ...interface{})) *MockInterface_Having_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Having_Call) Return(_a0 core.Interface) *MockInterface_Having_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Having_Call) RunAndReturn(run func(string, ...interface{}) core.Interface) *MockInterface_Having_Call {
	_c.Call.Return(run)
	return _c
}

// Joins provides a mock function with given fields: query, args
func (_m *MockInterface) Joins(query string, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Joins")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(string, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Joins_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Joins'
type MockInterface_Joins_Call struct {
	*mock.Call
}

// Joins is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Joins(query interface{}, args ...interface{}) *MockInterface_Joins_Call {
	return &MockInterface_Joins_Call{Call: _e.mock.On("Joins",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Joins_Call) Run(run func(query string, args ...interface{})) *MockInterface_Joins_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Joins_Call) Return(_a0 core.Interface) *MockInterface_Joins_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Joins_Call) RunAndReturn(run func(string, ...interface{}) core.Interface) *MockInterface_Joins_Call {
	_c.Call.Return(run)
	return _c
}

// Limit provides a mock function with given fields: limit
func (_m *MockInterface) Limit(limit int) core.Interface {
	ret := _m.Called(limit)

	if len(ret) == 0 {
		panic("no return value specified for Limit")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(int) core.Interface); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Limit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Limit'
type MockInterface_Limit_Call struct {
	*mock.Call
}

// Limit is a helper method to define mock.On call
//   - limit int
func (_e *MockInterface_Expecter) Limit(limit interface{}) *MockInterface_Limit_Call {
	return &MockInterface_Limit_Call{Call: _e.mock.On("Limit", limit)}
}

func (_c *MockInterface_Limit_Call) Run(run func(limit int)) *MockInterface_Limit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockInterface_Limit_Call) Return(_a0 core.Interface) *MockInterface_Limit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Limit_Call) RunAndReturn(run func(int) core.Interface) *MockInterface_Limit_Call {
	_c.Call.Return(run)
	return _c
}

// Not provides a mock function with given fields: query, args
func (_m *MockInterface) Not(query interface{}, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Not")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Not_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Not'
type MockInterface_Not_Call struct {
	*mock.Call
}

// Not is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *MockInterface_Expecter) Not(query interface{}, args ...interface{}) *MockInterface_Not_Call {
	return &MockInterface_Not_Call{Call: _e.mock.On("Not",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Not_Call) Run(run func(query interface{}, args ...interface{})) *MockInterface_Not_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Not_Call) Return(_a0 core.Interface) *MockInterface_Not_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Not_Call) RunAndReturn(run func(interface{}, ...interface{}) core.Interface) *MockInterface_Not_Call {
	_c.Call.Return(run)
	return _c
}

// Offset provides a mock function with given fields: offset
func (_m *MockInterface) Offset(offset int) core.Interface {
	ret := _m.Called(offset)

	if len(ret) == 0 {
		panic("no return value specified for Offset")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(int) core.Interface); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Offset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Offset'
type MockInterface_Offset_Call struct {
	*mock.Call
}

// Offset is a helper method to define mock.On call
//   - offset int
func (_e *MockInterface_Expecter) Offset(offset interface{}) *MockInterface_Offset_Call {
	return &MockInterface_Offset_Call{Call: _e.mock.On("Offset", offset)}
}

func (_c *MockInterface_Offset_Call) Run(run func(offset int)) *MockInterface_Offset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockInterface_Offset_Call) Return(_a0 core.Interface) *MockInterface_Offset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Offset_Call) RunAndReturn(run func(int) core.Interface) *MockInterface_Offset_Call {
	_c.Call.Return(run)
	return _c
}

// Or provides a mock function with given fields: query, args
func (_m *MockInterface) Or(query interface{}, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Or")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Or_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Or'
type MockInterface_Or_Call struct {
	*mock.Call
}

// Or is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *MockInterface_Expecter) Or(query interface{}, args ...interface{}) *MockInterface_Or_Call {
	return &MockInterface_Or_Call{Call: _e.mock.On("Or",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Or_Call) Run(run func(query interface{}, args ...interface{})) *MockInterface_Or_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Or_Call) Return(_a0 core.Interface) *MockInterface_Or_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Or_Call) RunAndReturn(run func(interface{}, ...interface{}) core.Interface) *MockInterface_Or_Call {
	_c.Call.Return(run)
	return _c
}

// Order provides a mock function with given fields: value
func (_m *MockInterface) Order(value interface{}) core.Interface {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Order")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(interface{}) core.Interface); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Order_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Order'
type MockInterface_Order_Call struct {
	*mock.Call
}

// Order is a helper method to define mock.On call
//   - value interface{}
func (_e *MockInterface_Expecter) Order(value interface{}) *MockInterface_Order_Call {
	return &MockInterface_Order_Call{Call: _e.mock.On("Order", value)}
}

func (_c *MockInterface_Order_Call) Run(run func(value interface{})) *MockInterface_Order_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockInterface_Order_Call) Return(_a0 core.Interface) *MockInterface_Order_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Order_Call) RunAndReturn(run func(interface{}) core.Interface) *MockInterface_Order_Call {
	_c.Call.Return(run)
	return _c
}

// Preload provides a mock function with given fields: query, args
func (_m *MockInterface) Preload(query string, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Preload")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(string, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Preload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Preload'
type MockInterface_Preload_Call struct {
	*mock.Call
}

// Preload is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Preload(query interface{}, args ...interface{}) *MockInterface_Preload_Call {
	return &MockInterface_Preload_Call{Call: _e.mock.On("Preload",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Preload_Call) Run(run func(query string, args ...interface{})) *MockInterface_Preload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Preload_Call) Return(_a0 core.Interface) *MockInterface_Preload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Preload_Call) RunAndReturn(run func(string, ...interface{}) core.Interface) *MockInterface_Preload_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: query, args
func (_m *MockInterface) Query(query string, args ...interface{}) core.Result {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 core.Result
	if rf, ok := ret.Get(0).(func(string, ...interface{}) core.Result); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Result)
		}
	}

	return r0
}

// MockInterface_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockInterface_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *MockInterface_Expecter) Query(query interface{}, args ...interface{}) *MockInterface_Query_Call {
	return &MockInterface_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Query_Call) Run(run func(query string, args ...interface{})) *MockInterface_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Query_Call) Return(_a0 core.Result) *MockInterface_Query_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Query_Call) RunAndReturn(run func(string, ...interface{}) core.Result) *MockInterface_Query_Call {
	_c.Call.Return(run)
	return _c
}

// Unscoped provides a mock function with given fields:
func (_m *MockInterface) Unscoped() core.Interface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Unscoped")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func() core.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Unscoped_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unscoped'
type MockInterface_Unscoped_Call struct {
	*mock.Call
}

// Unscoped is a helper method to define mock.On call
func (_e *MockInterface_Expecter) Unscoped() *MockInterface_Unscoped_Call {
	return &MockInterface_Unscoped_Call{Call: _e.mock.On("Unscoped")}
}

func (_c *MockInterface_Unscoped_Call) Run(run func()) *MockInterface_Unscoped_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_Unscoped_Call) Return(_a0 core.Interface) *MockInterface_Unscoped_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Unscoped_Call) RunAndReturn(run func() core.Interface) *MockInterface_Unscoped_Call {
	_c.Call.Return(run)
	return _c
}

// Where provides a mock function with given fields: query, args
func (_m *MockInterface) Where(query interface{}, args ...interface{}) core.Interface {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Where")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) core.Interface); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_Where_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Where'
type MockInterface_Where_Call struct {
	*mock.Call
}

// Where is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *MockInterface_Expecter) Where(query interface{}, args ...interface{}) *MockInterface_Where_Call {
	return &MockInterface_Where_Call{Call: _e.mock.On("Where",
		append([]interface{}{query}, args...)...)}
}

func (_c *MockInterface_Where_Call) Run(run func(query interface{}, args ...interface{})) *MockInterface_Where_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *MockInterface_Where_Call) Return(_a0 core.Interface) *MockInterface_Where_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Where_Call) RunAndReturn(run func(interface{}, ...interface{}) core.Interface) *MockInterface_Where_Call {
	_c.Call.Return(run)
	return _c
}

// WithTrashed provides a mock function with given fields:
func (_m *MockInterface) WithTrashed() core.Interface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for WithTrashed")
	}

	var r0 core.Interface
	if rf, ok := ret.Get(0).(func() core.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Interface)
		}
	}

	return r0
}

// MockInterface_WithTrashed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTrashed'
type MockInterface_WithTrashed_Call struct {
	*mock.Call
}

// WithTrashed is a helper method to define mock.On call
func (_e *MockInterface_Expecter) WithTrashed() *MockInterface_WithTrashed_Call {
	return &MockInterface_WithTrashed_Call{Call: _e.mock.On("WithTrashed")}
}

func (_c *MockInterface_WithTrashed_Call) Run(run func()) *MockInterface_WithTrashed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_WithTrashed_Call) Return(_a0 core.Interface) *MockInterface_WithTrashed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_WithTrashed_Call) RunAndReturn(run func() core.Interface) *MockInterface_WithTrashed_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
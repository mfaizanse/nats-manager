// Code generated by mockery v2.32.3. DO NOT EDIT.

package mocks

import (
	context "context"

	chart "github.com/kyma-project/nats-manager/pkg/k8s/chart"

	manager "github.com/kyma-project/nats-manager/pkg/manager"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/kyma-project/nats-manager/api/v1alpha1"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

type Manager_Expecter struct {
	mock *mock.Mock
}

func (_m *Manager) EXPECT() *Manager_Expecter {
	return &Manager_Expecter{mock: &_m.Mock}
}

// DeleteInstance provides a mock function with given fields: _a0, _a1
func (_m *Manager) DeleteInstance(_a0 context.Context, _a1 *chart.ReleaseInstance) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *chart.ReleaseInstance) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_DeleteInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteInstance'
type Manager_DeleteInstance_Call struct {
	*mock.Call
}

// DeleteInstance is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *chart.ReleaseInstance
func (_e *Manager_Expecter) DeleteInstance(_a0 interface{}, _a1 interface{}) *Manager_DeleteInstance_Call {
	return &Manager_DeleteInstance_Call{Call: _e.mock.On("DeleteInstance", _a0, _a1)}
}

func (_c *Manager_DeleteInstance_Call) Run(run func(_a0 context.Context, _a1 *chart.ReleaseInstance)) *Manager_DeleteInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*chart.ReleaseInstance))
	})
	return _c
}

func (_c *Manager_DeleteInstance_Call) Return(_a0 error) *Manager_DeleteInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_DeleteInstance_Call) RunAndReturn(run func(context.Context, *chart.ReleaseInstance) error) *Manager_DeleteInstance_Call {
	_c.Call.Return(run)
	return _c
}

// DeployInstance provides a mock function with given fields: _a0, _a1
func (_m *Manager) DeployInstance(_a0 context.Context, _a1 *chart.ReleaseInstance) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *chart.ReleaseInstance) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Manager_DeployInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeployInstance'
type Manager_DeployInstance_Call struct {
	*mock.Call
}

// DeployInstance is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *chart.ReleaseInstance
func (_e *Manager_Expecter) DeployInstance(_a0 interface{}, _a1 interface{}) *Manager_DeployInstance_Call {
	return &Manager_DeployInstance_Call{Call: _e.mock.On("DeployInstance", _a0, _a1)}
}

func (_c *Manager_DeployInstance_Call) Run(run func(_a0 context.Context, _a1 *chart.ReleaseInstance)) *Manager_DeployInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*chart.ReleaseInstance))
	})
	return _c
}

func (_c *Manager_DeployInstance_Call) Return(_a0 error) *Manager_DeployInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_DeployInstance_Call) RunAndReturn(run func(context.Context, *chart.ReleaseInstance) error) *Manager_DeployInstance_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateNATSResources provides a mock function with given fields: _a0, _a1
func (_m *Manager) GenerateNATSResources(_a0 *chart.ReleaseInstance, _a1 ...manager.Option) (*chart.ManifestResources, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *chart.ManifestResources
	var r1 error
	if rf, ok := ret.Get(0).(func(*chart.ReleaseInstance, ...manager.Option) (*chart.ManifestResources, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(*chart.ReleaseInstance, ...manager.Option) *chart.ManifestResources); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*chart.ManifestResources)
		}
	}

	if rf, ok := ret.Get(1).(func(*chart.ReleaseInstance, ...manager.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_GenerateNATSResources_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateNATSResources'
type Manager_GenerateNATSResources_Call struct {
	*mock.Call
}

// GenerateNATSResources is a helper method to define mock.On call
//   - _a0 *chart.ReleaseInstance
//   - _a1 ...manager.Option
func (_e *Manager_Expecter) GenerateNATSResources(_a0 interface{}, _a1 ...interface{}) *Manager_GenerateNATSResources_Call {
	return &Manager_GenerateNATSResources_Call{Call: _e.mock.On("GenerateNATSResources",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Manager_GenerateNATSResources_Call) Run(run func(_a0 *chart.ReleaseInstance, _a1 ...manager.Option)) *Manager_GenerateNATSResources_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]manager.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(manager.Option)
			}
		}
		run(args[0].(*chart.ReleaseInstance), variadicArgs...)
	})
	return _c
}

func (_c *Manager_GenerateNATSResources_Call) Return(_a0 *chart.ManifestResources, _a1 error) *Manager_GenerateNATSResources_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_GenerateNATSResources_Call) RunAndReturn(run func(*chart.ReleaseInstance, ...manager.Option) (*chart.ManifestResources, error)) *Manager_GenerateNATSResources_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateOverrides provides a mock function with given fields: _a0, _a1, _a2
func (_m *Manager) GenerateOverrides(_a0 *v1alpha1.NATSSpec, _a1 bool, _a2 bool) map[string]interface{} {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(*v1alpha1.NATSSpec, bool, bool) map[string]interface{}); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Manager_GenerateOverrides_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateOverrides'
type Manager_GenerateOverrides_Call struct {
	*mock.Call
}

// GenerateOverrides is a helper method to define mock.On call
//   - _a0 *v1alpha1.NATSSpec
//   - _a1 bool
//   - _a2 bool
func (_e *Manager_Expecter) GenerateOverrides(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Manager_GenerateOverrides_Call {
	return &Manager_GenerateOverrides_Call{Call: _e.mock.On("GenerateOverrides", _a0, _a1, _a2)}
}

func (_c *Manager_GenerateOverrides_Call) Run(run func(_a0 *v1alpha1.NATSSpec, _a1 bool, _a2 bool)) *Manager_GenerateOverrides_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*v1alpha1.NATSSpec), args[1].(bool), args[2].(bool))
	})
	return _c
}

func (_c *Manager_GenerateOverrides_Call) Return(_a0 map[string]interface{}) *Manager_GenerateOverrides_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Manager_GenerateOverrides_Call) RunAndReturn(run func(*v1alpha1.NATSSpec, bool, bool) map[string]interface{}) *Manager_GenerateOverrides_Call {
	_c.Call.Return(run)
	return _c
}

// IsNATSStatefulSetReady provides a mock function with given fields: _a0, _a1
func (_m *Manager) IsNATSStatefulSetReady(_a0 context.Context, _a1 *chart.ReleaseInstance) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *chart.ReleaseInstance) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *chart.ReleaseInstance) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *chart.ReleaseInstance) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Manager_IsNATSStatefulSetReady_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNATSStatefulSetReady'
type Manager_IsNATSStatefulSetReady_Call struct {
	*mock.Call
}

// IsNATSStatefulSetReady is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *chart.ReleaseInstance
func (_e *Manager_Expecter) IsNATSStatefulSetReady(_a0 interface{}, _a1 interface{}) *Manager_IsNATSStatefulSetReady_Call {
	return &Manager_IsNATSStatefulSetReady_Call{Call: _e.mock.On("IsNATSStatefulSetReady", _a0, _a1)}
}

func (_c *Manager_IsNATSStatefulSetReady_Call) Run(run func(_a0 context.Context, _a1 *chart.ReleaseInstance)) *Manager_IsNATSStatefulSetReady_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*chart.ReleaseInstance))
	})
	return _c
}

func (_c *Manager_IsNATSStatefulSetReady_Call) Return(_a0 bool, _a1 error) *Manager_IsNATSStatefulSetReady_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Manager_IsNATSStatefulSetReady_Call) RunAndReturn(run func(context.Context, *chart.ReleaseInstance) (bool, error)) *Manager_IsNATSStatefulSetReady_Call {
	_c.Call.Return(run)
	return _c
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

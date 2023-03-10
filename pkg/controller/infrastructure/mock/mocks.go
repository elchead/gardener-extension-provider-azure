// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-azure/pkg/controller/infrastructure (interfaces: Reconciler)

// Package infrastructure is a generated GoMock package.
package infrastructure

import (
	context "context"
	reflect "reflect"

	azure "github.com/gardener/gardener-extension-provider-azure/pkg/apis/azure"
	v1alpha1 "github.com/gardener/gardener-extension-provider-azure/pkg/apis/azure/v1alpha1"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	extensions "github.com/gardener/gardener/pkg/extensions"
	gomock "github.com/golang/mock/gomock"
)

// MockReconciler is a mock of Reconciler interface.
type MockReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReconcilerMockRecorder
}

// MockReconcilerMockRecorder is the mock recorder for MockReconciler.
type MockReconcilerMockRecorder struct {
	mock *MockReconciler
}

// NewMockReconciler creates a new mock instance.
func NewMockReconciler(ctrl *gomock.Controller) *MockReconciler {
	mock := &MockReconciler{ctrl: ctrl}
	mock.recorder = &MockReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconciler) EXPECT() *MockReconcilerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockReconciler) Delete(arg0 context.Context, arg1 *v1alpha10.Infrastructure, arg2 *azure.InfrastructureConfig, arg3 *extensions.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockReconcilerMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReconciler)(nil).Delete), arg0, arg1, arg2, arg3)
}

// GetState mocks base method.
func (m *MockReconciler) GetState(arg0 context.Context, arg1 *v1alpha1.InfrastructureStatus) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockReconcilerMockRecorder) GetState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockReconciler)(nil).GetState), arg0, arg1)
}

// Reconcile mocks base method.
func (m *MockReconciler) Reconcile(arg0 context.Context, arg1 *v1alpha10.Infrastructure, arg2 *azure.InfrastructureConfig, arg3 *extensions.Cluster) (*v1alpha1.InfrastructureStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.InfrastructureStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockReconcilerMockRecorder) Reconcile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockReconciler)(nil).Reconcile), arg0, arg1, arg2, arg3)
}
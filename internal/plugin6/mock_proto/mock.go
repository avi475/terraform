// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hashicorp/terraform/internal/tfplugin6 (interfaces: ProviderClient,Provider_InvokeActionClient)
//
// Generated by this command:
//
//	mockgen -destination mock.go github.com/hashicorp/terraform/internal/tfplugin6 ProviderClient,Provider_InvokeActionClient
//

// Package mock_tfplugin6 is a generated GoMock package.
package mock_tfplugin6

import (
	context "context"
	reflect "reflect"

	tfplugin6 "github.com/hashicorp/terraform/internal/tfplugin6"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockProviderClient is a mock of ProviderClient interface.
type MockProviderClient struct {
	ctrl     *gomock.Controller
	recorder *MockProviderClientMockRecorder
}

// MockProviderClientMockRecorder is the mock recorder for MockProviderClient.
type MockProviderClientMockRecorder struct {
	mock *MockProviderClient
}

// NewMockProviderClient creates a new mock instance.
func NewMockProviderClient(ctrl *gomock.Controller) *MockProviderClient {
	mock := &MockProviderClient{ctrl: ctrl}
	mock.recorder = &MockProviderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderClient) EXPECT() *MockProviderClientMockRecorder {
	return m.recorder
}

// ApplyResourceChange mocks base method.
func (m *MockProviderClient) ApplyResourceChange(arg0 context.Context, arg1 *tfplugin6.ApplyResourceChange_Request, arg2 ...grpc.CallOption) (*tfplugin6.ApplyResourceChange_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ApplyResourceChange", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ApplyResourceChange_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyResourceChange indicates an expected call of ApplyResourceChange.
func (mr *MockProviderClientMockRecorder) ApplyResourceChange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyResourceChange", reflect.TypeOf((*MockProviderClient)(nil).ApplyResourceChange), varargs...)
}

// CallFunction mocks base method.
func (m *MockProviderClient) CallFunction(arg0 context.Context, arg1 *tfplugin6.CallFunction_Request, arg2 ...grpc.CallOption) (*tfplugin6.CallFunction_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CallFunction", varargs...)
	ret0, _ := ret[0].(*tfplugin6.CallFunction_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallFunction indicates an expected call of CallFunction.
func (mr *MockProviderClientMockRecorder) CallFunction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallFunction", reflect.TypeOf((*MockProviderClient)(nil).CallFunction), varargs...)
}

// CloseEphemeralResource mocks base method.
func (m *MockProviderClient) CloseEphemeralResource(arg0 context.Context, arg1 *tfplugin6.CloseEphemeralResource_Request, arg2 ...grpc.CallOption) (*tfplugin6.CloseEphemeralResource_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloseEphemeralResource", varargs...)
	ret0, _ := ret[0].(*tfplugin6.CloseEphemeralResource_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseEphemeralResource indicates an expected call of CloseEphemeralResource.
func (mr *MockProviderClientMockRecorder) CloseEphemeralResource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseEphemeralResource", reflect.TypeOf((*MockProviderClient)(nil).CloseEphemeralResource), varargs...)
}

// ConfigureProvider mocks base method.
func (m *MockProviderClient) ConfigureProvider(arg0 context.Context, arg1 *tfplugin6.ConfigureProvider_Request, arg2 ...grpc.CallOption) (*tfplugin6.ConfigureProvider_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigureProvider", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ConfigureProvider_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureProvider indicates an expected call of ConfigureProvider.
func (mr *MockProviderClientMockRecorder) ConfigureProvider(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureProvider", reflect.TypeOf((*MockProviderClient)(nil).ConfigureProvider), varargs...)
}

// ConfigureStateStore mocks base method.
func (m *MockProviderClient) ConfigureStateStore(arg0 context.Context, arg1 *tfplugin6.ConfigureStateStore_Request, arg2 ...grpc.CallOption) (*tfplugin6.ConfigureStateStore_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConfigureStateStore", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ConfigureStateStore_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigureStateStore indicates an expected call of ConfigureStateStore.
func (mr *MockProviderClientMockRecorder) ConfigureStateStore(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureStateStore", reflect.TypeOf((*MockProviderClient)(nil).ConfigureStateStore), varargs...)
}

// DeleteState mocks base method.
func (m *MockProviderClient) DeleteState(arg0 context.Context, arg1 *tfplugin6.DeleteState_Request, arg2 ...grpc.CallOption) (*tfplugin6.DeleteState_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteState", varargs...)
	ret0, _ := ret[0].(*tfplugin6.DeleteState_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteState indicates an expected call of DeleteState.
func (mr *MockProviderClientMockRecorder) DeleteState(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteState", reflect.TypeOf((*MockProviderClient)(nil).DeleteState), varargs...)
}

// GetFunctions mocks base method.
func (m *MockProviderClient) GetFunctions(arg0 context.Context, arg1 *tfplugin6.GetFunctions_Request, arg2 ...grpc.CallOption) (*tfplugin6.GetFunctions_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunctions", varargs...)
	ret0, _ := ret[0].(*tfplugin6.GetFunctions_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunctions indicates an expected call of GetFunctions.
func (mr *MockProviderClientMockRecorder) GetFunctions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctions", reflect.TypeOf((*MockProviderClient)(nil).GetFunctions), varargs...)
}

// GetMetadata mocks base method.
func (m *MockProviderClient) GetMetadata(arg0 context.Context, arg1 *tfplugin6.GetMetadata_Request, arg2 ...grpc.CallOption) (*tfplugin6.GetMetadata_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetadata", varargs...)
	ret0, _ := ret[0].(*tfplugin6.GetMetadata_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockProviderClientMockRecorder) GetMetadata(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockProviderClient)(nil).GetMetadata), varargs...)
}

// GetProviderSchema mocks base method.
func (m *MockProviderClient) GetProviderSchema(arg0 context.Context, arg1 *tfplugin6.GetProviderSchema_Request, arg2 ...grpc.CallOption) (*tfplugin6.GetProviderSchema_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProviderSchema", varargs...)
	ret0, _ := ret[0].(*tfplugin6.GetProviderSchema_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProviderSchema indicates an expected call of GetProviderSchema.
func (mr *MockProviderClientMockRecorder) GetProviderSchema(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviderSchema", reflect.TypeOf((*MockProviderClient)(nil).GetProviderSchema), varargs...)
}

// GetResourceIdentitySchemas mocks base method.
func (m *MockProviderClient) GetResourceIdentitySchemas(arg0 context.Context, arg1 *tfplugin6.GetResourceIdentitySchemas_Request, arg2 ...grpc.CallOption) (*tfplugin6.GetResourceIdentitySchemas_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceIdentitySchemas", varargs...)
	ret0, _ := ret[0].(*tfplugin6.GetResourceIdentitySchemas_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceIdentitySchemas indicates an expected call of GetResourceIdentitySchemas.
func (mr *MockProviderClientMockRecorder) GetResourceIdentitySchemas(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceIdentitySchemas", reflect.TypeOf((*MockProviderClient)(nil).GetResourceIdentitySchemas), varargs...)
}

// GetStates mocks base method.
func (m *MockProviderClient) GetStates(arg0 context.Context, arg1 *tfplugin6.GetStates_Request, arg2 ...grpc.CallOption) (*tfplugin6.GetStates_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStates", varargs...)
	ret0, _ := ret[0].(*tfplugin6.GetStates_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStates indicates an expected call of GetStates.
func (mr *MockProviderClientMockRecorder) GetStates(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStates", reflect.TypeOf((*MockProviderClient)(nil).GetStates), varargs...)
}

// ImportResourceState mocks base method.
func (m *MockProviderClient) ImportResourceState(arg0 context.Context, arg1 *tfplugin6.ImportResourceState_Request, arg2 ...grpc.CallOption) (*tfplugin6.ImportResourceState_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportResourceState", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ImportResourceState_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportResourceState indicates an expected call of ImportResourceState.
func (mr *MockProviderClientMockRecorder) ImportResourceState(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportResourceState", reflect.TypeOf((*MockProviderClient)(nil).ImportResourceState), varargs...)
}

// InvokeAction mocks base method.
func (m *MockProviderClient) InvokeAction(arg0 context.Context, arg1 *tfplugin6.InvokeAction_Request, arg2 ...grpc.CallOption) (tfplugin6.Provider_InvokeActionClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InvokeAction", varargs...)
	ret0, _ := ret[0].(tfplugin6.Provider_InvokeActionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeAction indicates an expected call of InvokeAction.
func (mr *MockProviderClientMockRecorder) InvokeAction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeAction", reflect.TypeOf((*MockProviderClient)(nil).InvokeAction), varargs...)
}

// ListResource mocks base method.
func (m *MockProviderClient) ListResource(arg0 context.Context, arg1 *tfplugin6.ListResource_Request, arg2 ...grpc.CallOption) (tfplugin6.Provider_ListResourceClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResource", varargs...)
	ret0, _ := ret[0].(tfplugin6.Provider_ListResourceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResource indicates an expected call of ListResource.
func (mr *MockProviderClientMockRecorder) ListResource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResource", reflect.TypeOf((*MockProviderClient)(nil).ListResource), varargs...)
}

// MoveResourceState mocks base method.
func (m *MockProviderClient) MoveResourceState(arg0 context.Context, arg1 *tfplugin6.MoveResourceState_Request, arg2 ...grpc.CallOption) (*tfplugin6.MoveResourceState_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MoveResourceState", varargs...)
	ret0, _ := ret[0].(*tfplugin6.MoveResourceState_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MoveResourceState indicates an expected call of MoveResourceState.
func (mr *MockProviderClientMockRecorder) MoveResourceState(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveResourceState", reflect.TypeOf((*MockProviderClient)(nil).MoveResourceState), varargs...)
}

// OpenEphemeralResource mocks base method.
func (m *MockProviderClient) OpenEphemeralResource(arg0 context.Context, arg1 *tfplugin6.OpenEphemeralResource_Request, arg2 ...grpc.CallOption) (*tfplugin6.OpenEphemeralResource_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OpenEphemeralResource", varargs...)
	ret0, _ := ret[0].(*tfplugin6.OpenEphemeralResource_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenEphemeralResource indicates an expected call of OpenEphemeralResource.
func (mr *MockProviderClientMockRecorder) OpenEphemeralResource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenEphemeralResource", reflect.TypeOf((*MockProviderClient)(nil).OpenEphemeralResource), varargs...)
}

// PlanAction mocks base method.
func (m *MockProviderClient) PlanAction(arg0 context.Context, arg1 *tfplugin6.PlanAction_Request, arg2 ...grpc.CallOption) (*tfplugin6.PlanAction_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlanAction", varargs...)
	ret0, _ := ret[0].(*tfplugin6.PlanAction_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlanAction indicates an expected call of PlanAction.
func (mr *MockProviderClientMockRecorder) PlanAction(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlanAction", reflect.TypeOf((*MockProviderClient)(nil).PlanAction), varargs...)
}

// PlanResourceChange mocks base method.
func (m *MockProviderClient) PlanResourceChange(arg0 context.Context, arg1 *tfplugin6.PlanResourceChange_Request, arg2 ...grpc.CallOption) (*tfplugin6.PlanResourceChange_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PlanResourceChange", varargs...)
	ret0, _ := ret[0].(*tfplugin6.PlanResourceChange_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PlanResourceChange indicates an expected call of PlanResourceChange.
func (mr *MockProviderClientMockRecorder) PlanResourceChange(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlanResourceChange", reflect.TypeOf((*MockProviderClient)(nil).PlanResourceChange), varargs...)
}

// ReadDataSource mocks base method.
func (m *MockProviderClient) ReadDataSource(arg0 context.Context, arg1 *tfplugin6.ReadDataSource_Request, arg2 ...grpc.CallOption) (*tfplugin6.ReadDataSource_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadDataSource", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ReadDataSource_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDataSource indicates an expected call of ReadDataSource.
func (mr *MockProviderClientMockRecorder) ReadDataSource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDataSource", reflect.TypeOf((*MockProviderClient)(nil).ReadDataSource), varargs...)
}

// ReadResource mocks base method.
func (m *MockProviderClient) ReadResource(arg0 context.Context, arg1 *tfplugin6.ReadResource_Request, arg2 ...grpc.CallOption) (*tfplugin6.ReadResource_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadResource", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ReadResource_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadResource indicates an expected call of ReadResource.
func (mr *MockProviderClientMockRecorder) ReadResource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadResource", reflect.TypeOf((*MockProviderClient)(nil).ReadResource), varargs...)
}

// RenewEphemeralResource mocks base method.
func (m *MockProviderClient) RenewEphemeralResource(arg0 context.Context, arg1 *tfplugin6.RenewEphemeralResource_Request, arg2 ...grpc.CallOption) (*tfplugin6.RenewEphemeralResource_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenewEphemeralResource", varargs...)
	ret0, _ := ret[0].(*tfplugin6.RenewEphemeralResource_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewEphemeralResource indicates an expected call of RenewEphemeralResource.
func (mr *MockProviderClientMockRecorder) RenewEphemeralResource(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewEphemeralResource", reflect.TypeOf((*MockProviderClient)(nil).RenewEphemeralResource), varargs...)
}

// StopProvider mocks base method.
func (m *MockProviderClient) StopProvider(arg0 context.Context, arg1 *tfplugin6.StopProvider_Request, arg2 ...grpc.CallOption) (*tfplugin6.StopProvider_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopProvider", varargs...)
	ret0, _ := ret[0].(*tfplugin6.StopProvider_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopProvider indicates an expected call of StopProvider.
func (mr *MockProviderClientMockRecorder) StopProvider(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopProvider", reflect.TypeOf((*MockProviderClient)(nil).StopProvider), varargs...)
}

// UpgradeResourceIdentity mocks base method.
func (m *MockProviderClient) UpgradeResourceIdentity(arg0 context.Context, arg1 *tfplugin6.UpgradeResourceIdentity_Request, arg2 ...grpc.CallOption) (*tfplugin6.UpgradeResourceIdentity_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeResourceIdentity", varargs...)
	ret0, _ := ret[0].(*tfplugin6.UpgradeResourceIdentity_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeResourceIdentity indicates an expected call of UpgradeResourceIdentity.
func (mr *MockProviderClientMockRecorder) UpgradeResourceIdentity(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeResourceIdentity", reflect.TypeOf((*MockProviderClient)(nil).UpgradeResourceIdentity), varargs...)
}

// UpgradeResourceState mocks base method.
func (m *MockProviderClient) UpgradeResourceState(arg0 context.Context, arg1 *tfplugin6.UpgradeResourceState_Request, arg2 ...grpc.CallOption) (*tfplugin6.UpgradeResourceState_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeResourceState", varargs...)
	ret0, _ := ret[0].(*tfplugin6.UpgradeResourceState_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeResourceState indicates an expected call of UpgradeResourceState.
func (mr *MockProviderClientMockRecorder) UpgradeResourceState(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeResourceState", reflect.TypeOf((*MockProviderClient)(nil).UpgradeResourceState), varargs...)
}

// ValidateActionConfig mocks base method.
func (m *MockProviderClient) ValidateActionConfig(arg0 context.Context, arg1 *tfplugin6.ValidateActionConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateActionConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateActionConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateActionConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateActionConfig indicates an expected call of ValidateActionConfig.
func (mr *MockProviderClientMockRecorder) ValidateActionConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateActionConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateActionConfig), varargs...)
}

// ValidateDataResourceConfig mocks base method.
func (m *MockProviderClient) ValidateDataResourceConfig(arg0 context.Context, arg1 *tfplugin6.ValidateDataResourceConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateDataResourceConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateDataResourceConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateDataResourceConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateDataResourceConfig indicates an expected call of ValidateDataResourceConfig.
func (mr *MockProviderClientMockRecorder) ValidateDataResourceConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDataResourceConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateDataResourceConfig), varargs...)
}

// ValidateEphemeralResourceConfig mocks base method.
func (m *MockProviderClient) ValidateEphemeralResourceConfig(arg0 context.Context, arg1 *tfplugin6.ValidateEphemeralResourceConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateEphemeralResourceConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateEphemeralResourceConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateEphemeralResourceConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateEphemeralResourceConfig indicates an expected call of ValidateEphemeralResourceConfig.
func (mr *MockProviderClientMockRecorder) ValidateEphemeralResourceConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEphemeralResourceConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateEphemeralResourceConfig), varargs...)
}

// ValidateListResourceConfig mocks base method.
func (m *MockProviderClient) ValidateListResourceConfig(arg0 context.Context, arg1 *tfplugin6.ValidateListResourceConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateListResourceConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateListResourceConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateListResourceConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateListResourceConfig indicates an expected call of ValidateListResourceConfig.
func (mr *MockProviderClientMockRecorder) ValidateListResourceConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateListResourceConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateListResourceConfig), varargs...)
}

// ValidateProviderConfig mocks base method.
func (m *MockProviderClient) ValidateProviderConfig(arg0 context.Context, arg1 *tfplugin6.ValidateProviderConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateProviderConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateProviderConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateProviderConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateProviderConfig indicates an expected call of ValidateProviderConfig.
func (mr *MockProviderClientMockRecorder) ValidateProviderConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProviderConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateProviderConfig), varargs...)
}

// ValidateResourceConfig mocks base method.
func (m *MockProviderClient) ValidateResourceConfig(arg0 context.Context, arg1 *tfplugin6.ValidateResourceConfig_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateResourceConfig_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateResourceConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateResourceConfig_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateResourceConfig indicates an expected call of ValidateResourceConfig.
func (mr *MockProviderClientMockRecorder) ValidateResourceConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateResourceConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateResourceConfig), varargs...)
}

// ValidateStateStoreConfig mocks base method.
func (m *MockProviderClient) ValidateStateStoreConfig(arg0 context.Context, arg1 *tfplugin6.ValidateStateStore_Request, arg2 ...grpc.CallOption) (*tfplugin6.ValidateStateStore_Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateStateStoreConfig", varargs...)
	ret0, _ := ret[0].(*tfplugin6.ValidateStateStore_Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateStateStoreConfig indicates an expected call of ValidateStateStoreConfig.
func (mr *MockProviderClientMockRecorder) ValidateStateStoreConfig(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStateStoreConfig", reflect.TypeOf((*MockProviderClient)(nil).ValidateStateStoreConfig), varargs...)
}

// MockProvider_InvokeActionClient is a mock of Provider_InvokeActionClient interface.
type MockProvider_InvokeActionClient struct {
	ctrl     *gomock.Controller
	recorder *MockProvider_InvokeActionClientMockRecorder
}

// MockProvider_InvokeActionClientMockRecorder is the mock recorder for MockProvider_InvokeActionClient.
type MockProvider_InvokeActionClientMockRecorder struct {
	mock *MockProvider_InvokeActionClient
}

// NewMockProvider_InvokeActionClient creates a new mock instance.
func NewMockProvider_InvokeActionClient(ctrl *gomock.Controller) *MockProvider_InvokeActionClient {
	mock := &MockProvider_InvokeActionClient{ctrl: ctrl}
	mock.recorder = &MockProvider_InvokeActionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider_InvokeActionClient) EXPECT() *MockProvider_InvokeActionClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockProvider_InvokeActionClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockProvider_InvokeActionClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockProvider_InvokeActionClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockProvider_InvokeActionClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).Context))
}

// Header mocks base method.
func (m *MockProvider_InvokeActionClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockProvider_InvokeActionClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockProvider_InvokeActionClient) Recv() (*tfplugin6.InvokeAction_Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*tfplugin6.InvokeAction_Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockProvider_InvokeActionClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockProvider_InvokeActionClient) RecvMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockProvider_InvokeActionClientMockRecorder) RecvMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockProvider_InvokeActionClient) SendMsg(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockProvider_InvokeActionClientMockRecorder) SendMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockProvider_InvokeActionClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockProvider_InvokeActionClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockProvider_InvokeActionClient)(nil).Trailer))
}

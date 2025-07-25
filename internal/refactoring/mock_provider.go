package refactoring

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/terraform/internal/configs/configschema"
	"github.com/hashicorp/terraform/internal/providers"
	"github.com/hashicorp/terraform/internal/tfdiags"
)

var _ providers.Interface = (*mockProvider)(nil)

// mockProvider provides a mock implementation of providers.Interface that only
// implements the methods that are used by the refactoring package.
type mockProvider struct {
	moveResourceState bool
	moveResourceError error
}

func (provider *mockProvider) GetProviderSchema() providers.GetProviderSchemaResponse {
	return providers.GetProviderSchemaResponse{
		ResourceTypes: map[string]providers.Schema{
			"foo": {Body: &configschema.Block{}},
			"bar": {Body: &configschema.Block{}},
		},
		ServerCapabilities: providers.ServerCapabilities{
			MoveResourceState: provider.moveResourceState,
		},
	}
}

func (provider *mockProvider) GetResourceIdentitySchemas() providers.GetResourceIdentitySchemasResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateProviderConfig(providers.ValidateProviderConfigRequest) providers.ValidateProviderConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateResourceConfig(providers.ValidateResourceConfigRequest) providers.ValidateResourceConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateDataResourceConfig(providers.ValidateDataResourceConfigRequest) providers.ValidateDataResourceConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateListResourceConfig(providers.ValidateListResourceConfigRequest) providers.ValidateListResourceConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) UpgradeResourceState(providers.UpgradeResourceStateRequest) providers.UpgradeResourceStateResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) UpgradeResourceIdentity(providers.UpgradeResourceIdentityRequest) providers.UpgradeResourceIdentityResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ConfigureProvider(providers.ConfigureProviderRequest) providers.ConfigureProviderResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) Stop() error {
	panic("not implemented in mock")
}

func (provider *mockProvider) ReadResource(providers.ReadResourceRequest) providers.ReadResourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) PlanResourceChange(providers.PlanResourceChangeRequest) providers.PlanResourceChangeResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ApplyResourceChange(providers.ApplyResourceChangeRequest) providers.ApplyResourceChangeResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ImportResourceState(providers.ImportResourceStateRequest) providers.ImportResourceStateResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) MoveResourceState(providers.MoveResourceStateRequest) providers.MoveResourceStateResponse {
	if provider.moveResourceError != nil {
		return providers.MoveResourceStateResponse{
			Diagnostics: tfdiags.Diagnostics{
				tfdiags.Sourceless(tfdiags.Error, "expected error", provider.moveResourceError.Error()),
			},
		}
	}
	return providers.MoveResourceStateResponse{
		TargetState: cty.EmptyObjectVal,
	}
}

func (provider *mockProvider) ValidateEphemeralResourceConfig(providers.ValidateEphemeralResourceConfigRequest) providers.ValidateEphemeralResourceConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ReadDataSource(providers.ReadDataSourceRequest) providers.ReadDataSourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) OpenEphemeralResource(providers.OpenEphemeralResourceRequest) providers.OpenEphemeralResourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) RenewEphemeralResource(providers.RenewEphemeralResourceRequest) providers.RenewEphemeralResourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) CloseEphemeralResource(providers.CloseEphemeralResourceRequest) providers.CloseEphemeralResourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) CallFunction(providers.CallFunctionRequest) providers.CallFunctionResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ListResource(providers.ListResourceRequest) providers.ListResourceResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateStateStoreConfig(req providers.ValidateStateStoreConfigRequest) providers.ValidateStateStoreConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ConfigureStateStore(req providers.ConfigureStateStoreRequest) providers.ConfigureStateStoreResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) GetStates(req providers.GetStatesRequest) providers.GetStatesResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) DeleteState(req providers.DeleteStateRequest) providers.DeleteStateResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) PlanAction(providers.PlanActionRequest) providers.PlanActionResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) InvokeAction(_ providers.InvokeActionRequest) providers.InvokeActionResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) ValidateActionConfig(_ providers.ValidateActionConfigRequest) providers.ValidateActionConfigResponse {
	panic("not implemented in mock")
}

func (provider *mockProvider) Close() error {
	return nil // do nothing
}

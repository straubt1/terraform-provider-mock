// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure MockProvider satisfies various provider interfaces.
var _ provider.Provider = &MockProvider{}

// MockProvider defines the provider implementation.
type MockProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MockProviderModel describes the provider data model.
type MockProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *MockProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mock"
	resp.Version = p.version
}

func (p *MockProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Mock provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *MockProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// var data MockProviderModel

	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Mock client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *MockProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewMockResource,
	}
}

func (p *MockProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewMockDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MockProvider{
			version: version,
		}
	}
}

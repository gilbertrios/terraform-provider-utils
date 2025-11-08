package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var _ provider.Provider = &utilsProvider{}

// utilsProvider is the provider implementation.
type utilsProvider struct {
	version string
}

// New is a helper function to simplify provider server setup.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &utilsProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *utilsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "utils"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *utilsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A provider that provides utility functions for data manipulation and transformation in Terraform configurations.",
	}
}

// Configure prepares a provider for operation.
func (p *utilsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// This provider has no configuration
}

// DataSources defines the data sources implemented in the provider.
func (p *utilsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	// This is a function-only provider, so no data sources
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *utilsProvider) Resources(_ context.Context) []func() resource.Resource {
	// This is a function-only provider, so no resources
	return []func() resource.Resource{}
}

// Functions defines the functions implemented in the provider.
func (p *utilsProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewBase64EncodeFunction,
		NewBase64DecodeFunction,
		NewSHA256Function,
		NewMD5Function,
		NewUUIDv4Function,
		NewSlugifyFunction,
		NewTruncateFunction,
		NewReverseFunction,
		NewToUpperFunction,
		NewToLowerFunction,
		NewTrimFunction,
		NewJoinFunction,
		NewSplitFunction,
	}
}

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &cockpitProvider{}
)

type cockpitProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &cockpitProvider{
			version: version,
		}
	}
}

func (c cockpitProvider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "cockpit"
	response.Version = c.version
}

func (c cockpitProvider) Schema(ctx context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{}
}

func (c cockpitProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
}

func (c cockpitProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (c cockpitProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type cockpitProviderModel struct {
	APIEndpoint types.String `tfsdk:"api_endpoint"`
	APIToken    types.String `tfsdk:"api_token"`
}

func (c *cockpitProvider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "cockpit"
	response.Version = c.version
}

func (c *cockpitProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_endpoint": schema.StringAttribute{
				Required: true,
			},
			"api_token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (c *cockpitProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Cockpit CMS client")

	var config cockpitProviderModel
	diags := request.Config.Get(ctx, &config)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if config.APIEndpoint.IsUnknown() {
		response.Diagnostics.AddAttributeError(
			path.Root("api_endpoint"),
			"Unknown Cockpit CMS API Endpoint",
			"Provide a correct host endpoint or use the COCKPIT_API_ENDPOINT environment variable.",
		)
	}

	if config.APIToken.IsUnknown() {
		response.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown Cockpit CMS API Token",
			"Provide a correct token for the endpoint or use the COCKPIT_API_TOKEN environment variable.",
		)
	}

	if response.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("COCKPIT_API_ENDPOINT")
	token := os.Getenv("COCKPIT_API_TOKEN")

	if !config.APIEndpoint.IsNull() {
		endpoint = config.APIEndpoint.ValueString()
	}

	if !config.APIToken.IsNull() {
		token = config.APIToken.ValueString()
	}

	client, err := cockpitClient(&endpoint, &token)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to Create Cockpit CMS API Client",
			"An unexpected error occurred when creating the Cockpit CMS API client.\n\n"+
				"Client Error: "+err.Error(),
		)
		return
	}

	response.DataSourceData = client
	response.ResourceData = client
}

func (c *cockpitProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCollectionsDataSource,
	}
}

func (c *cockpitProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCollectionResource,
	}
}

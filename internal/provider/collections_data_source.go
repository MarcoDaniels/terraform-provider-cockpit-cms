package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &collectionsDataSource{}
	_ datasource.DataSourceWithConfigure = &collectionsDataSource{}
)

type collectionsDataSource struct {
	client *Client
}

type collectionsDataSourceModel struct {
	Collections []collectionModel `tfsdk:"collections"`
}

type collectionModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Label      types.String `tfsdk:"label"`
	Created    types.Int64  `tfsdk:"created"`
	Modified   types.Int64  `tfsdk:"modified"`
	ItemsCount types.Int64  `tfsdk:"items_count"`
	InMenu     types.Bool   `tfsdk:"in_menu"`
	Fields     []fieldModel `tfsdk:"fields"`
}

type fieldModel struct {
	Name  types.String `tfsdk:"name"`
	Label types.String `tfsdk:"label"`
	Type  types.String `tfsdk:"type"`
}

func NewCollectionsDataSource() datasource.DataSource {
	return &collectionsDataSource{}
}

func (c *collectionsDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*Client)
	if !ok {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *cockpit.Client, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)

		return
	}

	c.client = client
}

func (c *collectionsDataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_collections"
}

func (c *collectionsDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"collections": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id":          schema.StringAttribute{Computed: true},
						"name":        schema.StringAttribute{Computed: true},
						"label":       schema.StringAttribute{Computed: true},
						"created":     schema.Int64Attribute{Computed: true},
						"modified":    schema.Int64Attribute{Computed: true},
						"items_count": schema.Int64Attribute{Computed: true},
						"in_menu":     schema.BoolAttribute{Computed: true},
						"fields": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name":  schema.StringAttribute{Computed: true},
									"label": schema.StringAttribute{Computed: true},
									"type":  schema.StringAttribute{Computed: true},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (c *collectionsDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var state collectionsDataSourceModel
	collections, err := c.client.allCollections()
	if err != nil {
		response.Diagnostics.AddError("Unable to read Cockpit CMS collections", err.Error())
		return
	}

	for _, collection := range *collections {
		collectionState := collectionModel{
			ID:         types.StringValue(collection.Id),
			Name:       types.StringValue(collection.Name),
			Label:      types.StringValue(collection.Label),
			Created:    types.Int64Value(int64(collection.Created)),
			Modified:   types.Int64Value(int64(collection.Modified)),
			ItemsCount: types.Int64Value(int64(collection.ItemsCount)),
			InMenu:     types.BoolValue(collection.InMenu),
		}

		for _, field := range collection.Fields {
			collectionState.Fields = append(collectionState.Fields, fieldModel{
				Name:  types.StringValue(field.Name),
				Label: types.StringValue(field.Label),
				Type:  types.StringValue(field.Type),
			})
		}

		state.Collections = append(state.Collections, collectionState)
	}

	diags := response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

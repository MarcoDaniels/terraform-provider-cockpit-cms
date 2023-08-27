package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"
)

var (
	_ resource.Resource              = &collectionResource{}
	_ resource.ResourceWithConfigure = &collectionResource{}
)

type collectionResource struct {
	client *Client
}

type collectionResourceModel struct {
	Name   types.String         `tfsdk:"name"`
	Label  types.String         `tfsdk:"label"`
	Fields []fieldResourceModel `tfsdk:"fields"`
}

type fieldResourceModel struct {
	Name  types.String `tfsdk:"name"`
	Label types.String `tfsdk:"label"`
	Type  types.String `tfsdk:"type"`
}

func NewCollectionResource() resource.Resource {
	return &collectionResource{}
}

func (c *collectionResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (c *collectionResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_collection"
}

func (c *collectionResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name":  schema.StringAttribute{Required: true},
			"label": schema.StringAttribute{Optional: true},
			"fields": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name":  schema.StringAttribute{Required: true},
						"label": schema.StringAttribute{Optional: true},
						"type":  schema.StringAttribute{Required: true},
					},
				},
			},
		},
	}
}

func (c *collectionResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan collectionResourceModel
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	var fields []Field
	for _, field := range plan.Fields {
		fields = append(fields, Field{
			Name:  field.Name.ValueString(),
			Label: field.Label.ValueString(),
			Type:  field.Type.ValueString(),
		})
	}

	justNow := int(time.Now().Unix())
	name := plan.Name.ValueString()

	collection := Collection{
		Id:       name,
		Name:     plan.Name.ValueString(),
		Label:    plan.Label.ValueString(),
		Created:  justNow,
		Modified: justNow,
		Sort: Sort{
			Column: "_created",
			Dir:    -1,
		},
		Fields: fields,
	}

	// tflog.Info(ctx, fmt.Sprintf("Created new collection %v", collection))

	newCollection, err := c.client.createCollection(CreateCollection{Name: name, Data: collection})
	if err != nil {
		response.Diagnostics.AddError(
			"Error creating collection",
			"Could not create collection, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Name = types.StringValue(newCollection.Name)

	diags = response.State.Set(ctx, plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (c *collectionResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state collectionResourceModel
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	collection, err := c.client.getCollection(state.Name.ValueString())
	if err != nil {
		response.Diagnostics.AddError(
			"Error reading CockpitCMS Collection",
			"Could not read CockpitCMS collection id"+state.Name.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Name = types.StringValue(collection.Name)
	state.Label = types.StringValue(collection.Label)
	state.Fields = []fieldResourceModel{}
	for _, field := range collection.Fields {
		state.Fields = append(state.Fields, fieldResourceModel{
			Name:  types.StringValue(field.Name),
			Label: types.StringValue(field.Label),
			Type:  types.StringValue(field.Type),
		})
	}

	diags = response.State.Set(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (c *collectionResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan collectionResourceModel
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	var fields []Field
	for _, field := range plan.Fields {
		fields = append(fields, Field{
			Name:  field.Name.ValueString(),
			Label: field.Label.ValueString(),
			Type:  field.Type.ValueString(),
		})
	}

	justNow := int(time.Now().Unix())
	name := plan.Name.ValueString()

	collection := Collection{
		Id:       name,
		Name:     plan.Name.ValueString(),
		Label:    plan.Label.ValueString(),
		Modified: justNow,
		Fields:   fields,
	}

	// tflog.Info(ctx, fmt.Sprintf("Updating collection %v", collection))

	updated, err := c.client.updateCollection(plan.Name.ValueString(), UpdateCollection{Data: collection})
	if err != nil {
		response.Diagnostics.AddError(
			"Error updating collection",
			"Could not update collection, unexpected error: "+err.Error(),
		)
		return
	}

	plan.Name = types.StringValue(updated.Name)

	diags = response.State.Set(ctx, plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (c *collectionResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Info(ctx, "Deletion of collection is not available with Cockpit CMS API")

	var state collectionResourceModel
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	response.Diagnostics.AddError(
		"Error Deleting Cockpit CMS Collection",
		"Deletion of collection is not available with Cockpit CMS API",
	)

	return
}

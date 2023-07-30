package cockpit_cms

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCollectionCreate,
		ReadContext:   resourceCollectionRead,
		UpdateContext: resourceCollectionUpdate,
		DeleteContext: resourceCollectionDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	data := d.Get("data").(string)

	fields := make([]Field, 0, 1)

	field := Field{
		Name: data,
		Type: data,
	}

	fields = append(fields, field)

	enabled := false

	collection := Collection{
		Fields: fields,
		Sort: Sort{
			Column: "_created",
			Dir:    -1,
		},
		Rules: Rule{
			Create: RuleSet{Enabled: &enabled},
			Read:   RuleSet{Enabled: &enabled},
			Update: RuleSet{Enabled: &enabled},
			Delete: RuleSet{Enabled: &enabled},
		},
	}

	object, err := client.createCollection(CreateCollection{Name: name, Data: collection})
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to create collection.",
			Detail:   fmt.Sprintf("Unable to create collection with error: %s", err.Error()),
		})
	}

	d.SetId(object.Id)

	resourceCollectionRead(ctx, d, m)

	return diags
}

func flattenCollection(coll *Collection) interface{} {
	collection := make(map[string]interface{})
	fields := make([]interface{}, 0, len(coll.Fields))

	for _, f := range coll.Fields {
		field := make(map[string]interface{})
		field["name"] = f.Name
		field["type"] = f.Type

		fields = append(fields, field)
	}

	collection["_id"] = coll.Id
	collection["name"] = coll.Name
	collection["fields"] = fields
	collection["sort"] = coll.Sort

	/*
		rule := make(map[string]interface{})
		rule["enabled"] = false
		rules := make(map[string]interface{})
		rules["create"] = rule
		rules["read"] = rule
		rules["update"] = rule
		rules["delete"] = coll.Rules.Delete.Enabled
	*/

	collection["rules"] = coll.Rules

	return collection
}

func resourceCollectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	var diags diag.Diagnostics

	collection, err := client.getCollection(d.Id())
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to get collection.",
			Detail:   fmt.Sprintf("Unable to get collection with error: %s", err.Error()),
		})
	}

	if err := d.Set("name", &collection.Name); err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed set name.",
			Detail:   fmt.Sprintf("Unable to create collection with error: %s", err.Error()),
		})
	}

	if err := d.Set("data", &collection.Id); err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed set data collection.",
			Detail:   fmt.Sprintf("Unable to create collection with error: %s", err.Error()),
		})
	}

	return diags
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCollectionRead(ctx, d, m)
}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

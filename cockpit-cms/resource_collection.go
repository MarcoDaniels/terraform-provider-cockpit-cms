package cockpit_cms

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
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
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fields": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"label": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					}},
			},
		},
	}
}

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	data := d.Get("data").([]interface{})[0].(map[string]interface{})

	fields := make([]Field, 0, len(data["fields"].([]interface{})))

	for _, f := range data["fields"].([]interface{}) {
		field := f.(map[string]interface{})

		fields = append(fields, Field{
			Name:  field["name"].(string),
			Type:  field["type"].(string),
			Label: field["label"].(string),
		})
	}

	justNow := int(time.Now().Unix())

	collection := Collection{
		Id:       name,
		Name:     name,
		Fields:   fields,
		Created:  justNow,
		Modified: justNow,
		Sort: Sort{
			Column: "_created",
			Dir:    -1,
		},
	}

	newCollection, err := client.createCollection(CreateCollection{Name: name, Data: collection})
	if err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to create collection.",
			Detail:   fmt.Sprintf("Unable to create collection with error: %s", err.Error()),
		})
	}

	d.SetId(newCollection.Id)

	resourceCollectionRead(ctx, d, m)

	return diags
}

func flattenCollection(c *Collection) []interface{} {
	data := make([]interface{}, 0, 1)
	collection := make(map[string]interface{})
	fields := make([]interface{}, 0, len(c.Fields))

	for _, f := range c.Fields {
		field := make(map[string]interface{})
		field["name"] = f.Name
		field["type"] = f.Type
		field["label"] = f.Label

		fields = append(fields, field)
	}

	collection["fields"] = fields

	data = append(data, collection)

	return data
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

	if err := d.Set("data", flattenCollection(collection)); err != nil {
		return append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed set data collection.",
			Detail:   fmt.Sprintf("Unable to set data collection with error: %s", err.Error()),
		})
	}

	return diags
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	var diags diag.Diagnostics

	name := d.Id()

	if d.HasChange("data") {
		data := d.Get("data").([]interface{})[0].(map[string]interface{})

		fields := make([]Field, 0, len(data["fields"].([]interface{})))

		for _, f := range data["fields"].([]interface{}) {
			field := f.(map[string]interface{})

			fields = append(fields, Field{
				Name:  field["name"].(string),
				Type:  field["type"].(string),
				Label: field["label"].(string),
			})
		}

		justNow := int(time.Now().Unix())

		collection := Collection{
			Id:       name,
			Name:     name,
			Fields:   fields,
			Modified: justNow,
			Sort: Sort{
				Column: "_created",
				Dir:    -1,
			},
		}

		_, err := client.updateCollection(name, UpdateCollection{Data: collection})
		if err != nil {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to update collection.",
				Detail:   fmt.Sprintf("Unable to update collection with error: %s", err.Error()),
			})
		}
	}

	return resourceCollectionRead(ctx, d, m)
}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

package cockpit_cms

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceCollections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCollectionsRead,
		Schema: map[string]*schema.Schema{
			"collections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sortable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"_created": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"_modified": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"in_menu": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"items_count": {
							Type:     schema.TypeInt,
							Computed: true,
							Required: false,
						},
						"acl": {
							Type:     schema.TypeList,
							Computed: true,
							Required: false,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"rules": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"create": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"read": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"update": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"delete": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"sort": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"column": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"dir": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"fields": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"width": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"info": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"localize": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lst": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"options": {
										Type:     schema.TypeList,
										Computed: true,
										Required: false,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"acl": {
										Type:     schema.TypeList,
										Computed: true,
										Required: false,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func flattenCollections(nestedCollections *map[string]Collection) []interface{} {
	if nestedCollections != nil {
		collections := make([]interface{}, 0, len(*nestedCollections))

		for _, coll := range *nestedCollections {
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

			collections = append(collections, collection)
		}

		return collections
	}

	return make([]interface{}, 0)
}

func dataSourceCollectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	var diags diag.Diagnostics

	result, err := client.allCollections()
	if err != nil {
		return diag.FromErr(err)
	}

	collections := make(map[string]Collection, 0)
	if err := json.Unmarshal(result, &collections); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("collections", flattenCollections(&collections)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

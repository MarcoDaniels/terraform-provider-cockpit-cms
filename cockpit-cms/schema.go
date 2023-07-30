package cockpit_cms

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func fields() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"label": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"type": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"options": {
					Type:     schema.TypeList,
					Computed: true,
					Required: false,
					Elem:     &schema.Schema{Type: schema.TypeString},
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
				"acl": {
					Type:     schema.TypeList,
					Computed: true,
					Required: false,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

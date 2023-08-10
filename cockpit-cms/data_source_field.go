package cockpit_cms

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceField() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			var diags diag.Diagnostics

			// TODO: handle proper
			_ = d.Set("name", d.Get("name"))
			_ = d.Set("label", d.Get("label"))
			_ = d.Set("type", d.Get("type"))

			d.SetId(d.Get("name").(string))

			return diags
		},
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
	}
}

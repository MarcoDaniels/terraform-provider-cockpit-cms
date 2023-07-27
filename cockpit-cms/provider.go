package cockpit_cms

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COCKPIT_BASE_URL", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("COCKPIT_API_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"cockpit-cms_collections": dataSourceCollections(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
			baseUrl := d.Get("base_url").(string)
			token := d.Get("token").(string)

			var diags diag.Diagnostics

			client, err := cockpitClient(baseUrl, token)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			return client, diags
		},
	}
}

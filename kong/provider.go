package kong

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KONG_ADMIN_ENDPOINT", nil),
				Description: "On which the Admin API used to configure Kong listens.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"kong_api": resourceKongAPI(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Endpoint: d.Get("endpoint").(string),
	}

	return config.Client()
}
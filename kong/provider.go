package kong

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint":{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KONG_ADMIN_ENDPOINT", nil),
				Description: "On which the Admin API used to configure Kong listens.",
			},
		},
	}
}
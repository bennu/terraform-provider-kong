package kong

import (
	"context"
	"fmt"

	"github.com/bennu/kong-go/client"
	"github.com/bennu/kong-go/types"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKongAPI() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongAPICreate,
		Read:   resourceKongAPIRead,
		Delete: resourceKongAPIDelete,
		Update: resourceKongAPIUpdate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"uris": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: false,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"upstream_url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceKongAPICreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	req := types.APICreate{}

	if v, ok := d.GetOk("uris"); ok {
		var uris []string
		for _, tag := range v.([]interface{}) {
			uris = append(uris, tag.(string))
		}
		req.URIS = uris
	}
	req.UpstreamURL = d.Get("upstream_url").(string)

	api, err := client.APICreate(context.Background(), d.Get("name").(string), req)
	if err != nil {
		return fmt.Errorf("Error creating API: %s", err)
	}

	d.SetId(api.Id)

	return resourceKongAPIRead(d, meta)
}

func resourceKongAPIRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)

	_, err := client.APILookup(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error retrieving API: %s", err)
	}

	return nil
}

func resourceKongAPIDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)

	if err := client.APIDelete(context.Background(), d.Id()); err != nil {
		return err
	}
	return nil
}

func resourceKongAPIUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*client.Client)
	req := types.APIUpdate{}

	if v, ok := d.GetOk("uris"); ok {
		var uris []string
		for _, tag := range v.([]interface{}) {
			uris = append(uris, tag.(string))
		}
		req.URIS = uris
	}
	req.UpstreamURL = d.Get("upstream_url").(string)
	req.Name = d.Get("name").(string)


	api, err := client.APIUpdate(context.Background(), d.Id(), req)
	if err != nil {
		return fmt.Errorf("Error creating API: %s", err)
	}

	d.SetId(api.Id)

	return resourceKongAPIRead(d, meta)
}

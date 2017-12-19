package main

import (
	"github.com/bennu/terraform-provider-kong/kong"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kong.Provider})
}

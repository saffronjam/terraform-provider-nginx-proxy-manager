package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NPM_URL", ""),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NPM_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NPM_PASSWORD", ""),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"npm_proxy_host": resourceProxyHost(),
		},
		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	npmClient, err := client.New(&client.Config{
		ApiUrl:   url,
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return npmClient, nil
}

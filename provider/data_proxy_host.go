package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataProxyHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"forward_host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"certificate_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssl_forced": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_websocket_upgrade": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"forward_scheme": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		ReadContext: dataProxyHostRead,
	}
}

func dataProxyHostRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	domainName := d.Get("domain_name").(string)
	proxyHost, err := m.(*client.Client).ReadProxyHostByDomainName(domainName)
	if err != nil {
		return diag.FromErr(err)
	}

	if proxyHost == nil {
		return diag.FromErr(fmt.Errorf("proxy host with domain name %s not found", domainName))
	}

	d.SetId(fmt.Sprintf("%s/%s", ProxyHostPath, strconv.Itoa(proxyHost.ID)))
	_ = d.Set("domain_names", proxyHost.DomainNames)
	_ = d.Set("forward_host", proxyHost.ForwardHost)
	_ = d.Set("forward_port", proxyHost.ForwardPort)
	_ = d.Set("certificate_id", proxyHost.CertificateID)
	_ = d.Set("ssl_forced", client.IntToBool(proxyHost.SslForced))
	_ = d.Set("allow_websocket_upgrade", client.IntToBool(proxyHost.AllowWebsocketUpgrade))
	_ = d.Set("forward_scheme", proxyHost.ForwardScheme)
	_ = d.Set("enabled", client.IntToBool(proxyHost.Enabled))

	return nil
}

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"strconv"
)

func resourceProxyHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"domain_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"forward_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"forward_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"certificate_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssl_forced": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_websocket_upgrade": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"forward_scheme": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		CreateContext: resourceProxyHostCreate,
		ReadContext:   resourceProxyHostRead,
		UpdateContext: resourceProxyHostUpdate,
		DeleteContext: resourceProxyHostDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceProxyHostCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	proxyHost, err := m.(*client.Client).CreateProxyHost(d)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(proxyHost.ID))
	return nil
}

func resourceProxyHostRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	proxyHost, err := m.(*client.Client).ReadProxyHost(d)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(proxyHost.ID))
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

func resourceProxyHostUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := m.(*client.Client).UpdateProxyHost(d)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProxyHostRead(ctx, d, m)
}

func resourceProxyHostDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	err := m.(*client.Client).DeleteProxyHost(d)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"strconv"
)

const ProxyHostPath = "/nginx/proxy-hosts"

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
			"forward_scheme": {
				Type:     schema.TypeString,
				Required: true,
			},
			"forward_port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"certificate_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssl_forced": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_websocket_upgrade": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"block_exploits": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"caching_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"http2_support": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hsts_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"hsts_subdomains": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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

	d.SetId(fmt.Sprintf("%s/%s", ProxyHostPath, strconv.Itoa(proxyHost.ID)))
	return nil
}

func resourceProxyHostRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	proxyHost, err := m.(*client.Client).ReadProxyHost(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if proxyHost == nil {
		return diag.FromErr(fmt.Errorf("proxy host with id %s not found", d.Id()))
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
	_ = d.Set("block_exploits", client.IntToBool(proxyHost.BlockExploits))
	_ = d.Set("caching_enabled", client.IntToBool(proxyHost.CachingEnabled))
	_ = d.Set("http2_support", client.IntToBool(proxyHost.HTTP2Support))
	_ = d.Set("hsts_enabled", client.IntToBool(proxyHost.HstsEnabled))
	_ = d.Set("hsts_subdomains", client.IntToBool(proxyHost.HstsSubdomains))

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

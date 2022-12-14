package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"strconv"
)

func resourceProxyHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
		Create: resourceProxyHostCreate,
		Read:   resourceProxyHostRead,
		Update: resourceProxyHostUpdate,
		Delete: resourceProxyHostDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceProxyHostCreate(d *schema.ResourceData, m interface{}) error {
	proxyHost, err := m.(*client.Client).CreateProxyHost(d)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(proxyHost.ID))
	return resourceProxyHostRead(d, m)
}

func resourceProxyHostRead(d *schema.ResourceData, m interface{}) error {
	proxyHost, err := m.(*client.Client).ReadProxyHost(d)
	if err != nil {
		return err
	}

	_ = d.Set("id", proxyHost.ID)
	_ = d.Set("domain_names", proxyHost.DomainNames)
	_ = d.Set("certificate_id", proxyHost.CertificateID)
	_ = d.Set("enabled", proxyHost.Enabled)
	_ = d.Set("created_on", proxyHost.CreatedOn)
	_ = d.Set("forward_scheme", proxyHost.ForwardScheme)
	_ = d.Set("forward_port", proxyHost.ForwardPort)
	_ = d.Set("forward_host", proxyHost.ForwardHost)
	_ = d.Set("ssl_forced", proxyHost.SslForced)
	_ = d.Set("allow_websocket_upgrade", proxyHost.AllowWebsocketUpgrade)

	return nil
}

func resourceProxyHostUpdate(d *schema.ResourceData, m interface{}) error {
	err := m.(*client.Client).UpdateProxyHost(d)
	if err != nil {
		return err
	}

	return resourceProxyHostRead(d, m)
}

func resourceProxyHostDelete(d *schema.ResourceData, m interface{}) error {
	err := m.(*client.Client).DeleteProxyHost(d)
	if err != nil {
		return err
	}

	return nil
}

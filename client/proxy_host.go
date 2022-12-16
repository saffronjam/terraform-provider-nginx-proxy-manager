package client

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/models"
	"net/http"
)

func (client *Client) ReadProxyHostByDomainName(domainName string) (*models.ProxyHostRead, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch proxy host by domain name %s. details: %s", domainName, err)
	}

	res, err := client.doRequest("GET", fmt.Sprintf("/nginx/proxy-hosts/"))
	if err != nil {
		return nil, makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return nil, makeApiError(res.Body, makeError)
	}

	var proxyHosts []models.ProxyHostRead
	err = ParseBody(res.Body, &proxyHosts)
	if err != nil {
		return nil, makeError(err)
	}

	for _, proxyHost := range proxyHosts {
		for _, name := range proxyHost.DomainNames {
			if name == domainName {
				return &proxyHost, nil
			}
		}
	}

	return nil, nil
}

func (client *Client) ReadProxyHost(id string) (*models.ProxyHostRead, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch proxy host %s. details: %s", id, err)
	}

	res, err := client.doRequest("GET", id)
	if err != nil {
		return nil, makeError(err)
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return nil, makeApiError(res.Body, makeError)
	}

	var proxyHost models.ProxyHostRead
	err = ParseBody(res.Body, &proxyHost)
	if err != nil {
		return nil, makeError(err)
	}

	return &proxyHost, nil
}

func (client *Client) CheckIfProxyHostExists(id string) (bool, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to check if proxy host %s exists. details: %s", id, err)
	}

	res, err := client.doRequest("GET", id)
	if err != nil {
		return false, makeError(err)
	}

	if res.StatusCode == http.StatusNotFound {
		return false, nil
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return false, makeApiError(res.Body, makeError)
	}

	return true, nil
}

func (client *Client) CreateProxyHost(d *schema.ResourceData) (*models.ProxyHostCreated, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm proxy host. details: %s", err)
	}

	domainNamesInterface := d.Get("domain_names").([]interface{})
	domainNames := make([]string, len(domainNamesInterface))
	for i := range domainNames {
		domainNames[i] = domainNamesInterface[i].(string)
	}

	requestBody := models.ProxyHostCreate{
		DomainNames:           domainNames,
		ForwardHost:           d.Get("forward_host").(string),
		ForwardPort:           d.Get("forward_port").(int),
		AccessListID:          0,
		CertificateID:         d.Get("certificate_id").(int),
		SslForced:             d.Get("ssl_forced").(bool),
		CachingEnabled:        d.Get("caching_enabled").(bool),
		BlockExploits:         d.Get("block_exploits").(bool),
		AdvancedConfig:        "",
		AllowWebsocketUpgrade: d.Get("allow_websocket_upgrade").(bool),
		HTTP2Support:          d.Get("http2_support").(bool),
		ForwardScheme:         d.Get("forward_scheme").(string),
		Enabled:               d.Get("enabled").(bool),
		Locations:             []models.Location{},
		HstsEnabled:           d.Get("hsts_enabled").(bool),
		HstsSubdomains:        d.Get("hsts_subdomains").(bool),
	}

	res, err := client.doJSONRequest("POST", "/nginx/proxy-hosts", requestBody)
	if err != nil {
		return nil, makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return nil, makeApiError(res.Body, makeError)
	}

	var proxyHost models.ProxyHostCreated
	err = ParseBody(res.Body, &proxyHost)
	if err != nil {
		return nil, makeError(err)
	}

	return &proxyHost, nil
}

func (client *Client) DeleteProxyHost(d *schema.ResourceData) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete npm proxy host. details: %s", err)
	}

	res, err := client.doRequest("DELETE", d.Id())
	if err != nil {
		return makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return makeApiError(res.Body, makeError)
	}

	return nil
}

func (client *Client) UpdateProxyHost(d *schema.ResourceData) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update npm proxy host. details: %s", err)
	}

	meta := models.MetaData{DNSChallenge: false, LetsencryptAgree: false}

	domainNamesInterface := d.Get("domain_names").([]interface{})
	domainNames := make([]string, len(domainNamesInterface))
	for i := range domainNames {
		domainNames[i] = domainNamesInterface[i].(string)
	}

	requestBody := models.ProxyHostUpdate{
		ForwardScheme:         d.Get("forward_scheme").(string),
		ForwardHost:           d.Get("forward_host").(string),
		ForwardPort:           d.Get("forward_port").(int),
		AdvancedConfig:        "",
		DomainNames:           domainNames,
		AccessListID:          0,
		CertificateID:         d.Get("certificate_id").(int),
		SslForced:             d.Get("ssl_forced").(bool),
		Meta:                  meta,
		Locations:             []models.Location{},
		BlockExploits:         d.Get("block_exploits").(bool),
		CachingEnabled:        d.Get("caching_enabled").(bool),
		AllowWebsocketUpgrade: d.Get("allow_websocket_upgrade").(bool),
		HTTP2Support:          d.Get("http2_support").(bool),
		HstsEnabled:           d.Get("hsts_enabled").(bool),
		HstsSubdomains:        d.Get("hsts_subdomains").(bool),
	}

	res, err := client.doJSONRequest("PUT", d.Id(), requestBody)
	if err != nil {
		return makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return makeApiError(res.Body, makeError)
	}

	return nil
}

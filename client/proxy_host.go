package client

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/models"
	"net/http"
)

func (client *Client) ReadProxyHost(d *schema.ResourceData) (*models.ProxyHostRead, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch proxy host %s. details: %s", d.Id(), err)
	}

	res, err := client.doRequest("GET", fmt.Sprintf("/nginx/proxy-hosts/%s", d.Id()))
	if err != nil {
		return nil, makeError(err)
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

	res, err := client.doRequest("GET", fmt.Sprintf("/nginx/proxy-hosts/%s", id))
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

func (client *Client) CreateProxyHost(d *schema.ResourceData) (*models.ProxyHostRead, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm proxy host. details: %s", err)
	}

	domainNamesInteface := d.Get("domain_names").([]interface{})
	domainNames := make([]string, len(domainNamesInteface))
	for i := range domainNames {
		domainNames[i] = domainNamesInteface[i].(string)
	}

	requestBody := models.ProxyHostCreate{
		DomainNames:           domainNames,
		ForwardHost:           d.Get("forward_host").(string),
		ForwardPort:           d.Get("forward_port").(int),
		AccessListID:          0,
		CertificateID:         d.Get("certificate_id").(int),
		SslForced:             d.Get("ssl_forced").(bool),
		CachingEnabled:        false,
		BlockExploits:         false,
		AdvancedConfig:        "",
		AllowWebsocketUpgrade: d.Get("allow_websocket_upgrade").(bool),
		HTTP2Support:          false,
		ForwardScheme:         d.Get("forward_scheme").(string),
		Enabled:               d.Get("enabled").(bool),
		Locations:             nil,
		HstsEnabled:           false,
		HstsSubdomains:        0,
	}
	res, err := client.doJSONRequest("POST", "/nginx/proxy-hosts", requestBody)
	if err != nil {
		return nil, makeError(err)
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

func (client *Client) DeleteProxyHost(d *schema.ResourceData) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete npm proxy host. details: %s", err)
	}

	res, err := client.doRequest("DELETE", fmt.Sprintf("/nginx/proxy-hosts/%s", d.Id()))
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

	requestBody := models.ProxyHostUpdate{
		ForwardScheme:         d.Get("forward_scheme").(string),
		ForwardHost:           d.Get("forward_host").(string),
		ForwardPort:           d.Get("forward_port").(int),
		AdvancedConfig:        "",
		DomainNames:           d.Get("domain_names").([]string),
		AccessListID:          "",
		CertificateID:         d.Get("certificate_id").(int),
		SslForced:             d.Get("ssl_forced").(bool),
		Meta:                  meta,
		Locations:             nil,
		BlockExploits:         d.Get("block_exploits").(bool),
		CachingEnabled:        d.Get("caching_enabled").(bool),
		AllowWebsocketUpgrade: d.Get("allow_websocket_upgrade").(bool),
		HTTP2Support:          d.Get("http2_support").(bool),
		HstsEnabled:           d.Get("hsts_enabled").(bool),
		HstsSubdomains:        false,
	}

	res, err := client.doJSONRequest("PUT", fmt.Sprintf("/nginx/proxy-hosts/%s", d.Id()), requestBody)
	if err != nil {
		return makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return makeApiError(res.Body, makeError)
	}

	return nil
}

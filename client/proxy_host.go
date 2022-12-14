package client

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/models"
)

func (client *Client) ReadProxyHost(d *schema.ResourceData) (*models.ProxyHost, error) {
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

	var proxyHost models.ProxyHost
	err = ParseBody(res.Body, &proxyHost)
	if err != nil {
		return nil, makeError(err)
	}

	return &proxyHost, nil
}

func (client *Client) CreateProxyHost(d *schema.ResourceData) (*models.ProxyHost, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm proxy host. details: %s", err)
	}

	requestBody := models.ProxyHostCreateReq{
		DomainNames:           d.Get("domain_names").([]string),
		ForwardHost:           d.Get("forward_host").(string),
		ForwardPort:           d.Get("forward_port").(int),
		AccessListID:          0,
		CertificateID:         d.Get("certificate_id").(int),
		SslForced:             boolToInt(d.Get("ssl_forced").(bool)),
		CachingEnabled:        0,
		BlockExploits:         0,
		AdvancedConfig:        "",
		AllowWebsocketUpgrade: 1,
		HTTP2Support:          0,
		ForwardScheme:         d.Get("forward_scheme").(string),
		Enabled:               boolToInt(d.Get("enabled").(bool)),
		Locations:             nil,
		HstsEnabled:           0,
		HstsSubdomains:        0,
	}
	res, err := client.doJSONRequest("POST", "/nginx/proxy-hosts", requestBody)
	if err != nil {
		return nil, makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return nil, makeApiError(res.Body, makeError)
	}

	var proxyHost models.ProxyHost
	err = ParseBody(res.Body, &proxyHost)
	if err != nil {
		return nil, makeError(err)
	}

	return &proxyHost, nil
}

func (client *Client) DeleteProxyHost(d *schema.ResourceData) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm proxy host. details: %s", err)
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
	return nil
}

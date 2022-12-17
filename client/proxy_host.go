package client

import (
	"fmt"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/models"
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

func (client *Client) ReadProxyHost(id int) (*models.ProxyHostRead, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to fetch proxy host %d. details: %s", id, err)
	}

	res, err := client.doRequest("GET", fmt.Sprintf("/nginx/proxy-hosts/%d", id))
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

func (client *Client) CreateProxyHost(requestBody *models.ProxyHostCreate) (*models.ProxyHostCreated, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm proxy host. details: %s", err)
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

func (client *Client) DeleteProxyHost(id int) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to delete npm proxy host. details: %s", err)
	}

	res, err := client.doRequest("DELETE", fmt.Sprintf("/nginx/proxy-hosts/%d", id))
	if err != nil {
		return makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return makeApiError(res.Body, makeError)
	}

	return nil
}

func (client *Client) UpdateProxyHost(id int, requestBody *models.ProxyHostUpdate) error {
	makeError := func(err error) error {
		return fmt.Errorf("failed to update npm proxy host. details: %s", err)
	}

	res, err := client.doJSONRequest("PUT", fmt.Sprintf("/nginx/proxy-hosts/%d", id), requestBody)
	if err != nil {
		return makeError(err)
	}

	if !IsGoodStatusCode(res.StatusCode) {
		return makeApiError(res.Body, makeError)
	}

	return nil
}

package client

import (
	"fmt"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/models"
)

type Client struct {
	apiUrl string
	token  string
}

type Config struct {
	ApiUrl   string
	Username string
	Password string
}

func New(config *Config) (*Client, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create npm client. details: %s", err)
	}

	client := Client{
		apiUrl: config.ApiUrl,
	}

	token, err := client.createToken(config.Username, config.Password)
	if err != nil {
		return nil, makeError(err)
	}

	client.token = token

	return &client, nil
}

func (client *Client) createToken(username, password string) (string, error) {
	makeError := func(err error) error {
		return fmt.Errorf("failed to create token. details: %s", err)
	}

	tokenReq := models.TokenReq{Identity: username, Secret: password}
	res, err := client.doJSONRequestUnauthorized("POST", "/tokens", tokenReq)
	if err != nil {
		return "", makeError(err)
	}

	// check if good request
	if !IsGoodStatusCode(res.StatusCode) {
		return "", makeApiError(res.Body, makeError)
	}

	var token models.Token
	err = ParseBody(res.Body, &token)
	if err != nil {
		return "", makeError(err)
	}

	return token.Token, nil

}

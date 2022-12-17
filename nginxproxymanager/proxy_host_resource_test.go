package nginxproxymanager

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const resourceName = "nginxproxymanager_proxy_host.web"

func TestCreateProxyHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckProxyHost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "forward_host", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "forward_scheme", "http"),
				),
			},
			{
				Config: testAccCheckProxyHostUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "forward_host", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "forward_scheme", "http"),
				),
			},
		},
	})
}

func testAccCheckProxyHost() string {
	return providerConfig +
		`resource "nginxproxymanager_proxy_host" "web" {
			domain_names            = ["test.com"]
			forward_host            = "1.1.1.1"
			forward_port            = 8080
			certificate_id          = 0
			ssl_forced              = false
			allow_websocket_upgrade = false
			forward_scheme          = "http"
			enabled                 = false
		}`
}

func testAccCheckProxyHostUpdate() string {
	return providerConfig +
		`resource "nginxproxymanager_proxy_host" "web" {
			domain_names            = ["test.com"]
			forward_host            = "2.2.2.2"
			forward_port            = 8080
			certificate_id          = 0
			ssl_forced              = false
			allow_websocket_upgrade = false
			forward_scheme          = "http"
			enabled                 = false
		}`
}

func testAccCheckProxyHostDataSource() string {
	return providerConfig +
		`resource "nginxproxymanager_proxy_host" "web" {
			domain_names            = ["test.com"]
			forward_host            = "1.1.1.1"
			forward_port            = 8080
			certificate_id          = 0
			ssl_forced              = false
			allow_websocket_upgrade = false
			forward_scheme          = "http"
			enabled                 = false
		}

		data "npm_proxy_host" "imported" {
			domain_name            = "test.com"
		}`
}

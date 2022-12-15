package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"testing"
)

const resourceName = "npm_proxy_host.web"

func testAccCheckProjectDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		exists, err := apiClient.CheckIfProxyHostExists(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("failed to check if resource was not deleted, assuming not. details: %s", err)

		}
		if exists {
			return fmt.Errorf("resource was not deleted: %s", rs.Primary.ID)
		}
	}

	return nil
}

func TestCreateProxyHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckProxyHost(),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "forward_host", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "forward_scheme", "http"),
				),
			},
		},
	})
}

func testAccCheckProxyHost() string {
	return `resource "npm_proxy_host" "web" {
				domain_names            = ["onlyfortesting.dev.kthcloud.com"]
				forward_host            = "1.1.1.1"
				forward_port            = 8080
				certificate_id          = 0
				ssl_forced              = false
				allow_websocket_upgrade = false
				forward_scheme          = "http"
				enabled                 = true
			}`
}

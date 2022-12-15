package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/client"
	"os"
	"testing"
)

var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()

	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"npm": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func testAccPreCheck(t *testing.T) {
	if _, result := os.LookupEnv("NPM_URL"); !result {
		t.Fatal("NPM_URL must be set for acceptance test")
	}
	if _, result := os.LookupEnv("NPM_USERNAME"); !result {
		t.Fatal("NPM_USERNAME must be set for acceptance test")
	}
	if _, result := os.LookupEnv("NPM_PASSWORD"); !result {
		t.Fatal("NPM_PASSWORD must be set for acceptance test")
	}
}

func testAccResourceExists(res string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[res]
		if !ok {
			return fmt.Errorf("not found %s", res)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no record ID is set")
		}

		name := rs.Primary.ID
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err := apiClient.CheckIfProxyHostExists(name)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", res, err)
		}

		return nil
	}
}

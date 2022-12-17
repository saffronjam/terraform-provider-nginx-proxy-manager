package nginxproxymanager

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"os"
	"testing"
)

const providerConfig = "provider \"nginxproxymanager\" {}\n"

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"nginxproxymanager": providerserver.NewProtocol6WithError(New()),
}

func testAccPreCheck(t *testing.T) {
	requiredEnvs := []string{
		"NGINX_PROXY_MANAGER_URL", "NGINX_PROXY_MANAGER_USERNAME", "NGINX_PROXY_MANAGER_PASSWORD",
	}

	for _, env := range requiredEnvs {
		if _, result := os.LookupEnv(env); !result {
			t.Fatalf("%s must be set for acceptance test", env)
		}
	}
}

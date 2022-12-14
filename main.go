package terraform_provider_nginx_proxy_manager

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/saffronjam/terraform-provider-nginx-proxy-manager/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider})
}

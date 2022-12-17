package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/nginxproxymanager"
)

func main() {
	_ = providerserver.Serve(context.Background(), nginxproxymanager.New, providerserver.ServeOpts{
		Address: "nginxproxymanager",
	})
}

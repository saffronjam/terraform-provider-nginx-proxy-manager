package nginxproxymanager

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/client"
	"os"
)

type nginxproxymanagerProviderModel struct {
	Url      types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &nginxproxymanagerProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &nginxproxymanagerProvider{}
}

// nginxproxymanagerProvider is the provider implementation.
type nginxproxymanagerProvider struct{}

// Metadata returns the provider type name.
func (p *nginxproxymanagerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nginxproxymanager"
}

// Schema defines the provider-level schema for configuration data.
func (p *nginxproxymanagerProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Nginx Proxy Manager.",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Description: "URL for Nginx Proxy Manager API. May also be provided via NGINX_PROXY_MANAGER_URL environment variable.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for Nginx Proxy Manager API. May also be provided via NGINX_PROXY_MANAGER_USERNAME environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for Nginx Proxy Manager API. May also be provided via NGINX_PROXY_MANAGER_PASSWORD environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a Nginx Proxy Manager API client for data sources and resources.
func (p *nginxproxymanagerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var config nginxproxymanagerProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	url := os.Getenv("NGINX_PROXY_MANAGER_URL")
	username := os.Getenv("NGINX_PROXY_MANAGER_USERNAME")
	password := os.Getenv("NGINX_PROXY_MANAGER_PASSWORD")

	if !config.Url.IsNull() {
		url = config.Url.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	npmClient, err := client.New(&client.Config{
		ApiUrl:   url,
		Username: username,
		Password: password,
	})

	if err != nil {
		resp.Diagnostics.AddError("failed to create api client", fmt.Sprintf("%s", err))
		return
	}

	ctx = tflog.SetField(ctx, "nginxproxymanager_host", url)
	ctx = tflog.SetField(ctx, "nginxproxymanager_username", username)
	ctx = tflog.SetField(ctx, "nginxproxymanager_password", password)

	resp.DataSourceData = npmClient
	resp.ResourceData = npmClient
}

// DataSources defines the data sources implemented in the provider.
func (p *nginxproxymanagerProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *nginxproxymanagerProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProxyHostResource,
	}
}

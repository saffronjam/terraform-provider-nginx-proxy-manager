package nginxproxymanager

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/client"
	resourceModels "github.com/saffronjam/terraform-provider-nginxproxymanager/nginxproxymanager/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &proxyHostResource{}
	_ resource.ResourceWithConfigure = &proxyHostResource{}
)

// NewProxyHostResource is a helper function to simplify the provider implementation.
func NewProxyHostResource() resource.Resource {
	return &proxyHostResource{}
}

// proxyHostResource is the resource implementation.
type proxyHostResource struct {
	apiClient *client.Client
}

// Configure adds the provider configured client to the resource.
func (resource *proxyHostResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	resource.apiClient = req.ProviderData.(*client.Client)
}

// Metadata returns the resource type name.
func (resource *proxyHostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_order"
}

// Schema defines the schema for the resource.
func (resource *proxyHostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"domain_names": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"forward_host": schema.StringAttribute{
				Description: "internal forward host",
				Required:    true,
			},
			"forward_scheme": schema.StringAttribute{
				Required: true,
			},
			"forward_port": schema.Int64Attribute{
				Required: true,
			},
			"certificate_id": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"ssl_forced": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"allow_websocket_upgrade": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"block_exploits": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"caching_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"http2_support": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"hsts_enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"hsts_subdomains": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (resource *proxyHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceModels.ProxyHostResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	requestBody := resourceModels.CreateProxyHostCreateModel(&plan)

	proxyHost, err := resource.apiClient.CreateProxyHost(requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating proxy host",
			"Could not create proxy-host, unexpected error: "+err.Error(),
		)
		return
	}

	resourceModel := resourceModels.CreateResourceModelCreated(proxyHost)

	diags = resp.State.Set(ctx, resourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (resource *proxyHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resourceModels.ProxyHostResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	proxyHost, err := resource.apiClient.ReadProxyHost(int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading proxy host",
			fmt.Sprintf("Could not read proxy-host %d, unexpected error: %s", state.ID, err.Error()),
		)
		return
	}

	resourceModel := resourceModels.CreateResourceModelRead(proxyHost)

	diags = resp.State.Set(ctx, resourceModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (resource *proxyHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceModels.ProxyHostResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := int(plan.ID.ValueInt64())
	requestBody := resourceModels.CreateProxyHostUpdateModel(&plan)
	err := resource.apiClient.UpdateProxyHost(id, requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating proxy host",
			fmt.Sprintf("Could not update proxy-host %d, unexpected error: %s", id, err.Error()),
		)
		return
	}

	proxyHost, err := resource.apiClient.ReadProxyHost(id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading proxy host when updating proxy host",
			fmt.Sprintf("Could not read proxy-host %d, unexpected error: %s", id, err.Error()),
		)
		return
	}

	diags = resp.State.Set(ctx, proxyHost)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (resource *proxyHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceModels.ProxyHostResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := resource.apiClient.DeleteProxyHost(int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting proxy host",
			fmt.Sprintf("Could not delete proxy-host %d, unexpected error: %s", state.ID, err.Error()),
		)
		return
	}
}

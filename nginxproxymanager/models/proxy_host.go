package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/client"
	"github.com/saffronjam/terraform-provider-nginxproxymanager/models"
)

type Location struct {
	Path           string `tfsdk:"path"`
	AdvancedConfig string `tfsdk:"advanced_config"`
	ForwardScheme  string `tfsdk:"forward_scheme"`
	ForwardHost    string `tfsdk:"forward_host"`
	ForwardPort    string `tfsdk:"forward_port"`
}

type ProxyHostResourceModel struct {
	ID                    types.Int64  `tfsdk:"id"`
	CreatedOn             types.String `tfsdk:"created_on"`
	ModifiedOn            types.String `tfsdk:"modified_on"`
	OwnerUserID           types.Int64  `tfsdk:"owner_user_id"`
	DomainNames           types.List   `tfsdk:"domain_names"`
	ForwardHost           types.String `tfsdk:"forward_host"`
	ForwardPort           types.Int64  `tfsdk:"forward_port"`
	AccessListID          types.Int64  `tfsdk:"access_list_id"`
	CertificateID         types.Int64  `tfsdk:"certificate_id"`
	SslForced             types.Bool   `tfsdk:"ssl_forced"`
	CachingEnabled        types.Bool   `tfsdk:"caching_enabled"`
	BlockExploits         types.Bool   `tfsdk:"block_exploits"`
	AdvancedConfig        types.String `tfsdk:"advanced_config"`
	AllowWebsocketUpgrade types.Bool   `tfsdk:"allow_websocket_upgrade"`
	HTTP2Support          types.Bool   `tfsdk:"http_2_support"`
	ForwardScheme         types.String `tfsdk:"forward_scheme"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	Locations             []Location   `tfsdk:"locations"`
	HstsEnabled           types.Bool   `tfsdk:"hsts_enabled"`
	HstsSubdomains        types.Bool   `tfsdk:"hsts_subdomains"`
}

func CreateResourceModelRead(proxyHost *models.ProxyHostRead) *ProxyHostResourceModel {
	resourceModel := ProxyHostResourceModel{
		ID:                    types.Int64Value(int64(proxyHost.ID)),
		CreatedOn:             types.StringValue(proxyHost.CreatedOn),
		ModifiedOn:            types.StringValue(proxyHost.ModifiedOn),
		OwnerUserID:           types.Int64Value(int64(proxyHost.OwnerUserID)),
		DomainNames:           types.List{},
		ForwardHost:           types.StringValue(proxyHost.ForwardHost),
		ForwardPort:           types.Int64Value(int64(proxyHost.ForwardPort)),
		AccessListID:          types.Int64Value(int64(proxyHost.AccessListID)),
		CertificateID:         types.Int64Value(int64(proxyHost.CertificateID)),
		SslForced:             types.BoolValue(client.IntToBool(proxyHost.SslForced)),
		CachingEnabled:        types.BoolValue(client.IntToBool(proxyHost.CachingEnabled)),
		BlockExploits:         types.BoolValue(client.IntToBool(proxyHost.BlockExploits)),
		AdvancedConfig:        types.StringValue(proxyHost.AdvancedConfig),
		AllowWebsocketUpgrade: types.BoolValue(client.IntToBool(proxyHost.AllowWebsocketUpgrade)),
		HTTP2Support:          types.BoolValue(client.IntToBool(proxyHost.HTTP2Support)),
		ForwardScheme:         types.StringValue(proxyHost.ForwardScheme),
		Enabled:               types.BoolValue(client.IntToBool(proxyHost.Enabled)),
		Locations:             []Location{},
		HstsEnabled:           types.BoolValue(client.IntToBool(proxyHost.HstsEnabled)),
		HstsSubdomains:        types.BoolValue(client.IntToBool(proxyHost.HstsSubdomains)),
	}
	return &resourceModel
}

func CreateResourceModelCreated(proxyHost *models.ProxyHostCreated) *ProxyHostResourceModel {
	resourceModel := ProxyHostResourceModel{
		ID:                    types.Int64Value(int64(proxyHost.ID)),
		CreatedOn:             types.StringValue(proxyHost.CreatedOn),
		ModifiedOn:            types.StringValue(proxyHost.ModifiedOn),
		OwnerUserID:           types.Int64Value(int64(proxyHost.OwnerUserID)),
		DomainNames:           types.List{},
		ForwardHost:           types.StringValue(proxyHost.ForwardHost),
		ForwardPort:           types.Int64Value(int64(proxyHost.ForwardPort)),
		AccessListID:          types.Int64Value(int64(proxyHost.AccessListID)),
		CertificateID:         types.Int64Value(int64(proxyHost.CertificateID)),
		SslForced:             types.BoolValue(client.IntToBool(proxyHost.SslForced)),
		CachingEnabled:        types.BoolValue(client.IntToBool(proxyHost.CachingEnabled)),
		BlockExploits:         types.BoolValue(client.IntToBool(proxyHost.BlockExploits)),
		AdvancedConfig:        types.StringValue(proxyHost.AdvancedConfig),
		AllowWebsocketUpgrade: types.BoolValue(client.IntToBool(proxyHost.AllowWebsocketUpgrade)),
		HTTP2Support:          types.BoolValue(client.IntToBool(proxyHost.HTTP2Support)),
		ForwardScheme:         types.StringValue(proxyHost.ForwardScheme),
		Enabled:               types.BoolValue(client.IntToBool(proxyHost.Enabled)),
		Locations:             []Location{},
		HstsEnabled:           types.BoolValue(client.IntToBool(proxyHost.HstsEnabled)),
		HstsSubdomains:        types.BoolValue(client.IntToBool(proxyHost.HstsSubdomains)),
	}
	return &resourceModel
}

func CreateProxyHostCreateModel(resourceModel *ProxyHostResourceModel) *models.ProxyHostCreate {

	domainNames := make([]string, len(resourceModel.DomainNames.Elements()))
	for idx := range domainNames {
		domainNames[idx] = resourceModel.DomainNames.Elements()[idx].String()
	}

	requestBody := models.ProxyHostCreate{
		Enabled:               resourceModel.Enabled.ValueBool(),
		ForwardScheme:         resourceModel.ForwardScheme.String(),
		ForwardHost:           resourceModel.ForwardHost.String(),
		ForwardPort:           int(resourceModel.ForwardPort.ValueInt64()),
		AdvancedConfig:        resourceModel.AdvancedConfig.String(),
		DomainNames:           domainNames,
		AccessListID:          int(resourceModel.AccessListID.ValueInt64()),
		CertificateID:         int(resourceModel.CertificateID.ValueInt64()),
		SslForced:             resourceModel.SslForced.ValueBool(),
		Locations:             []models.Location{},
		BlockExploits:         resourceModel.BlockExploits.ValueBool(),
		CachingEnabled:        resourceModel.CachingEnabled.ValueBool(),
		AllowWebsocketUpgrade: resourceModel.AllowWebsocketUpgrade.ValueBool(),
		HTTP2Support:          resourceModel.HTTP2Support.ValueBool(),
		HstsEnabled:           resourceModel.HstsEnabled.ValueBool(),
		HstsSubdomains:        resourceModel.HstsSubdomains.ValueBool(),
	}

	return &requestBody
}

func CreateProxyHostUpdateModel(resourceModel *ProxyHostResourceModel) *models.ProxyHostUpdate {

	domainNames := make([]string, len(resourceModel.DomainNames.Elements()))
	for idx := range domainNames {
		domainNames[idx] = resourceModel.DomainNames.Elements()[idx].String()
	}

	requestBody := models.ProxyHostUpdate{
		Enabled:               resourceModel.Enabled.ValueBool(),
		ForwardScheme:         resourceModel.ForwardScheme.String(),
		ForwardHost:           resourceModel.ForwardHost.String(),
		ForwardPort:           int(resourceModel.ForwardPort.ValueInt64()),
		AdvancedConfig:        resourceModel.AdvancedConfig.String(),
		DomainNames:           domainNames,
		AccessListID:          int(resourceModel.AccessListID.ValueInt64()),
		CertificateID:         int(resourceModel.CertificateID.ValueInt64()),
		SslForced:             resourceModel.SslForced.ValueBool(),
		Meta:                  models.MetaData{},
		Locations:             []models.Location{},
		BlockExploits:         resourceModel.BlockExploits.ValueBool(),
		CachingEnabled:        resourceModel.CachingEnabled.ValueBool(),
		AllowWebsocketUpgrade: resourceModel.AllowWebsocketUpgrade.ValueBool(),
		HTTP2Support:          resourceModel.HTTP2Support.ValueBool(),
		HstsEnabled:           resourceModel.HstsEnabled.ValueBool(),
		HstsSubdomains:        resourceModel.HstsSubdomains.ValueBool(),
	}

	return &requestBody
}

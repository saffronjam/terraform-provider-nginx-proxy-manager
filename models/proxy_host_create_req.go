package models

type ProxyHostCreateReq struct {
	DomainNames           []string `json:"domain_names,omitempty"`
	ForwardHost           string   `json:"forward_host,omitempty"`
	ForwardPort           int      `json:"forward_port,omitempty"`
	AccessListID          int      `json:"access_list_id,omitempty"`
	CertificateID         int      `json:"certificate_id,omitempty"`
	SslForced             int      `json:"ssl_forced,omitempty"`
	CachingEnabled        int      `json:"caching_enabled,omitempty"`
	BlockExploits         int      `json:"block_exploits,omitempty"`
	AdvancedConfig        string   `json:"advanced_config,omitempty"`
	AllowWebsocketUpgrade int      `json:"allow_websocket_upgrade,omitempty"`
	HTTP2Support          int      `json:"http2_support,omitempty"`
	ForwardScheme         string   `json:"forward_scheme,omitempty"`
	Enabled               int      `json:"enabled,omitempty"`
	Locations             []string `json:"locations,omitempty"`
	HstsEnabled           int      `json:"hsts_enabled,omitempty"`
	HstsSubdomains        int      `json:"hsts_subdomains,omitempty"`
}

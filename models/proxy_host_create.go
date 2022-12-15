package models

type ProxyHostCreate struct {
	DomainNames           []string `json:"domain_names,omitempty"`
	ForwardHost           string   `json:"forward_host,omitempty"`
	ForwardPort           int      `json:"forward_port,omitempty"`
	AccessListID          int      `json:"access_list_id,omitempty"`
	CertificateID         int      `json:"certificate_id,omitempty"`
	SslForced             bool     `json:"ssl_forced,omitempty"`
	CachingEnabled        bool     `json:"caching_enabled,omitempty"`
	BlockExploits         bool     `json:"block_exploits,omitempty"`
	AdvancedConfig        string   `json:"advanced_config,omitempty"`
	AllowWebsocketUpgrade bool     `json:"allow_websocket_upgrade,omitempty"`
	HTTP2Support          bool     `json:"http2_support,omitempty"`
	ForwardScheme         string   `json:"forward_scheme,omitempty"`
	Enabled               bool     `json:"enabled,omitempty"`
	Locations             []string `json:"locations,omitempty"`
	HstsEnabled           bool     `json:"hsts_enabled,omitempty"`
	HstsSubdomains        int      `json:"hsts_subdomains,omitempty"`
}

package models

type ProxyHostCreate struct {
	DomainNames           []string   `json:"domain_names"`
	ForwardHost           string     `json:"forward_host"`
	ForwardPort           int        `json:"forward_port"`
	AccessListID          int        `json:"access_list_id"`
	CertificateID         int        `json:"certificate_id"`
	SslForced             bool       `json:"ssl_forced"`
	CachingEnabled        bool       `json:"caching_enabled"`
	BlockExploits         bool       `json:"block_exploits"`
	AdvancedConfig        string     `json:"advanced_config"`
	AllowWebsocketUpgrade bool       `json:"allow_websocket_upgrade"`
	HTTP2Support          bool       `json:"http2_support"`
	ForwardScheme         string     `json:"forward_scheme"`
	Enabled               bool       `json:"enabled"`
	Locations             []Location `json:"locations"`
	HstsEnabled           bool       `json:"hsts_enabled"`
	HstsSubdomains        bool       `json:"hsts_subdomains"`
}

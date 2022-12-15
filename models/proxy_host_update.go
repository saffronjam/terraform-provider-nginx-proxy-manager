package models

type MetaData struct {
	LetsencryptAgree bool `json:"letsencrypt_agree,omitempty"`
	DNSChallenge     bool `json:"dns_challenge,omitempty"`
}
type Location struct {
	Path           string `json:"path,omitempty"`
	AdvancedConfig string `json:"advanced_config,omitempty"`
	ForwardScheme  string `json:"forward_scheme,omitempty"`
	ForwardHost    string `json:"forward_host,omitempty"`
	ForwardPort    string `json:"forward_port,omitempty"`
}

type ProxyHostUpdate struct {
	ForwardScheme         string     `json:"forward_scheme,omitempty"`
	ForwardHost           string     `json:"forward_host,omitempty"`
	ForwardPort           int        `json:"forward_port,omitempty"`
	AdvancedConfig        string     `json:"advanced_config,omitempty"`
	DomainNames           []string   `json:"domain_names,omitempty"`
	AccessListID          string     `json:"access_list_id,omitempty"`
	CertificateID         int        `json:"certificate_id,omitempty"`
	SslForced             bool       `json:"ssl_forced,omitempty"`
	Meta                  MetaData   `json:"meta,omitempty"`
	Locations             []Location `json:"locations,omitempty"`
	BlockExploits         bool       `json:"block_exploits,omitempty"`
	CachingEnabled        bool       `json:"caching_enabled,omitempty"`
	AllowWebsocketUpgrade bool       `json:"allow_websocket_upgrade,omitempty"`
	HTTP2Support          bool       `json:"http2_support,omitempty"`
	HstsEnabled           bool       `json:"hsts_enabled,omitempty"`
	HstsSubdomains        bool       `json:"hsts_subdomains,omitempty"`
}

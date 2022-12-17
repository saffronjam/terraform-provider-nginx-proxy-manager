package models

type MetaData struct {
	LetsencryptAgree bool `json:"letsencrypt_agree"`
	DNSChallenge     bool `json:"dns_challenge"`
}
type Location struct {
	Path           string `json:"path"`
	AdvancedConfig string `json:"advanced_config"`
	ForwardScheme  string `json:"forward_scheme"`
	ForwardHost    string `json:"forward_host"`
	ForwardPort    string `json:"forward_port"`
}

type ProxyHostUpdate struct {
	ForwardScheme         string     `json:"forward_scheme"`
	ForwardHost           string     `json:"forward_host"`
	ForwardPort           int        `json:"forward_port"`
	AdvancedConfig        string     `json:"advanced_config"`
	DomainNames           []string   `json:"domain_names"`
	AccessListID          int        `json:"access_list_id"`
	CertificateID         int        `json:"certificate_id"`
	SslForced             bool       `json:"ssl_forced"`
	Meta                  MetaData   `json:"meta"`
	Locations             []Location `json:"locations"`
	BlockExploits         bool       `json:"block_exploits"`
	CachingEnabled        bool       `json:"caching_enabled"`
	AllowWebsocketUpgrade bool       `json:"allow_websocket_upgrade"`
	HTTP2Support          bool       `json:"http2_support"`
	HstsEnabled           bool       `json:"hsts_enabled"`
	HstsSubdomains        bool       `json:"hsts_subdomains"`
	Enabled               bool       `json:"enabled"`
}

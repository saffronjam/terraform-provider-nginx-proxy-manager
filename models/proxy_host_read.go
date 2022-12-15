package models

type ProxyHostRead struct {
	ID             int      `json:"id,omitempty"`
	CreatedOn      string   `json:"created_on,omitempty"`
	ModifiedOn     string   `json:"modified_on,omitempty"`
	OwnerUserID    int      `json:"owner_user_id,omitempty"`
	DomainNames    []string `json:"domain_names,omitempty"`
	ForwardHost    string   `json:"forward_host,omitempty"`
	ForwardPort    int      `json:"forward_port,omitempty"`
	AccessListID   int      `json:"access_list_id,omitempty"`
	CertificateID  int      `json:"certificate_id,omitempty"`
	SslForced      int      `json:"ssl_forced,omitempty"`
	CachingEnabled int      `json:"caching_enabled,omitempty"`
	BlockExploits  int      `json:"block_exploits,omitempty"`
	AdvancedConfig string   `json:"advanced_config,omitempty"`
	Meta           struct {
		NginxOnline bool        `json:"nginx_online,omitempty"`
		NginxErr    interface{} `json:"nginx_err,omitempty"`
	} `json:"meta,omitempty"`
	AllowWebsocketUpgrade int           `json:"allow_websocket_upgrade,omitempty"`
	HTTP2Support          int           `json:"http2_support,omitempty"`
	ForwardScheme         string        `json:"forward_scheme,omitempty"`
	Enabled               int           `json:"enabled,omitempty"`
	Locations             []interface{} `json:"locations,omitempty"`
	HstsEnabled           int           `json:"hsts_enabled,omitempty"`
	HstsSubdomains        int           `json:"hsts_subdomains,omitempty"`
}

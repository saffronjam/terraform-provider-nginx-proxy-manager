package models

type ProxyHostCreated struct {
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
		LetsencryptAgree bool `json:"letsencrypt_agree,omitempty"`
		DNSChallenge     bool `json:"dns_challenge,omitempty"`
	} `json:"meta,omitempty"`
	AllowWebsocketUpgrade int           `json:"allow_websocket_upgrade,omitempty"`
	HTTP2Support          int           `json:"http2_support,omitempty"`
	ForwardScheme         string        `json:"forward_scheme,omitempty"`
	Enabled               int           `json:"enabled,omitempty"`
	Locations             []interface{} `json:"locations,omitempty"`
	HstsEnabled           int           `json:"hsts_enabled,omitempty"`
	HstsSubdomains        int           `json:"hsts_subdomains,omitempty"`
	Certificate           struct {
		OwnerUserID int      `json:"owner_user_id,omitempty"`
		Provider    string   `json:"provider,omitempty"`
		NiceName    string   `json:"nice_name,omitempty"`
		DomainNames []string `json:"domain_names,omitempty"`
		ExpiresOn   string   `json:"expires_on,omitempty"`
		Meta        struct {
		} `json:"meta,omitempty"`
	} `json:"certificate,omitempty"`
	Owner struct {
		IsDisabled int    `json:"is_disabled,omitempty"`
		Name       string `json:"name,omitempty"`
		Nickname   string `json:"nickname,omitempty"`
		Avatar     string `json:"avatar,omitempty"`
	} `json:"owner,omitempty"`
	AccessList         interface{} `json:"access_list,omitempty"`
	UseDefaultLocation bool        `json:"use_default_location,omitempty"`
	Ipv6               bool        `json:"ipv6,omitempty"`
}

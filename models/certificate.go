package models

type Certificate struct {
	ID          int      `json:"id"`
	CreatedOn   string   `json:"created_on"`
	ModifiedOn  string   `json:"modified_on"`
	OwnerUserID int      `json:"owner_user_id"`
	Provider    string   `json:"provider"`
	NiceName    string   `json:"nice_name"`
	DomainNames []string `json:"domain_names"`
	ExpiresOn   string   `json:"expires_on"`
	Meta        struct {
		LetsencryptEmail       string `json:"letsencrypt_email"`
		DNSChallenge           bool   `json:"dns_challenge"`
		DNSProvider            string `json:"dns_provider"`
		DNSProviderCredentials string `json:"dns_provider_credentials"`
		LetsencryptAgree       bool   `json:"letsencrypt_agree"`
	} `json:"meta"`
}

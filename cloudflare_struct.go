package cloudflare

import "net/http"

//ConnectInformation - Base struct with all information about the connection
type ConnectInformation struct {
	AuthEmail      string
	AuthKey        string
	URL            string
	ZoneIdentifier string
	HTTPClient     *http.Client
}

// ZoneResponse - A struct with the basic response information for zones
type ZoneResponse struct {
	Success    bool       `json:"success"`
	Errors     []errors   `json:"errors"`
	Messages   []string   `json:"messages"`
	Result     []zones    `json:"result"`
	ResultInfo resultInfo `json:"result_info"`
}
type zones struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	DevelopmentMode    int      `json:"development_mode"`
	OriginalNameServer []string `json:"original_name_servers"`
	OriginalRegistrar  string   `json:"original_registrar"`
	OriginalDNSHost    string   `json:"original_dnshost"`
	CreatedOn          string   `json:"created_on"`
	ModifiedOn         string   `json:"modified_on"`
	NameServers        []string `json:"name_servers"`
	Owner              struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		OwnerType string `json:"owner_type"`
	} `json:"owner"`
	Permissions []string `json:"permissions"`
	Plan        struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Price        int    `json:"price"`
		Currency     string `json:"currency"`
		Frequency    string `json:"frequency"`
		LegacyID     string `json:"legacy_id"`
		IsSubscribed bool   `json:"is_subscribed"`
		CanSubscribe bool   `json:"can_subscribe"`
	} `json:"plan"`
	PlanPending struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Price        int    `json:"price"`
		Currency     string `json:"currency"`
		Frequency    string `json:"frequency"`
		LegacyID     string `json:"legacy_id"`
		IsSubscribed bool   `json:"is_subscribed"`
		CanSubscribe bool   `json:"can_subscribe"`
	} `json:"plan_pending"`
	Status    string `json:"status"`
	Paused    bool   `json:"paused"`
	Type      string `json:"type"`
	CheckedOn string `json:"checked_on"`
}

// DNSRecordsResponse - A struct with the response information for dns records
type DNSRecordsResponse struct {
	Success    bool        `json:"success"`
	Errors     []errors    `json:"errors"`
	Messages   []string    `json:"messages"`
	Result     []DNSRecord `json:"result"`
	ResultInfo resultInfo  `json:"result_info"`
}

// DNSRecordResponse - A struct with the respone information for a dns record
type DNSRecordResponse struct {
	Success    bool       `json:"success"`
	Errors     []errors   `json:"errors"`
	Messages   []string   `json:"messages"`
	Result     DNSRecord  `json:"result"`
	ResultInfo resultInfo `json:"result_info"`
}

type errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// DNSRecord - A struct with all the information for a dns record
type DNSRecord struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Content     string   `json:"content"`
	Proxiable   bool     `json:"proxiable"`
	Proxied     bool     `json:"proxied"`
	TTL         int      `json:"ttl"`
	Locked      bool     `json:"locked"`
	DisplayName string   `json:"display_name"`
	ZoneID      string   `json:"zone_id"`
	ZoneName    string   `json:"zone_name"`
	Priority    string   `json:"prio"`
	CreatedOn   string   `json:"created_on"`
	ModifiedOn  string   `json:"modified_on"`
	Data        struct{} `json:"data"`
}
type resultInfo struct {
	Page        int `json:"page"`
	PerPage     int `json:"per_page"`
	Count       int `json:"count"`
	TotaltCount int `json:"total_count"`
}

// CreateDNSRecord - A struct with all the information for creating a dns record
type CreateDNSRecord struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

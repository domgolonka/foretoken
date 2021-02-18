package entity

type EmailResponse struct {
	Valid      bool   `json:"valid"`
	Disposable bool   `json:"disposable"`
	RecentSpam bool   `json:"recent_spam"`
	Free       bool   `json:"free"`
	SMTPScore  uint8  `json:"smtp_score,omitempty"`
	Generic    bool   `json:"generic"`
	Score      uint8  `json:"score"`
	Leaked     bool   `json:"leaked,omitempty"`
	DomainAge  string `json:"domain_age,omitempty"`
	DNSValid   bool   `json:"dns_valid,omitempty"`
	HoneyPot   bool   `json:"honeypot,omitempty"`
}

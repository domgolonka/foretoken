package entity

type IPAddressResponse struct {
	Success         bool    `json:"success"`
	Proxy           bool    `json:"proxy"`
	ISP             string  `json:"ISP"`
	Organization    string  `json:"organization"`
	ASN             int     `json:"ASN"`
	Host            string  `json:"host"`
	CountryCode     string  `json:"country_code"`
	City            string  `json:"city"`
	Region          string  `json:"region"`
	IsCrawler       bool    `json:"is_crawler"`
	ConnectionType  string  `json:"connection_type"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Timezone        string  `json:"timezone"`
	Vpn             bool    `json:"vpn"`
	Tor             bool    `json:"tor"`
	RecentAbuse     bool    `json:"recent_abuse"`
	AbuseVelocity   string  `json:"abuse_velocity"`
	BotStatus       bool    `json:"bot_status"`
	Mobile          bool    `json:"mobile"`
	Score           uint8   `json:"score"`
	OperatingSystem string  `json:"operating_system"`
	Browser         string  `json:"browser"`
	DeviceModel     string  `json:"device_model"`
	DeviceBrand     string  `json:"device_brand"`
}

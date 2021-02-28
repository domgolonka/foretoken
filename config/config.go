package config

type Config struct {
	PublicAddress    string
	GRPCAddress      string
	Env              string
	AutoTLS          bool
	Proxy            Proxy
	External         External
	SMTP             SMTP
	Debug            bool
	Database         Database
	ErrorReporter    ErrorReporter
	Email            Email
	IP               IP
	RateLimit        RateLimit
	Crontab          Crontab
	ServiceDiscovery ServiceDiscovery
	Resource         Resource
}

type External struct {
	MaxmindDest string
	IP2Location string
}

type ServiceDiscovery struct {
	Service  string
	NodeID   string
	Endpoint string
}

type Crontab struct {
	VPN     int
	Tor     int
	Proxy   int
	Spam    int
	Maxmind int
}

type RateLimit struct {
	Enabled    bool
	Max        int
	Expiration int
}

type Email struct {
	Score EmailScore
}
type IP struct {
	Score IPScore
}

type EmailScore struct {
	Disposable Statement
	Free       Statement
	Spam       Statement
	Valid      Statement
	Generic    Statement
	CatchAll   Statement
	Leaked     Statement
	Domain     DomainStatement
}

type DomainStatement struct {
	Week     int8
	Month    int8
	Year     int8
	YearPlus int8
}

type IPScore struct {
	Proxy Statement
	Spam  Statement
	Tor   Statement
	VPN   Statement
}

type Statement struct {
	Yes int8
	No  int8
}
type SMTP struct {
	Hostname    string
	MailAddress string
}

type Proxy struct {
	CacheDurationMinutes int
	Workers              int
}

type ErrorReporter struct {
	Default             string
	AirbrakeCredentials string
}

type Database struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	TimeZone string
	SSL      bool
}

type Resource struct {
	EmailDisposalList []string
	EmailFreeList     []string
	EmailSpamList     []string
	IPVPNList         []string
	IPOpenVPNList     []string
	IPTorList         []string
	IPProxyList       []string
	IPSpamList        []string
	ExpressionList    []string
}

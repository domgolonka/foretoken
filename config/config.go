package config

type Config struct {
	Rooturl       string
	PublicPort    int
	ServerPort    int
	GRPCPort      int
	Env           string
	AutoTLS       bool
	Proxy         Proxy
	PwnedKey      string
	SMTP          SMTP
	Debug         bool
	Database      Database
	ErrorReporter ErrorReporter
	Email         Email
	IP            IP
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
}

type IPScore struct {
	Proxy Statement
	Spam  Statement
	Tor   Statement
	VPN   Statement
}

type Statement struct {
	Yes uint8
	No  uint8
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

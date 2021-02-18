package config

type Config struct {
	Appname       string
	Appversion    string
	Rooturl       string
	PublicPort    int
	ServerPort    int
	GRPCPort      int
	Env           string
	AutoTLS       bool
	Proxy         Proxy
	SMTP          SMTP
	Debug         bool
	DatabaseName  string
	Database      Database
	ErrorReporter ErrorReporter
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
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	TimeZone string
	SSL      bool
}

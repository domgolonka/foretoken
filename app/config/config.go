package config

type Config struct {
	Appname        string
	Appversion     string
	Rooturl        string
	PublicPort     int
	ServerPort     int
	Env            string
	DisposableFile string
	Proxy          Proxy
	Debug          bool
	DatabaseName   string
	Database       Database
	ErrorReporter  ErrorReporter
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

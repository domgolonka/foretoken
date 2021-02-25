package ip

import (
	"context"
	"github.com/oschwald/geoip2-golang"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"regexp"
)

var maxmindFilenameToDedupRegexp = regexp.MustCompile(`(.*/).*/.*`)

type Maxmind struct {
	license string
	paid    bool
	reader  *geoip2.Reader
}

type config struct {
	URL         string         // The URL of the file to download
	Store       file.Store     // The file.Store in which to place the file
	PathPrefix  string         // The prefix to attach to the file's path after it's downloaded
	CurrentName string         // The name to give the most recent version of the file.
	FilePrefix  string         // The prefix to attach to the filename after it's downloaded
	URLRegexp   *regexp.Regexp // The regular expression to apply to the URL to create the filename.
	// The first matching group will go before the timestamp, the second after.
	DedupRegexp   *regexp.Regexp // The regexp to apply to the filename to determine the directory to dedupe in.
	FixedFilename string         // The saved file could have fixed filename.
	MaxDuration   time.Duration  // The longest we allow the download process to go on before we consider it failed.
}

var maxmindDownloadInfo = []struct {
	url      string
	filename string
	current  string
}{
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz&license_key=",
		filename: "GeoLite2-ASN.tar.gz",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN-CSV&suffix=zip&license_key=",
		filename: "GeoLite2-ASN-CSV.zip",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz&license_key=",
		filename: "GeoLite2-City.tar.gz",
		current:  "Maxmind/current/GeoLite2-City.tar.gz",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City-CSV&suffix=zip&license_key=",
		filename: "GeoLite2-City-CSV.zip",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz&license_key=",
		filename: "GeoLite2-Country.tar.gz",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country-CSV&suffix=zip&license_key=",
		filename: "GeoLite2-Country-CSV.zip",
	},
}

func NewMaxmind(license string, paid bool) *Maxmind {
	return &Maxmind{
		license: license,
		paid:    paid,
	}
}

func MaxmindFiles(ctx context.Context, timestamp string, store file.Store, maxmindLicenseKey string) error {
	var lastErr error
	for _, info := range maxmindDownloadInfo {
		dc := config{
			URL:           info.url + maxmindLicenseKey,
			Store:         store,
			PathPrefix:    "Maxmind/" + timestamp,
			CurrentName:   info.current,
			FilePrefix:    time.Now().UTC().Format("20060102T150405Z-"),
			FixedFilename: info.filename,
			DedupRegexp:   maxmindFilenameToDedupRegexp,
			MaxDuration:   *downloadTimeout,
		}
		if err := runFunctionWithRetry(ctx, download, dc, *waitAfterFirstDownloadFailure, *maximumWaitBetweenDownloadAttempts); err != nil {
			lastErr = err
			metrics.FailedDownloadCount.With(prometheus.Labels{"download_type": "Maxmind"}).Inc()
		}
	}
	return lastErr

}

func (m *Maxmind) GetIPdata(IP net.IP) {

	asn, err := m.reader.ASN(IP)
	city, err := m.reader.City(IP)
	country, err := m.reader.Country(IP)

}

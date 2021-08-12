package ip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/domgolonka/foretoken/config"
	"github.com/oschwald/geoip2-golang"
)

type Maxmind struct {
	license string
	logger  logrus.FieldLogger
	cfg     *config.Config
	asn     *geoip2.Reader
	city    *geoip2.Reader
	country *geoip2.Reader
}

type maxmindDownload struct {
	url      string
	md5      string
	filename string
	file     string
	filemd5  string
}

var maxmindDownloadInfo = []maxmindDownload{
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-ASN.mmdb",
		file:     "assets/Maxmind/download/GeoLite2-ASN.tar.gz",
		filemd5:  "assets/Maxmind/download/GeoLite2-ASN.tar.gz.md5",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-City.mmdb",
		file:     "assets/Maxmind/download/GeoLite2-City.tar.gz",
		filemd5:  "assets/Maxmind/download/GeoLite2-City.tar.gz.md5",
	},

	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-Country.mmdb",
		file:     "assets/Maxmind/download/GeoLite2-Country.tar.gz",
		filemd5:  "assets/Maxmind/download/GeoLite2-Country.tar.gz.md5",
	},
}

func NewMaxmind(cfg *config.Config, license string, logger logrus.FieldLogger) *Maxmind {
	return &Maxmind{
		logger:  logger,
		license: license,
		cfg:     cfg,
	}
}

func (m *Maxmind) GetCountry(ip net.IP) (country string, err error) {

	count, err := m.country.Country(ip)
	if err != nil {
		m.logger.Errorf("cannot download maxmind country file with err: %s", err.Error())
		return "", err
	}
	return count.Country.IsoCode, err
}

func (m *Maxmind) GetCityData(ip net.IP) (postalcode string, timezone string, city string, latitude, longitude float64, err error) {
	fmtStr := "en-US"
	data, err := m.city.City(ip)
	if err != nil {
		m.logger.Errorf("cannot download maxmind city file with err: %s", err.Error())
		return "", "", "", 0, 0, err
	}
	return data.Postal.Code, data.Location.TimeZone, data.City.Names[fmtStr], data.Location.Latitude, data.Location.Longitude, err
}

func (m *Maxmind) GetASN(ip net.IP) (company string, asn uint, err error) {
	data, err := m.asn.ASN(ip)
	if err != nil {
		m.logger.Errorf("cannot download maxmind asn file with err: %s", err.Error())
		return "", 0, err
	}
	return data.AutonomousSystemOrganization, data.AutonomousSystemNumber, err
}

// Download a file
func (m *Maxmind) download(url, dest string, wg *sync.WaitGroup) error {
	defer wg.Done()

	resp, err := http.Get(url + m.license)
	if err != nil {
		return err
	}
	if resp.StatusCode%200 > 99 {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(dest)
	if err != nil {
		return err
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if err = out.Close(); err != nil {
		return err
	}
	return nil
}

func (m *Maxmind) DownloadAndUpdate() error {
	var wg sync.WaitGroup
	wg.Add(6)
	err := os.MkdirAll(m.cfg.External.MaxmindDest, os.ModePerm)
	if err != nil {
		m.logger.Errorf("cannot make directory %s", m.cfg.External.MaxmindDest)
		return err
	}
	for _, dd := range maxmindDownloadInfo {
		go func(maxD maxmindDownload) {
			err := m.download(maxD.url, maxD.file, &wg)
			if err != nil {
				m.logger.Error(err)
				return
			}
			err = m.download(maxD.md5, maxD.filemd5, &wg)
			if err != nil {
				m.logger.Error(err)
				return
			}
		}(dd)
	}
	wg.Wait()

	for _, d := range maxmindDownloadInfo {
		if err := VerifyMD5HashFromFile(d.file, d.filemd5); err != nil {
			return err
		}
		r, err := os.Open(d.file)
		if err != nil {
			return err
		}
		err = ExtractTarGz(r, m.cfg.External.MaxmindDest)
		if err != nil {
			return err
		}

		// Move mmdb to MAXMIND_DB_LOCATION
		geoCityDBPath, _, err := FindFile(m.cfg.External.MaxmindDest, "mmdb$")
		if err != nil {
			return err
		}

		if err = MoveFile(geoCityDBPath, d.filename); err != nil {
			return err
		}

		db, err := geoip2.Open(d.filename)
		if err != nil {
			return fmt.Errorf("cannot open maxmind file %s, error: %w", d.filename, err)
		}
		if strings.Contains(d.filename, "Country") {
			m.country = db
		} else if strings.Contains(d.filename, "City") {
			m.city = db
		} else if strings.Contains(d.filename, "ASN") {
			m.asn = db
		} else {
			return fmt.Errorf("cannot match maxmind filename: %s with db", d.filename)
		}
	}

	matches, err := filepath.Glob(m.cfg.External.MaxmindDest + "*")
	if err != nil {
		return err
	}
	for _, v := range matches {
		if err := os.RemoveAll(v); err != nil {
			return err
		}
	}

	return nil

}

package ip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/domgolonka/threatdefender/config"
	"github.com/oschwald/geoip2-golang"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Maxmind struct {
	license string
	cfg     config.Config
	reader  *geoip2.Reader
}

var maxmindDownloadInfo = []struct {
	url      string
	filename string
	file     string
}{
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz&license_key=",
		filename: "GeoLite2-ASN.tar.gz",
		file:     "assets/Maxmind/current/GeoLite2-ASN.tar.gz",
	},
	//{
	//	url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN-CSV&suffix=zip&license_key=",
	//	filename: "GeoLite2-ASN-CSV.zip",
	//},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz&license_key=",
		filename: "GeoLite2-City.tar.gz",
		file:     "assets/Maxmind/current/GeoLite2-City.tar.gz",
	},
	//{
	//	url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City-CSV&suffix=zip&license_key=",
	//	filename: "GeoLite2-City-CSV.zip",
	//},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz&license_key=",
		filename: "GeoLite2-Country.tar.gz",
		file:     "assets/Maxmind/current/GeoLite2-Country.tar.gz",
	},
	//{
	//	url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country-CSV&suffix=zip&license_key=",
	//	filename: "GeoLite2-Country-CSV.zip",
	//},
}

func NewMaxmind(license string, cfg config.Config) *Maxmind {
	return &Maxmind{
		license: license,
		cfg:     cfg,
	}
}

// Download a file
func download(url, dest string, wg *sync.WaitGroup) error {
	defer wg.Done()

	resp, err := http.Get(url)
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

func (m *Maxmind) GetIPdata(IP net.IP) {

	asn, err := m.reader.ASN(IP)
	city, err := m.reader.City(IP)
	country, err := m.reader.Country(IP)

}

func (m *Maxmind) downloadAndUpdate() error {
	var wg sync.WaitGroup
	wg.Add(2)
	for _, d := range maxmindDownloadInfo {
		go download(d.url, d.file, &wg)
	}
	wg.Wait()
	for _, d := range maxmindDownloadInfo {
		r, err := os.Open(d.file)
		if err != nil {
			return err
		}
		err = ExtractTarGz(r, m.cfg.APIKeys.MaxmindDest)
		if err != nil {
			return err
		}
	}

	matches, err := filepath.Glob(m.cfg.APIKeys.MaxmindDest + "*.tar.gz")
	if err != nil {
		return err
	}
	for _, v := range matches {
		if err := os.RemoveAll(v); err != nil {
			return err
		}
	}

}

// ExtractTarGz extracts a gzipped stream to dest
func ExtractTarGz(r io.Reader, dest string) error {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("Stream requires gzip-compressed body: %v", err)
	}

	tr := tar.NewReader(zr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Tar error: %v", err)
		}

		switch f.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(dest+f.Name, 0750); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %v", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(dest + f.Name)
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %v", err)
			}
			// For our purposes, we don't expect any files larger than 100MiB
			limited := &io.LimitedReader{R: tr, N: 100 << 20}
			if _, err := io.Copy(outFile, limited); err != nil {
				return fmt.Errorf("ExtractTarGz: Copy() failed: %v", err)
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf(
				"ExtractTarGz: %s has uknown type: %v",
				f.Name,
				f.Typeflag)
		}
	}

	return nil
}

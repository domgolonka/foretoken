package ip

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5" //nolint
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/domgolonka/threatdefender/config"
	"github.com/oschwald/geoip2-golang"
)

type Maxmind struct {
	license string
	logger  logrus.FieldLogger
	cfg     config.Config
	asn     *geoip2.Reader
	city    *geoip2.Reader
	country *geoip2.Reader
}

var maxmindDownloadInfo = []struct {
	url      string
	md5      string
	filename string
	file     string
	filemd5  string
}{
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-ASN.mmdb",
		file:     "assets/Maxmind/GeoLite2-ASN.tar.gz",
		filemd5:  "assets/Maxmind/GeoLite2-ASN.tar.gz.md5",
	},
	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-City.mmdb",
		file:     "assets/Maxmind/GeoLite2-City.tar.gz",
		filemd5:  "assets/Maxmind/GeoLite2-City.tar.gz.md5",
	},

	{
		url:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz&license_key=",
		md5:      "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-Country&suffix=tar.gz.md5&license_key=",
		filename: "assets/Maxmind/GeoLite2-Country.mmdb",
		file:     "assets/Maxmind/GeoLite2-Country.tar.gz",
		filemd5:  "assets/Maxmind/GeoLite2-Country.tar.gz.md5",
	},
}

func NewMaxmind(cfg config.Config, logger logrus.FieldLogger) *Maxmind {
	return &Maxmind{
		logger:  logger,
		license: cfg.APIKeys.Maxmind,
		cfg:     cfg,
	}
}

func (m *Maxmind) GetIPdata(IP net.IP) error {

	//asn, err := m.asn.ASN(IP)
	//if err != nil {
	//
	//}
	//city, err := m.city.City(IP)
	//if err != nil {
	//
	//}
	//country, err := m.country.Country(IP)
	//if err != nil {
	//
	//}
	return nil
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
	wg.Add(2)
	for _, d := range maxmindDownloadInfo {
		go func() {
			err := m.download(d.url, d.file, &wg)
			if err != nil {
				m.logger.Error(err)
				return
			}
		}()
		go m.download(d.md5, d.filemd5, &wg) //nolint
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
		err = ExtractTarGz(r, m.cfg.APIKeys.MaxmindDest)
		if err != nil {
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
		} else if strings.Contains(d.filename, "Asn") {
			m.asn = db
		} else {
			return fmt.Errorf("cannot match maxmind filename: %s with db", d.filename)
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

	return nil

}

// ExtractTarGz extracts a gzipped stream to dest
func ExtractTarGz(r io.Reader, dest string) error {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("stream requires gzip-compressed body: %v", err)
	}

	tr := tar.NewReader(zr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar error: %v", err)
		}

		switch f.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(dest+f.Name, 0750); err != nil {
				return fmt.Errorf("extractTarGz: Mkdir() failed: %v", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(dest + f.Name)
			if err != nil {
				return fmt.Errorf("extractTarGz: Create() failed: %v", err)
			}
			// For our purposes, we don't expect any files larger than 100MiB
			limited := &io.LimitedReader{R: tr, N: 100 << 20}
			if _, err := io.Copy(outFile, limited); err != nil {
				return fmt.Errorf("extractTarGz: Copy() failed: %v", err)
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf(
				"extractTarGz: %s has uknown type: %v",
				f.Name,
				f.Typeflag)
		}
	}

	return nil
}
func md5Hash(file string) ([]byte, error) {
	filePath := filepath.Clean(file)

	// We know exactly where this file and path is
	// #nosec G304
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// #nosec G401
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), f.Close()
}

func VerifyMD5HashFromFile(file, md5sumFile string) error {
	actual, err := md5Hash(file)
	if err != nil {
		return err
	}

	cleanMD5SumFile := filepath.Clean(md5sumFile)

	// We know exactly where this file and path is
	// #nosec G304
	expected, err := ioutil.ReadFile(cleanMD5SumFile)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", actual) != fmt.Sprintf("%s", expected) {
		return errors.New("checksum error")
	}

	return nil
}

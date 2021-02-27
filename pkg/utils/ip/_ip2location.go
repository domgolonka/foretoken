package ip

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// BIN: DB11LITEBIN
//const ip2locationUrl = "https://www.ip2location.com/download/?token={DOWNLOAD_TOKEN}&file={DATABASE_CODE}"

type IP2Location struct {
	license string
}

func NewIP2Location(license string) *IP2Location {
	return &IP2Location{
		license: license,
	}
}

func DownloadDBToFile(token, dbcode string, to string) error {
	f, err := os.Create(to)
	if err != nil {
		return err
	}
	defer f.Close()

	return downloadDBToWriter(token, dbcode, f)
}

func downloadDBToWriter(token, dbcode string, to io.Writer) error {
	resp, err := http.Get(fmt.Sprintf("https://www.ip2location.com/download/?token=%s&file=%s", token, dbcode))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status code %v", resp.StatusCode)
	}

	if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") || resp.ContentLength < 100 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("server response error '%v'", string(body))
	}

	_, err = io.Copy(to, resp.Body)
	return err
}

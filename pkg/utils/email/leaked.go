package email

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/domgolonka/threatdefender/app"
)

const (
	apiURL            = "https://haveibeenpwned.com/api/v3/breachedaccount/"
	clientTimeoutSecs = 2
	userAgent         = "threatdefender"
)

type respError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func Leaked(app *app.App, email, since string) (*bool, error) {

	hibpClient := http.Client{
		Timeout: time.Second * clientTimeoutSecs,
	}
	asTruncated := true
	if since != "" {
		asTruncated = false
	}
	request, err := http.NewRequest(http.MethodGet, fullURL(email, asTruncated), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("hibp-api-key", app.Config.PwnedKey)

	response, err := hibpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respErr := validateHTTPResponse(response.StatusCode, body)
	if respErr != nil {
		return nil, errors.New(respErr.Message)
	}

	return parseResponseBody(body)
}

func parseResponseBody(body []byte) (*bool, error) {
	breached := true
	if len(body) == 0 {
		breached = false
		// If the body is empty then there's no breaches
		return &breached, nil
	}

	return &breached, nil
}

func fullURL(account string, truncated bool) string {
	truncStr := "false"
	if truncated {
		truncStr = "true"
	}

	return apiURL + account + fmt.Sprintf("?truncateResponse=%s", truncStr)
}

func validateHTTPResponse(responseCode int, body []byte) *respError {
	hibpErr := &respError{}

	switch responseCode {
	case 401, 402:
		err := json.Unmarshal(body, hibpErr)
		if err != nil {
			return nil
		}
	default:
		hibpErr = nil
	}

	return hibpErr
}

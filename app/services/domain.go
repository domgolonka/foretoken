package services

import (
	"github.com/domainr/whois"
	whoisparser "github.com/likexian/whois-parser-go"
)

func domainAge(domain string) (string, error) {
	request, err := whois.NewRequest(domain)
	if err != nil {
		return "", err
	}
	response, err := whois.DefaultClient.Fetch(request)
	if err != nil {
		return "", err
	}
	result, err := whoisparser.Parse(string(response.Body))

	if err != nil {
		return "", err
	}

	return result.Domain.CreatedDate, nil
}

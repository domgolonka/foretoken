package email

import (
	"github.com/domainr/whois"
	"github.com/domgolonka/threatdefender/app/entity"
	whoisparser "github.com/likexian/whois-parser-go"
)

func DomainAge(domain string) (*entity.Domain, error) {
	request, err := whois.NewRequest(domain)
	if err != nil {
		return nil, err
	}
	response, err := whois.DefaultClient.Fetch(request)
	if err != nil {
		return nil, err
	}
	result, err := whoisparser.Parse(string(response.Body))
	if err != nil {
		return nil, err
	}
	return &entity.Domain{

		CreatedDate:    result.Domain.CreatedDate,
		ExpirationDate: result.Domain.ExpirationDate,
	}, nil
}

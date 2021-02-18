package services

import (
	"github.com/Boostport/address"
	"github.com/domgolonka/threatdefender/app/entity"
	"github.com/hashicorp/go-multierror"
)

func ValidateAddress(addy *entity.Address) (bool, error) {
	_, err := address.NewValid(
		address.WithCountry(addy.CountryCode), // Must be an ISO 3166-1 country code
		address.WithName(addy.Name),
		address.WithOrganization(addy.Organization),
		address.WithStreetAddress(addy.StreetAddress),
		address.WithLocality(addy.Locality),
		address.WithAdministrativeArea(addy.AdministrativeArea), // If the country has a pre-defined list of admin areas (like here), you must use the key and not the name
		address.WithPostCode(addy.PostalCode),
	)
	if err != nil {
		// If there was an error and you want to find out which validations failed,
		// type switch it as a *multierror.Error to access the list of errors
		if merr, ok := err.(*multierror.Error); ok {
			for _, subErr := range merr.Errors {
				if subErr == address.ErrInvalidCountryCode {
					return false, subErr
				}
			}
		}
	}
	return true, nil

}

package entity

type Address struct {
	CountryCode        string   `json:"country"`
	Name               string   `json:"name"`
	Organization       string   `json:"organization"`
	StreetAddress      []string `json:"street_address"`
	Locality           string   `json:"locality"`
	AdministrativeArea string   `json:"admin_area"`
	PostalCode         string   `json:"postal_code"`
}

package companies

import "errors"

const (
	endpointAll = "companies/getAll"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.AccessToken

	if s.Client.AccessToken == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointAll)
	if err != nil {
		return nil, err
	}

	responseBody := &AllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string `json:"AccessToken"`
}

type AllResponse struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	ID                          string  `json:"Id"`                          // Unique identifier of the company.
	Name                        string  `json:"Name"`                        // Name of the company.
	Identifier                  string  `json:"Identifier"`                  // Identifier of the company (e.g. legal identifier).
	TaxIdentificationNumber     string  `json:"TaxIdentificationNumber"`     // Tax identification number of the company.
	ElectronicInvoiceIdentifier string  `json:"ElectronicInvoiceIdentifier"` // Electronic invoice identifer of the company.
	Address                     Address `json:"Address"`
}

type Address struct {
	City        string `json:"City"`
	CountryCode string `json:"CountryCode"` // ISO 3166-1 alpha-2 country code (two letter country code).
	Line1       string `json:"Line1"`
	Line2       string `json:"Line2"`
	PostalCode  string `json:"PostalCode"`
}

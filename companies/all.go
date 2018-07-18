package companies

import "github.com/tim-online/go-mews/json"

const (
	endpointAll = "companies/getAll"
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
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
	json.BaseRequest
}

type AllResponse struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	ID                          string  `json:"Id"`                          // Unique identifier of the company.
	Name                        string  `json:"Name"`                        // Name of the company.
	Identifier                  string  `json:"Identifier"`                  // Identifier of the company (e.g. legal identifier).
	TaxIdentificationNumber     string  `json:"TaxIdentificationNumber"`     // Tax identification number of the company.
	AdditionalTaxIdentifier     string  `json:"AdditionalTaxIdentifier"`     // Additional tax identifer of the company.
	ElectronicInvoiceIdentifier string  `json:"ElectronicInvoiceIdentifier"` // Electronic invoice identifer of the company.
	Address                     Address `json:"Address"`
	AccountingCode              string  `json:"AccountingCode"` // NEW
	TaxIdentifier               string  `json:"TaxIdentifier"`  // NEW
}

type Address struct {
	City        string `json:"City"`
	CountryCode string `json:"CountryCode"` // ISO 3166-1 alpha-2 country code (two letter country code).
	Line1       string `json:"Line1"`
	Line2       string `json:"Line2"`
	PostalCode  string `json:"PostalCode"`
}

package taxenvironments

import "github.com/tim-online/go-mews/json"

const (
	endpointAll = "taxenvironments/getAll"
)

var (
	TaxRateStrategyDiscriminatorFlat     TaxRateStrategyDiscriminator = "Flat"
	TaxRateStrategyDiscriminatorRelative TaxRateStrategyDiscriminator = "Relative"
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

type AllResponse struct {
	TaxEnvironments TaxEnvironments `json:"TaxEnvironments"` // The supported tax environments.
	Taxations       Taxations       `json:"Taxations"`       // The supported taxations.
	TaxRates        TaxRates        `json:"TaxRates"`        // The supported tax rates.
}

type TaxEnvironments []TaxEnvironment

type TaxEnvironment struct {
	Code        string `json:"Code"`        // Code of the tax environment.
	CountryCode string `json:"CountryCode"` // ISO 3166-1 alpha-3 code, e.g. USA or GBR.

}

type Taxations []Taxation

type Taxation struct {
	Code               string `json:"Code"`               // Code of the taxation.
	TaxEnvironmentCode string `json:"TaxEnvironmentCode"` // Code of the tax environment.
	Name               string `json:"Name"`               // Name of the taxation.
	LocalName          string `json:"LocalName"`          // Local name of the taxation.
}

type TaxRates []TaxRate

type TaxRate struct {
	Code         string          `json:"Code"`         // Code of the tax rate.
	TaxationCode string          `json:"TaxationCode"` // Code of the taxation.
	Strategy     TaxRateStrategy `json:"Strategy"`     // Tax strategy type, e.g. relative or flat.
}

type TaxRateStrategy struct {
	Discriminator TaxRateStrategyDiscriminator `json:"TaxRateStrategyDiscriminator"` // If tax rate is flat or relative.
	// Value         interface{}                  `json:"Value"`                        // Structure of the object depends on Tax rate strategy discriminator.
	Value FlatTaxRateStrategyData `json:"Value"` // Structure of the object depends on Tax rate strategy discriminator.
}

type FlatTaxRateStrategyData struct {
	Value        float64 `json:"Value"`        // Absolute value of tax.
	CurrencyCode string  `json:"CurrencyCode"` // Code of Currency.
}

type RelativeTaxRateStrategyData struct {
	Value float64 `json:"Value"` // Tax rate, e.g. 0.21 in case of 21% tax rate.
}

type TaxRateStrategyDiscriminator string

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
}

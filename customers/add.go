package customers

import (
	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/json"
)

const (
	endpointAdd = "customers/add"
)

// Add customer
func (s *Service) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest

	// Overwrite existing
	OverwriteExisting bool `json:"OverwriteExisting"`
	// New first name.
	FirstName string `json:"FirstName"`
	// New last name.
	LastName string `json:"LastName"`
	// New second last name.
	SecondLastName string `json:"SecondLastName"`
	// New title.
	Title string `json:"Title"`
	// Sex
	Sex string `json:"Sex"`
	// ISO 3166-1 code of the Country.
	NationalityCode string `json:"NationalityCode"`
	// New birth date in ISO 8601 format.
	BirthDate *json.Date `json:"BirthDate"`
	// New birth place.
	BirthPlace string `json:"BirthPlace"`
	// New email address.
	Email string `json:"Email"`
	// New phone number.
	Phone string `json:"Phone"`
	// Loyalty code of the customer.
	LoyaltyCode string `json:"LocaltyCode"`
	// Internal notes about the customer. Old value will be overwritten.
	Notes string `json:"Notes"`
	// New identity card details.
	IdentityCard *Document `json:"IdentityCard"`
	// New passport details.
	Passport *Document `json:"Passport"`
	// New visa details.
	Visa *Document `json:"Visa"`
	// New drivers license details.
	DriversLicense *Document `json:"DriversLicense"`
	// New address details.
	Address *configuration.Address `json:"Address"`
	// Italian Destination Code
	ItalianDestinationCode string `json:"ItalianDestinationCode"`
}

type AddResponse Customer

package configuration

import (
	"errors"
	"time"
)

const (
	endpointGet = "configuration/get"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// Returns configuration of the enterprise and the client.
func (s *Service) Get(requestBody *GetRequest) (*GetResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.Token

	if s.Client.Token == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointGet)
	if err != nil {
		return nil, err
	}

	responseBody := &GetResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewGetRequest() *GetRequest {
	return &GetRequest{}
}

type GetRequest struct {
	AccessToken string `json:"AccessToken"`
}

type GetResponse struct {
	// The closed bills.
	Enterprise Enterprise `json:"Enterprise"`
}

type Enterprise struct {
	ID                      string     `json:"Id"`                      // Unique identifier of the enterprise.
	CreatedUTC              time.Time  `json:"CreatedUtc"`              // Creation date and time of the enterprise in UTC timezone in ISO 8601 format.
	Name                    string     `json:"Name"`                    // Name of the enterprise.
	WebsiteURL              string     `json:"WebsiteUrl"`              // URL of the enterprise website.
	Email                   string     `json:"Email"`                   // Email address of the enterprise.
	Phone                   string     `json:"Phone"`                   // Phone number of the enterprise.
	TimeZoneIdentifier      string     `json:"TimeZoneIdentifier"`      // IANA timezone identifier of the enterprise.
	EditableHistoryInterval string     `json:"EditableHistoryInterval"` // Editable history interval in ISO 8601 duration format.
	Address                 Address    `json:"Address"`                 // Address of the enterprise.
	Currencies              Currencies `json:"Currencies"`              // Currencies accepted by the enterprise.
}

type Currencies []Currency

type Currency struct {
	Currency  string `json:"Currency"`
	IsDefault bool   `json:"IsDefault"`
	IsEnabled bool   `json:"IsEnabled"`
}

type Address struct {
	// First line of the address.
	Line1 string `json:"Line1"`

	// Second line of the address.
	Line2 string `json:"Line2"`

	// The city.
	City string `json:"City"`

	// Postal code.
	PostalCode string `json:"PostalCode"`

	// ISO 3166-1 alpha-2 country code (two letter country code).
	CountryCode string `json:"CountryCode"`
}

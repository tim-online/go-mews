package configuration

import (
	"errors"
	"time"
)

const (
	endpointGetCountries = "configuration/countries/getAll"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// Returns configuration of the enterprise and the client.
func (s *Service) GetCountries(requestBody *GetCountriesRequest) (*GetCountriesResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.Token

	if s.Client.Token == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetCountriesApiURL(endpointGetCountries)
	if err != nil {
		return nil, err
	}

	responseBody := &GetCountriesResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewGetCountriesRequest() *GetCountriesRequest {
	return &GetCountriesRequest{}
}

type GetCountriesRequest struct {
	AccessToken string `json:"AccessToken"`
}

type GetCountriesResponse struct {
	// The closed bills.
	Counties Countries `json:"Countries"`
}

type Countries struct {
	EnglishName string `json:"EnglishName"` // English name of Country.
	Code        string `json:"Code"`        // Iso Country Code.
}

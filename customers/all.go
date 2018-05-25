package customers

import "time"

const (
	endpointAll = "customers/getAll"
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	// Set request tokens
	requestBody.AccessToken = s.Client.AccessToken
	requestBody.ClientToken = s.Client.ClientToken

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
	AccessToken string             `json:"AccessToken"`
	ClientToken string             `json:"ClientToken,omitempty"`
	TimeFilter  CustomerTimeFilter `json:"TimeFilter,omitempty"`
	StartUTC    *time.Time         `json:"StartUtc,omitempty"`
	EndUTC      *time.Time         `json:"EndUtc,omitempty"`
}

type AllResponse struct {
	Customers Customers `json:"customers"`
}

type CustomerTimeFilter string

const (
	CustomerTimeFilterCreated CustomerTimeFilter = "Created"
	CustomerTimeFilterUpdated CustomerTimeFilter = "Updated"
)

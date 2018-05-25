package businesssegments

import "errors"

const (
	endpointAll = "businesssegments/getAll"
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
	BusinessSegments BusinessSegments `json:"BusinessSegments"`
}

type BusinessSegments []BusinessSegment

type BusinessSegment struct {
	ID       string `json:"Id"`       // Unique identifier of the segment.
	IsActive bool   `json:"IsActive"` // Whether the business segment is still active.
	Name     string `json:"Name"`     // Name of the segment.
}

package reservations

const (
	endpointAllByIDs = "reservations/getAllByIds"
)

// List all products
func (s *Service) AllByIDs(requestBody *AllByIDsRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.Token

	if s.Client.Token == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointAllByIDs)
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

func (s *Service) NewAllByIDsRequest() *AllByIDsRequest {
	return &AllByIDsRequest{}
}

type AllByIDsRequest struct {
	AccessToken    string            `json:"AccessToken"`
	ReservationIDs []string          `json:"ReservationIds"`
	Extent         ReservationExtent `json:"Extent,omitempty"`
}

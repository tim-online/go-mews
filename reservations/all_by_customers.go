package reservations

const (
	endpointAllByCustomers = "reservations/getAllByCustomers"
)

// List all products
func (s *APIService) AllByCustomers(requestBody *AllByCustomersRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	// Set request tokens
	requestBody.AccessToken = s.Client.AccessToken
	requestBody.ClientToken = s.Client.ClientToken

	apiURL, err := s.Client.GetApiURL(endpointAllByCustomers)
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

func (s *APIService) NewAllByCustomersRequest() *AllByCustomersRequest {
	return &AllByCustomersRequest{}
}

type AllByCustomersRequest struct {
	AccessToken string            `json:"AccessToken"`
	ClientToken string            `json:"ClientToken,omitempty"`
	CustomerIDs []string          `json:"CustomerIds"`
	Extent      ReservationExtent `json:"Extent,omitempty"`
}

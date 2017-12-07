package bills

const (
	endpointAllByIDs = "bills/getAllByIds"
)

// List all products
func (s *Service) AllByIDs(requestBody *AllByIDsRequest) (*AllByIDsResponse, error) {
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

	responseBody := &AllByIDsResponse{}
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
	AccessToken string   `json:"AccessToken"`
	BillIDs     []string `json:"BillIds"`
}

type AllByIDsResponse struct {
	Bills Bills `json:"Bills"` // The closed bills.
}

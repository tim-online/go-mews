package commands

const (
	endpointUpdate = "commands/update"
)

// List all products
func (s *Service) Update(requestBody *UpdateRequest) (*UpdateResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.AccessToken

	if s.Client.AccessToken == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointUpdate)
	if err != nil {
		return nil, err
	}

	responseBody := &UpdateResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewUpdateRequest() *UpdateRequest {
	return &UpdateRequest{}
}

type UpdateRequest struct {
	AccessToken string       `json:"AccessToken"`
	CommandID   string       `json:"CommandId"`
	State       CommandState `json:"State"`
}

type UpdateResponse struct {
}

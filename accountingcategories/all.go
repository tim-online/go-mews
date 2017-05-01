package accountingcategories

import "errors"

const (
	endpointAll = "accountingCategories/getAll"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.Token

	if s.Client.Token == "" {
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

type AllResponse struct {
	AccountingCategories []AccountingCategory
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string `json:"AccessToken"`
}

type AccountingCategory struct {
	ID                 string `json:"ID"`
	Code               string `json:"Code"`
	IsActive           bool   `json:"IsActive"`
	Name               string `json:"Name"`
	ExternalCode       string `json:"ExternalCode"`
	CostCenterCode     string `json:"CostCenterCode"`
	LedgerAccountCode  string `json:"LedgerAccountCode"`
	PostingAccountCode string `json:"PostingAccountCode"`
}

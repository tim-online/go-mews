package bills

import (
	"time"

	"github.com/tim-online/go-mews/json"
)

const (
	endpointAll = "bills/getAll"
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
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
	json.BaseRequest
	// Unique identifiers of the Bills.
	BillIDs []string `json:"BillIds,omitempty"`
	// Unique identifiers of the Customers.
	CustomerIDs []string `json:"CustomerIds"`
	// Bill state the bills should be in. If not specified Open and Closed bills are returned.
	State string `json:"State,omitempty"`
	// Time filter of the interval.
	TimeFilter string `json:"TimeFilter,omitempty"`
	// Start of the interval in UTC timezone in ISO 8601 format.
	StartUTC *time.Time `json:"StartUtc,omitempty"`
	// End of the interval in UTC timezone in ISO 8601 format.
	EndUTC *time.Time `json:"EndUtc,omitempty"`
	// Extent of data to be returned. E.g. it is possible to specify that together with the bills, payments and revenue items should be also returned. If not specified, no extent is used.
	Extent BillExtent `json:"Extent,omitempty"`
}

type AllResponse struct {
	Bills Bills `json:"Bills"` // The closed bills.
}

type BillExtent struct {
	Items bool `json:"Items,omitempty"`
}

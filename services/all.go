package services

import (
	"github.com/tim-online/go-mews/json"
)

const (
	endpointAll = "services/getAll"

	Reservable ServiceType = "Reservable"
	Orderable  ServiceType = "Orderable"
)

// List all products
func (s *APIService) All(requestBody *AllRequest) (*AllResponse, error) {
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

type AllResponse struct {
	Services Services `json:"Services"` // Services offered by the enterprise.
}

type Services []Service

type Service struct {
	ID         string      `json:"Id"`          // Unique identifier of the service.
	IsActive   bool        `json:"IsActive"`    // Whether the service is still active.
	Name       string      `json:"Name"`        // Name of the service.
	StartTime  string      `json:"StartTime"`   // Default start time of the service orders in ISO 8601 duration format.
	EndTime    string      `json:"EndTime"`     // Default end time of the service orders in ISO 8601 duration format.
	Promotions Promotions  `json:"Promotions"`  // Promotions of the service.
	Type       ServiceType `json:"ServiceType"` // Type of the service
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	json.BaseRequest
}

type Promotions struct {
	BeforeCheckIn  bool `json:"BeforeCheckIn"`  // Whether it can be promoted before check-in.
	AfterCheckIn   bool `json:"AfterCheckIn"`   // Whether it can be promoted after check-in.
	DuringStay     bool `json:"DuringStay"`     // Whether it can be promoted during stay.
	BeforeCheckOut bool `json:"BeforeCheckOut"` // Whether it can be promoted before check-out.
	AfterCheckOut  bool `json:"AfterCheckOut"`  // Whether it can be promoted after check-out.
}

type ServiceType string

const (
	ServiceReservable ServiceType = "Reservable"
	ServiceOrderable  ServiceType = "Orderable"
)

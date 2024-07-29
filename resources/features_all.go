package resources

import (
	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/omitempty"
)

const (
	endpointResourceFeaturesAll = "resourceFeatures/getAll"
)

// List all products
func (s *APIService) FeaturesAll(requestBody *FeaturesAllRequest) (*FeaturesAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourceFeaturesAll)
	if err != nil {
		return nil, err
	}

	responseBody := &FeaturesAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewFeaturesAllRequest() *FeaturesAllRequest {
	return &FeaturesAllRequest{}
}

type FeaturesAllRequest struct {
	json.BaseRequest

	ResourceCategoryIDs []string                   `json:"ResourceCategoryIds,omitempty"` // Unique identifiers of Resource categories.
	ServiceIDs          []string                   `json:"ServiceIds,omitempty"`          // Unique identifiers of Services to which the resource categories belong.
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc,omitempty"`          // Interval in which the resource categories were updated.
	ActivityStates      ActivityStates             `json:"ActivityStates,omitempty"`      // Whether to return only active, only deleted or both records.
}

func (r FeaturesAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type FeaturesAllResponse struct {
	ResourceFeatures ResourceFeatures `json:"ResourceCategories"`
	Cursor           string           `json:"Cursor"` // Unique identifier of the last and hence oldest resource category returned. This can be used in Limitation in a subsequent request to fetch the next batch of older resource categories.
}

type ResourceFeatures []ResourceFeature

type ResourceFeature struct {
	ID             string                 `json:"Id"`             // Unique identifier of the category.
	ServiceID      string                 `json:"ServiceId"`      // Unique identifier of the Service.
	IsActive       bool                   `json:"IsActive"`       // Whether the resource feature is still active.
	Classification ResourceClassification `json:"Classification"` // Classification of the feature.
	Names          map[string]string      `json:"Names"`          // All translations of the name.
	ShortNames     map[string]string      `json:"ShortNames"`     // All translations of the short name.
	Descriptions   map[string]string      `json:"Descriptions"`   // All translations of the description.
}

type ResourceClassification string

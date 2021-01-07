package enterprises

import (
	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/omitempty"
)

const (
	endpointResourceBlocksGetAll = "resourceBlocks/getAll"
)

// List all products
func (s *APIService) ResourceBlocksGetAll(requestBody *ResourceBlocksGetAllRequest) (*ResourceBlocksGetAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourceBlocksGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &ResourceBlocksGetAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewResourceBlocksGetAllRequest() *ResourceBlocksGetAllRequest {
	return &ResourceBlocksGetAllRequest{}
}

type ResourceBlocksGetAllRequest struct {
	json.BaseRequest

	ResourceBlockIDs    []string                   `json:"ResourceBlockIds"`    // Unique identifiers of the requested Resource blocks.
	AssignedResourceIDs []string                   `json:"AssignedResourceIds"` // Unique identifiers of the requested Assigned Resources.
	CollidingUTC        configuration.TimeInterval `json:"CollidingUtc"`        // Interval in which the Resource block is active.
	CreatedUTC          configuration.TimeInterval `json:"CreatedUtc"`          // Interval in which the Resource block was created.
	UpdatedUTC          configuration.TimeInterval `json:"UpdatedUtc"`          // Interval in which the Resource block was updated.
	Extent              ResourceBlockExtent        `json:"Extent,omitempty"`    // Extent of data to be returned.
}

func (r ResourceBlocksGetAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type ResourceBlocksGetAllResponse struct {
}

type ResourceBlockExtent struct {
	Inactive bool `json:"Inactive"` // Whether the response should contain inactive entities.
}

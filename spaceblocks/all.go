package spaceblocks

import (
	"errors"
	"time"
)

const (
	endpointAll = "spaceBlocks/getAll"
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

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string     `json:"AccessToken"`
	StartUTC    *time.Time `json:"StartUtc,omitempty"`
	EndUTC      *time.Time `json:"EndUtc,omitempty"`
}

type AllResponse struct {
	SpaceBlocks SpaceBlocks `json:"SpaceBlocks"` // The space blocks colliding with the interval.
}

type SpaceBlocks []SpaceBlock

type SpaceBlock struct {
	ID              string    `json:"Id"`              // Unique identifier of the block.
	AssignedSpaceId string    `json:"AssignedSpaceId"` // Unique identifier of the assigned Space.
	Type            string    `json:"Type"`            // Type of the space block.
	StartUTC        time.Time `json:"startUtc"`        // Start of the block in UTC timezone in ISO 8601 format.
	EndUTC          time.Time `json:"endUtc"`          // End of the block in UTC timezone in ISO 8601 format.
	CreatedUTC      time.Time `json:"createdUtc"`      // Creation date and time of the block in UTC timezone in ISO 8601 format.
	UpdatedUTC      time.Time `json:"updatedUtc"`      // Last update date and time of the block in UTC timezone in ISO 8601 format.
}

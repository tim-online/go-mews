package spaces

import "errors"

const (
	endpointAll = "spaces/getAll"
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
	Spaces                  Spaces                  `json:"Spaces"`                  // The spaces of the enterprise
	SpaceCategories         SpaceCategories         `json:"SpaceCategories"`         // Categories of spaces in the enterprise.
	SpaceFeatures           SpaceFeatures           `json:"SpaceFeatures"`           // Features of spaces in the enterprise.
	SpaceFeatureAssignments SpaceFeatureAssignments `json:"SpaceFeatureAssignments"` // Assignments of space features to spaces.
}

type Spaces []Space
type SpaceCategories []SpaceCategory
type SpaceFeatures []SpaceFeatures
type SpaceFeatureAssignments []SpaceFeatureAssignment

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string    `json:"AccessToken"`
	Extent      AllExtent `json:"Extent,omitempty"`
}

type AllExtent struct {
	Spaces          bool `json:"Spaces"`          // Whether the response should contain spaces.
	SpaceCategories bool `json:"SpaceCategories"` // Whether the response should contain space categories.
	SpaceFeatures   bool `json:"SpaceFeatures"`   // Whether the response should contain space features and their assignments.
}

type Space struct {
	ID            string     `json:"Id"`            // Unique identifier of the reservation.
	Type          SpaceType  `json:"Type"`          // 	Type of the space.
	Number        string     `json:"number"`        // Number of the space (e.g. room number).
	ParentSpaceID string     `json:"ParentSpaceId"` // dentifier of the parent Space (e.g. room of a bed).
	CategoryID    string     `json:"CategoryId"`    // Identifier of the Space Category assigned to the space.
	State         SpaceState `json:"State"`         // State of the room.
}

type SpaceCategory struct {
}

type SpaceFeature struct {
	ID          string `json:"Id"`          // Unique identifier of the feature.
	Name        string `json:"Name"`        // Name of the feature.
	Description string `json:"Description"` // Description of the feature.
}

type SpaceFeatureAssignment struct {
	SpaceID        string `json:"SpaceId"`        // Unique identifier Space.
	SpaceFeatureID string `json:"SpaceFeatureId"` // Unique identifier Space Feature.
}

type SpaceType string

const (
	SpaceTypeRoom SpaceType = "Room"
	SpaceTypeDorm SpaceType = "Dorm"
	SpaceTypeBed  SpaceType = "Bed"
)

type SpaceState string

const (
	SpaceStateDirty        SpaceState = "Dirty"
	SpaceStateClean        SpaceState = "Clean"
	SpaceStateInspected    SpaceState = "Inspected"
	SpaceStateOutOfService SpaceState = "OutOfService"
	SpaceStateOutOfOrder   SpaceState = "OutOfOrder"
)
package enterprises

import (
	"time"

	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/omitempty"
)

const (
	endpointResourcesGetAll = "resources/getAll"
)

// List all products
func (s *APIService) ResourcesGetAll(requestBody *ResourcesGetAllRequest) (*ResourcesGetAllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointResourcesGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &ResourcesGetAllResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewResourcesGetAllRequest() *ResourcesGetAllRequest {
	return &ResourcesGetAllRequest{}
}

type ResourcesGetAllRequest struct {
	json.BaseRequest

	// Extent of data to be returned.
	Extent ResourceExtent `json:"Extent,omitempty"`
}

func (r ResourcesGetAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type ResourcesGetAllResponse struct {
	Resources                   Resources                   `json:"Resources"`
	ResourceCategories          ResourceCategories          `json:"ResourceCategories"`
	ResourceCategoryAssignments ResourceCategoryAssignments `json:"ResourceCategoryAssignments"`
	// ResourceCategoryImageAssignments ResourceCategoryImageAssignments `json:"ResourceCategoryImageAssignments"`
	// ResourceFeatures                 ResourceFeatures                 `json:"ResourceFeatures"`
	// ResourceFeatureAssignments       ResourceFeatureAssignments       `json:"ResourceFeatureAssignments"`
}

type Resources []Resource

type Resource struct {
	ID               string        `json:"Id"`               // Unique identifier of the resource.
	IsActive         bool          `json:"IsActive"`         // Whether the resource is still active.
	Name             string        `json:"Name"`             // Name of the resource (e.g. room number).
	ParentResourceID string        `json:"ParentResourceId"` // Identifier of the parent Resource (e.g. room of a bed).
	State            ResourceState `json:"State"`            // State of the resource.
	Data             ResourceData  `json:"Data"`             // Additional data of the resource.
	CreatedUTC       time.Time     `json:"CreatedUtc"`       // Creation date and time of the resource in UTC timezone in ISO 8601 format.
	UpdatedUTC       time.Time     `json:"UpdatedUtc"`       // Last update date and time of the resource in UTC timezone in ISO 8601 format.
}

type ResourceCategories []ResourceCategory

type ResourceCategory struct {
	ID            string                      `json:"Id"`            // Unique identifier of the category.
	IsActive      bool                        `json:"IsActive"`      // Whether the category is still active.
	Type          ResourceCategoryType        `json:"Type"`          // Type of the category.
	Names         configuration.LocalizedText `json:"Names"`         // All translations of the name.
	ShortNames    configuration.LocalizedText `json:"ShortNames"`    // All translations of the short name.
	Descriptions  configuration.LocalizedText `json:"Descriptions"`  // All translations of the description.
	Ordering      int                         `json:"Ordering"`      // Ordering of the category, lower number corresponds to lower category (note that uniqueness nor continuous sequence is guaranteed).
	Capacity      int                         `json:"Capacity"`      // Capacity that can be served (e.g. bed count).
	ExtraCapacity int                         `json:"ExtraCapacity"` // Extra capacity that can be served (e.g. extra bed count).
}

type ResourceCategoryType string

type ResourceExtent struct {
	Resources                        bool `json:"Resources"`                        // Whether the response should contain resources.
	ResourceCategories               bool `json:"ResourceCategories"`               // Whether the response should contain categories.
	ResourceCategoryAssignments      bool `json:"ResourceCategoryAssignments"`      // Whether the response should contain assignments of the resources to categories.
	ResourceCategoryImageAssignments bool `json:"ResourceCategoryImageAssignments"` // Whether the response should contain assignments of the images to categories.
	ResourceFeatures                 bool `json:"ResourceFeatures"`                 // Whether the response should contain resource features.
	ResourceFeatureAssignments       bool `json:"ResourceFeatureAssignments"`       // Whether the response should contain assignments of the resources to features.
	Inactive                         bool `json:"Inactive"`                         // Whether the response should contain inactive entities.
}

type ResourceCategoryAssignments []ResourceCategoryAssignment

type ResourceCategoryAssignment struct {
	ID         string `json:"Id"`         // Unique identifier of the assignment.
	IsActive   bool   `json:"IsActive"`   // Whether the assignment is still active.
	CategoryID string `json:"CategoryId"` // Unique identifier of the Resource category.
	ResourceID string `json:"ResourceId"` // Unique identifier of the Resource assigned to the Resource category.
	CreatedUTC string `json:"CreatedUtc"` // Creation date and time of the assignment in UTC timezone in ISO 8601 format.
	UpdatedUTC string `json:"UpdatedUtc"` // Last update date and time of the assignment in UTC timezone in ISO 8601 format.
}

type ResourceState string

type ResourceData struct {
	Discriminator ResourceDataDiscriminator `json:"Discriminator"` // If resource is space, object or person.
	Value         interface{}               `json:"Value"`         // Based on Resource data discriminator, e.g. Space resource data
}

type ResourceDataDiscriminator string

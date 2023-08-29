package reservations

import (
	"fmt"
	"time"

	"github.com/tim-online/go-mews/configuration"
	base "github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/omitempty"
)

const (
	endpointGetAll = "reservations/getAll/2023-06-06"
)

// List all products
func (s *APIService) GetAll(requestBody *GetAllRequest) (*AllResponsev2, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointGetAll)
	if err != nil {
		return nil, err
	}

	responseBody := &AllResponsev2{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}
	fmt.Println(httpReq)

	_, err = s.Client.Do(httpReq, responseBody)
	fmt.Println("Response : ", responseBody)
	return responseBody, err
}

type AllResponsev2 struct {
	Reservations Reservationsv2
	Cursor       string `json:"Cursor"`
}

func (s *APIService) NewGetAllRequest() *GetAllRequest {
	return &GetAllRequest{}
}

type GetAllRequest struct {
	base.BaseRequest
	Limitation          base.Limitation            `json:"Limitation,omitempty"`
	EnterpriseIds       []string                   `json:"EnterpriseIds,omitempty"`
	ReservationIds      []string                   `json:"ReservationIds,omitempty"`
	AccountIds          []string                   `json:"AccountIds,omitempty"`
	CustomerIDs         []string                   `json:"CustomerIds,omitempty"`
	ReservationGroupIds []string                   `json:"ReservationGroupIds,omitempty"`
	States              []ReservationStates        `json:"States,omitempty"`
	UpdatedUtc          configuration.TimeInterval `json:"UpdatedUtc,omitempty"`
	CollidingUtc        configuration.TimeInterval `json:"CollidingUtc,omitempty"`
}

func (r GetAllRequest) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type Reservationsv2 []Reservationv2

type Reservationv2 struct {
	ID                          string            `json:"ID,omitempty"`
	EnterpriseID                string            `json:"EnterpriseID,omitempty"`
	ServiceID                   string            `json:"ServiceID,omitempty"`
	AccountID                   string            `json:"AccountID,omitempty"`
	AccountType                 string            `json:"AccountType,omitempty"`
	CreatorProfileID            string            `json:"CreatorProfileID,omitempty"`
	UpdaterProfileID            string            `json:"UpdaterProfileID,omitempty"`
	BookerId                    string            `json:"BookerId,omitempty"`
	StartUTC                    time.Time         `json:"StartUtc,omitempty"`
	EndUTC                      time.Time         `json:"EndUtc,omitempty"`
	Number                      string            `json:"Number,omitempty"`
	State                       ReservationStates `json:"State"`
	Origin                      string            `json:"Origin,omitempty"`
	OriginDetails               string            `json:"OriginDetails,omitempty"`
	CommanderOrigin             CommanderOrigin   `json:"CommanderOrigin,omitempty"`
	CreatedUTC                  time.Time         `json:"CreatedUtc,omitempty"`
	UpdatedUTC                  time.Time         `json:"UpdatedUtc,omitempty"`
	ReleasedUTC                 time.Time         `json:"ReleasedUtc,omitempty"`
	CancelledUTC                time.Time         `json:"CancelledUtc,omitempty"`
	VoucherId                   string            `json:"VoucherId,omitempty"`
	BusinessSegmentID           string            `json:"BusinessSegmentId,omitempty"`
	RateID                      string            `json:"RateId,omitempty"`
	CreditCardId                string            `json:"CreditCardId,omitempty"`
	GroupID                     string            `json:"GroupId,omitempty"`
	RequestedResourceCategoryId string            `json:"RequestedResourceCategoryId,omitempty"`
	AssignedResourceID          string            `json:"AssignedResourceId,omitempty"`
	AvailabilityBlockId         string            `json:"AvailabilityBlockId,omitempty"`
	CompanyID                   string            `json:"CompanyId,omitempty"`
	TravelAgencyID              string            `json:"TravelAgencyId,omitempty"`
	AssignedResourceLocked      bool              `json:"AssignedSpaceLocked,omitempty"`
	ChannelNumber               string            `json:"ChannelNumber,omitempty"`
	ChannelManagerNumber        string            `json:"ChannelManagerNumber,omitempty"`
	CancellationReason          string            `json:"CancellationReason,omitempty"` // ??
	Purpose                     string            `json:"Purpose,omitempty"`            // ??
	Options                     string            `json:"Options,omitempty"`            // ??
	PersonCounts                []PersonCounts    `json:"PersonCounts,omitempty"`       // ??
}

type CommanderOrigin string

const (
	CommanderOriginInPerson   CommanderOrigin = "InPerson"
	CommanderOriginChannel    CommanderOrigin = "Channel"
	CommanderOriginPhone      CommanderOrigin = "Phone"
	CommanderOriginEmail      CommanderOrigin = "Email"
	CommanderOriginWebsite    CommanderOrigin = "Website"
	CommanderOriginMessage    CommanderOrigin = "Message"
	CommanderOriginCallCenter CommanderOrigin = "CallCenter"
)

type ReservationStates string

const (
	ReservationStatesEnquired  ReservationStates = "Enquired"
	ReservationStatesRequested ReservationStates = "Requested"
	ReservationStatesOptional  ReservationStates = "Optional"
	ReservationStatesConfirmed ReservationStates = "Confirmed"
	ReservationStatesStarted   ReservationStates = "Started"
	ReservationStatesProcessed ReservationStates = "Processed"
	ReservationStatesCanceled  ReservationStates = "Canceled"
)

type PersonCounts struct {
	AgeCategoryId string `json:"AgeCategoryId,omitempty"`
	Count         string `json:"Count,omitempty"`
}

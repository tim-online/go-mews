package reservations

import (
	"encoding/json"
	"time"

	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/customers"
)

const (
	endpointAll = "reservations/getAll"
)

// List all products
func (s *APIService) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	// Set request tokens
	requestBody.AccessToken = s.Client.AccessToken
	requestBody.ClientToken = s.Client.ClientToken

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
	BusinessSegments  BusinessSegments                `json:"BusinessSegments"`  //  Business segments of the reservations.
	Customers         customers.Customers             `json:"Customers"`         // Customers that are members of the reservations.
	Items             accountingitems.AccountingItems `json:"Items"`             // Revenue items of the reservations.
	Products          Products                        `json:"Products"`          // Products orderable with reservations.
	RateGroups        RateGroups                      `json:"RateGroups"`        // Rate groups of the reservation rates.
	Rates             Rates                           `json:"Rates"`             // Rates of the reservations.
	ReservationGroups ReservationGroups               `json:"ReservationGroups"` // Reservation groups that the reservations are members of.
	Reservations      Reservations                    `json:"Reservations"`      // The reservations that collide with the specified interval.
	Services          Services                        `json:"Services"`          // Services that have been reserved.
	SpaceCategories   SpaceCategories                 `json:"SpaceCategories"`   // Space categories of the spaces.
	Spaces            Spaces                          `json:"Spaces"`            // Assigned spaces of the reservations.
}

type Reservations []Reservation
type Services []Service

type Service struct {
	ID         string     `json:"Id"`         // Unique identifier of the service.
	IsActive   bool       `json:"IsActive"`   // Whether the service is still active.
	Name       string     `json:"Name"`       // Whether the service is still active.
	StartTime  string     `json:"StartTime"`  // Default start time of the service orders in ISO 8601 duration format.
	EndTime    string     `json:"EndTime"`    // Default end time of the service orders in ISO 8601 duration format.
	Promotions Promotions `json:"Promotions"` // Promotions of the service.
}

func (s *APIService) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string                `json:"AccessToken"`
	ClientToken string                `json:"ClientToken,omitempty"`
	TimeFilter  ReservationTimeFilter `json:"TimeFilter,omitempty"`
	StartUTC    *time.Time            `json:"StartUtc,omitempty"`
	EndUTC      *time.Time            `json:"EndUtc,omitempty"`
	States      []ReservationState    `json:"States"`
	Extent      ReservationExtent     `json:"Extent,omitempty"`
}

type ReservationExtent struct {
	BusinessSegments  bool `json:"BusinessSegments"`  // Whether the response should contain business segmentation.
	Customers         bool `json:"Customers"`         // Whether the response should contain customers of the reservations.
	Items             bool `json:"Items"`             // Whether the response should contain reservation items.
	Products          bool `json:"Products"`          // Whether the response should contain products orderable with the reservations.
	Rates             bool `json:"Rates"`             // Whether the response should contain rates and rate groups.
	Reservations      bool `json:"Reservations"`      // Whether the response should contain reservations.
	ReservationGroups bool `json:"ReservationGroups"` // Whether the response should contain groups of the reservations.
	Services          bool `json:"Services"`          // Whether the response should contain services reserved by the reservations.
	Spaces            bool `json:"Spaces"`            // Whether the response should contain spaces and space categories.
}

type Reservation struct {
	ID                        string           `json:"Id"`                        // Unique identifier of the reservation.
	ServiceID                 string           `json:"ServiceId"`                 // Unique identifier of the Service that is reserved.
	GroupID                   string           `json:"GroupId"`                   // Unique identifier of the Reservation Group.
	Number                    string           `json:"Number"`                    // Confirmation number of the reservation in Mews.
	ChannelNumber             string           `json:"ChannelNumber"`             // Number of the reservation within the Channel (i.e. OTA, GDS, CRS, etc) in case the reservation group originates there (e.g. Booking.com confirmation number).
	ChannelManagerNumber      string           `json:"ChannelManagerNumber"`      //  Unique number of the reservation within the reservation group.
	ChannelManagerGroupNumber string           `json:"ChannelManagerGroupNumber"` // Number of the reservation group within a Channel manager that transferred the reservation from Channel to Mews.
	ChannelManager            string           `json:"ChannelManager"`            // Name of the Channel manager (e.g. AvailPro, SiteMinder, TravelClick, etc).
	State                     ReservationState `json:"State"`                     // State of the reservation.
	Origin                    string           `json:"Origin"`                    // Origin of the reservation.
	CreatedUTC                time.Time        `json:"CreatedUtc"`                // Creation date and time of the reservation in UTC timezone in ISO 8601 format.
	UpdatedUTC                time.Time        `json:"UpdatedUtc"`                // Last update date and time of the reservation in UTC timezone in ISO 8601 format.
	StartUTC                  time.Time        `json:"StartUtc"`                  // Start of the reservation (arrival) in UTC timezone in ISO 8601 format.
	EndUTC                    time.Time        `json:"EndUtc"`                    // End of the reservation (departure) in UTC timezone in ISO 8601 format.
	CancelledUTC              time.Time        `json:"CancelledUtc"`              // Cancellation date and time in UTC timezone in ISO 8601 format.
	RequestedCategoryID       string           `json:"RequestedCategoryId"`       // Identifier of the requested Space Category.
	AssignedSpaceID           string           `json:"AssignedSpaceId"`           // Identifier of the assigned Space.
	AssignedSpaceLocked       bool             `json:"AssignedSpaceLocked"`       // Whether the reservation is locked in the assigned Space and cannot be moved.
	BusinessSegmentID         string           `json:"BusinessSegmentId"`         // Identifier of the reservation Business Segment.
	CompanyID                 string           `json:"CompanyId"`                 // Identifier of the Company on behalf of which the reservation was made.
	TravelAgencyID            string           `json:"TravelAgencyId"`            // Identifier of the Company that mediated the reservation.
	RateID                    string           `json:"RateId"`                    // Identifier of the reservation Rate.
	AdultCount                int              `json:"AdultCount"`                // Count of adults the reservation was booked for.
	ChildCount                int              `json:"ChildCount"`                // Count of children the reservation was booked for.
	CustomerID                string           `json:"Customerid"`                // required	Unique identifier of the Customer who owns the reservation.
	CompanionIDs              []string         `json:"CompanionIds"`              // Unique identifiers of Customers that will occupy the space.
	ChannelManagerID          string           `json:"ChannelManagerId"`          // ??
}

type ReservationState string

const (
	ReservationStateEnquired  ReservationState = "Enquired"
	ReservationStateRequested ReservationState = "Requested"
	ReservationStateOptional  ReservationState = "Optional"
	ReservationStateConfirmed ReservationState = "Confirmed"
	ReservationStateStarted   ReservationState = "Started"
	ReservationStateProcessed ReservationState = "Processed"
	ReservationStateCanceled  ReservationState = "Canceled"
)

type ReservationTimeFilter string

const (
	ReservationTimeFilterColliding   ReservationTimeFilter = "Colliding"
	ReservationTimeFilterCreated     ReservationTimeFilter = "Created"
	ReservationTimeFilterUpdated     ReservationTimeFilter = "Updated"
	ReservationTimeFilterStart       ReservationTimeFilter = "Start"
	ReservationTimeFilterEnd         ReservationTimeFilter = "End"
	ReservationTimeFilterOverlapping ReservationTimeFilter = "Overlapping"
	ReservationTimeFilterCancelled   ReservationTimeFilter = "Cancelled"
)

type Title string

const (
	TitleMister Title = "Mister"
	TitleMiss   Title = "Miss"
	TitleMisses Title = "Missed"
)

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)

type BusinessSegments []BusinessSegment

type BusinessSegment struct {
	ID       string `json:"Id"`       // Unique identifier of the segment.
	IsActive bool   `json:"IsActive"` // Whether the business segment is still active.
	Name     string `json:"Name"`     // Name of the segment.
}

type Document struct {
	Number     string `json:"Number"`     // Number of the document (e.g. passport number).
	Issuance   string `json:"Issuance"`   // Date of issuance in ISO 8601 format.
	Expiration Date   `json:"Expiration"` // Expiration date in ISO 8601 format.
}

type Address struct {
	Line1       string `json:"Line1"`       // First line of the address.
	Line2       string `json:"Line2"`       // Second line of the address.
	City        string `json:"City"`        // The city.
	PostalCode  string `json:"PostalCode"`  // Postal code.
	CountryCode string `json:"CountryCode"` // ISO 3166-1 alpha-2 country code (two letter country code).
}

type CurrencyValue struct {
	Currency string  `json:"Currency"` // ISO-4217 currency code, e.g. EUR or USD.
	Value    float64 `json:"Value"`    // Amount in the currency (including tax if taxed).
	TaxRate  float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Tax      float64 `json:"Tax"`      // Tax value in case the item is taxed.
}

type Products []Product

type Product struct {
	ID         string                 `json:"Id"`         // Unique identifier of the product.
	ServiceID  string                 `json:"ServiceId"`  // Unique identifier of the Service.
	IsActive   bool                   `json:"IsActive"`   // Whether the product is still active.
	Name       string                 `json:"Name"`       // Name of the product.
	ShortName  string                 `json:"ShortName"`  // Short name of the product.
	Promotions Promotions             `json:"Promotions"` // Promotions of the service.
	Price      accountingitems.Amount `json:"Price"`      // Price of the product.
}

type RateGroups []RateGroup

type RateGroup struct {
	Id       string `json:"Id"`       // Unique identifier of the group.
	IsActive bool   `json:"IsActive"` // Whether the rate group is still active.
	Name     string `json:"Name"`     // Name of the rate group.
}

type Rates []Rate

type Rate struct {
	ID         string `json:"Id"`         // Unique identifier of the rate.
	GroupID    string `json:"GroupId"`    // Unique identifier of Rate Group where the rate belongs.
	BaseRateID string `json:"BaseRateId"` // Unique identifier of the base Rate.
	IsActive   bool   `json:"IsActive"`   // Whether the rate is still active.
	IsPublic   bool   `json:"IsPublic"`   // Whether the rate is publicly available.
	Name       string `json:"Name"`       // Name of the rate.
	ShortName  string `json:"ShortName"`  // Short name of the rate.
}

type ReservationGroups []ReservationGroup

type ReservationGroup struct {
	ID   string `json:"Id"`   // Unique identifier of the reservation group.
	Name string `json:"Name"` // Name of the reservation group, might be empty or same for multiple groups.
}

type SpaceCategories []SpaceCategory

type SpaceCategory struct {
	ID          string   `json:"Id"`          // Unique identifier of the category.
	Name        string   `json:"Name"`        // Name of the category.
	ShortName   string   `json:"ShortName"`   // Short name (e.g. code) of the category.
	Description string   `json:"Description"` // Description of the category.
	Ordering    int      `json:"Ordering"`    // Ordering of the category, lower number corresponds to lower category (note that uniqueness nor continuous sequence is guaranteed).
	UnitCount   int      `json:"Unitcount"`   // Count of units that can be accommodated (e.g. bed count).
	ImageIDs    []string `json:"ImageIds"`    // Unique identifiers of the space category images.
}

type Spaces []Space

type Space struct {
	ID             string     `json:"Id"`             // Unique identifier of the space.
	Type           SpaceType  `json:"Type"`           // Type of the space.
	Number         string     `json:"Number"`         // Number of the space (e.g. room number).
	FloorNumber    string     `json:"FloorNumber"`    // Number of the floor the space is on.
	BuildingNumber string     `json:"BuildingNumber"` // Number of the building the space is in.
	ParentSpaceID  string     `json:"ParentSpaceId"`  // Identifier of the parent Space (e.g. room of a bed).
	CategoryID     string     `json:"CategoryId"`     // Identifier of the Space Category assigned to the space.
	State          SpaceState `json:"State"`          // State of the room.
}

type SpaceType string

type SpaceState string

const (
	SpaceStateDirty        SpaceState = "Dirty"
	SpaceStateClean        SpaceState = "Clean"
	SpaceStateInspected    SpaceState = "Inspected"
	SpaceStateOutOfService SpaceState = "OutOfService"
	SpaceStateOutOfOrder   SpaceState = "OutOfOrder"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return err
	}

	d.Time, err = time.Parse(time.RFC3339, value)
	return err
}

type Promotions struct {
	BeforeCheckIn  bool `json:"BeforeCheckIn"`  // Whether it can be promoted before check-in.
	AfterCheckIn   bool `json:"AfterCheckIn"`   // Whether it can be promoted after check-in.
	DuringStay     bool `json:"DuringStay"`     // Whether it can be promoted during stay.
	BeforeCheckOut bool `json:"BeforeCheckout"` // Whether it can be promoted before check-out.
	AfterCheckOut  bool `json:"AfterCheckout"`  // Whether it can be promoted after check-out.
}

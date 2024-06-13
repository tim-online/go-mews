package reservations

import (
	"time"

	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/orderitems"
)

const (
	endpointAdd = "reservations/add"
)

// Add customer
func (s *APIService) Add(requestBody *AddRequest) (*AddResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAdd)
	if err != nil {
		return nil, err
	}

	responseBody := &AddResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *APIService) NewAddRequest() *AddRequest {
	return &AddRequest{}
}

type AddRequest struct {
	json.BaseRequest

	ServiceID             string     `json:"ServiceId"`             // Unique identifier of the Service to be reserved.
	GroupID               string     `json:"GroupId"`               // Unique identifier of the Reservation group where the reservations are added. If not specified, a new group is created.
	GroupName             string     `json:"GroupName"`             // Name of the Reservation group which the reservations are added to. If GroupId is specified, this field is ignored. If not specified, the group name is automatically created.
	SendConfirmationEmail bool       `json:"SendConfirmationEmail"` // Whether the confirmation email is sent. Default value is true.
	Reservations          []struct { // Parameters of the new reservations.
		Identifier          string            `json:"Identifier"`          // Identifier of the reservation within the transaction.
		State               string            `json:"State"`               // State of the newly created reservation (either Optional, Enquired or Confirmed). If not specified, Confirmed is used.
		StartUtc            time.Time         `json:"StartUtc"`            // Reservation start in UTC timezone in ISO 8601 format.
		EndUtc              time.Time         `json:"EndUtc"`              // Reservation end in UTC timezone in ISO 8601 format.
		ReleasedUtc         time.Time         `json:"ReleasedUtc"`         // Release date and time of an unconfirmed reservation in UTC timezone in ISO 8601 format.
		CustomerID          string            `json:"CustomerId"`          // Unique identifier of the Customer who owns the reservation.
		BookerID            string            `json:"BookerId"`            // Unique identifier of the Customer on whose behalf the reservation was made.
		RequestedCategoryID string            `json:"RequestedCategoryId"` // Identifier of the requested Resource category.
		RateID              string            `json:"RateId"`              // Identifier of the reservation Rate.
		TravelAgencyID      string            `json:"TravelAgencyId"`      // Identifier of the Company that mediated the reservation.
		CompanyID           string            `json:"CompanyId"`           // Identifier of the Company on behalf of which the reservation was made.
		Notes               string            `json:"Notes"`               // Additional notes.
		TimeUnitAmount      orderitems.Amount `json:"TimeUnitAmount"`      // Amount of each night of the reservation.
		PersonCounts        []struct {
			AgeCategoryID string `json:"AgeCategoryId"` // Unique identifier of the Age category.
			Count         int    `json:"Count"`         // Number of people of a given age category. Only positive value is accepted.
		} `json:"PersonCounts"` // Number of people per age category the reservation was booked for. At least one category with valid count must be provided.
		TimeUnitPrices []struct {
			Index  int               `json:"Index"`  // Index of the unit. Indexing starts with 0. E.g the first night of the reservation has index 0.
			Amount orderitems.Amount `json:"Amount"` // Amount of the unit.
		} `json:"TimeUnitPrices"` // Prices for time units of the reservation. E.g. prices for the first or second night.
		ProductOrders []struct {
			ProductID string    `json:"ProductId"` // Unique identifier of the Product whose prices should be returned.
			StartUtc  time.Time `json:"StartUtc"`  // Start of the time interval, expressed as the timestamp for the start of the first time unit, in UTC timezone ISO 8601 format. See Time units.
			EndUtc    time.Time `json:"EndUtc"`    // End of the time interval, expressed as the timestamp for the start of the last time unit, in UTC timezone ISO 8601 format. See Time units. The maximum size of time interval depends on the service's time unit: 100 hours if hours, 100 days if days, or 24 months if months.
		} `json:"ProductOrders"` // Parameters of the products ordered together with the reservation.
		CreditCardID        string `json:"CreditCardId"`        // Identifier of Credit card belonging to Customer who owns the reservation.
		AvailabilityBlockID string `json:"AvailabilityBlockId"` // Unique identifier of the Availability block the reservation is assigned to.
		VoucherCode         string `json:"VoucherCode"`         // Voucher code value providing access to specified private Rate applied to this reservation.0
	} `json:"Reservations"`
}

type AddResponse struct {
	Reservations []struct {
		Identifier  string      `json:"Identifier"`  // Identifier of the reservation within the transaction.
		Reservation Reservation `json:"Reservation"` // The added reservations.
	} `json:"Reservations"`
}

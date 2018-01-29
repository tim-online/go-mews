package accountingitems

import (
	"errors"
	"time"
)

const (
	endpointAll = "accountingItems/getAll"

	ServiceRenue      AccountingItemType = "ServiceRevenue"
	ProductRevenue    AccountingItemType = "ProductRevenue"
	AdditionalRevenue AccountingItemType = "AdditionalRevenue"
	Payment           AccountingItemType = "Payment"
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
	AccountingItems []AccountingItem
}

func (s *Service) NewAllRequest() *AllRequest {
	return &AllRequest{}
}

type AllRequest struct {
	AccessToken string                    `json:"AccessToken"`
	StartUtc    *time.Time                `json:"StartUtc,omitempty"`
	EndUtc      *time.Time                `json:"EndUtc,omitempty`
	TimeFilter  AccountingItemsTimeFilter `json:"TimeFilter"`
}

type AccountingItemsTimeFilter string

const (
	TimeFilterClosed  AccountingItemsTimeFilter = "Closed"
	TimeFilterUpdated AccountingItemsTimeFilter = "Updated"
)

// 	"AccountingCategoryId": "4ac8ce68-5732-4f1d-bf0d-e557072c926f",
// 	"Amount": {
// 		"Currency": "GBP",
// 		"Tax": 0.42,
// 		"TaxRate": 0.2,
// 		"Value": 2.5
// 	},
// 	"BillId": null,
// 	"ConsumptionUtc": "2016-07-27T12:48:39Z",
// 	"Id": "89b93f7c-5c63-4de2-bd17-ec5fee5e3120",
// 	"Name": "Caramel, Pepper & Chilli Popcorn",
// 	"Notes": null,
// 	"OrderId": "810b8c3a-d358-4378-84a9-534c830016fc",
// 	"ProductId": null,
// 	"Type": "ServiceRevenue"
// }

type AccountingItem struct {
	ID                   string             `json:"Id"`                   // Unique identifier of the item.
	CustomerID           string             `json:"CustomerId"`           // Unique identifier of the Customer whose account the item belongs to.
	ProductID            string             `json:"ProductId"`            // Unique identifier of the Product.
	ServiceID            string             `json:"ServiceId"`            // Unique identifier of the Service the item belongs to.
	OrderID              string             `json:"OrderId"`              // Unique identifier of the order (or Reservation) the item belongs to.
	BillID               string             `json:"BillId"`               // Unique identifier of the bill the item is assigned to.
	AccountingCategoryID string             `json:"AccountingCategoryId"` // Unique identifier of the Accounting Category the item belongs to.
	Amount               Amount             `json:"Amount"`               // Amount the item costs, negative amount represents either rebate or a payment.
	Type                 AccountingItemType `json:"Type"`                 // Type of the item.
	Name                 string             `json:"Name"`                 // Name of the item.
	Notes                string             `json:"Notes"`                // Additional notes.
	ConsumptionUtc       time.Time          `json:"ConsumptionUtc"`       // Date and time of the item consumption in UTC timezone in ISO 8601 format.
}

type Amount struct {
	Currency string  `json:"Currency"`
	Tax      float64 `json:"Tax"`
	TaxRate  float64 `json:"TaxRate"`
	Value    float64 `json:"Value"`
}

type AccountingItemType string

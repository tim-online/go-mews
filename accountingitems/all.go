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
	AccessToken string     `json:"AccessToken"`
	StartUtc    *time.Time `json:"StartUtc,omitempty"`
	EndUtc      *time.Time `json:"EndUtc,omitempty`
}

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
	AccountingCategoryId string             `json:"AccountingCategoryId"`
	Amount               Amount             `json:"Amount"`
	BillID               string             `json:"BillId"`
	ConsumptionUtc       time.Time          `json:"ConsumptionUtc"`
	ID                   string             `json:"Id"`
	Name                 string             `json:"Name"`
	Notes                string             `json:"Notes"`
	OrderID              string             `json:"OrderId"`
	ProductID            string             `json:"ProductId"`
	Type                 AccountingItemType `json:"Type"`
}

type Amount struct {
	Currency string  `json:"Currency"`
	Tax      float64 `json:"Tax"`
	TaxRate  float64 `json:"TaxRate"`
	Value    float64 `json:"Value"`
}

type AccountingItemType string

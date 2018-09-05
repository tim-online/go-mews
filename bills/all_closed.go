package bills

import (
	"time"

	"github.com/tim-online/go-mews/json"
)

const (
	endpointAllClosed = "bills/getAllClosed"
)

// List all products
func (s *Service) All(requestBody *AllRequest) (*AllResponse, error) {
	// @TODO: create wrapper?
	if err := s.Client.CheckTokens(); err != nil {
		return nil, err
	}

	apiURL, err := s.Client.GetApiURL(endpointAllClosed)
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
	json.BaseRequest
	StartUTC *time.Time `json:"StartUtc,omitempty"`
	EndUTC   *time.Time `json:"EndUtc,omitempty"`
}

type AllResponse struct {
	Bills Bills `json:"Bills"` // The closed bills.
}

type Bills []Bill

type Bill struct {
	ID         string    `json:"Id"`         // Unique identifier of the bill.
	CustomerID string    `json:"CustomerId"` // Unique identifier of the Customer the bill is issued to.
	CounterID  string    `json:"CounterId"`  // Unique identifier of the Customer the bill is issued to.ompanyID  string    `json:"CompanyId"`  // Unique identifier of the Company the bill is issued to.
	Type       BillType  `json:"Type"`       // Type of the bill.
	Number     string    `json:"Number"`     // Number of the bill.
	IssuedUTC  time.Time `json:"IssuedUtc"`  // Date and time of the bill issuance in UTC timezone in ISO 8601 format.
	DueUTC     time.Time `json:"DueUtc"`     // Bill due date and time in UTC timezone in ISO 8601 format.
	Notes      string    `json:"Notes"`      // Additional notes.
	Revenue    Revenue   `json:"Revenue"`    // The revenue items on the bill.
	Payments   Payments  `json:"Payments"`   // The payments on the bill.
	State      BillState `json:"State"`      // State of the bill.
}

type BillType string

const (
	BillTypeReceipt BillType = "Receipt"
	BillTypeInvoice BillType = "Invoice"
)

type BillState string

const (
	BillStateOpen   BillState = "open"
	BillStateClosed BillState = "closed"
)

type Revenue []AccountingItem

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
	InvoiceID            string             `json:"InvoiceId"`            // Unique identifier of the invoiced Bill the item is receivable for.
	Notes                string             `json:"Notes"`                // Additional notes.
	ConsumptionUTC       time.Time          `json:"ConsumptionUtc"`       // Date and time of the item consumption in UTC timezone in ISO 8601 format.
	ClosedUTC            time.Time          `json:"ClosedUtc"`            // Date and time of the item bill closure in UTC timezone in ISO 8601 format.
}

type AccountingItemType string

const (
	AccountingItemTypeServiceRevenue    AccountingItemType = "ServiceRevenue"
	AccountingItemTypeProductRevenue    AccountingItemType = "ProductRevenue"
	AccountingItemTypeAdditionalRevenue AccountingItemType = "AdditionalRevenue"
	AccountingItemTypePayment           AccountingItemType = "Payment"
)

type Payments []AccountingItem

type Amount struct {
	Currency string  `json:"Currency"` // ISO-4217 currency code, e.g. EUR or USD.
	Net      float64 `json:"Net"`      // Net value in case the item is taxed.
	Tax      float64 `json:"Tax"`      // Tax value in case the item is taxed.
	TaxRate  float64 `json:"TaxRate"`  // Tax rate in case the item is taxed (e.g. 0.21).
	Value    float64 `json:"Value"`    // Amount in the currency (including tax if taxed).
}

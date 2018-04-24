package customers

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	endpointAllByIDs = "customers/getAllByIds"
)

var (
	ErrNoToken = errors.New("No token specified")
)

// List all products
func (s *Service) AllByIDs(requestBody *AllByIDsRequest) (*AllByIDsResponse, error) {
	// @TODO: create wrapper?
	// Set request token
	requestBody.AccessToken = s.Client.Token

	if s.Client.Token == "" {
		return nil, ErrNoToken
	}

	apiURL, err := s.Client.GetApiURL(endpointAllByIDs)
	if err != nil {
		return nil, err
	}

	responseBody := &AllByIDsResponse{}
	httpReq, err := s.Client.NewRequest(apiURL, requestBody)
	if err != nil {
		return nil, err
	}

	_, err = s.Client.Do(httpReq, responseBody)
	return responseBody, err
}

func (s *Service) NewAllByIDsRequest() *AllByIDsRequest {
	return &AllByIDsRequest{}
}

type AllByIDsRequest struct {
	AccessToken string   `json:"AccessToken"`
	CustomerIDs []string `json:"CustomerIds"`
}

type AllByIDsResponse struct {
	Customers Customers `json:"Customers"`
}

type Customers []Customer

type Customer struct {
	ID                      string           `json:"Id"`                      // Unique identifier of the customer.
	Number                  string           `json:"Number"`                  // Number of the customer.
	FirstName               string           `json:"FirstName"`               // First name of the customer.
	LastName                string           `json:"LastName"`                // Last name of the customer.
	SecondLastName          string           `json:"SecondLastName"`          // Second last name of the customer.
	Title                   Title            `json:"Title"`                   // Title prefix of the customer.
	Gender                  Gender           `json:"Gender"`                  // Gender of the customer.
	NationalityCode         string           `json:"NationalityCode"`         // ISO 3166-1 alpha-2 country code (two letter country code) of the nationality.
	LanguageCode            string           `json:"LanguageCode"`            // Language and culture code of the customers preferred language. E.g. en-US or fr-FR.
	BirthDate               string           `json:"BirthDate"`               // Date of birth in ISO 8601 format.
	BirthDateUTC            time.Time        `json:"BirthDateUtc"`            // ??
	BirthPlace              string           `json:"BirthPlace"`              // Place of birth.
	Email                   string           `json:"Email"`                   // Email address of the customer.
	Phone                   string           `json:"Phone"`                   // Phone number of the customer (possibly mobile).
	LoyaltyCode             string           `json:"LoyaltyCode"`             // Loyalty code of the customer.
	Classifications         []Classification `json:"Classifications"`         // Classifications of the customer.
	Passport                Document         `json:"Passport"`                // Passport details of the customer.
	IdentityCard            Document         `json:"IdentityCard"`            // IdentityCard details for Customer.
	Visa                    Document         `json:"Visa"`                    // Visa details for Customer.
	Address                 Address          `json:"Address"`                 // Address of the customer.
	CreatedUTC              time.Time        `json:"CreatedUtc"`              // Creation date and time of the customer in UTC timezone in ISO 8601 format.
	UpdatedUTC              time.Time        `json:"UpdatedUtc"`              // Last update date and time of the customer in UTC timezone in ISO 8601 format.
	TaxIdentificationNumber string           `json:"TaxIdentificationNumber"` // tax id customer
	CategoryID              string           `json:"CategoryId"`              // ??
	CitizenNumber           string           `json:"CitizenNumber"`           // ??
	FatherName              string           `json:"FatherName"`              // ??
	MotherName              string           `json:"MotherName"`              // ??
	Notes                   string           `json:"Notes"`                   // ??
	Occupation              string           `json:"Occupation"`              // ??
}

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

type Document struct {
	Number        string    `json:"Number"`        // Number of the document (e.g. passport number).
	Issuance      Date      `json:"Issuance"`      // Date of issuance in ISO 8601 format.
	Expiration    Date      `json:"Expiration"`    // Expiration date in ISO 8601 format.
	ExpirationUTC time.Time `json:"ExpirationUtc"` // ??
	IssuanceUTC   time.Time `json:"IssuanceUtc"`   // ??
}

type Classification string

type Address struct {
	Line1       string `json:"Line1"`       // First line of the address.
	Line2       string `json:"Line2"`       // Seconds line of the address.
	City        string `json:"City"`        // The City.
	PostalCode  string `json:"PostalCode"`  // Postal code.
	CountryCode string `json:"CountryCode"` // ISO 3166-1 alpha-2 country code (two letter country code).
}

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

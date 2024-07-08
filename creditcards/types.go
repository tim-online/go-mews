package creditcards

import "time"

type CreditCards []CreditCard

type CreditCard struct {
	// Unique identifier of the credit card.
	ID         string `json:"Id"`
	// Unique identifier of the Enterprise.
	EnterpriseID string `json:"EnterpriseId"`
	// Unique identifier of the credit card owner.
	CustomerID string `json:"CustomerId"`
	// Creation date and time of the credit card in UTC timezone in ISO 8601
	// format.
	CreatedUTC time.Time `json:"CreatedUtc"`
	// Last update date and time of the credit card in UTC timezone in ISO 8601 format.
	UpdatedUTC time.Time `json:"UpdatedUtc"`
	// Expiration of the credit card in format MM/YYYY.
	Expiration string `json:"Expiration"`
	// Whether the credit card is still active.
	IsActive bool `json:"IsActive"`
	// Obfuscated credit card number. At most first six digits and last four
	// digits can be specified, otherwise the digits are replaced with *.
	ObfuscatedNumber string `json:"ObfuscatedNumber"`
	// Format of the credit card.
	Format CreditCardFormat `json:"Format"`
	// Kind of the credit card.
	Kind CreditCardKind `json:"Kind"`
	// State of the credit card.
	State CreditCardState `json:"State"`
	// Type of the credit card.
	Type CreditCardType `json:"Type"`
}

// Physical
// Virtual
type CreditCardFormat string

// Terminal
// Gateway
type CreditCardKind string

// Enabled
// Disabled
type CreditCardState string

// MasterCard
// Visa
// Amex
// Maestro
// Discover
// VPay
// ...
type CreditCardType string

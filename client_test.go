package mews_test

import (
	"os"
	"testing"
	"time"

	mews "github.com/tim-online/go-mews"
	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/bills"
	"github.com/tim-online/go-mews/companies"
	"github.com/tim-online/go-mews/customers"
	"github.com/tim-online/go-mews/reservations"
)

func getClient() *mews.Client {
	// get username & password
	token := os.Getenv("MEWS_TOKEN")

	// build client
	client := mews.NewClient(nil, token)
	client.SetDebug(true)
	client.SetBaseURL(mews.BaseURLDemo)
	client.SetDisallowUnknownFields(true)

	return client
}

func TestBillsAll(t *testing.T) {
	client := getClient()
	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &bills.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	_, err := client.Bills.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestAccountingItems(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &accountingitems.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	_, err := client.AccountingItems.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestCompanies(t *testing.T) {
	client := getClient()

	requestBody := &companies.AllRequest{}
	_, err := client.Companies.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestCustomers(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &customers.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	requestBody.TimeFilter = customers.CustomerTimeFilterCreated
	_, err := client.Customers.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestReservations(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, -1, 0)
	endUTC := time.Now()

	requestBody := &reservations.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	requestBody.Extent = reservations.ReservationExtent{
		BusinessSegments:  true,
		Customers:         true,
		Items:             true,
		Products:          true,
		Rates:             true,
		Reservations:      true,
		ReservationGroups: true,
		Services:          true,
		Spaces:            true,
	}
	requestBody.TimeFilter = reservations.ReservationTimeFilterCreated
	_, err := client.Reservations.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

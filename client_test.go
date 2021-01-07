package mews_test

import (
	"os"
	"testing"
	"time"

	mews "github.com/tim-online/go-mews"
	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/reservations"
)

func getClient() *mews.Client {
	// get username & password
	accessToken := os.Getenv("MEWS_ACCESS_TOKEN")
	clientToken := os.Getenv("MEWS_CLIENT_TOKEN")

	// build client
	client := mews.NewClient(nil, accessToken, clientToken)
	client.SetDebug(true)
	// client.SetBaseURL(mews.BaseURLDemo)
	client.SetDisallowUnknownFields(true)

	return client
}

func TestAccountingItems(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, 0, -1)
	endUTC := time.Now()

	requestBody := &accountingitems.AllRequest{}
	requestBody.StartUTC = &startUTC
	requestBody.EndUTC = &endUTC
	requestBody.Extent.AccountingItems = true
	_, err := client.AccountingItems.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestReservations(t *testing.T) {
	client := getClient()

	startUTC := time.Now().AddDate(0, 0, -1)
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
		Resources:         true,
	}
	requestBody.TimeFilter = reservations.ReservationTimeFilterCreated
	_, err := client.Reservations.All(requestBody)
	if err != nil {
		t.Error(err)
	}
}

func TestConfig(t *testing.T) {
	client := getClient()

	requestBody := client.Configuration.NewGetRequest()
	_, err := client.Configuration.Get(requestBody)
	if err != nil {
		t.Error(err)
	}
}

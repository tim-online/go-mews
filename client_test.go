package mews

import (
	"fmt"
	"os"
	"testing"

	"github.com/tim-online/go-mews/bills"
)

func TestDing(t *testing.T) {
	// get username & password
	token := os.Getenv("MEWS_TOKEN")

	// build client
	client := NewClient(nil, token)
	client.SetDebug(true)
	client.SetBaseURL(BaseURL)

	requestBody := &bills.AllByIDsRequest{}
	requestBody.BillIDs = []string{"79931cf5-e483-4738-bbc9-8835009db19c"}
	resp, err := client.Bills.AllByIDs(requestBody)
	if err != nil {
		panic(err)
	}

	bills := resp.Bills
	for _, bill := range bills {
		fmt.Printf("%+v\n", bill)
	}

	// // request all companies this token has access to
	// format := "2006-01-02"

	// startUtc, err := time.Parse(format, "2016-01-01")
	// if err != nil {
	// 	panic(err)
	// }

	// endUtc, err := time.Parse(format, "2017-01-01")
	// if err != nil {
	// 	panic(err)
	// }

	// requestBody := &accountingitems.AllRequest{
	// 	StartUtc: &startUtc,
	// 	EndUtc:   &endUtc,
	// }
	// resp, err := client.AccountingItems.All(requestBody)
	// if err != nil {
	// 	panic(err)
	// }

	// items := resp.AccountingItems
	// for _, item := range items {
	// 	fmt.Printf("%+v\n", item)
	// }
}

package mews

import (
	"fmt"
	"os"
	"testing"

	"github.com/tim-online/go-mews/accountingcategories"
)

func TestDing(t *testing.T) {
	// get username & password
	token := os.Getenv("MEWS_TOKEN")

	// build client
	client := NewClient(nil, token)
	client.SetDebug(true)
	client.SetBaseURL(baseURLDemo)

	requestBody := &accountingcategories.AllRequest{}
	resp, err := client.AccountingCategories.All(requestBody)
	if err != nil {
		panic(err)
	}

	categories := resp.AccountingCategories
	for _, cat := range categories {
		fmt.Printf("%+v\n", cat)
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

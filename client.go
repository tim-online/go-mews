package mews

import (
	"net/http"
	"net/url"

	"github.com/tim-online/go-mews/accountingcategories"
	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/bills"
	"github.com/tim-online/go-mews/businesssegments"
	"github.com/tim-online/go-mews/companies"
	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/customers"
	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/reservations"
	"github.com/tim-online/go-mews/spaceblocks"
	"github.com/tim-online/go-mews/spaces"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-mews/" + libraryVersion
)

var (
	BaseURL = &url.URL{
		Scheme: "https",
		Host:   "www.mews.li",
		Path:   "/api/connector/v1/",
	}
	BaseURLDemo = &url.URL{
		Scheme: "https",
		Host:   "demo.mews.li",
		Path:   "/api/connector/v1/",
	}
)

// NewClient returns a new MEWS API client
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	jsonClient := json.NewClient(httpClient, token)
	jsonClient.UserAgent = userAgent
	jsonClient.Token = token
	jsonClient.Debug = false

	c := &Client{
		client: jsonClient,
	}

	c.SetBaseURL(BaseURL)

	// Services
	c.AccountingItems = accountingitems.NewService()
	c.AccountingItems.Client = c.client
	c.AccountingCategories = accountingcategories.NewService()
	c.AccountingCategories.Client = c.client
	c.Companies = companies.NewService()
	c.Companies.Client = c.client
	c.Customers = customers.NewService()
	c.Customers.Client = c.client
	c.Reservations = reservations.NewAPIService()
	c.Reservations.Client = c.client
	c.Spaces = spaces.NewService()
	c.Spaces.Client = c.client
	c.Bills = bills.NewService()
	c.Bills.Client = c.client
	c.Commands = commands.NewService()
	c.Commands.Client = c.client
	c.Configuration = configuration.NewService()
	c.Configuration.Client = c.client
	c.BusinessSegments = businesssegments.NewService()
	c.BusinessSegments.Client = c.client

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	client *json.Client

	// Services used for communicating with the API
	AccountingItems      *accountingitems.Service
	AccountingCategories *accountingcategories.Service
	Companies            *companies.Service
	Customers            *customers.Service
	Reservations         *reservations.APIService
	Spaces               *spaces.Service
	Bills                *bills.Service
	Commands             *commands.Service
	Configuration        *configuration.Service
	BusinessSegments     *businesssegments.Service
}

func (c *Client) SetDebug(debug bool) {
	c.client.Debug = debug
}

func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.client.BaseURL = baseURL
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.client.DisallowUnknownFields = disallowUnknownFields
}

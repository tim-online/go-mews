package mews

import (
	"context"
	"net/http"
	"net/url"

	"github.com/tim-online/go-mews/accountingcategories"
	"github.com/tim-online/go-mews/accountingitems"
	"github.com/tim-online/go-mews/bills"
	"github.com/tim-online/go-mews/businesssegments"
	"github.com/tim-online/go-mews/commands"
	"github.com/tim-online/go-mews/companies"
	"github.com/tim-online/go-mews/configuration"
	"github.com/tim-online/go-mews/customers"
	"github.com/tim-online/go-mews/json"
	"github.com/tim-online/go-mews/outletitems"
	"github.com/tim-online/go-mews/outlets"
	"github.com/tim-online/go-mews/products"
	"github.com/tim-online/go-mews/reservations"
	"github.com/tim-online/go-mews/services"
	"github.com/tim-online/go-mews/spaceblocks"
	"github.com/tim-online/go-mews/spaces"
	"github.com/tim-online/go-mews/tasks"
	"github.com/tim-online/go-mews/taxenvironments"
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
func NewClient(httpClient *http.Client, accessToken string, clientToken string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	jsonClient := json.NewClient(httpClient, accessToken, clientToken)
	jsonClient.UserAgent = userAgent
	jsonClient.AccessToken = accessToken
	jsonClient.ClientToken = clientToken
	jsonClient.Debug = false

	c := &Client{
		client: jsonClient,
	}

	c.SetBaseURL(BaseURL)

	// Services
	c.AccountingItems = accountingitems.NewService()
	c.AccountingItems.Client = c.client
	c.OutletItems = outletitems.NewService()
	c.OutletItems.Client = c.client
	c.AccountingCategories = accountingcategories.NewService()
	c.AccountingCategories.Client = c.client
	c.Companies = companies.NewService()
	c.Companies.Client = c.client
	c.Customers = customers.NewService()
	c.Customers.Client = c.client
	c.Outlets = outlets.NewAPIService()
	c.Outlets.Client = c.client
	c.Products = products.NewAPIService()
	c.Products.Client = c.client
	c.Reservations = reservations.NewAPIService()
	c.Reservations.Client = c.client
	c.Spaces = spaces.NewService()
	c.Spaces.Client = c.client
	c.Services = services.NewAPIService()
	c.Services.Client = c.client
	c.SpaceBlocks = spaceblocks.NewService()
	c.SpaceBlocks.Client = c.client
	c.Bills = bills.NewService()
	c.Bills.Client = c.client
	c.Commands = commands.NewService()
	c.Commands.Client = c.client
	c.Configuration = configuration.NewService()
	c.Configuration.Client = c.client
	c.BusinessSegments = businesssegments.NewService()
	c.BusinessSegments.Client = c.client
	c.Tasks = tasks.NewService()
	c.Tasks.Client = c.client
	c.TaxEnvironments = taxenvironments.NewService()
	c.TaxEnvironments.Client = c.client

	return c
}

// Client manages communication with MEWS API
type Client struct {
	// HTTP client used to communicate with the API.
	client *json.Client

	// Services used for communicating with the API
	AccountingItems      *accountingitems.APIService
	OutletItems          *outletitems.Service
	AccountingCategories *accountingcategories.Service
	Companies            *companies.Service
	Customers            *customers.Service
	Outlets              *outlets.APIService
	Products             *products.APIService
	Reservations         *reservations.APIService
	Services             *services.APIService
	Spaces               *spaces.Service
	SpaceBlocks          *spaceblocks.Service
	Bills                *bills.Service
	Commands             *commands.Service
	Configuration        *configuration.Service
	BusinessSegments     *businesssegments.Service
	Tasks                *tasks.Service
	TaxEnvironments      *taxenvironments.Service
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

func (c *Client) SetLanguageCode(code string) {
	c.client.SetLanguageCode(code)
}

func (c *Client) SetCultureCode(code string) {
	c.client.SetCultureCode(code)
}

func (c *Client) GetWebsocket(ctx context.Context) *Websocket {
	ws := NewWebsocket(c.client.Client, c.client.AccessToken, c.client.ClientToken)
	url := &url.URL{
		Scheme: WebsocketURL.Scheme,
		Host:   c.client.BaseURL.Host,
		Path:   WebsocketURL.Path,
	}
	ws.SetBaseURL(url)
	ws.SetDebug(c.client.Debug)
	return ws
}

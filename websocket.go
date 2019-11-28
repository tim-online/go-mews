package mews

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tim-online/go-mews/commands"
	"github.com/tim-online/go-mews/reservations"
	"github.com/tim-online/go-mews/spaces"
)

var (
	WebsocketURL = &url.URL{
		Scheme: "wss",
		Host:   "www.mews.li",
		Path:   "/ws/connector",
	}
	WebsocketURLDemo = &url.URL{
		Scheme: "wss",
		Host:   "demo.mews.li",
		Path:   "/ws/connector",
	}
)

type Websocket struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// Base URL for API requests
	baseURL *url.URL

	// Debugging flag
	debug bool

	// Disallow unknown json fields
	disallowUnknownFields bool

	accessToken string
	clientToken string

	connection *websocket.Conn

	doneChan chan struct{}
	msgChan  chan []byte
	errChan  chan error

	cmdChan         chan CommandEvent
	resChan         chan ReservationEvent
	spaceChan       chan SpaceEvent
	priceUpdateChan chan PriceUpdateEvent
}

func NewWebsocket(httpClient *http.Client, accessToken string, clientToken string) *Websocket {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	ws := &Websocket{}
	ws.SetAccessToken(accessToken)
	ws.SetClientToken(clientToken)
	ws.SetDebug(false)
	ws.SetBaseURL(WebsocketURL)

	ws.doneChan = make(chan struct{})
	ws.msgChan = make(chan []byte)
	ws.errChan = make(chan error)
	ws.cmdChan = make(chan CommandEvent)
	ws.resChan = make(chan ReservationEvent)
	ws.spaceChan = make(chan SpaceEvent)
	ws.priceUpdateChan = make(chan PriceUpdateEvent)

	return ws
}

func (ws Websocket) AccessToken() string {
	return ws.accessToken
}

func (ws *Websocket) SetAccessToken(accessToken string) {
	ws.accessToken = accessToken
}

func (ws Websocket) ClientToken() string {
	return ws.clientToken
}

func (ws *Websocket) SetClientToken(clientToken string) {
	ws.clientToken = clientToken
}

func (ws *Websocket) BaseURL() *url.URL {
	return ws.baseURL
}

func (ws *Websocket) SetBaseURL(baseURL *url.URL) {
	ws.baseURL = baseURL
	ws.baseURL.Scheme = "wss"
}

func (ws *Websocket) Debug() bool {
	return ws.debug
}

func (ws *Websocket) SetDebug(debug bool) {
	ws.debug = debug
}

func (ws *Websocket) CommandEvents() chan (CommandEvent) {
	return ws.cmdChan
}

func (ws *Websocket) ReservationEvents() chan (ReservationEvent) {
	return ws.resChan
}

func (ws *Websocket) SpaceEvents() chan (SpaceEvent) {
	return ws.spaceChan
}

func (ws *Websocket) PriceUpdateEvents() chan (PriceUpdateEvent) {
	return ws.priceUpdateChan
}

func (ws Websocket) Connect() error {
	var err error

	u := ws.BaseURL()
	q := u.Query()
	q.Add("ClientToken", ws.ClientToken())
	q.Add("AccessToken", ws.AccessToken())
	u.RawQuery = q.Encode()

	resp := &http.Response{}
	ws.connection, resp, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	log.Println(string(dump))

	ws.connection.SetCloseHandler(func(code int, text string) error {
		log.Println(code)
		log.Println(text)
		return nil
	})

	ws.connection.SetPingHandler(func(appData string) error {
		log.Println(appData)
		return nil
	})

	done := make(chan struct{})

	// read messages
	go func() {
		defer close(done)
		for {
			_, msg, err := ws.connection.ReadMessage()
			if err != nil {
				ws.errChan <- err
				return
			}

			events := Events{}
			err = json.Unmarshal(msg, &events)
			log.Println(string(msg))
			continue
			if err != nil {
				ws.errChan <- err
				return
			}

			for _, event := range events {
				if event.Type == EventTypeDeviceCommand && ws.cmdChan != nil {
					cmdEvent := CommandEvent{}
					err := json.Unmarshal(msg, &cmdEvent)
					if err != nil {
						ws.errChan <- err
						return
					}
					ws.cmdChan <- cmdEvent
				}

				if event.Type == EventTypeReservation && ws.resChan != nil {
					resEvent := ReservationEvent{}
					err := json.Unmarshal(msg, &resEvent)
					if err != nil {
						ws.errChan <- err
						return
					}
					ws.resChan <- resEvent
				} else if event.Type == EventTypeSpace && ws.spaceChan != nil {
					spaceEvent := SpaceEvent{}
					err := json.Unmarshal(msg, &spaceEvent)
					if err != nil {
						ws.errChan <- err
						return
					}
					ws.spaceChan <- spaceEvent
				} else if event.Type == EventTypePriceUpdate && ws.priceUpdateChan != nil {
					priceUpdateEvent := PriceUpdateEvent{}
					err := json.Unmarshal(msg, &priceUpdateEvent)
					if err != nil {
						ws.errChan <- err
						return
					}
					ws.priceUpdateChan <- priceUpdateEvent
				}
			}
		}
	}()

	return nil
}

func (ws *Websocket) ReadMessages() {
}

func (ws *Websocket) Stop() {
	ws.doneChan <- struct{}{}
	ws.connection.Close()
}

type Events []Event

type Event struct {
	Type  EventType             `json:"Type"`  // Type of the event.
	ID    string                `json:"Id"`    // Unique identifier of the Command.
	State commands.CommandState `json:"State"` // State of the command.
}

type EventType string

const (
	EventTypeDeviceCommand EventType = "DeviceCommand"
	EventTypeReservation   EventType = "Reservation"
	EventTypeSpace         EventType = "Space"
	EventTypePriceUpdate   EventType = "PriceUpdate"
)

type CommandEvent struct {
	Event
}

type ReservationEvent struct {
	Event

	ID              string                        `json:"Id"`              // Unique identifier of the Reservation.
	State           reservations.ReservationState `json:"State"`           // State of the reservation.
	StartUTC        time.Time                     `json:"StartUtc"`        // Start of the reservation (arrival) in UTC timezone in ISO 8601 format.
	EndUTC          time.Time                     `json:"EndUtc"`          // End of the reservation (departure) in UTC timezone in ISO 8601 format.
	AssignedSpaceID string                        `json:"AssignedSpaceId"` // Unique identifier of the operations/enterprises#space assigned to the reservation.
}

type SpaceEvent struct {
	Event

	State spaces.SpaceState `json:"State"` // State of the space.

}

type PriceUpdateEvent struct {
	Event

	StartUTC        time.Time `json:"StartUtc"`        // Start of the price update interval in UTC timezone in ISO 8601 format.
	EndUtc          time.Time `json:"EndUtc"`          // End of the price update interval in UTC timezone in ISO 8601 format.
	RateID          string    `json:"RateId"`          // Unique identifier of the Rate assigned to the update price event.
	SpaceCategoryID string    `json:"SpaceCategoryId"` // Unique identifier of the Space category assigned to the update price event.
}

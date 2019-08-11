package numbers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const (
	// OrderStatusComplete is a completed order.
	OrderStatusComplete = "COMPLETE"
	// OrderStatusFailed is a failed order.
	OrderStatusFailed = "FAILED"
	// OrderStatusReceived is a received order.
	OrderStatusReceived = "RECEIVED"
)

// ExistingTelephoneNumberOrderType is an order type defining existing phone numbers normally found
// using the GetAvailableNumbers.
type ExistingTelephoneNumberOrderType struct {
	XMLName          xml.Name `xml:"ExistingTelephoneNumberOrderType"`
	TelephoneNumbers []string `xml:"TelephoneNumberList>TelephoneNumber"`
	ReservationIDs   []string `xml:"ReservationIdList>ReservationId"`
}

// AreaCodeSearchAndOrderType is an order type defining an order based on an area code.
type AreaCodeSearchAndOrderType struct {
	XMLName  xml.Name `xml:"AreaCodeSearchAndOrderType"`
	AreaCode int      `xml:"AreaCode"`
	Quantity int      `xml:"Quantity"`
}

// RateCenterSearchAndOrderType is an order type defining an order based on rate center.
type RateCenterSearchAndOrderType struct {
	XMLName    xml.Name `xml:"RateCenterSearchAndOrderType"`
	RateCenter string   `xml:"RateCenter"`
	State      string   `xml:"State"`
	Quantity   int      `xml:"Quantity"`
}

// TollFreeVanitySearchAndOrderType is an order type defining an order based on toll-free vanity search.
type TollFreeVanitySearchAndOrderType struct {
	XMLName        xml.Name `xml:"TollFreeVanitySearchAndOrderType"`
	TollFreeVanity string   `xml:"TollFreeVanity"`
	Quantity       int      `xml:"Quantity"`
}

// TollFreeWildCharSearchAndOrderType is an order type defining an order based on toll-free wildcard search.
type TollFreeWildCharSearchAndOrderType struct {
	XMLName                 xml.Name `xml:"TollFreeWildCharSearchAndOrderType"`
	TollFreeWildCardPattern string   `xml:"TollFreeWildCardPattern"`
	Quantity                int      `xml:"Quantity"`
}

// StateSearchAndOrderType is an order type defining an order based on state.
type StateSearchAndOrderType struct {
	XMLName  xml.Name `xml:"StateSearchAndOrderType"`
	State    string   `xml:"State"`
	Quantity int      `xml:"Quantity"`
}

// CitySearchAndOrderType is an order type defining an order based on city.
type CitySearchAndOrderType struct {
	XMLName  xml.Name `xml:"CitySearchAndOrderType"`
	City     string   `xml:"City"`
	State    string   `xml:"State"`
	Quantity int      `xml:"Quantity"`
}

// ZIPSearchAndOrderType is an order type defining an order based on zip code.
type ZIPSearchAndOrderType struct {
	XMLName  xml.Name `xml:"ZIPSearchAndOrderType"`
	Zip      string   `xml:"Zip"`
	Quantity int      `xml:"Quantity"`
}

// LATASearchAndOrderType is an order type defining an order based on lata.
type LATASearchAndOrderType struct {
	XMLName  xml.Name `xml:"LATASearchAndOrderType"`
	Lata     string   `xml:"Lata"`
	Quantity int      `xml:"Quantity"`
}

// CombinedSearchAndOrderType is an order type defining an order based on a combination of search parameters.
type CombinedSearchAndOrderType struct {
	XMLName    xml.Name `xml:"CombinedSearchAndOrderType"`
	Quantity   int      `xml:"Quantity"`
	AreaCode   int      `xml:"AreaCode,omitempty"`
	RateCenter string   `xml:"RateCenter,omitempty"`
	State      string   `xml:"State,omitempty"`
	Lata       string   `xml:"Lata,omitempty"`
	City       string   `xml:"City,omitempty"`
	Zip        string   `xml:"Zip,omitempty"`
}

// Order is an order for a telephone number.
type Order struct {
	XMLName            xml.Name   `xml:"Order"`
	ID                 string     `xml:"id,omitempty"`
	CustomerOrderID    string     `xml:"CustomerOrderId"`
	Name               string     `xml:"Name"`
	OrderCreatedDate   *time.Time `xml:"OrderCreateDate,omitempty"`
	PeerID             string     `xml:"PeerId,omitempty"`
	BackOrderRequested bool       `xml:"BackOrderRequested"`
	TnAttributes       []string   `xml:"TnAttributes,omitempty>TnAttribute,omitempty"`
	PartialAllowed     bool       `xml:"PartialAllowed"`
	SiteID             string     `xml:"SiteId"`

	// Different order types for the order, only one of the following should be
	// and will be defined for an Order.
	ExistingTelephoneNumberOrder   *ExistingTelephoneNumberOrderType   `xml:"ExistingTelephoneNumberOrderType,omitempty"`
	AreaCodeSearchAndOrder         *AreaCodeSearchAndOrderType         `xml:"AreaCodeSearchAndOrderType,omitempty"`
	RateCenterSearchAndOrder       *RateCenterSearchAndOrderType       `xml:"RateCenterSearchAndOrderType,omitempty"`
	TollFreeVanitySearchAndOrder   *TollFreeVanitySearchAndOrderType   `xml:"TollFreeVanitySearchAndOrderType,omitempty"`
	TollFreeWildCharSearchAndOrder *TollFreeWildCharSearchAndOrderType `xml:"TollFreeWildCharSearchAndOrderType,omitempty"`
	StateSearchAndOrder            *StateSearchAndOrderType            `xml:"StateSearchAndOrderType,omitempty"`
	CitySearchAndOrder             *CitySearchAndOrderType             `xml:"CitySearchAndOrderType,omitempty"`
	ZIPSearchAndOrder              *ZIPSearchAndOrderType              `xml:"ZIPSearchAndOrderType,omitempty"`
	LATASearchAndOrder             *LATASearchAndOrderType             `xml:"LATASearchAndOrderType,omitempty"`
	CombinedSearchAndOrder         *CombinedSearchAndOrderType         `xml:"CombinedSearchAndOrderType,omitempty"`
}

// GetOrderResponse is the response for GetOrder.
type GetOrderResponse struct {
	XMLName           xml.Name   `xml:"OrderResponse"`
	CompletedQuantity int        `xml:"CompletedQuantity"`
	CreatedByUser     string     `xml:"CreatedByUser"`
	LastModifiedDate  *time.Time `xml:"LastModifiedDate"`
	OrderCompleteDate *time.Time `xml:"OrderCompleteDate"`
	Order             Order      `xml:"Order"`
	OrderStatus       string     `xml:"OrderStatus"`
	CompletedNumbers  []string   `xml:"CompletedNumbers>TelephoneNumber>FullNumber"`
	FailedQuantity    int        `xml:"FailedQuantity"`
}

// GetOrder gets an order.
func (c *Client) GetOrder(orderID string) (*GetOrderResponse, error) {
	respBody := &GetOrderResponse{}
	err := c.makeRequest(http.MethodGet, fmt.Sprintf("orders/%s", orderID), nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// CreateOrderResponse is the response for CreateOrder.
type CreateOrderResponse struct {
	XMLName xml.Name `xml:"OrderResponse"`
	Order   Order    `xml:"Order"`
}

// CreateOrder creates a new order.
func (c *Client) CreateOrder(data *Order) (*CreateOrderResponse, error) {
	respBody := &CreateOrderResponse{}
	err := c.makeRequest(http.MethodPost, "orders", data, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

package numbers

import (
	"encoding/xml"
	"net/http"
	"time"
)

const (
	// DisconnectModeNormal used on disconnect to set priority of disconnect.
	DisconnectModeNormal = "normal"
	// ProtectedTrue used on disconnected to mark the telephone numbers protected.
	ProtectedTrue = "TRUE"
	// ProtectedFalse used on disconnected to mark the telephone numbers as not protected.
	ProtectedFalse = "FALSE"
	// ProtectedUnchanged used on disconnect to to leave telephone numbers protected status unchanged.
	ProtectedUnchanged = "UNCHANGED"
)

// DisconnectTelephoneNumberOrderType is an order type to disconnect a list of telephone numbers.
type DisconnectTelephoneNumberOrderType struct {
	XMLName          xml.Name `xml:"DisconnectTelephoneNumberOrderType"`
	TelephoneNumbers []string `xml:"TelephoneNumberList>TelephoneNumber"`
	DisconnectMode   string   `xml:"DisconnectMode,omitempty"`
	Protected        string   `xml:"UNCHANGED,omitempty"`
}

// DisconnectTelephoneNumberOrder is an order to disconnect a list of telephone numbers.
type DisconnectTelephoneNumberOrder struct {
	XMLName                        xml.Name                            `xml:"DisconnectTelephoneNumberOrder"`
	ID                             string                              `xml:"id,omitempty"`
	Name                           string                              `xml:"name,omitempty"`
	CustomerOrderID                string                              `xml:"CustomerOrderID,omitempty"`
	OrderCreateDate                *time.Time                          `xml:"OrderCreateDate,omitempty"`
	DisconnectTelephoneNumberOrder *DisconnectTelephoneNumberOrderType `xml:"DisconnectTelephoneNumberOrderType"`
}

// DisconnectTelephoneNumberOrderResponse is the response for CreateDisconnectOrder.
type DisconnectTelephoneNumberOrderResponse struct {
	XMLName      xml.Name                       `xml:"DisconnectTelephoneNumberOrderResponse"`
	OrderRequest DisconnectTelephoneNumberOrder `xml:"orderRequest"`
}

// CreateDisconnectOrder creates a new disconnect order.
func (c *Client) CreateDisconnectOrder(data *DisconnectTelephoneNumberOrder) (*DisconnectTelephoneNumberOrderResponse, error) {
	respBody := &DisconnectTelephoneNumberOrderResponse{}
	err := c.makeRequest(http.MethodPost, "disconnects", data, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

package numbers

import (
	"encoding/xml"
	"net/http"
)

// Address for an Account.
type Address struct {
	XMLName          xml.Name `xml:"Address"`
	HouseNumber      int      `xml:"HouseNumber"`
	HouseSuffix      string   `xml:"HouseSuffix"`
	PreDirectional   string   `xml:"PreDirectional"`
	StreetName       string   `xml:"StreetName"`
	StreetSuffix     string   `xml:"StreetSuffix"`
	PostDirectional  string   `xml:"PostDirectional"`
	AddressLine2     string   `xml:"AddressLine2"`
	City             string   `xml:"City"`
	StateCode        string   `xml:"StateCode"`
	Zip              int      `xml:"Zip"`
	PlusFour         int      `xml:"PlusFour"`
	County           string   `xml:"County"`
	Country          string   `xml:"Country"`
	AddressType      string   `xml:"AddressType"`
	EndpointCount    int      `xml:"EndpointCount"`
	ValidationStatus string   `xml:"ValidationStatus"`
}

// AddressesResponse is the response for GetAddresses.
type AddressesResponse struct {
	XMLName    xml.Name   `xml:"AddressesResponse"`
	TotalCount int        `xml:"TotalCount"`
	Addresses  []*Address `xml:"Addresses"`
}

// GetAddressesQuery is query parameters of GetAddresses.
type GetAddressesQuery struct {
	E911LocationID string `url:"e911locationid,omitempty"`
	Page           int    `url:"page,omitempty"`
	Size           int    `url:"size,omitempty"`
	Suggestions    string `url:"suggestions,omitempty"`
	Type           string `url:"type,omitempty"`
}

// GetAddresses returns addresses for the account.
func (c *Client) GetAddresses(query *GetAddressesQuery) (*AddressesResponse, error) {
	respBody := &AddressesResponse{}
	err := c.makeRequest(http.MethodGet, "addresses", query, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

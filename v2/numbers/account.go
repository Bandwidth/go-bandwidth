package numbers

import (
	"encoding/xml"
	"net/http"
)

// Account information
type Account struct {
	XMLName                   xml.Name `xml:"Account"`
	AccountID                 string   `xml:"AccountId"`
	AssociatedCatapultAccount string   `xml:"AssociatedCatapultAccount"`
	GlobalAccountNumber       string   `xml:"GlobalAccountNumber"`
	CompanyName               string   `xml:"CompanyName"`
	AccountType               string   `xml:"AccountType"`
	NenaID                    string   `xml:"NenaId"`
	CustomerSegment           string   `xml:"CustomerSegment"`
	Tiers                     []int    `xml:"Tiers>Tier"`
	Address                   Address  `xml:"Address"`
	Contact                   Contact  `xml:"Contact"`
	AltSpid                   string   `xml:"AltSpid"`
	SPID                      string   `xml:"SPID"`
	PortCarrierType           string   `xml:"PortCarrierType"`
}

// AccountResponse is the response for GetAccount.
type AccountResponse struct {
	XMLName xml.Name `xml:"AccountResponse"`
	Account Account  `xml:"Account"`
}

// GetAccount returns account information.
func (c *Client) GetAccount() (*AccountResponse, error) {
	respBody := &AccountResponse{}
	err := c.makeRequest(http.MethodGet, "", nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

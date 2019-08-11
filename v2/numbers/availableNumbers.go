package numbers

import (
	"encoding/xml"
	"net/http"
)

// AvailableNumbersResponse is the response for GetAvailableNumbers.
type AvailableNumbersResponse struct {
	XMLName          xml.Name `xml:"SearchResult"`
	ResultCount      int      `xml:"ResultCount"`
	TelephoneNumbers []string `xml:"TelephoneNumberList>TelephoneNumber"`
}

// GetAvailableNumbersQuery is the query parameters for GetAvailableNumbers.
type GetAvailableNumbersQuery struct {
	LCA                     bool   `url:"LCA,omitempty"`
	AreaCode                int    `url:"areaCode,omitempty"`
	City                    string `url:"city,omitempty"`
	EndsIn                  bool   `url:"endsIn,omitempty"`
	Lata                    int    `url:"lata,omitempty"`
	LocalVanity             string `url:"localVanity,omitempty"`
	OrderBy                 string `url:"orderBy,omitempty"`
	Quantity                int    `url:"quantity"`
	RateCenter              string `url:"rateCenter,omitempty"`
	State                   string `url:"state,omitempty"`
	TollFreeVanity          string `url:"tollFreeVanity,omitempty"`
	TollFreeWildCardPattern string `url:"tollFreeWildCardPattern,omitempty"`
	Zip                     int    `url:"zip,omitempty"`
}

// GetAvailableNumbers returns a list of available numbers based on the provided query.
func (c *Client) GetAvailableNumbers(query *GetAvailableNumbersQuery) (*AvailableNumbersResponse, error) {
	respBody := &AvailableNumbersResponse{}
	err := c.makeRequest(http.MethodGet, "availableNumbers", query, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

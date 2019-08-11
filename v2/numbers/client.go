package numbers

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	// DefaultEndpoint is the default endpoint to communicate with Bandwidth Numbers V2 API.
	DefaultEndpoint = "https://dashboard.bandwidth.com"
)

// Client provides a client to communicate with the Bandwidth Numbers V2 API.
type Client struct {
	accountID string
	username  string
	password  string

	endpoint string
	client   *http.Client
}

// NewClient makes a new client to communicate with the Bandwidth Numbers V2 API.
func NewClient(accountID string, username string, password string, endpoint string, client *http.Client) (*Client, error) {
	if accountID == "" || username == "" || password == "" || endpoint == "" || client == nil {
		return nil, errors.New(`Missing auth data. Please use api := numbers.NewClient("accountId", "username", "password", "endpoint", http.DefaultClient)`)
	}
	return &Client{
		accountID: accountID,
		username:  username,
		password:  password,
		endpoint:  strings.TrimRight(endpoint, "/"),
		client:    client,
	}, nil
}

// NewDefaultClient makes a new client communicating with the Bandwidth Numbers V2 API at the
// DefaultEndpoint and with the http.DefaultClient.
func NewDefaultClient(accountID string, username string, password string) (*Client, error) {
	if accountID == "" || username == "" || password == "" {
		return nil, errors.New(`Missing auth data. Please use api := numbers.NewDefaultClient("accountId", "username", "password")`)
	}
	return NewClient(accountID, username, password, DefaultEndpoint, http.DefaultClient)
}

// ClientError is an error returned when performing a client operation and the
// API responded with an error.
type ClientError struct {
	XMLName     xml.Name `xml:"Error"`
	StatusCode  int      `xml:"-"`
	Code        string   `xml:"Code"`
	Description string   `xml:"Description"`
}

func (e *ClientError) Error() string {
	if e.Code != "" && e.Description != "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Description)
	} else if e.Code != "" {
		return fmt.Sprintf("%s", e.Code)
	}
	return fmt.Sprintf("HTTP error %d", e.StatusCode)
}

// prepareURL constructs the request URL.
func (c *Client) prepareURL(path string) string {
	return strings.TrimRight(fmt.Sprintf("%s/api/accounts/%s/%s", c.endpoint, c.accountID, strings.TrimLeft(path, "/")), "/")
}

// createRequest creates the HTTP request.
func (c *Client) createRequest(method string, path string, data interface{}) (*http.Request, error) {
	request, err := http.NewRequest(method, c.prepareURL(path), nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.username, c.password)
	request.Header.Set("Accept", "application/xml")
	request.Header.Set("User-Agent", "go-bandwidth/v2/numbers")
	if method == http.MethodGet && data != nil {
		urlValues, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		request.URL.RawQuery = urlValues.Encode()
	} else if (method == http.MethodPost || method == http.MethodPut) && data != nil {
		body, err := xml.Marshal(&data)
		if err != nil {
			return nil, err
		}
		request.Body = nopCloser{bytes.NewReader(body)}
		request.Header.Set("Content-Type", "application/xml")
	}
	return request, nil
}

type errorResponse struct {
	Error ClientError `xml:"Error"`
}

// parseResponse parsed the HTTP response into the responseBody
func (c *Client) parseResponse(response *http.Response, responseBody interface{}) error {
	defer response.Body.Close()
	rawXML, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		if len(rawXML) > 0 {
			return xml.Unmarshal(rawXML, &responseBody)
		}
		return nil
	}
	var errorResp errorResponse
	if len(rawXML) > 0 {
		err = xml.Unmarshal(rawXML, &errorResp)
		if err != nil {
			return err
		}
		errorResp.Error.StatusCode = response.StatusCode
		return &errorResp.Error
	}
	errorResp.Error.StatusCode = response.StatusCode
	return &errorResp.Error
}

// makeRequest is a shortcut to createRequest, http client.Do, and parseResponse.
func (c *Client) makeRequest(method string, path string, data interface{}, responseBody interface{}) error {
	request, err := c.createRequest(method, path, data)
	if err != nil {
		return err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	return c.parseResponse(response, responseBody)
}

// nopCloser is an io.ReaderCloser that does nothing for the close.
type nopCloser struct {
	io.Reader
}

// Close does nothing.
func (nopCloser) Close() error {
	return nil
}

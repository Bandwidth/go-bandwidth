package bandwidth

import (
	"fmt"
	"net/http"
	"net/url"
	"errors"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"bytes"
)

// Client is main API object
type Client struct{
	UserID, APIToken, APISecret string
	APIVersion string
	APIEndPoint string
}

// New creates new instances of api
func New(userID, apiToken, apiSecret string, other... string) (*Client, error){
	apiVersion := "v1"
	apiEndPoint := "https://api.catapult.inetwork.com"
	if userID == "" || apiToken == "" || apiSecret == "" {
		return nil, errors.New("Missing auth data. Please use api := bandwidth.New(\"user-id\", \"api-token\", \"api-secret\")")
	}
	l := len(other)
	if l > 1 {
		apiEndPoint = other[1]
	}
	if l > 0 {
		apiVersion = other[0]
	}
	client := &Client {userID, apiToken, apiSecret, apiVersion, apiEndPoint}
	return client, nil
}


func (c *Client) concatUserPath(path string) string{
	if path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("/users/%s%s", c.UserID, path)
}

func (c *Client) prepareURL(path string) string{
	if path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("/%s/%s%s", c.APIEndPoint, c.APIVersion, path)
}

func (c* Client) createRequest(method, path string) (*http.Request, error){
	request, err := http.NewRequest(method, c.prepareURL(path), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.APIToken, c.APISecret)))))

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", fmt.Sprintf("go-bandwidth-v%s", Version))
	return request, nil
}

func (c* Client) checkResponse(response *http.Response)(interface{}, error){
	defer response.Body.Close()
	var body interface{}
	rawJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if len(rawJSON) > 0 {
		err = json.Unmarshal([]byte(rawJSON), &body)
		if err != nil {
			return nil, err
		}
	}
	if response.StatusCode >= 200 && response.StatusCode < 400 {
		return body, nil
	}
	errorBody := body.(map[string]interface{})
	message := errorBody["message"].(string)
	if message == "" {
		message = errorBody["code"].(string)
	}
	if message == "" {
		return nil, fmt.Errorf("Http code %d", response.StatusCode)
	}
	return nil, errors.New(message)
}


func (c *Client) makeRequest(method, path string, data... interface{}) (interface{}, error){
    request, err := c.createRequest(method, path)
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		if method == "GET" {
			item := data[0].(map[string] interface{})
			query := make(url.Values)
			for key, value := range item {
				query[key] = []string{value.(string)}
			}
			request.URL.RawQuery = query.Encode()
		} else {
			request.Header.Set("Content-Type", "application/json")
			rawJSON, err := json.Marshal(data[0])
			if err != nil {
				return nil, err
			}
			request.Body = nopCloser{bytes.NewReader(rawJSON)}
		}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return c.checkResponse(response)
}

type nopCloser struct {
    io.Reader
}

func (nopCloser) Close() error { return nil }
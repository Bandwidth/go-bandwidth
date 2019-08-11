package numbers

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	// ApplicationServiceTypeMessagingV2 is the messaging V2 service type.
	ApplicationServiceTypeMessagingV2 = "Messaging-V2"
	// ApplicationServiceTypeVoiceV2 is the voice V2 service type.
	ApplicationServiceTypeVoiceV2 = "Voice-V2"
)

// ApplicationCallbackCreds is callback credentials for an application.
type ApplicationCallbackCreds struct {
	XMLName  xml.Name `xml:"CallbackCreds"`
	UserID   string   `xml:"UserId"`
	Password string   `xml:"Password"`
}

// Application is an application in an Account.
type Application struct {
	XMLName                  xml.Name                  `xml:"Application"`
	ApplicationID            string                    `xml:"ApplicationId"`
	ServiceType              string                    `xml:"ServiceType"`
	AppName                  string                    `xml:"AppName"`
	MsgCallbackURL           string                    `xml:"MsgCallbackUrl"`
	CallInitiatedCallbackURL string                    `xml:"CallInitiatedCallbackUrl"`
	CallInitiatedMethod      string                    `xml:"CallInitiatedMethod"`
	CallStatusCallbackURL    string                    `xml:"CallStatusCallbackUrl"`
	CallStatusMethod         string                    `xml:"CallStatusMethod"`
	ApplicationCallbackCreds *ApplicationCallbackCreds `xml:"CallbackCreds"`
}

// ApplicationsResponse is the response to GetApplications.
type ApplicationsResponse struct {
	XMLName      xml.Name      `xml:"ApplicationProvisioningResponse"`
	Applications []Application `xml:"ApplicationList>Application"`
}

// GetApplications returns applications for the account.
func (c *Client) GetApplications() (*ApplicationsResponse, error) {
	respBody := &ApplicationsResponse{}
	err := c.makeRequest(http.MethodGet, "applications", nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// ApplicationResponse is the response to GetApplication, CreateApplication, and UpdateApplication.
type ApplicationResponse struct {
	XMLName     xml.Name    `xml:"ApplicationProvisioningResponse"`
	Application Application `xml:"Application"`
}

// GetApplication gets an application.
func (c *Client) GetApplication(appID string) (*ApplicationResponse, error) {
	respBody := &ApplicationResponse{}
	err := c.makeRequest(http.MethodGet, fmt.Sprintf("applications/%s", appID), nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// CreateApplication creates a new application.
func (c *Client) CreateApplication(data *Application) (*ApplicationResponse, error) {
	respBody := &ApplicationResponse{}
	err := c.makeRequest(http.MethodPost, "applications", data, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// UpdateApplication updates an application.
func (c *Client) UpdateApplication(data *Application) (*ApplicationResponse, error) {
	respBody := &ApplicationResponse{}
	err := c.makeRequest(http.MethodPut, fmt.Sprintf("applications/%s", data.ApplicationID), data, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// DeleteApplication deletes an application.
func (c *Client) DeleteApplication(appID string) error {
	return c.makeRequest(http.MethodDelete, fmt.Sprintf("applications/%s", appID), nil, nil)
}

// AssociatedSipPeer is a SIP peer associated to an application.
type AssociatedSipPeer struct {
	XMLName  xml.Name `xml:"AssociatedSipPeer"`
	SiteID   string   `xml:"SiteId"`
	SiteName string   `xml:"SiteName"`
	PeerID   string   `xml:"PeerId"`
	PeerName string   `xml:"PeerName"`
}

// AssociatedSipPeersResponse is the response to GetApplicationAssociatedSipPeers.
type AssociatedSipPeersResponse struct {
	XMLName            xml.Name             `xml:"AssociatedSipPeersResponse"`
	AssociatedSipPeers []*AssociatedSipPeer `xml:"AssociatedSipPeers"`
}

// GetApplicationAssociatedSipPeers gets the associated SIP peers for an application.
func (c *Client) GetApplicationAssociatedSipPeers(appID string) (*AssociatedSipPeersResponse, error) {
	respBody := &AssociatedSipPeersResponse{}
	err := c.makeRequest(http.MethodPut, fmt.Sprintf("applications/%s/associatedsippeers", appID), nil, respBody)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

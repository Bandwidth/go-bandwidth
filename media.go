package bandwidth

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const mediaPath = "media"

// MediaFile struct
type MediaFile struct {
	ContentLength int64  `json:"contentLength"`
	Content       string `json:"content"`
	MediaName     string `json:"mediaName"`
}

// GetMediaFiles returns  a list of your media files
func (api *Client) GetMediaFiles() ([]*MediaFile, error) {
	result, _, err := api.makeRequest(http.MethodGet, api.concatUserPath(mediaPath), &[]*MediaFile{})
	if err != nil {
		return nil, err
	}
	return *(result.(*[]*MediaFile)), nil
}

// DeleteMediaFile removes a media file
func (api *Client) DeleteMediaFile(name string) error {
	_, _, err := api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(mediaPath), url.QueryEscape(name)))
	return err
}

// UploadMediaFile creates a new media
func (api *Client) UploadMediaFile(name string, file interface{}, contentType ...string) error {
	request, err := api.createRequest(http.MethodPut, fmt.Sprintf("%s/%s", api.concatUserPath(mediaPath), url.QueryEscape(name)))
	if err != nil {
		return err
	}
	if len(contentType) > 0 {
		request.Header.Set("Content-Type", contentType[0])
	} else {
		request.Header.Set("Content-Type", "application/octet-stream")
	}
	switch file.(type) {
	case string:
		request.Body, err = os.Open(file.(string))
		if err != nil {
			return err
		}
	default:
		request.Body = file.(io.ReadCloser)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	_, _, err = api.checkResponse(response, nil)
	return err
}

// DownloadMediaFile creates a new media
func (api *Client) DownloadMediaFile(name string) (io.ReadCloser, string, error) {
	request, err := api.createRequest(http.MethodGet, fmt.Sprintf("%s/%s", api.concatUserPath(mediaPath), url.QueryEscape(name)))
	if err != nil {
		return nil, "", err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, "", err
	}
	if response.StatusCode >= 400 {
		text, _ := ioutil.ReadAll(response.Body)
		return nil, "", fmt.Errorf("Http code %d: %s", response.StatusCode, text)
	}
	return response.Body, response.Header.Get("Content-Type"), nil
}

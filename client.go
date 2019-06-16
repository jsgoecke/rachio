package rachio

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

// Provides the client and associated elements for interacting with the
// Rachio API
type Client struct {
	Token string `json:"token"`
	HTTP  *http.Client
}

var (
	BaseURL      = "https://api.rach.io/1/public"
	ActiveClient *Client
)

// Generates a new client for the Rachio API
func NewClient(token string) (*Client, error) {
	ActiveClient := &Client{
		Token: token,
		HTTP:  &http.Client{},
	}
	return ActiveClient, nil
}

// // Calls an HTTP DELETE
func (c Client) delete(url string) error {
	req, _ := http.NewRequest("DELETE", url, nil)
	_, err := c.processRequest(req)
	return err
}

// Calls an HTTP GET
func (c Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	return c.processRequest(req)
}

// Calls an HTTP POST with a JSON body
func (c Client) post(url string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// Calls an HTTP PUT
func (c Client) put(resource string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("PUT", BaseURL+resource, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// Processes a HTTP POST/PUT request
func (c Client) processRequest(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Sets the required headers for calls to the Rachio API
func (c Client) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

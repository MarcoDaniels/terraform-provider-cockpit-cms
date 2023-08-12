package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	Endpoint   string
	Token      string
}

func cockpitClient(endpoint, token *string) (*Client, error) {
	if (*endpoint == "") || (*token == "") {
		return nil, fmt.Errorf("empty endpoint or token")
	}

	client := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      *token,
		Endpoint:   *endpoint,
	}

	return &client, nil
}

func (c *Client) makeRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Cockpit-Token", c.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}

func (c *Client) allCollections() (*map[string]Collection, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/collections/listCollections/extended", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	result := map[string]Collection{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, err
}

package cockpit_cms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string
}

func cockpitClient(baseUrl, token string) (*Client, error) {
	client := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token:      token,
		BaseURL:    baseUrl,
	}

	return &client, nil
}

func (c *Client) allCollections() ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/collections/listCollections/extended", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cockpit-Token", c.Token)
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

	return body, err
}

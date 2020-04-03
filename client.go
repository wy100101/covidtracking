package covidtracking

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{HttpClient: httpClient, UserAgent: "RestClient"}
	return c
}

func (c *Client) GetStates() ([]StateData, error) {
	var states []StateData
	rel := &url.URL{Path: "/api/states"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&states)
	return states, err
}

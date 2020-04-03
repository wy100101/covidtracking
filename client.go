package covidtracking

import (
	"encoding/json"
	"fmt"
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

func (c *Client) Request(method, endpoint, query string) (*http.Response, error) {
	rel := &url.URL{Path: endpoint, RawQuery: query}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return c.HttpClient.Do(req)
}

func (c *Client) GetStates(states ...string) ([]StateData, error) {
	var states_data []StateData
	var resp *http.Response
	var err error

	api_path := "/api/states"

	if states == nil || len(states) == 0 {
		resp, err = c.Request("GET", api_path, "")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&states_data)
		if err != nil {
			return nil, err
		}
	} else {
		for _, state := range states {
			state_data := StateData{}
			query := fmt.Sprintf("state=%s", state)
			resp, err = c.Request("GET", api_path, query)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			//body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Print(string(body))
			err = json.NewDecoder(resp.Body).Decode(&state_data)
			if err != nil {
				return nil, err
			}
			states_data = append(states_data, state_data)
		}
	}
	return states_data, err
}

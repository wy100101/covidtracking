package covidtracking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type StateData struct {
	CheckTimeEt          string `json:"checkTimeEt"`
	CommercialScore      int64  `json:"commercialScore"`
	DateChecked          string `json:"dateChecked"`
	DateModified         string `json:"dateModified"`
	Death                int64  `json:"death"`
	Fips                 string `json:"fips"`
	Grade                string `json:"grade"`
	Hash                 string `json:"hash"`
	Hospitalized         int64  `json:"hospitalized"`
	LastUpdateEt         string `json:"lastUpdateEt"`
	Negative             int64  `json:"negative"`
	NegativeRegularScore int64  `json:"negativeRegularScore"`
	NegativeScore        int64  `json:"negativeScore"`
	Notes                string `json:"notes"`
	Pending              int64  `json:"pending"`
	Positive             int64  `json:"positive"`
	PositiveScore        int64  `json:"positiveScore"`
	Score                int64  `json:"score"`
	State                string `json:"state"`
	Total                int64  `json:"total"`
	TotalTestResults     int64  `json:"totalTestResults"`
}

func (s StateData) String() string {
	return fmt.Sprintf("State: %s Negative: %d Positive: %d Total: %d", s.State, s.Negative, s.Positive, s.Total)
}

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{httpClient: httpClient, UserAgent: "RestClient"}
	return c
}

func (c *Client) GetStates() ([]StateData, error) {
	var states []StateData
	rel := &url.URL{Path: "/states"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&states)
	return states, err
}

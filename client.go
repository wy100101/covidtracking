package covidtracking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client basic HTTP client struct for fixed base URL used for REST endpoints
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HTTPClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{HTTPClient: httpClient, UserAgent: "RestClient"}
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

	return c.HTTPClient.Do(req)
}

func (c *Client) GetUS() (USData, error) {
	var usData USData
	var resp *http.Response
	var err error

	apiPath := "/api/us"
	resp, err = c.Request("GET", apiPath, "")
	if err != nil {
		return USData{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&usData)
	if err != nil {
		return USData{}, err
	}
	return usData, nil
}

// GetStatesDaily test comment
func (c *Client) GetStatesDaily(states []string, dates []int) (statesDailyData []StateDailyData, err error) {
	var resp *http.Response

	apiPath := "/api/states/daily"
	validDates := []int{}
	minDate := 20200311
	maxDate, err := strconv.Atoi(time.Now().Format("20060201"))
	if err != nil {
		return nil, err
	}

	for _, d := range dates {
		if minDate <= d && d <= maxDate {
			validDates = append(validDates, d)
		}
	}
	if len(validDates) == 0 {
		return nil, fmt.Errorf("No valid dates specified in dates: %v", dates)
	}

	for _, d := range validDates {
		statesDailyDataForDay := []StateDailyData{}
		if states == nil || len(states) == 0 {
			query := fmt.Sprintf("date=%d", d)
			resp, err = c.Request("GET", apiPath, query)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			err = json.NewDecoder(resp.Body).Decode(&statesDailyDataForDay)
			if err != nil {
				return nil, err
			}
			statesDailyData = append(statesDailyData, statesDailyDataForDay...)
		} else {
			for _, state := range states {
				stateDailyDataForDay := []StateDailyData{}
				query := fmt.Sprintf("state=%s&date=%d", state, d)
				resp, err = c.Request("GET", apiPath, query)
				if err != nil {
					return nil, err
				}
				defer resp.Body.Close()
				//body, _ := ioutil.ReadAll(resp.Body)
				//fmt.Print(string(body))
				err = json.NewDecoder(resp.Body).Decode(&stateDailyDataForDay)
				if err != nil {
					return nil, err
				}
				statesDailyData = append(statesDailyData, stateDailyDataForDay...)
			}
		}
	}
	return
}

func (c *Client) GetStates(states []string) (statesData []StateData, err error) {
	var resp *http.Response

	apiPath := "/api/states"

	if states == nil || len(states) == 0 {
		resp, err = c.Request("GET", apiPath, "")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&statesData)
		if err != nil {
			return nil, err
		}
	} else {
		for _, state := range states {
			stateData := StateData{}
			query := fmt.Sprintf("state=%s", state)
			resp, err = c.Request("GET", apiPath, query)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			//body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Print(string(body))
			err = json.NewDecoder(resp.Body).Decode(&stateData)
			if err != nil {
				return nil, err
			}
			statesData = append(statesData, stateData)
		}
	}
	return
}

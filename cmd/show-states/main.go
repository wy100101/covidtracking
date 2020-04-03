package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"

	"github.com/wy100101/covidtracking"
)

func main() {
	baseURL, err := url.Parse("https://covidtracking.com/")
	if err != nil {
		panic(err)
	}
	c := covidtracking.Client{
		BaseURL:    baseURL,
		UserAgent:  "RestClient",
		HttpClient: http.DefaultClient,
	}
	states, err := c.GetStates()
	if err != nil {
		panic(err)
	}
	sort.Sort(covidtracking.ByDeath(states))
	for _, s := range states {
		fmt.Printf("%s Death: %-8d Hospitalized: %-8d\n", s, s.Death, s.Hospitalized)
	}

}

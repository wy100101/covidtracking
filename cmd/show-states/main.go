package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/wy100101/covidtracking"
)

func main() {
	c := covidtracking.Client{
		BaseURL:    &url.URL{},
		UserAgent:  "RestClient",
		HttpClient: http.DefaultClient,
	}
	states, err := c.GetStates()
	for s := range states {
		fmt.Println(s)
	}

}

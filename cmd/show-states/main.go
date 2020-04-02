package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	c := convidtracking.Client{
		BaseURL:    &url.URL{},
		UserAgent:  "RestClient",
		httpClient: http.DefaultClient,
	}
	states, err := c.GetStates()
	for s := range states {
		fmt.Println(s)
	}

}

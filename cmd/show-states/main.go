package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/wy100101/covidtracking"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	statesFlag = kingpin.Flag("states", "states to show (comma separated").Short('s').String()
	sortFlag   = kingpin.Flag("sort", "Sort key").Short('k').String()
)

func main() {
	kingpin.Parse()
	statesList := []string{}
	if *statesFlag != "" {
		statesList = strings.Split(*statesFlag, ",")
	}

	baseURL, err := url.Parse("https://covidtracking.com/")
	if err != nil {
		panic(err)
	}
	c := covidtracking.Client{
		BaseURL:    baseURL,
		UserAgent:  "RestClient",
		HTTPClient: http.DefaultClient,
	}
	states, err := c.GetStates(statesList)
	if err != nil {
		panic(err)
	}
	switch *sortFlag {
	case "D":
		sort.Sort(covidtracking.StateDataByDeath(states))
	case "T":
		sort.Sort(covidtracking.StateDataByTotal(states))
	case "P":
		sort.Sort(covidtracking.StateDataByPositive(states))
	default:
		sort.Sort(covidtracking.StateDataByState(states))
	}

	for _, s := range states {
		fmt.Printf("%s Death: %-8d Hospitalized: %-8d\n", s, s.Death, s.Hospitalized)
	}

}

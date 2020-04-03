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
	states_flag = kingpin.Flag("states", "states to show (comma separated").Short('s').String()
	sort_flag   = kingpin.Flag("sort", "Sort key").Short('k').String()
)

func main() {
	kingpin.Parse()
	states_list := []string{}
	if *states_flag != "" {
		states_list = strings.Split(*states_flag, ",")
	}

	baseURL, err := url.Parse("https://covidtracking.com/")
	if err != nil {
		panic(err)
	}
	c := covidtracking.Client{
		BaseURL:    baseURL,
		UserAgent:  "RestClient",
		HttpClient: http.DefaultClient,
	}
	states, err := c.GetStates(states_list...)
	if err != nil {
		panic(err)
	}
	switch *sort_flag {
	case "D":
		sort.Sort(covidtracking.ByDeath(states))
	case "T":
		sort.Sort(covidtracking.ByTotal(states))
	case "P":
		sort.Sort(covidtracking.ByPositive(states))
	default:
		sort.Sort(covidtracking.ByState(states))
	}

	for _, s := range states {
		fmt.Printf("%s Death: %-8d Hospitalized: %-8d\n", s, s.Death, s.Hospitalized)
	}

}

package covidtracking

import (
	"fmt"
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
	TotalTestResults     int64  `json:"totalTestResults"`
}

func (s StateData) Total() int64 {
	return s.Positive + s.Negative + s.Pending
}

func (s StateData) String() string {
	return fmt.Sprintf("State: %s Negative: %-8d Positive: %-8d Pending: %-8d Total: %-8d", s.State, s.Negative, s.Positive, s.Pending, s.Total())
}

type ByTotal []StateData

func (s ByTotal) Len() int           { return len(s) }
func (s ByTotal) Less(i, j int) bool { return s[i].Total() < s[j].Total() }
func (s ByTotal) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type ByDeath []StateData

func (s ByDeath) Len() int           { return len(s) }
func (s ByDeath) Less(i, j int) bool { return s[i].Death < s[j].Death }
func (s ByDeath) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type ByPositive []StateData

func (s ByPositive) Len() int           { return len(s) }
func (s ByPositive) Less(i, j int) bool { return s[i].Positive < s[j].Positive }
func (s ByPositive) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type ByState []StateData

func (s ByState) Len() int           { return len(s) }
func (s ByState) Less(i, j int) bool { return s[i].State < s[j].State }
func (s ByState) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

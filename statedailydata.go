package covidtracking

import "time"

type StateDailyData struct {
	Date                     int       `json:"date"`
	State                    string    `json:"state"`
	Positive                 int       `json:"positive"`
	Negative                 int       `json:"negative"`
	Pending                  int       `json:"pending"`
	HospitalizedCurrently    int       `json:"hospitalizedCurrently"`
	HospitalizedCumulative   int       `json:"hospitalizedCumulative"`
	InIcuCurrently           int       `json:"inIcuCurrently"`
	InIcuCumulative          int       `json:"inIcuCumulative"`
	OnVentilatorCurrently    int       `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   int       `json:"onVentilatorCumulative"`
	Recovered                int       `json:"recovered"`
	Hash                     string    `json:"hash"`
	DateChecked              time.Time `json:"dateChecked"`
	Death                    int       `json:"death"`
	Hospitalized             int       `json:"hospitalized"`
	TotalTestResults         int       `json:"totalTestResults"`
	PosNeg                   int       `json:"posNeg"`
	Fips                     string    `json:"fips"`
	DeathIncrease            int       `json:"deathIncrease"`
	HospitalizedIncrease     int       `json:"hospitalizedIncrease"`
	NegativeIncrease         int       `json:"negativeIncrease"`
	PositiveIncrease         int       `json:"positiveIncrease"`
	TotalTestResultsIncrease int       `json:"totalTestResultsIncrease"`
}

func (s *StateDailyData) Total() int {
	return s.Positive + s.Negative + s.Pending
}

type StateDailyDataByDate []StateDailyData

func (s StateDailyDataByDate) Len() int           { return len(s) }
func (s StateDailyDataByDate) Less(i, j int) bool { return s[i].Date < s[j].Date }
func (s StateDailyDataByDate) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

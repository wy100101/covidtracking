package covidtracking

import "time"

type USData struct {
	Positive               int64     `json:"positive"`
	Negative               int64     `json:"negative"`
	Pending                int64     `json:"pending"`
	HospitalizedCurrently  int64     `json:"hospitalizedCurrently"`
	HospitalizedCumulative int64     `json:"hospitalizedCumulative"`
	InIcuCurrently         int64     `json:"inIcuCurrently"`
	InIcuCumulative        int64     `json:"inIcuCumulative"`
	OnVentilatorCurrently  int64     `json:"onVentilatorCurrently"`
	OnVentilatorCumulative int64     `json:"onVentilatorCumulative"`
	Recovered              int64     `json:"recovered"`
	Hash                   string    `json:"hash"`
	LastModified           time.Time `json:"lastModified"`
	Death                  int64     `json:"death"`
	Hospitalized           int64     `json:"hospitalized"`
	TotalTestResults       int64     `json:"totalTestResults"`
	Notes                  string    `json:"notes"`
}

type USDataByTotal []USData

func (u USData) Total() int64 { return u.Positive + u.Negative + u.Pending }

func (u USDataByTotal) Len() int           { return len(u) }
func (u USDataByTotal) Less(i, j int) bool { return u[i].Total() < u[j].Total() }
func (u USDataByTotal) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type USDataByDeath []USData

func (u USDataByDeath) Len() int           { return len(u) }
func (u USDataByDeath) Less(i, j int) bool { return u[i].Death < u[j].Death }
func (u USDataByDeath) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type USDataByRecovered []USData

func (u USDataByRecovered) Len() int           { return len(u) }
func (u USDataByRecovered) Less(i, j int) bool { return u[i].Recovered < u[j].Recovered }
func (u USDataByRecovered) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type USDataByPositive []USData

func (u USDataByPositive) Len() int           { return len(u) }
func (u USDataByPositive) Less(i, j int) bool { return u[i].Positive < u[j].Positive }
func (u USDataByPositive) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type USDataByNegative []USData

func (u USDataByNegative) Len() int           { return len(u) }
func (u USDataByNegative) Less(i, j int) bool { return u[i].Negative < u[j].Negative }
func (u USDataByNegative) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type USDataByPending []USData

func (u USDataByPending) Len() int           { return len(u) }
func (u USDataByPending) Less(i, j int) bool { return u[i].Pending < u[j].Pending }
func (u USDataByPending) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

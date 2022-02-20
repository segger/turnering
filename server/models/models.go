package models

type Participant struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	DogName   string `json:"dogName"`
	ClassNbr  int8   `json:"classNbr,string"`
}

type Contest struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	EventList []Event
}

type Event struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ClassNbr  int8   `json:"classNbr,string"`
	SortOrder int    `json:"sortOrder"`
}

type Protocol struct {
	Participant     Participant
	ContestId       string `json:"contestId"`
	EventResultList []EventResult
}

type EventResult struct {
	EventId  string  `json:"eventId"`
	ResultId string  `json:"resultId"`
	Points   float32 `json:"points"`
	Errors   int32   `json:"errors"`
	Time     int64   `json:"time"`
	Sse      bool    `json:"sse"`
}

type ContestRegistered struct {
	FirstName string `json:"firstName"`
	DogName   string `json:"dogName"`
	ClassNbr  int8   `json:"classNbr,string"`
	EventName string `json:"eventName"`
}

type ContestResult struct {
	FirstName string  `json:"firstName"`
	DogName   string  `json:"dogName"`
	ClassNbr  int8    `json:"classNbr,string"`
	EventName string  `json:"eventName"`
	Points    float32 `json:"points"`
	Errors    int32   `json:"errors"`
	Time      int64   `json:"time"`
	Sse       bool    `json:"sse"`
}

type ContestResultSorted struct {
	Placement int     `json:"placement"`
	FirstName string  `json:"firstName"`
	DogName   string  `json:"dogName"`
	Points    float32 `json:"points"`
	Errors    int32   `json:"errors"`
	Time      int64   `json:"time"`
	Sse       bool    `json:"sse"`
	TP        float32 `json:"tournamentPoints"`
}

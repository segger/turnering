package models

type Participant struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	DogName   string `json:"dogName"`
	ClassNbr  int8   `json:"classNbr,string"`
}

type Contest struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

type Protocol struct {
	Participant     Participant
	ContestId       string `json:"contestId"`
	EventResultList []EventResult
}

type EventResult struct {
	EventName string  `json:"eventName"`
	Points    float32 `json:"points"`
	Errors    int32   `json:"errors"`
	Time      int64   `json:"time"`
	Sse       bool    `json:"sse"`
}
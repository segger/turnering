package models

type Participant struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	DogName   string `json:"dogName"`
	ClassNbr  int8   `json:"classNbr"`
}

type Contest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

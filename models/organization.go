package models

type Organization struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:description`
	Speakers    []Speaker
	Events      []Event
}

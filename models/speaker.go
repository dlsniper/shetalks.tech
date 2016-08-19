//package models

package models

type Speaker struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Events []Event
}

type Speakers []Speaker

package models

import "database/sql"

// Event event
type Event struct {
	Model
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:description`
	Speakers    []Speaker
}

// Events all events
type Events []Event

// Create blah
func (e *Event) Create(db *sql.DB) {
	// SQL STUFF
}

// Read blah
func (e *Event) Read() {

}

// Update blah
func (e *Event) Update() {

}

// Delete blah
func (e *Event) Delete() {

}

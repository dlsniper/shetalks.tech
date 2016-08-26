package models

import "database/sql"
import "fmt"

// Event event
type Event struct {
	Name        string   `json:"name"`
	Description string   `json:description`
	Speakers    []string
}

// Events all events
type Events []Event

/**
 * @TODO Should these functions be methods instead? 
 */

// Create Event
func CreateEvent(db *sql.DB, name string, description string, speakers []string) {
	id,err := db.Exec(
		"INSERT INTO events (name, description) OUTPUT INSERTED.id VALUES ($1, $2)",
	    name,
	    description,
	)

	for _,speaker := range speakers {
		
		//@TODO check if speaker identifier exists

		_,err := db.Exec(
			"INSERT INTO speaker_event (speaker, event) VALUES ($1, $2)",
		    speaker,
		    id,
		)

		if err != nil {
			fmt.Println("Error creating speaker_event entry %v", err)
		}	
	}


	if err != nil {
		fmt.Println("Error in creating event entry: %v", err)
		// @TODO Log error 
	}
}

// Read blah
//func ReadEvent(db *sql.DB, id int, name string, description string) {

//}

// Update blah
func UpdateEvent(db *sql.DB, id int, name string, description string) {

	// @TODO figure out how you want to separate only updating one column
	// 2 transactions? If statement logic? 

	_,err := db.Exec(
		"UPDATE event SET name=$1, description=$2, WHERE id=$3",
		name,
		description,
		id,
	)

	if err != nil {
		fmt.Println("Error in updating event entry: %v", err)
		// @TODO Log error 
	}
}

// Delete blah
func DeleteEvent(db *sql.DB, id int) {

	_,err := db.Exec(
		"DELETE FROM events WHERE id=$1",
		id,
	)

	if err != nil {
		fmt.Println("Error in deleting event entry: %v", err)
		// @TODO Log error 
	}

}

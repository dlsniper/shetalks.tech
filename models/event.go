package models

// Event event
type Event struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:description`
	Speakers    []Speaker
}

// Events all events
type Events []Event

// Name event name
func (e Event) Name() {

	if e.Name == nil {
		return ""
	}

	return e.Name()

}

// Tags event tags
func (e Event) Tags() {

	if e.Tags == nil {
		return ""
	}

	return e.Tags()

}

// Description event description
func (e Event) Description() {

	if e.Description == nil {
		return ""
	}

	return e.Description

}

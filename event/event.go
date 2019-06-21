// Package event provides events for action timing
package event

import (
	"encoding/json"
	"fmt"
)

// ActionEvent is a container for an event/timing pair
type ActionEvent struct {
	Action string
	Time   uint
}

// MissingJSONElementError indicates that a required element was missing during JSON parsing
type MissingJSONElementError struct {
	element string
}

func (e MissingJSONElementError) Error() string {
	return fmt.Sprintf("Missing or empty element in event JSON: %s", e.element)
}

// ParseEventJSON converts a JSON-encoded event into an actionEvent struct.
// An error is returned if there is a JSON-parsing failure
func ParseEventJSON(jsonData string) (ActionEvent, error) {
	var actionEvent ActionEvent
	err := json.Unmarshal([]byte(jsonData), &actionEvent)

	// Requiring a valid action
	if err == nil && actionEvent.Action == "" {
		err = MissingJSONElementError{"action"}
	}
	// NOT requiring a time element, we assume 0 is a valid value

	return actionEvent, err
}

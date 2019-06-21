// Package stats provides services for tracking statistics about events
package stats

import (
	"encoding/json"
	"sync"

	"github.com/akjc/jumpcloud/event"
)

// trackerData provides a generic interface for tracking different statistics.
// This interface makes it very easy to add additional statistics.
type trackerData interface {
	addTime(time uint) // addTime provides a value to track stats against
	name() string      // name returns the JSON field name for the statistic
	value() uint       // value returns the calculated statistic
}

// Tracker provides a framework for tracking statistics. It must be provided
// with a function for creating new trackerData implementations.
type Tracker struct {
	data     map[string]trackerData
	lock     sync.RWMutex
	newEntry func() trackerData
}

// newTracker is a basic constructor for Trackers
func newTracker(newEntryFunc func() trackerData) Tracker {
	return Tracker{data: make(map[string]trackerData), newEntry: newEntryFunc}
}

// AddAction takes a JSON encoded action/time event and maintains statistics
func (tracker *Tracker) AddAction(jsonSource string) error {
	var actionEvent event.ActionEvent
	actionEvent, err := event.ParseEventJSON(jsonSource)
	if err == nil {
		var entry trackerData

		// We need a write lock since we are modifying the map
		tracker.lock.Lock()
		defer tracker.lock.Unlock()

		// Instead of allowing the map to init an entry, we need to check for the key
		// so we can call our newEntry method to create the correct trackerData struct
		entry, k := tracker.data[actionEvent.Action]
		if k == false {
			entry = tracker.newEntry()
		}

		// We still have our lock, update the entry
		entry.addTime(actionEvent.Time)

		// Store the entry
		tracker.data[actionEvent.Action] = entry
	}
	return err
}

// GetStats returns a JSON string containing aggregated statistics about actions
func (tracker *Tracker) GetStats() string {
	var stats []map[string]interface{}

	// We only need a read lock
	tracker.lock.RLock()

	// Create a slice of all the different actions and the resulting value
	for action, entry := range tracker.data {
		stats = append(stats, map[string]interface{}{"action": action, entry.name(): entry.value()})
	}
	tracker.lock.RUnlock()

	// The only potential error would be an empty list, handle that
	if len(stats) == 0 {
		return "[]"
	}

	// No other error is expected and thus not checked
	output, _ := json.Marshal(stats)

	return string(output)
}

package event

import "testing"

// Tests that a normal event JSON parse works as expected
func TestSimpleParse(t *testing.T) {
	json := `{"action":"jump", "time":100}`
	parsed, _ := ParseEventJSON(json)
	expected := ActionEvent{"jump", 100}
	if parsed != expected {
		t.Errorf("JSON parsing of %v resulted in invalid action: %v", json, parsed)
	}
}

// Tests that invalid JSON throws an error during parse
func TestInvalidJSON(t *testing.T) {
	json := `{"action":jump", "time":100}`
	_, err := ParseEventJSON(json)
	if err == nil {
		t.Errorf("JSON parsing of %v should have generated error, but didn't", json)
	}
}

// Tests that a missing or misspelled "action" element throws the correct error
func TestActionMisname(t *testing.T) {
	json := `{"acton":"jump", "time":100}`
	_, err := ParseEventJSON(json)
	if err == nil {
		t.Errorf("JSON parsing of %v should have generated error, but didn't", json)
	} else if _, ok := err.(MissingJSONElementError); !ok {
		t.Errorf("JSON parsing of %v should have generated a MissingJSONElementError error, but didn't", json)
	} else if e, _ := err.(MissingJSONElementError); e.element != "action" {
		t.Errorf("JSON parsing of %v generated a MissingJSONElementError with the wrong element: %s", json, e.element)
	}
}

// Tests that a misspelled or missing "time" element uses 0 as a default
func TestTimeMisname(t *testing.T) {
	json := `{"action":"jump", "tie":100}`
	_, err := ParseEventJSON(json)
	if err != nil {
		// A missing time element is OK, we can use a default of 0
		t.Errorf("JSON parsing of %v generated an error: %s", json, err)
	}
}

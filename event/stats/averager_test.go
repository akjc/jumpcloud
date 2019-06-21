package stats

import "testing"

// Test that an empty averageData instance behaves as expected
func TestEmptyData(t *testing.T) {
	averager := averagerData{}
	if averager.count != 0 {
		t.Errorf("averagerData initilized with non-zero count")
	}
	if averager.total != 0 {
		t.Errorf("averagerData initilized with non-zero total")
	}

	// Check that divide by zero is handled
	if averager.value() != 0 {
		t.Errorf("averagerData expected value 0 but got %d", averager.value())
	}
}

type averageDataTestData struct {
	times         []uint
	expectedValue uint
}

// Test that average calculations work as expected
func TestAverages(t *testing.T) {
	// Build a set of cases
	var tests = []averageDataTestData{
		averageDataTestData{make([]uint, 0), uint(0)},                           // Empty case
		averageDataTestData{[]uint{uint(100)}, uint(100)},                       // Single item
		averageDataTestData{[]uint{uint(100), uint(200)}, uint(150)},            // Two items
		averageDataTestData{[]uint{uint(100), uint(100), uint(100)}, uint(100)}, // Many of the same item
		averageDataTestData{[]uint{uint(100), uint(200), uint(300)}, uint(200)}, // Many items
	}

	for _, test := range tests {
		averagerData := averagerData{}
		for _, time := range test.times {
			averagerData.addTime(time)
		}
		if averagerData.value() != test.expectedValue {
			t.Errorf("averagerData expected value of %v but got %v", test.expectedValue, averagerData.value())
		}
	}
}

// Test that the JSON field is named properly
func TestName(t *testing.T) {
	if name := (&averagerData{}).name(); name != "avg" {
		t.Errorf("averagerData expected name 'avg' but got '%v'", name)
	}
}

type averagerTestData struct {
	jsons          []string
	expectedResult string
}

func TestAverager(t *testing.T) {
	// Build test cases
	var tests = []averagerTestData{
		averagerTestData{make([]string, 0), `[]`},                                                        // Empty
		averagerTestData{[]string{`{"action": "test1", "time": 100}`}, `[{"action":"test1","avg":100}]`}, // Single item

		// Two different items
		averagerTestData{
			[]string{`{"action": "test1", "time": 100}`, `{"action": "test2", "time": 100}`},
			`[{"action":"test1","avg":100},{"action":"test2","avg":100}]`,
		},

		// Two of the same item
		averagerTestData{
			[]string{`{"action": "test1", "time": 100}`, `{"action": "test1", "time": 200}`},
			`[{"action":"test1","avg":150}]`,
		},
	}

	for _, test := range tests {
		averager := Averager()
		for _, json := range test.jsons {
			averager.AddAction(json)
		}
		if result := averager.GetStats(); result != test.expectedResult {
			t.Errorf("averager expected result of %v but got %v", test.expectedResult, result)
		}
	}
}

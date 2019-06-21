package stats

// averageData uses two unsigned ints but could easily be changed to use math/big
// to avoid overflow issues if the event list is very large.
type averagerData struct {
	total uint
	count uint
}

func (data *averagerData) addTime(time uint) {
	data.total += time
	data.count++
}

func (data *averagerData) name() string {
	return "avg"
}

func (data *averagerData) value() uint {
	// Avoid divide by zero
	if data.count == 0 {
		return 0
	}
	return data.total / data.count
}

// Averager creates a Tracker that calculates the mean of the events added
func Averager() Tracker {
	return newTracker(func() trackerData { return &averagerData{} })
}

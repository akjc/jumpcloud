package stats

type counterData struct {
	count uint
}

func (data *counterData) addTime(time uint) {
	data.count++
}

func (data *counterData) name() string {
	return "count"
}

func (data *counterData) value() uint {
	return data.count
}

// Counter creates a Tracker that counts events
func Counter() Tracker {
	return newTracker(func() trackerData { return &counterData{^uint(0)} })
}

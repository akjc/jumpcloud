package stats

type maxerData struct {
	max uint
}

func (data *maxerData) addTime(time uint) {
	if data.max < time {
		data.max = time
	}
}

func (data *maxerData) name() string {
	return "max"
}

func (data *maxerData) value() uint {
	return data.max
}

// Maxer creates a Tracker that provides the maximum time
func Maxer() Tracker {
	return newTracker(func() trackerData { return &maxerData{0} })
}

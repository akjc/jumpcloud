package stats

type minerData struct {
	min uint
}

func (data *minerData) addTime(time uint) {
	if data.min > time {
		data.min = time
	}
}

func (data *minerData) name() string {
	return "min"
}

func (data *minerData) value() uint {
	return data.min
}

// Minner creates a Tracker that provides the minimum time
func Minner() Tracker {
	return newTracker(func() trackerData { return &minerData{^uint(0)} }) // Init to max uint
}

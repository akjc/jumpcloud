package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/akjc/jumpcloud/event/stats"
)

// A simple application to demonstrate parallelism and usage
func main() {
	// A few different actions to pick from
	actions := [...]string{"jump", "run", "crouch", "shoot"}

	// A list of different statistics to calculate
	statTrackers := [...]stats.Tracker{stats.Averager(), stats.Minner(), stats.Maxer(), stats.Counter()}

	// We need a channel for output from all the threads below or it will get really ugly...
	outputChannel := make(chan string)

	go func() {
		for output := range outputChannel {
			fmt.Println(output)
		}
	}()

	// Could accept this as a parameter for the application, but hardcoding is OK for this demo
	numberOfOperations := 1000

	// Use a WaitGroup to know when calculations are all done
	var waitGroup sync.WaitGroup

	// We will do numberOfOperations random actions
	for i := 0; i < numberOfOperations; i++ {
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup) {
			// Let the WaitGroup know we are done at the end
			defer waitGroup.Done()

			// 1 in 10 chance of doing a read
			if rand.Intn(10) == 0 {
				for i := 0; i < len(statTrackers); i++ {
					tracker := &statTrackers[i]
					outputChannel <- tracker.GetStats()
				}
			} else {
				// Pick a random action and time
				action := actions[rand.Intn(len(actions))]
				time := rand.Intn(1000)

				json := fmt.Sprintf(`{"action":"%v", "time":%v}`, action, time)

				for i := 0; i < len(statTrackers); i++ {
					tracker := &statTrackers[i]
					tracker.AddAction(json)
				}
			}
		}(&waitGroup)
	}

	// Wait for all the threads to finish
	waitGroup.Wait()

	// One final read
	for i := 0; i < len(statTrackers); i++ {
		tracker := &statTrackers[i]
		fmt.Println(tracker.GetStats())
	}
}

# Jumpcloud Assignment

My primary goal in design of this solution to the assignment was to provide a clean Go implementation
that provides an easy way to expand the scope into additional statistics beyond a simple average.

Calculating statistics on a set of action/timing data (which I refer to as an "event" in the code)
requires the same basic steps for any basic stats. This implementation focuses on those needs and
abstracts the actual calculations into a simple interface.

I chose to demonstrate the flexibility by implementing additional statistics beyond just an average.

#### Installation

To install:

```
$ go get -u -v github.com/akjc/jumpcloud
```

#### Usage

Basic usage is demonstrated in a small application in the `examples/random` package.

```
$ cd $GOPATH/src/github.com/akjc/jumpcloud
$ go run examples/random/random.go
```

#### Unit Tests

Unit tests were written before much of the code was completed to verify accuracy and
provide a basic test-driven workflow.

I chose not to implement unit tests for the additional statistics beyond average for the sake of
brevity. This is the reason for the lower than 100% coverage.

```
$ go test -cover github.com/akjc/jumpcloud/event
$ go test -cover github.com/akjc/jumpcloud/event/stats
```

#### Ignored Improvement

In order to keep the scope small enough for a simple exercise I left out some things that
would likely be part of a true implementation:
* Basic stats might be desireable for more than simple action/time pairs
  * The ActionEvent struct could be abstracted into something more generic to accomplish this
* Potential overflow is ignored, but noted
* Scale is assumed to be small enough to be handled within a single application
  * I had to fight my background in "big data" big time
* If "web scale" were desired, the implementation for Tracker could include a number of improvements
  * Sharding of keys (actions) into separate maps with separate mutexes to minimize waiting on locks
  * Clustering of multiple instances to support a large number of distinct actions
  * A map/reduce style approach to storing events in order to scale writes and calculations for a particular action type
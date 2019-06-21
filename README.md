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
$ go run examples/random/random.go
```
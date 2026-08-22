//go:build darwin

package signpost_test

import (
	"fmt"

	"github.com/tmc/apple/x/signpost"
)

// Emit a named interval around a unit of work. Run the program under
// "log stream --signpost" or record it with Instruments to see the interval.
func Example() {
	log := signpost.New("com.example.app", signpost.PointsOfInterest)

	id := log.NewID()
	log.IntervalBegin(id, "load")
	// ... do work ...
	log.IntervalEnd(id, "load")

	fmt.Println("signpost interval logged")
	// Output:
	// signpost interval logged
}

// A point-in-time event marks a moment of interest rather than a duration.
func ExampleLogger_Event() {
	log := signpost.New("com.example.app", signpost.PointsOfInterest)
	log.Event(log.NewID(), "cache miss")

	fmt.Println("signpost event logged")
	// Output:
	// signpost event logged
}

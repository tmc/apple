//go:build darwin

package signpostmsg_test

import (
	"fmt"

	"github.com/tmc/apple/x/signpost/signpostmsg"
)

// Emit an interval around a block of work. The name groups intervals in
// Instruments; the message carries the identity of this particular span. Run
// the program under "log stream --signpost" or record it with Instruments to
// see the interval.
func Example() {
	log := signpostmsg.New("com.example.app", signpostmsg.PointsOfInterest)

	id := log.NewID()
	log.IntervalBegin(id, signpostmsg.Layer, "TransformerBlock_0")
	// ... do work ...
	log.IntervalEnd(id, signpostmsg.Layer, "TransformerBlock_0")

	fmt.Println("signpostmsg interval logged")
	// Output:
	// signpostmsg interval logged
}

// A point-in-time event marks a moment of interest rather than a duration.
func ExampleLogger_Event() {
	log := signpostmsg.New("com.example.app", signpostmsg.PointsOfInterest)
	log.Event(log.NewID(), signpostmsg.Op, "MatMul_512x4096")

	fmt.Println("signpostmsg event logged")
	// Output:
	// signpostmsg event logged
}

// Names are a fixed set, because the underlying os_signpost macros require a
// compile-time literal.
func ExampleName() {
	fmt.Println(signpostmsg.Model, signpostmsg.Layer, signpostmsg.Op)
	// Output: Model Layer Op
}

//go:build darwin && !native_signpost

package main

import (
	"flag"

	"github.com/tmc/apple/x/signpost"
)

var traceSignposts = flag.Bool("signpost", false, "emit send and receive signposts for xctrace")

var traceLog *signpost.Logger

func enableSignposts() {
	if *traceSignposts {
		traceLog = signpost.New("github.com.tmc.apple.netperfbench", signpost.PointsOfInterest)
	}
}

// signpostInterval ties the events that describe one asynchronous operation
// to its interval. The zero value is an inactive interval.
type signpostInterval struct {
	id   signpost.ID
	name string
}

// beginSignpost returns an interval that can be ended and annotated. It
// returns nil outside trace runs so ordinary benchmark measurements pay
// nothing.
func beginSignpost(name string) *signpostInterval {
	if traceLog == nil || !traceLog.Enabled() {
		return nil
	}
	id := traceLog.NewID()
	traceLog.IntervalBegin(id, name)
	return &signpostInterval{id: id, name: name}
}

func (s *signpostInterval) end() {
	if s != nil {
		traceLog.IntervalEnd(s.id, s.name)
	}
}

func (s *signpostInterval) event(name string) {
	if s != nil {
		traceLog.Event(s.id, name)
	}
}

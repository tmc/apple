// Command logdemo exercises every capability of the x/oslog and x/signpost
// packages so the output can be inspected in Console.app or `log stream`.
//
// All records use the subsystem "com.tmc.appledemo". To watch them:
//
//	log stream --predicate 'subsystem == "com.tmc.appledemo"' --level debug --signpost
//
// or in Console.app, search: subsystem:com.tmc.appledemo
//
// Signpost intervals appear in Instruments' Points of Interest track when you
// record this process with the os_signpost or Points of Interest instrument.
//
// Run once, or with -loop to emit continuously (^C to stop):
//
//	go run ./examples/os/logdemo
//	go run ./examples/os/logdemo -loop
package main

import (
	"errors"
	"flag"
	"log/slog"
	"time"

	"github.com/tmc/apple/x/oslog"
	"github.com/tmc/apple/x/signpost"
)

const subsystem = "com.tmc.appledemo"

func main() {
	loop := flag.Bool("loop", false, "emit continuously until interrupted")
	flag.Parse()

	for {
		runCycle()
		if !*loop {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// runCycle wraps one pass of the demo in an outer activity, so every section's
// own activity nests under it. In Console.app the Activity column (and the
// parentActivityIdentifier field) then shows the parent/child tree.
func runCycle() {
	cycle := oslog.NewActivity("demo-cycle")
	defer cycle.Enter()()

	log := oslog.New(subsystem, "cycle")
	log.Info("cycle start — outer activity id %#x", cycle.ID())

	demoOSLogLevels()
	demoOSLogFormatting()
	demoActivities()
	demoSlog()
	demoSignposts()

	log.Info("cycle end")
}

// demoOSLogLevels emits one message at each os_log level. In Console.app the
// level shows in the "Type" column (Debug, Info, Default, Error, Fault).
func demoOSLogLevels() {
	log := oslog.New(subsystem, "levels")
	log.Debug("debug message — most verbose, usually filtered")
	log.Info("info message — diagnostic detail")
	log.Default("default message — standard level")
	log.Error("error message — something went wrong")
	log.Fault("fault message — a bug in execution")
}

// demoOSLogFormatting shows the format-buffer builder: scalars, widths,
// strings, and public/private visibility. Private values are redacted in
// Console.app unless the device is configured to show private data.
func demoOSLogFormatting() {
	log := oslog.New(subsystem, "formatting")
	log.Default("int %d, hex %x, wide %ld", 42, 255, int64(1)<<40)
	log.Default("public string %{public}s", "visible-in-console")
	log.Default("private string %{private}s", "redacted-by-default")
	log.Default("mixed: %d bytes from %{public}s in %dms", 4096, "cdn.example.com", 12)
	log.Default("bool %d, uint %u, pointer %p", true, uint32(7), uintptr(0xdeadbeef))
	log.Error("error value: %{public}s", errors.New("connection reset"))
}

// demoActivities logs inside an activity scope so the entries are correlated by
// a shared activity id. In Console.app, group by "Activity" (or look at the
// ActivityID column) to see the two "inside" messages share an id while the
// "outside" one does not.
func demoActivities() {
	log := oslog.New(subsystem, "activity")

	// This activity is created while the outer "demo-cycle" activity is current,
	// so it nests under it (Console shows demo-cycle as its parent).
	act := oslog.NewActivity("handle-request")
	leave := act.Enter()
	log.Info("inside child activity — id %#x, nested under the cycle", act.ID())
	log.Default("still inside the same child activity")

	// A second level of nesting under handle-request.
	sub := oslog.NewActivity("db-query")
	leaveSub := sub.Enter()
	log.Default("inside grandchild activity — id %#x", sub.ID())
	leaveSub()

	leave()
	log.Default("back in the cycle activity, child left")
}

// demoSlog routes Go's structured logger to os_log, showing levels, attributes,
// With, and groups.
func demoSlog() {
	handler := oslog.NewHandler(
		oslog.New(subsystem, "slog"),
		&oslog.HandlerOptions{Level: slog.LevelDebug},
	)
	log := slog.New(handler)

	log.Debug("slog debug", "detail", "cache miss")
	log.Info("slog info", "port", 8080, "tls", true)
	log.Warn("slog warn maps to os_log Default", "latency_ms", 340)
	log.Error("slog error", "err", "timeout", "retries", 3)

	// With + nested group: request_id stays top-level, http.* is grouped.
	req := log.With("request_id", "abc123").WithGroup("http")
	req.Info("request handled", "method", "GET", "status", 200)
}

// demoSignposts emits a point event and a timed interval. Both show up in
// Instruments' Points of Interest track; intervals render as spans.
func demoSignposts() {
	sp := signpost.New(subsystem, signpost.PointsOfInterest)

	sp.Event(sp.NewID(), "checkpoint")

	// A timed interval with a name that varies per round (dynamic names — a
	// thing only the runtime buffer builder can do, not a compiled C macro).
	id := sp.NewID()
	name := "work-batch"
	sp.IntervalBegin(id, name)
	time.Sleep(15 * time.Millisecond)
	sp.IntervalEnd(id, name)

	// Nested intervals under the same batch.
	outer := sp.NewID()
	sp.IntervalBegin(outer, "outer")
	for range 3 {
		inner := sp.NewID()
		sp.IntervalBegin(inner, "inner-step")
		time.Sleep(5 * time.Millisecond)
		sp.IntervalEnd(inner, "inner-step")
	}
	sp.IntervalEnd(outer, "outer")
}

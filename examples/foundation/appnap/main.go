// Command appnap shows the effect of an NSProcessInfo activity on a timer.
//
// Run two copies, one with -protect and one without, then leave the Mac idle.
// The protected copy asks macOS not to App Nap the ticker while the unprotected
// copy is an intentionally ordinary background process.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

func main() {
	protect := flag.Bool("protect", false, "disable App Nap for the ticker activity")
	period := flag.Duration("period", time.Second, "ticker period")
	flag.Parse()
	if *period <= 0 {
		fmt.Fprintln(os.Stderr, "appnap: period must be positive")
		os.Exit(2)
	}

	var activity objectivec.NSObjectObject
	if *protect {
		info := foundation.GetProcessInfoClass().ProcessInfo()
		obj := info.BeginActivityWithOptionsReason(
			foundation.NSActivityBackground|foundation.NSActivityLatencyCritical,
			"appnap timer example")
		activity = objectivec.NSObjectObjectFromID(obj.ID)
		defer info.EndActivity(activity)
	}

	last := time.Now()
	ticker := time.NewTicker(*period)
	defer ticker.Stop()
	state := "ordinary"
	if *protect {
		state = "protected"
	}
	fmt.Printf("pid=%d mode=%s period=%s\n", os.Getpid(), state, period)
	for now := range ticker.C {
		fmt.Printf("mode=%s elapsed=%s\n", state, now.Sub(last).Round(time.Millisecond))
		last = now
	}
}

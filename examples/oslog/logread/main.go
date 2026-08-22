// Command logread emits log entries at several levels and reads them back from
// the unified logging store.
//
// It writes one message per os_log level under a subsystem, then opens an
// OSLogStore scoped to the current process and enumerates the entries recorded
// in the recent past, printing the ones that match the subsystem.
//
// Reading the current process's own entries needs no entitlement; the
// system-wide store (OSLogStoreSystem) does.
//
// The composed message often reads "<compose failure [UUID]>": os_log stores
// only the format string's address, and the logging system cannot recover it
// from a binary it has no UUID record for. The entry metadata — date, level,
// subsystem, category, process — is still accurate.
//
// Usage:
//
//	go run ./examples/oslog/logread [-subsystem name] [-since duration]
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/oslog"
	xoslog "github.com/tmc/apple/x/oslog"
)

func main() {
	subsystem := flag.String("subsystem", "com.tmc.appledemo.logread", "subsystem to log under and read back")
	since := flag.Duration("since", 30*time.Second, "how far back to read")
	flag.Parse()

	if err := run(*subsystem, *since); err != nil {
		fmt.Fprintln(os.Stderr, "logread:", err)
		os.Exit(1)
	}
}

func run(subsystem string, since time.Duration) error {
	emit(subsystem)

	store, err := oslog.NewOSLogStoreWithScopeError(oslog.OSLogStoreCurrentProcessIdentifier)
	if err != nil {
		return fmt.Errorf("open log store: %w", err)
	}
	pos := store.PositionWithTimeIntervalSinceEnd(foundation.NSTimeInterval(-since.Seconds()))
	entries, err := store.EntriesEnumeratorWithOptionsPositionPredicateError(0, pos, foundation.NSPredicate{})
	if err != nil {
		return fmt.Errorf("enumerate entries: %w", err)
	}

	logClass := oslog.GetOSLogEntryLogClass().Class()
	n := 0
	for {
		obj := entries.NextObject()
		id := obj.GetID()
		if id == 0 {
			break
		}
		if !objc.Send[bool](id, objc.Sel("isKindOfClass:"), objc.ID(logClass)) {
			continue
		}
		entry := oslog.OSLogEntryLogFromID(id)
		if entry.Subsystem() != subsystem {
			continue
		}
		t := time.Unix(0, int64(entry.Date().TimeIntervalSince1970()*float64(time.Second)))
		fmt.Printf("%s [%s] %s: %s\n", t.Format(time.RFC3339), entry.Level(), entry.Category(), entry.ComposedMessage())
		n++
	}
	fmt.Printf("%d entries for subsystem %s in the last %s\n", n, subsystem, since)
	return nil
}

// emit writes one message per level. Debug and info entries are discarded
// unless the subsystem is configured to keep them, so they may not read back.
func emit(subsystem string) {
	log := xoslog.New(subsystem, "levels")
	log.Debug("debug message")
	log.Info("info message")
	log.Default("default message")
	log.Error("error message")
	log.Fault("fault message")
}

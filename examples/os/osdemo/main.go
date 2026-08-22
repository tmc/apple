// Command osdemo exercises the generated os and oslog bindings from Go:
//   - os_unfair_lock: a working low-level mutex,
//   - os_activity + os_log_create: entering an activity scope and creating a log,
//   - OSLogStore: reading log entries back out of the unified log.
//
// Run it directly:
//
//	go run ./examples/os/osdemo
package main

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/os"
	"github.com/tmc/apple/oslog"
)

func main() {
	tryUnfairLock()
	tryActivityAndLog()
	tryLogStoreReader()
}

// tryUnfairLock exercises the low-level os_unfair_lock primitive. The lock is a
// single 4-byte word; we own the storage and pass its address.
func tryUnfairLock() {
	fmt.Println("== os_unfair_lock ==")
	var lock uint32 // os_unfair_lock_s is one uint32, OS_UNFAIR_LOCK_INIT == 0
	p := unsafe.Pointer(&lock)

	if !os.OSUnfairLockTrylock(p) {
		fmt.Println("  trylock failed on a fresh lock (unexpected)")
		return
	}
	fmt.Println("  trylock: acquired")
	// A second trylock must fail while held.
	if os.OSUnfairLockTrylock(p) {
		fmt.Println("  second trylock unexpectedly succeeded")
	} else {
		fmt.Println("  second trylock: correctly blocked while held")
	}
	os.OSUnfairLockUnlock(p)
	fmt.Println("  unlock: released")
	// Now a plain lock/unlock round trip.
	os.OSUnfairLockLock(p)
	os.OSUnfairLockUnlock(p)
	fmt.Println("  lock/unlock round trip: ok")
}

// tryActivityAndLog enters an os_activity scope and emits an os_log message
// inside it, so the log entry is tagged with the activity id.
func tryActivityAndLog() {
	fmt.Println("== os_activity + os_log ==")
	// os_log_create returns an os_log_t, which these bindings spell as
	// kernel.Os_log_t, an alias for objectivec.Object, so a nil log is a zero
	// ID rather than a zero value.
	logHandle := os.OSLogCreate("com.example.traceexample", "demo")
	if logHandle.ID == 0 {
		fmt.Println("  os_log_create returned nil")
		return
	}
	fmt.Println("  os_log_create: ok")

	// OS_ACTIVITY_NONE (a null activity object) is always valid to enter.
	// OS_ACTIVITY_CURRENT and os_activity_create are macros with no bound
	// symbol, so NONE is the reachable path from these bindings.
	var none os.OSActivity // zero object == OS_ACTIVITY_NONE

	var state [2]uint64 // os_activity_scope_state_s { uint64 opaque[2] } == 16 bytes
	sp := os.OSActivityScopeState(unsafe.Pointer(&state[0]))
	os.OSActivityScopeEnter(none, sp)

	var parent os.OSActivityID
	id := os.OSActivityGetIdentifier(none, &parent)
	fmt.Printf("  entered activity scope: id=%#x parent=%#x\n", id, parent)

	os.OSActivityScopeLeave(sp)
	fmt.Println("  left activity scope: ok")
}

// tryLogStoreReader opens the current-process log store and counts entries.
func tryLogStoreReader() {
	fmt.Println("== oslog OSLogStore reader ==")
	store, err := oslog.NewOSLogStoreWithScopeError(oslog.OSLogStoreCurrentProcessIdentifier)
	if err != nil {
		fmt.Println("  open store:", err)
		return
	}
	if store.ID == 0 {
		fmt.Println("  open store: nil")
		return
	}
	fmt.Println("  opened OSLogStore (current process)")

	en, err := store.EntriesEnumeratorAndReturnError()
	if err != nil {
		fmt.Println("  enumerator:", err)
		return
	}

	// OSLogEnumerator is an NSEnumerator; walk it with -nextObject and print the
	// composed message of the first few entries.
	printed := 0
	for printed < 5 {
		next := objc.Send[objc.ID](en.GetID(), objc.Sel("nextObject"))
		if next == 0 {
			break
		}
		entry := oslog.OSLogEntryFromID(next)
		msg := entry.ComposedMessage()
		if msg == "" {
			continue
		}
		fmt.Printf("  entry %d: %q\n", printed+1, truncate(msg, 80))
		printed++
	}
	if printed == 0 {
		fmt.Println("  (no readable entries for this process scope)")
	}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

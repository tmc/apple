package oslog

import (
	"fmt"
	"runtime"
	"strings"
)

// Stack returns the current goroutine's stack trace, formatted the way the Go
// runtime prints it. It is a convenience wrapper around runtime.Stack for
// including a Go call stack in a log message, since the backtrace os_log
// captures natively refers to the C/purego call site, not Go frames.
func Stack() string {
	buf := make([]byte, 8192)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return string(buf[:n])
		}
		buf = make([]byte, 2*len(buf))
	}
}

// LogStack logs the message at the given type with the current goroutine's Go
// stack trace appended. The stack is logged as a public string so it is not
// redacted.
func (l *Logger) LogStack(t Type, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Log(t, "%{public}s\n%{public}s", msg, Stack())
}

// Recover recovers a panicking goroutine and logs the panic value and stack at
// [TypeFault], then returns the recovered value (nil if there was no panic).
// Use it with defer at a goroutine boundary:
//
//	func worker(log *oslog.Logger) {
//		defer log.Recover()
//		// ... work that might panic ...
//	}
//
// Recover swallows the panic. To log and then re-panic (preserving a crash),
// use [Logger.RecoverAndRepanic].
func (l *Logger) Recover() (recovered any) {
	if r := recover(); r != nil {
		l.logPanic(r)
		return r
	}
	return nil
}

// RecoverAndRepanic logs a panic at [TypeFault] with its stack, then re-panics
// so the process still crashes (and any outer recover still sees it). Use it
// when you want the panic recorded in the system log but not suppressed:
//
//	defer log.RecoverAndRepanic()
func (l *Logger) RecoverAndRepanic() {
	if r := recover(); r != nil {
		l.logPanic(r)
		panic(r)
	}
}

func (l *Logger) logPanic(r any) {
	// runtime.Stack captures the stack at the point of recovery, which still
	// includes the panicking frames below this deferred call. Trim the leading
	// frames belonging to this package so the trace starts at panic().
	l.Log(TypeFault, "panic: %{public}s\n%{public}s", fmt.Sprint(r), trimStack(Stack()))
}

// trimStack drops this package's own recovery frames from the top of a stack
// trace so it begins at the runtime panic frame, keeping the goroutine header.
func trimStack(s string) string {
	lines := strings.SplitAfter(s, "\n")
	if len(lines) == 0 {
		return s
	}
	header := lines[0] // "goroutine N [running]:"
	rest := lines[1:]
	// A frame is two lines: "pkg.func(...)" then "\tfile:line +0x..". Drop leading
	// frame pairs whose function is in this package, up to the panic frame.
	for len(rest) >= 2 {
		fn := rest[0]
		if strings.Contains(fn, "github.com/tmc/apple/x/oslog.") {
			rest = rest[2:]
			continue
		}
		break
	}
	// Also drop the runtime panic() frame pair itself for a cleaner start.
	if len(rest) >= 2 && strings.HasPrefix(rest[0], "panic(") {
		rest = rest[2:]
	}
	return header + strings.Join(rest, "")
}

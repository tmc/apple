// Package oslog writes messages to the unified logging system (os_log).
//
// The os_log family (os_log, os_log_error, os_log_with_type, …) is a set of C
// macros in <os/log.h>, not exported functions: clang parses the format string
// at compile time and builds a companion argument buffer via
// __builtin_os_log_format. Applegen cannot generate bindings for them, so this
// package is a hand-written overlay. It builds the same argument buffer at run
// time and calls the underlying exported _os_log_impl symbol through purego, so
// it needs no cgo.
//
// Unlike the compiler, this package interprets the format string at run time.
// It supports the os_log format specifiers %d, %u, %x, %ld, %lu, %lx, %p, %s,
// and %@ (the last logged as %s). Public and private markers (%{public}s,
// %{private}s) are honored; os_log redacts private arguments unless the
// consumer is entitled to see them.
//
// Basic usage:
//
//	log := oslog.New("com.example.app", "network")
//	log.Info("connected to %{public}s in %dms", host, elapsed)
//	log.Error("request failed: %s", err)
//
// Messages appear in Console.app and in a live `log stream`:
//
//	log stream --predicate 'subsystem == "com.example.app"'
//
// Live readers (Console.app, `log stream`) compose the message by reading the
// format string from the running process, so they render fully. Offline readers
// (`log show`, a saved `.logarchive`) resolve format strings by (image UUID,
// section offset) into the Mach-O __oslogstring section. A compiler writes the
// format there; this package's format string is an ordinary Go value with no
// image identity, so offline readers cannot find it and display
// "<compose failure>" with an empty format. This is a fundamental property of
// emitting os_log without cgo, not a bug: read these logs live.
package oslog

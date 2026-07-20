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
// Messages appear in Console.app and the `log` command:
//
//	log stream --predicate 'subsystem == "com.example.app"'
package oslog

// Package signpost emits os_signpost intervals and events for use with
// Instruments and the unified logging system.
//
// The os_signpost interval and event operations are C macros in
// <os/signpost.h>, not exported functions, so applegen cannot generate
// bindings for them: there is no symbol to resolve at runtime. This package
// is a hand-written overlay that reproduces what the macros expand to,
// calling the underlying exported _os_signpost_emit_with_name_impl symbol
// directly through purego.
//
// Signposts carry no formatted message here: the format buffer that the C
// macros build at compile time via __builtin_os_log_format is omitted. Names
// and interval pairing are preserved, which is what interval timing in
// Instruments and "log stream --signpost" rely on.
//
// Basic usage:
//
//	log := signpost.New("com.example.app", signpost.PointsOfInterest)
//	id := log.NewID()
//	log.IntervalBegin(id, "load")
//	// ... work ...
//	log.IntervalEnd(id, "load")
package signpost

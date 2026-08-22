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
// The name-only operations ([Logger.IntervalBegin], [Logger.IntervalEnd],
// [Logger.Event]) pass an empty format buffer. The Message variants attach a
// run-time string as a "%{public}s" argument, and the f variants
// ([Logger.IntervalBeginf], [Logger.IntervalEndf], [Logger.Eventf]) take an
// os_log format string with typed arguments; both build the argument buffer
// the C macros get from __builtin_os_log_format (shared with x/oslog).
// Instruments shows the result as the interval's message.
//
// Signposts always emit and pair. Whether their names and messages decode in
// trace output depends on the strings being present in a loaded image's
// __TEXT,__oslogstring section (the log tools resolve them by offset from
// the on-disk image; heap strings render as "<missing name>"). The
// signpostnames tool (cmd/signpostnames) generates that pool, and the first
// unpooled emit prints a one-time warning. Measured build-mode matrix:
//
//   - CGO_ENABLED=1 (macOS default) + names_darwin.syso: decodes. The syso
//     forces external linking, which lays the section out; no cgo source is
//     required.
//   - CGO_ENABLED=0 + syso: does NOT decode. The internal linker drops the
//     section. Use signpostnames -dylib and [LoadNames] instead, which works
//     in every build mode.
//   - cgo builds of this package also pool the format strings automatically
//     (oslogstrings_cgo.go), so Message output decodes without a syso.
//
// Basic usage:
//
//	log := signpost.New("com.example.app", signpost.PointsOfInterest)
//	id := log.NewID()
//	log.IntervalBegin(id, "load")
//	// ... work ...
//	log.IntervalEnd(id, "load")
package signpost

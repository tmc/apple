// Package telemetry provides diagnostic and performance measurement types
// for the Apple Neural Engine.
//
// Telemetry fields may return zeros on production hardware. Many ANE
// performance counters and diagnostic selectors are gated by
// ANEDevicePropertyIsInternalBuild in the IOKit registry. Each struct
// field has a Known bool that distinguishes "unavailable" from "zero":
// if Known is false, the underlying selector was not available.
//
// Use [EvalWithStats] to collect hardware execution time and performance
// counters after evaluation. Use [ProbeDiagnostics] to inspect model
// queue depth, program state, and async request counts. Use
// [ProbeClientInfo], [ProbeCacheInfo], and [Snapshot] to capture the
// host environment.
//
// Use [StartDRAM] to measure the bytes the engine moves to and from
// DRAM over a region. These are whole-engine IOReport channels in the
// "AMC Stats" / "Perf Counters" group, and they are system-wide: they
// count every process's ANE traffic, not only the caller's.
//
// On macOS 26.6.1 (25G76) an unprivileged process can enumerate these
// channels but cannot sample them, so [StartDRAM] returns an error. See
// the comment on [StartDRAM] for what was measured. Whether root lifts
// this is untested. Bryngelson, arXiv 2606.22283 §33.3 reports the
// whole-engine channels as readable on the unentitled path, under
// different channel names, on M1 and M5; that does not reproduce here.
//
// All telemetry types have an Available method that reports whether any
// data was collected, and a ReportMetrics method that emits available
// data to a [testing.B]-compatible reporter.
package telemetry

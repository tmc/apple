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
// All telemetry types have an Available method that reports whether any
// data was collected, and a ReportMetrics method that emits available
// data to a [testing.B]-compatible reporter.
package telemetry

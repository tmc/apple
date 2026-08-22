//go:build darwin

package telemetry_test

import (
	"fmt"

	"github.com/tmc/apple/x/ane/telemetry"
)

type metricPrinter struct{}

func (m metricPrinter) ReportMetric(value float64, name string) {
	fmt.Printf("%s: %.0f\n", name, value)
}

// Report ANE hardware execution metrics in a benchmark.
func ExampleEvalStats_ReportMetrics() {
	stats := telemetry.EvalStats{
		HWExecutionNS:    125000,
		PerfStatsEntries: 1,
		PerfCounters: []telemetry.PerfCounter{
			{Index: 0, Name: "kANE_FP16_CYCLES", Value: 4200},
			{Index: 1, Name: "kANE_NE_COMPUTE_CYCLES", Value: 3800},
		},
	}

	stats.ReportMetrics(metricPrinter{})

	// Output:
	// hw-ns/op: 125000
	// kane-fp16-cycles/op: 4200
	// kane-ne-compute-cycles/op: 3800
	// perf-stats-entries/op: 1
}

// Collect a client snapshot and per-eval telemetry after evaluation.
func ExampleModelTelemetry() {
	tel := telemetry.EvalTelemetry{
		Stats: telemetry.EvalStats{
			HWExecutionNS: 50000,
		},
		Diagnostics: telemetry.Diagnostics{
			ProgramClass:      "ANEInMemoryModel",
			ProgramClassKnown: true,
		},
	}

	fmt.Printf("available: %v\n", tel.Available())
	fmt.Println(tel.String())

	// Output:
	// available: true
	// EvalTelemetry{hw=50000ns counters=0 Diagnostics{class=ANEInMemoryModel}}
}

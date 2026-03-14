//go:build darwin

package telemetry_test

import (
	"fmt"
	"log"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/telemetry"
)

// Report ANE hardware execution metrics in a benchmark.
func ExampleEvalStats_ReportMetrics() {
	c, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	const channels = 4
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	m, err := c.Compile(ane.CompileOptions{
		ModelType:     ane.ModelTypeMIL,
		MILText:       []byte(milText),
		WeightBlob:    blob,
		PerfStatsMask: 0xF,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	input := []float32{1, 2, 3, 4}
	if err := m.WriteInputF32(0, input); err != nil {
		log.Fatal(err)
	}

	stats, err := telemetry.EvalWithStats(m)
	if err != nil {
		log.Fatal(err)
	}

	// In a real benchmark, pass testing.B:
	//   stats.ReportMetrics(b)
	fmt.Printf("HWExecutionNS=%d PerfCounterBytes=%d\n", stats.HWExecutionNS, len(stats.PerfCounterData))
}

// Collect a client snapshot and per-eval telemetry after evaluation.
func ExampleModelTelemetry() {
	c, err := ane.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Snapshot captures the host environment once.
	snap := telemetry.Snapshot(c)
	fmt.Printf("snapshot available: %v\n", snap.Available())
	fmt.Printf("device has ANE: %v\n", snap.Device.HasANE)
	fmt.Printf("client available: %v\n", snap.Client.Available())

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		log.Fatal(err)
	}

	m, err := c.Compile(ane.CompileOptions{
		ModelType:     ane.ModelTypeMIL,
		MILText:       []byte(milText),
		WeightBlob:    blob,
		PerfStatsMask: 0xF,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if err := m.WriteInputF32(0, []float32{42}); err != nil {
		log.Fatal(err)
	}

	// Evaluate and collect stats, then build telemetry post-hoc.
	stats, err := telemetry.EvalWithStats(m)
	if err != nil {
		log.Fatal(err)
	}
	tel := telemetry.ModelTelemetry(m, stats)

	fmt.Printf("telemetry available: %v\n", tel.Available())
	fmt.Printf("diagnostics available: %v\n", tel.Diagnostics.Available())
	fmt.Printf("program class: %s\n", tel.Diagnostics.ProgramClass)
	fmt.Printf("timing available: %v\n", tel.Timing.Available())
	// Output:
	// snapshot available: true
	// device has ANE: true
	// client available: true
	// telemetry available: true
	// diagnostics available: true
	// program class: ANEInMemoryModel
	// timing available: false
}

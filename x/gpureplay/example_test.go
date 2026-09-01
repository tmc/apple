package gpureplay_test

import (
	"fmt"
	"time"

	"github.com/tmc/apple/x/gpureplay"
)

func ExampleReplayReport() {
	report := &gpureplay.ReplayReport{
		TotalGPUActiveTime: 12500 * time.Microsecond,
		PeakRegisterUsage:  32,
		KernelTimings: []gpureplay.KernelTiming{
			{
				Name:          "FastGEMV",
				Duration:      4500 * time.Microsecond,
				DispatchCount: 16,
			},
			{
				Name:          "ComputeRMSNorm",
				Duration:      800 * time.Microsecond,
				DispatchCount: 32,
			},
		},
		ShaderMetrics: []gpureplay.ShaderMetric{
			{
				Name:             "FastGEMV",
				RegisterCount:    32,
				ThreadgroupBytes: 1024,
				ALUInstructions:  256,
				MemInstructions:  64,
				ExecutionCostPct: 65.5,
			},
		},
	}

	fmt.Printf("Active Time: %v\n", report.TotalGPUActiveTime)
	fmt.Printf("Peak GPRs: %d\n", report.PeakRegisterUsage)
	fmt.Printf("Timings: %d\n", len(report.KernelTimings))

	// Output:
	// Active Time: 12.5ms
	// Peak GPRs: 32
	// Timings: 2
}

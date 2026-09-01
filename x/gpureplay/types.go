package gpureplay

import (
	"fmt"
	"time"
)

// ShaderMetric holds static and dynamic performance counters for a compiled Metal compute shader.
type ShaderMetric struct {
	Name             string  `json:"name"`
	RegisterCount    int     `json:"register_count"`    // GPRs per thread
	ThreadgroupBytes int     `json:"threadgroup_bytes"` // Shared memory bytes
	ALUInstructions  int     `json:"alu_instructions"`  // Arithmetic ops
	MemInstructions  int     `json:"mem_instructions"`  // Load/store ops
	ExecutionCostPct float64 `json:"execution_cost_pct"`// % of total GPU active time
}

// KernelTiming describes the measured GPU hardware execution duration of a kernel dispatch.
type KernelTiming struct {
	Name            string        `json:"name"`
	Duration        time.Duration `json:"duration"`
	DispatchCount   int           `json:"dispatch_count"`
	AverageDuration time.Duration `json:"average_duration"`
}

// ReplayReport summarizes GPU profiling results produced by replaying a Metal capture.
type ReplayReport struct {
	TotalGPUActiveTime time.Duration   `json:"total_gpu_active_time"`
	KernelTimings      []KernelTiming  `json:"kernel_timings"`
	ShaderMetrics      []ShaderMetric  `json:"shader_metrics"`
	PeakRegisterUsage  int             `json:"peak_register_usage"`
	RawOutputPath      string          `json:"raw_output_path,omitempty"`
}

// String returns a human-readable summary of the replay report.
func (r *ReplayReport) String() string {
	if r == nil {
		return "<nil replay report>"
	}
	res := fmt.Sprintf("GPU Replay Report (Active Time: %v, Peak GPR: %d):\n", r.TotalGPUActiveTime, r.PeakRegisterUsage)
	for _, k := range r.KernelTimings {
		res += fmt.Sprintf("  - %-30s: %10v (%d dispatches)\n", k.Name, k.Duration, k.DispatchCount)
	}
	return res
}

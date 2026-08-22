//go:build darwin

package coremlconfig

import (
	"testing"

	"github.com/tmc/apple/coreml"
)

func TestComputeDeviceString(t *testing.T) {
	tests := []struct {
		dev  ComputeDevice
		want string
	}{
		{DeviceCPU, "CPU"},
		{DeviceGPU, "GPU"},
		{DeviceNeuralEngine, "ANE"},
		{DeviceUnknown, "unknown"},
	}
	for _, tt := range tests {
		if got := tt.dev.String(); got != tt.want {
			t.Errorf("%d.String() = %q, want %q", tt.dev, got, tt.want)
		}
	}
}

func TestPlanOptionsDefaults(t *testing.T) {
	opts := PlanOptions{}
	if opts.ComputeUnits != 0 {
		t.Errorf("default ComputeUnits = %d, want 0", opts.ComputeUnits)
	}
	if opts.Timeout != 0 {
		t.Errorf("default Timeout = %v, want 0", opts.Timeout)
	}
}

func TestPlanCostFractionEmpty(t *testing.T) {
	p := &Plan{}
	if got := p.ANEFraction(); got != 0 {
		t.Errorf("ANEFraction() on empty = %v, want 0", got)
	}
	if got := p.GPUFraction(); got != 0 {
		t.Errorf("GPUFraction() on empty = %v, want 0", got)
	}
	if got := p.CPUFraction(); got != 0 {
		t.Errorf("CPUFraction() on empty = %v, want 0", got)
	}
}

func TestPlanFilterByDevice(t *testing.T) {
	p := &Plan{
		operations: []OperationPlan{
			{Name: "conv", Device: DeviceNeuralEngine, Cost: 0.5},
			{Name: "softmax", Device: DeviceCPU, Cost: 0.1},
			{Name: "linear", Device: DeviceGPU, Cost: 0.3},
			{Name: "relu", Device: DeviceNeuralEngine, Cost: 0.1},
		},
	}

	ane := p.ANEOperations()
	if len(ane) != 2 {
		t.Errorf("ANEOperations() len = %d, want 2", len(ane))
	}
	gpu := p.GPUOperations()
	if len(gpu) != 1 {
		t.Errorf("GPUOperations() len = %d, want 1", len(gpu))
	}
	cpu := p.CPUOperations()
	if len(cpu) != 1 {
		t.Errorf("CPUOperations() len = %d, want 1", len(cpu))
	}
}

func TestPlanCostFraction(t *testing.T) {
	p := &Plan{
		operations: []OperationPlan{
			{Name: "conv", Device: DeviceNeuralEngine, Cost: 0.6},
			{Name: "softmax", Device: DeviceCPU, Cost: 0.4},
		},
	}
	if got := p.ANEFraction(); got != 0.6 {
		t.Errorf("ANEFraction() = %v, want 0.6", got)
	}
	if got := p.CPUFraction(); got != 0.4 {
		t.Errorf("CPUFraction() = %v, want 0.4", got)
	}
	if got := p.GPUFraction(); got != 0 {
		t.Errorf("GPUFraction() = %v, want 0", got)
	}
}

func TestLoadPlanInvalidPath(t *testing.T) {
	_, err := LoadPlan("/nonexistent/path.mlmodelc", PlanOptions{})
	if err == nil {
		t.Fatal("expected error for invalid path")
	}
}

// TestComputeUnitsResolve pins the mapping LoadPlan applies to a
// configuration. The CPUOnly row is the regression control: coreml's own
// MLComputeUnitsCPUOnly is 0, so an earlier "if opts.ComputeUnits != 0"
// treated an explicit CPU-only request as unset and enabled every unit
// including the Neural Engine. That row fails against the old behavior.
func TestComputeUnitsResolve(t *testing.T) {
	tests := []struct {
		name string
		in   ComputeUnits
		want coreml.MLComputeUnits
	}{
		{"default", ComputeUnitsDefault, coreml.MLComputeUnitsAll},
		{"cpu only", ComputeUnitsCPUOnly, coreml.MLComputeUnitsCPUOnly},
		{"cpu and gpu", ComputeUnitsCPUAndGPU, coreml.MLComputeUnitsCPUAndGPU},
		{"cpu and ane", ComputeUnitsCPUAndNeuralEngine, coreml.MLComputeUnitsCPUAndNeuralEngine},
		{"all", ComputeUnitsAll, coreml.MLComputeUnitsAll},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.resolve()
			if err != nil {
				t.Fatalf("resolve() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestComputeUnitsCPUOnlyIsNotDefault is the narrowest statement of the bug:
// the zero value and an explicit CPU-only request must not resolve alike.
func TestComputeUnitsCPUOnlyIsNotDefault(t *testing.T) {
	zero, err := ComputeUnits(0).resolve()
	if err != nil {
		t.Fatalf("resolve() error = %v", err)
	}
	cpu, err := ComputeUnitsCPUOnly.resolve()
	if err != nil {
		t.Fatalf("resolve() error = %v", err)
	}
	if zero == cpu {
		t.Fatalf("zero value and CPUOnly both resolve to %v; an explicit "+
			"CPU-only request is indistinguishable from an unset field", zero)
	}
	if cpu != coreml.MLComputeUnitsCPUOnly {
		t.Errorf("CPUOnly resolved to %v, want %v", cpu, coreml.MLComputeUnitsCPUOnly)
	}
}

func TestComputeUnitsResolveRejectsUnknown(t *testing.T) {
	if _, err := ComputeUnits(99).resolve(); err == nil {
		t.Fatal("expected error for unknown compute units, got nil")
	}
}

func TestComputeUnitsString(t *testing.T) {
	tests := []struct {
		in   ComputeUnits
		want string
	}{
		{ComputeUnitsDefault, "default"},
		{ComputeUnitsCPUOnly, "CPUOnly"},
		{ComputeUnitsCPUAndGPU, "CPUAndGPU"},
		{ComputeUnitsCPUAndNeuralEngine, "CPUAndNeuralEngine"},
		{ComputeUnitsAll, "All"},
		{ComputeUnits(99), "ComputeUnits(99)"},
	}
	for _, tt := range tests {
		if got := tt.in.String(); got != tt.want {
			t.Errorf("ComputeUnits(%d).String() = %q, want %q", int(tt.in), got, tt.want)
		}
	}
}

// TestLoadPlanRejectsUnknownComputeUnits proves the validation runs on the
// LoadPlan path, not only on resolve. It uses a path that does not exist, so
// a nil error would mean the units were never checked.
func TestLoadPlanRejectsUnknownComputeUnits(t *testing.T) {
	_, err := LoadPlan(t.TempDir()+"/absent.mlmodelc", PlanOptions{ComputeUnits: ComputeUnits(99)})
	if err == nil {
		t.Fatal("expected error for unknown compute units, got nil")
	}
}

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

func TestLoadPlanComputeUnits(t *testing.T) {
	units := []coreml.MLComputeUnits{
		coreml.MLComputeUnitsAll,
		coreml.MLComputeUnitsCPUOnly,
		coreml.MLComputeUnitsCPUAndGPU,
		coreml.MLComputeUnitsCPUAndNeuralEngine,
	}
	for _, u := range units {
		if u.String() == "" {
			t.Errorf("MLComputeUnits(%d).String() is empty", u)
		}
	}
}

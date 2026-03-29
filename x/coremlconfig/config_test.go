//go:build darwin

package coremlconfig

import (
	"testing"

	"github.com/tmc/apple/coreml"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	if cfg == nil {
		t.Fatal("NewConfig returned nil")
	}
	if got := cfg.ComputeUnits(); got != coreml.MLComputeUnitsAll {
		t.Errorf("default ComputeUnits = %v, want MLComputeUnitsAll", got)
	}
}

func TestConfigComputeUnits(t *testing.T) {
	cfg := NewConfig()
	cfg.SetComputeUnits(coreml.MLComputeUnitsCPUOnly)
	if got := cfg.ComputeUnits(); got != coreml.MLComputeUnitsCPUOnly {
		t.Errorf("ComputeUnits = %v, want MLComputeUnitsCPUOnly", got)
	}
}

func TestConfigDeviceMask(t *testing.T) {
	cfg := NewConfig()
	want := DeviceMaskCPU | DeviceMaskNeuralEngine
	cfg.SetDeviceMask(want)
	if got := cfg.DeviceMask(); got != want {
		t.Errorf("DeviceMask = %d, want %d", got, want)
	}
}

func TestConfigANECompilerOptions(t *testing.T) {
	cfg := NewConfig()
	cfg.SetANECompilerOptions("--fast-math")
	if got := cfg.ANECompilerOptions(); got != "--fast-math" {
		t.Errorf("ANECompilerOptions = %q, want %q", got, "--fast-math")
	}
}

func TestConfigPreparesLazily(t *testing.T) {
	cfg := NewConfig()
	cfg.SetPreparesLazily(true)
	if !cfg.PreparesLazily() {
		t.Error("PreparesLazily = false, want true")
	}
	cfg.SetPreparesLazily(false)
	if cfg.PreparesLazily() {
		t.Error("PreparesLazily = true, want false")
	}
}

func TestConfigRaw(t *testing.T) {
	cfg := NewConfig()
	raw := cfg.Raw()
	if raw.GetID() == 0 {
		t.Fatal("Raw() returned nil ID")
	}
}

func TestDeviceMaskConstants(t *testing.T) {
	if DeviceMaskCPU != 1 {
		t.Errorf("DeviceMaskCPU = %d, want 1", DeviceMaskCPU)
	}
	if DeviceMaskGPU != 2 {
		t.Errorf("DeviceMaskGPU = %d, want 2", DeviceMaskGPU)
	}
	if DeviceMaskNeuralEngine != 4 {
		t.Errorf("DeviceMaskNeuralEngine = %d, want 4", DeviceMaskNeuralEngine)
	}
}

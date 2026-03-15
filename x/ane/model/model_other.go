//go:build !darwin

package model

import (
	"fmt"

	"github.com/tmc/apple/coregraphics"
	xane "github.com/tmc/apple/x/ane"
	xanetelemetry "github.com/tmc/apple/x/ane/telemetry"
)

type CompileOptions struct {
	MILText       string
	WeightBlob    []byte
	WeightPath    string
	WeightFiles   []WeightFile
	PackagePath   string
	ModelKey      string
	SharedModel   bool
	QoS           uint32
	PerfStatsMask uint32
}

type WeightFile struct {
	Path string
	Blob []byte
}

type CompileStats struct {
	CompileNS int64
	LoadNS    int64
	TotalNS   int64
}

type EvalStats struct {
	HWExecutionNS uint64
	Metrics       map[string]float64
}

type Kernel struct{}

func Compile(CompileOptions) (*Kernel, error) { return nil, fmt.Errorf("ane model requires darwin") }
func CompileWithStats(CompileOptions) (*Kernel, CompileStats, error) {
	return nil, CompileStats{}, fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) Model() *xane.Model  { return nil }
func (k *Kernel) InputBytes(int) int  { return 0 }
func (k *Kernel) NumInputs() int      { return 0 }
func (k *Kernel) OutputBytes(int) int { return 0 }
func (k *Kernel) NumOutputs() int     { return 0 }
func (k *Kernel) InputSurface(int) coregraphics.IOSurfaceRef {
	return 0
}
func (k *Kernel) InputLayout(int) xane.TensorLayout { return xane.TensorLayout{} }
func (k *Kernel) OutputSurface(int) coregraphics.IOSurfaceRef {
	return 0
}
func (k *Kernel) OutputLayout(int) xane.TensorLayout { return xane.TensorLayout{} }
func (k *Kernel) WriteInput(int, []byte) error       { return fmt.Errorf("ane model requires darwin") }
func (k *Kernel) WriteInputF32(int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) WriteInputFP16(int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) WriteInputFP16Channels(int, int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) ReadOutput(int, []byte) error { return fmt.Errorf("ane model requires darwin") }
func (k *Kernel) ReadOutputF32(int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) ReadOutputFP16(int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) ReadOutputFP16Channels(int, int, []float32) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) Eval() error { return fmt.Errorf("ane model requires darwin") }
func (k *Kernel) EvalWithStats() (EvalStats, error) {
	return EvalStats{}, fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) EvalHWExecutionNS() (uint64, error) {
	return 0, fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) Diagnostics() xanetelemetry.Diagnostics { return xanetelemetry.Diagnostics{} }
func (k *Kernel) EvalWithSignalEvent(uint32, uint64, xane.SharedEventEvalOptions) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) EvalBidirectional(uint32, uint64, uint32, uint64, xane.SharedEventEvalOptions) error {
	return fmt.Errorf("ane model requires darwin")
}
func (k *Kernel) Close() {}

func CopyOutputChannelsToInput(*Kernel, int, int, *Kernel, int, int, int) error {
	return fmt.Errorf("ane model requires darwin")
}
func CopyOutputRangeToInput(*Kernel, int, int, int, *Kernel, int, int, int, int, int) error {
	return fmt.Errorf("ane model requires darwin")
}

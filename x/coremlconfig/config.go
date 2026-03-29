//go:build darwin

package coremlconfig

import (
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	privatecoreml "github.com/tmc/apple/private/coreml"
)

// Config wraps MLModelConfiguration with access to both public and
// private ANE compiler settings.
type Config struct {
	raw coreml.MLModelConfiguration
}

// NewConfig creates a default model configuration targeting all compute units.
func NewConfig() *Config {
	cfg := coreml.NewMLModelConfiguration()
	cfg.SetComputeUnits(coreml.MLComputeUnitsAll)
	return &Config{raw: cfg}
}

// SetComputeUnits sets which processing units the model may use.
func (c *Config) SetComputeUnits(units coreml.MLComputeUnits) {
	c.raw.SetComputeUnits(units)
}

// ComputeUnits returns the current compute unit setting.
func (c *Config) ComputeUnits() coreml.MLComputeUnits {
	return c.raw.ComputeUnits()
}

// DeviceMask bits for E5rtComputeDeviceTypeMask.
const (
	DeviceMaskCPU          uint64 = 1 << 0
	DeviceMaskGPU          uint64 = 1 << 1
	DeviceMaskNeuralEngine uint64 = 1 << 2
)

// SetDeviceMask sets the E5rt compute device type mask (private API).
// Combine DeviceMask* constants with bitwise OR.
func (c *Config) SetDeviceMask(mask uint64) {
	objc.Send[struct{}](c.raw.ID, objc.Sel("setE5rtComputeDeviceTypeMask:"), mask)
}

// DeviceMask returns the E5rt compute device type mask.
func (c *Config) DeviceMask() uint64 {
	return objc.Send[uint64](c.raw.ID, objc.Sel("e5rtComputeDeviceTypeMask"))
}

// SetANECompilerOptions sets the E5rt custom ANE compiler options string (private API).
func (c *Config) SetANECompilerOptions(opts string) {
	objc.Send[struct{}](c.raw.ID, objc.Sel("setE5rtCustomANECompilerOptions:"), objc.String(opts))
}

// ANECompilerOptions returns the current E5rt custom ANE compiler options.
func (c *Config) ANECompilerOptions() string {
	rv := objc.Send[objc.ID](c.raw.ID, objc.Sel("e5rtCustomANECompilerOptions"))
	if rv == 0 {
		return ""
	}
	return objc.Send[string](rv, objc.Sel("UTF8String"))
}

// NeuralEngineCompilerOptions returns the neural engine compiler options object (private API).
// Returns nil if not set.
func (c *Config) NeuralEngineCompilerOptions() objectivec.Object {
	rv := objc.Send[objc.ID](c.raw.ID, objc.Sel("neuralEngineCompilerOptions"))
	return objectivec.Object{ID: rv}
}

// SetPreparesLazily controls whether model preparation is deferred.
func (c *Config) SetPreparesLazily(lazy bool) {
	objc.Send[struct{}](c.raw.ID, objc.Sel("setPreparesLazily:"), lazy)
}

// PreparesLazily returns whether lazy preparation is enabled.
func (c *Config) PreparesLazily() bool {
	return objc.Send[bool](c.raw.ID, objc.Sel("preparesLazily"))
}

// Raw returns the underlying MLModelConfiguration for use with
// low-level CoreML APIs.
func (c *Config) Raw() coreml.MLModelConfiguration {
	return c.raw
}

// SetPreferredMTLDevice routes GPU work to a specific Metal device (private API).
func (c *Config) SetPreferredMTLDevice(device objectivec.IObject) {
	pcfg := privatecoreml.MLModelConfigurationFromID(c.raw.ID)
	pcfg.SetPreferredMTLDevice(device)
}

// SetUsePrecompiledE5Bundle controls whether the E5 precompiled bundle is used (private API).
func (c *Config) SetUsePrecompiledE5Bundle(use bool) {
	pcfg := privatecoreml.MLModelConfigurationFromID(c.raw.ID)
	pcfg.SetUsePrecompiledE5Bundle(use)
}

// UsePrecompiledE5Bundle returns whether the E5 precompiled bundle is used.
func (c *Config) UsePrecompiledE5Bundle() bool {
	pcfg := privatecoreml.MLModelConfigurationFromID(c.raw.ID)
	return pcfg.UsePrecompiledE5Bundle()
}

// SetSerializesMILTextForDebugging enables MIL text serialization for debugging (private API).
func (c *Config) SetSerializesMILTextForDebugging(serialize bool) {
	pcfg := privatecoreml.MLModelConfigurationFromID(c.raw.ID)
	pcfg.SetSerializesMILTextForDebugging(serialize)
}

// SerializesMILTextForDebugging returns whether MIL text debugging is enabled.
func (c *Config) SerializesMILTextForDebugging() bool {
	pcfg := privatecoreml.MLModelConfigurationFromID(c.raw.ID)
	return pcfg.SerializesMILTextForDebugging()
}

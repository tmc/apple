//go:build darwin

package coremlconfig

import (
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	privatecoreml "github.com/tmc/apple/private/coreml"
)

// Config wraps MLModelConfiguration with access to both public and
// private ANE compiler settings.
type Config struct {
	raw coreml.MLModelConfiguration
}

func (c *Config) priv() privatecoreml.MLModelConfiguration {
	return privatecoreml.MLModelConfigurationFromID(c.raw.ID)
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
	c.priv().SetPreferredMTLDevice(device)
}

// SetUsePrecompiledE5Bundle controls whether the E5 precompiled bundle is used (private API).
func (c *Config) SetUsePrecompiledE5Bundle(use bool) {
	c.priv().SetUsePrecompiledE5Bundle(use)
}

// UsePrecompiledE5Bundle returns whether the E5 precompiled bundle is used.
func (c *Config) UsePrecompiledE5Bundle() bool {
	return c.priv().UsePrecompiledE5Bundle()
}

// SetSerializesMILTextForDebugging enables MIL text serialization for debugging (private API).
func (c *Config) SetSerializesMILTextForDebugging(serialize bool) {
	c.priv().SetSerializesMILTextForDebugging(serialize)
}

// SerializesMILTextForDebugging returns whether MIL text debugging is enabled.
func (c *Config) SerializesMILTextForDebugging() bool {
	return c.priv().SerializesMILTextForDebugging()
}

// SetPredictionConcurrencyHint sets the model preparation/prediction concurrency hint.
func (c *Config) SetPredictionConcurrencyHint(hint int64) {
	c.priv().SetPredictionConcurrencyHint(hint)
}

// PredictionConcurrencyHint returns the configured concurrency hint.
func (c *Config) PredictionConcurrencyHint() int64 {
	return c.priv().PredictionConcurrencyHint()
}

// SetAllowsInstrumentation enables private runtime instrumentation.
func (c *Config) SetAllowsInstrumentation(enable bool) {
	c.priv().SetAllowsInstrumentation(enable)
}

// AllowsInstrumentation reports whether instrumentation is enabled.
func (c *Config) AllowsInstrumentation() bool {
	return c.priv().AllowsInstrumentation()
}

// SetProfilingOptions sets private profiling flags.
func (c *Config) SetProfilingOptions(options int64) {
	c.priv().SetProfilingOptions(options)
}

// ProfilingOptions returns the private profiling flags.
func (c *Config) ProfilingOptions() int64 {
	return c.priv().ProfilingOptions()
}

// SetRootModelURL sets the root model URL for nested-model loads.
func (c *Config) SetRootModelURL(url foundation.INSURL) {
	c.priv().SetRootModelURL(url)
}

// RootModelURL returns the configured root model URL.
func (c *Config) RootModelURL() foundation.INSURL {
	return c.priv().RootModelURL()
}

// SetParentModelName sets the parent model name for nested-model loads.
func (c *Config) SetParentModelName(name string) {
	c.priv().SetParentModelName(name)
}

// ParentModelName returns the configured parent model name.
func (c *Config) ParentModelName() string {
	return c.priv().ParentModelName()
}

// SetDynamicCallableFunctions sets the private dynamic callable functions dictionary.
func (c *Config) SetDynamicCallableFunctions(functions foundation.INSDictionary) {
	c.priv().SetE5rtDynamicCallableFunctions(functions)
}

// DynamicCallableFunctions returns the private dynamic callable functions dictionary.
func (c *Config) DynamicCallableFunctions() foundation.INSDictionary {
	return c.priv().E5rtDynamicCallableFunctions()
}

// SetMutableMILWeightURLs sets the private mutable MIL weight URL mapping.
func (c *Config) SetMutableMILWeightURLs(urls foundation.INSDictionary) {
	c.priv().SetE5rtMutableMILWeightURLs(urls)
}

// MutableMILWeightURLs returns the mutable MIL weight URL mapping.
func (c *Config) MutableMILWeightURLs() foundation.INSDictionary {
	return c.priv().E5rtMutableMILWeightURLs()
}

// SetAllowFloat16AccumulationOnGPU toggles float16 accumulation for GPU execution.
func (c *Config) SetAllowFloat16AccumulationOnGPU(enable bool) {
	c.priv().SetAllowFloat16AccumulationOnGPU(enable)
}

// AllowFloat16AccumulationOnGPU reports whether float16 GPU accumulation is enabled.
func (c *Config) AllowFloat16AccumulationOnGPU() bool {
	return c.priv().AllowFloat16AccumulationOnGPU()
}

// SetExperimentalMLE5EngineUsage sets the experimental private engine usage mode.
func (c *Config) SetExperimentalMLE5EngineUsage(value int64) {
	c.priv().SetExperimentalMLE5EngineUsage(value)
}

// ExperimentalMLE5EngineUsage returns the experimental private engine usage mode.
func (c *Config) ExperimentalMLE5EngineUsage() int64 {
	return c.priv().ExperimentalMLE5EngineUsage()
}

// SetExperimentalMLE5BNNSGraphBackendUsage sets the experimental BNNS graph backend mode.
func (c *Config) SetExperimentalMLE5BNNSGraphBackendUsage(value int64) {
	c.priv().SetExperimentalMLE5BNNSGraphBackendUsage(value)
}

// ExperimentalMLE5BNNSGraphBackendUsage returns the experimental BNNS graph backend mode.
func (c *Config) ExperimentalMLE5BNNSGraphBackendUsage() int64 {
	return c.priv().ExperimentalMLE5BNNSGraphBackendUsage()
}

// SetExperimentalMLProgramEncryptedCacheUsage sets the experimental encrypted cache mode.
func (c *Config) SetExperimentalMLProgramEncryptedCacheUsage(value int64) {
	c.priv().SetExperimentalMLProgramEncryptedCacheUsage(value)
}

// ExperimentalMLProgramEncryptedCacheUsage returns the experimental encrypted cache mode.
func (c *Config) ExperimentalMLProgramEncryptedCacheUsage() int64 {
	return c.priv().ExperimentalMLProgramEncryptedCacheUsage()
}

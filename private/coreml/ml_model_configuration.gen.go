// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelConfiguration] class.
var (
	_MLModelConfigurationClass     MLModelConfigurationClass
	_MLModelConfigurationClassOnce sync.Once
)

func getMLModelConfigurationClass() MLModelConfigurationClass {
	_MLModelConfigurationClassOnce.Do(func() {
		_MLModelConfigurationClass = MLModelConfigurationClass{class: objc.GetClass("MLModelConfiguration")}
	})
	return _MLModelConfigurationClass
}

// GetMLModelConfigurationClass returns the class object for MLModelConfiguration.
func GetMLModelConfigurationClass() MLModelConfigurationClass {
	return getMLModelConfigurationClass()
}

type MLModelConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelConfigurationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelConfigurationClass) Alloc() MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelConfiguration.AllowBackgroundGPUCompute]
//   - [MLModelConfiguration.SetAllowBackgroundGPUCompute]
//   - [MLModelConfiguration.AllowBackgroundGPUComputeSetting]
//   - [MLModelConfiguration.SetAllowBackgroundGPUComputeSetting]
//   - [MLModelConfiguration.AllowFloat16AccumulationOnGPU]
//   - [MLModelConfiguration.SetAllowFloat16AccumulationOnGPU]
//   - [MLModelConfiguration.AllowsInstrumentation]
//   - [MLModelConfiguration.SetAllowsInstrumentation]
//   - [MLModelConfiguration.BnnsGraphBackendUsageToString]
//   - [MLModelConfiguration.ComputeUnitsToString]
//   - [MLModelConfiguration.E5rtComputeDeviceTypeMask]
//   - [MLModelConfiguration.SetE5rtComputeDeviceTypeMask]
//   - [MLModelConfiguration.E5rtCustomANECompilerOptions]
//   - [MLModelConfiguration.SetE5rtCustomANECompilerOptions]
//   - [MLModelConfiguration.E5rtDynamicCallableFunctions]
//   - [MLModelConfiguration.SetE5rtDynamicCallableFunctions]
//   - [MLModelConfiguration.E5rtMutableMILWeightURLs]
//   - [MLModelConfiguration.SetE5rtMutableMILWeightURLs]
//   - [MLModelConfiguration.EnableTestVectorMode]
//   - [MLModelConfiguration.SetEnableTestVectorMode]
//   - [MLModelConfiguration.ExperimentalMLE5BNNSGraphBackendUsage]
//   - [MLModelConfiguration.SetExperimentalMLE5BNNSGraphBackendUsage]
//   - [MLModelConfiguration.ExperimentalMLE5BNNSGraphBackendUsageMultiSegment]
//   - [MLModelConfiguration.SetExperimentalMLE5BNNSGraphBackendUsageMultiSegment]
//   - [MLModelConfiguration.ExperimentalMLE5EngineUsage]
//   - [MLModelConfiguration.SetExperimentalMLE5EngineUsage]
//   - [MLModelConfiguration.ExperimentalMLE5EngineUsageToString]
//   - [MLModelConfiguration.ExperimentalMLProgramEncryptedCacheUsage]
//   - [MLModelConfiguration.SetExperimentalMLProgramEncryptedCacheUsage]
//   - [MLModelConfiguration.ExperimentalMLProgramEncryptedCacheUsageToString]
//   - [MLModelConfiguration.IsEqualToModelConfiguration]
//   - [MLModelConfiguration.NeuralEngineCompilerOptions]
//   - [MLModelConfiguration.ParentModelName]
//   - [MLModelConfiguration.SetParentModelName]
//   - [MLModelConfiguration.PredictionConcurrencyHint]
//   - [MLModelConfiguration.SetPredictionConcurrencyHint]
//   - [MLModelConfiguration.PreferredMTLDevice]
//   - [MLModelConfiguration.SetPreferredMTLDevice]
//   - [MLModelConfiguration.PreparesLazily]
//   - [MLModelConfiguration.SetPreparesLazily]
//   - [MLModelConfiguration.ProfilingOptions]
//   - [MLModelConfiguration.SetProfilingOptions]
//   - [MLModelConfiguration.RootModelURL]
//   - [MLModelConfiguration.SetRootModelURL]
//   - [MLModelConfiguration.SerializesMILTextForDebugging]
//   - [MLModelConfiguration.SetSerializesMILTextForDebugging]
//   - [MLModelConfiguration.SpecializationUsesMPSGraphExecutable]
//   - [MLModelConfiguration.SetSpecializationUsesMPSGraphExecutable]
//   - [MLModelConfiguration.TrainWithMLCompute]
//   - [MLModelConfiguration.SetTrainWithMLCompute]
//   - [MLModelConfiguration.UsePrecompiledE5Bundle]
//   - [MLModelConfiguration.SetUsePrecompiledE5Bundle]
//   - [MLModelConfiguration.UsePreloadedKey]
//   - [MLModelConfiguration.SetUsePreloadedKey]
//   - [MLModelConfiguration.UseWatchSPIForScribble]
//   - [MLModelConfiguration.SetUseWatchSPIForScribble]
//   - [MLModelConfiguration.UsesCompileTimeMPSGraphTypeInferenceForModelVersion]
//   - [MLModelConfiguration.InitWithCoder]
//   - [MLModelConfiguration.InitWithComputeUnits]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration
type MLModelConfiguration struct {
	objectivec.Object
}

// MLModelConfigurationFromID constructs a [MLModelConfiguration] from an objc.ID.
func MLModelConfigurationFromID(id objc.ID) MLModelConfiguration {
	return MLModelConfiguration{objectivec.Object{ID: id}}
}

// Ensure MLModelConfiguration implements IMLModelConfiguration.
var _ IMLModelConfiguration = MLModelConfiguration{}

// An interface definition for the [MLModelConfiguration] class.
//
// # Methods
//
//   - [IMLModelConfiguration.AllowBackgroundGPUCompute]
//   - [IMLModelConfiguration.SetAllowBackgroundGPUCompute]
//   - [IMLModelConfiguration.AllowBackgroundGPUComputeSetting]
//   - [IMLModelConfiguration.SetAllowBackgroundGPUComputeSetting]
//   - [IMLModelConfiguration.AllowFloat16AccumulationOnGPU]
//   - [IMLModelConfiguration.SetAllowFloat16AccumulationOnGPU]
//   - [IMLModelConfiguration.AllowsInstrumentation]
//   - [IMLModelConfiguration.SetAllowsInstrumentation]
//   - [IMLModelConfiguration.BnnsGraphBackendUsageToString]
//   - [IMLModelConfiguration.ComputeUnitsToString]
//   - [IMLModelConfiguration.E5rtComputeDeviceTypeMask]
//   - [IMLModelConfiguration.SetE5rtComputeDeviceTypeMask]
//   - [IMLModelConfiguration.E5rtCustomANECompilerOptions]
//   - [IMLModelConfiguration.SetE5rtCustomANECompilerOptions]
//   - [IMLModelConfiguration.E5rtDynamicCallableFunctions]
//   - [IMLModelConfiguration.SetE5rtDynamicCallableFunctions]
//   - [IMLModelConfiguration.E5rtMutableMILWeightURLs]
//   - [IMLModelConfiguration.SetE5rtMutableMILWeightURLs]
//   - [IMLModelConfiguration.EnableTestVectorMode]
//   - [IMLModelConfiguration.SetEnableTestVectorMode]
//   - [IMLModelConfiguration.ExperimentalMLE5BNNSGraphBackendUsage]
//   - [IMLModelConfiguration.SetExperimentalMLE5BNNSGraphBackendUsage]
//   - [IMLModelConfiguration.ExperimentalMLE5BNNSGraphBackendUsageMultiSegment]
//   - [IMLModelConfiguration.SetExperimentalMLE5BNNSGraphBackendUsageMultiSegment]
//   - [IMLModelConfiguration.ExperimentalMLE5EngineUsage]
//   - [IMLModelConfiguration.SetExperimentalMLE5EngineUsage]
//   - [IMLModelConfiguration.ExperimentalMLE5EngineUsageToString]
//   - [IMLModelConfiguration.ExperimentalMLProgramEncryptedCacheUsage]
//   - [IMLModelConfiguration.SetExperimentalMLProgramEncryptedCacheUsage]
//   - [IMLModelConfiguration.ExperimentalMLProgramEncryptedCacheUsageToString]
//   - [IMLModelConfiguration.IsEqualToModelConfiguration]
//   - [IMLModelConfiguration.NeuralEngineCompilerOptions]
//   - [IMLModelConfiguration.ParentModelName]
//   - [IMLModelConfiguration.SetParentModelName]
//   - [IMLModelConfiguration.PredictionConcurrencyHint]
//   - [IMLModelConfiguration.SetPredictionConcurrencyHint]
//   - [IMLModelConfiguration.PreferredMTLDevice]
//   - [IMLModelConfiguration.SetPreferredMTLDevice]
//   - [IMLModelConfiguration.PreparesLazily]
//   - [IMLModelConfiguration.SetPreparesLazily]
//   - [IMLModelConfiguration.ProfilingOptions]
//   - [IMLModelConfiguration.SetProfilingOptions]
//   - [IMLModelConfiguration.RootModelURL]
//   - [IMLModelConfiguration.SetRootModelURL]
//   - [IMLModelConfiguration.SerializesMILTextForDebugging]
//   - [IMLModelConfiguration.SetSerializesMILTextForDebugging]
//   - [IMLModelConfiguration.SpecializationUsesMPSGraphExecutable]
//   - [IMLModelConfiguration.SetSpecializationUsesMPSGraphExecutable]
//   - [IMLModelConfiguration.TrainWithMLCompute]
//   - [IMLModelConfiguration.SetTrainWithMLCompute]
//   - [IMLModelConfiguration.UsePrecompiledE5Bundle]
//   - [IMLModelConfiguration.SetUsePrecompiledE5Bundle]
//   - [IMLModelConfiguration.UsePreloadedKey]
//   - [IMLModelConfiguration.SetUsePreloadedKey]
//   - [IMLModelConfiguration.UseWatchSPIForScribble]
//   - [IMLModelConfiguration.SetUseWatchSPIForScribble]
//   - [IMLModelConfiguration.UsesCompileTimeMPSGraphTypeInferenceForModelVersion]
//   - [IMLModelConfiguration.InitWithCoder]
//   - [IMLModelConfiguration.InitWithComputeUnits]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration
type IMLModelConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	AllowBackgroundGPUCompute() bool
	SetAllowBackgroundGPUCompute(value bool)
	AllowBackgroundGPUComputeSetting() bool
	SetAllowBackgroundGPUComputeSetting(value bool)
	AllowFloat16AccumulationOnGPU() bool
	SetAllowFloat16AccumulationOnGPU(value bool)
	AllowsInstrumentation() bool
	SetAllowsInstrumentation(value bool)
	BnnsGraphBackendUsageToString(string_ int64) objectivec.IObject
	ComputeUnitsToString(string_ int64) objectivec.IObject
	E5rtComputeDeviceTypeMask() uint64
	SetE5rtComputeDeviceTypeMask(value uint64)
	E5rtCustomANECompilerOptions() string
	SetE5rtCustomANECompilerOptions(value string)
	E5rtDynamicCallableFunctions() foundation.INSDictionary
	SetE5rtDynamicCallableFunctions(value foundation.INSDictionary)
	E5rtMutableMILWeightURLs() foundation.INSDictionary
	SetE5rtMutableMILWeightURLs(value foundation.INSDictionary)
	EnableTestVectorMode() bool
	SetEnableTestVectorMode(value bool)
	ExperimentalMLE5BNNSGraphBackendUsage() int64
	SetExperimentalMLE5BNNSGraphBackendUsage(value int64)
	ExperimentalMLE5BNNSGraphBackendUsageMultiSegment() int64
	SetExperimentalMLE5BNNSGraphBackendUsageMultiSegment(value int64)
	ExperimentalMLE5EngineUsage() int64
	SetExperimentalMLE5EngineUsage(value int64)
	ExperimentalMLE5EngineUsageToString(string_ int64) objectivec.IObject
	ExperimentalMLProgramEncryptedCacheUsage() int64
	SetExperimentalMLProgramEncryptedCacheUsage(value int64)
	ExperimentalMLProgramEncryptedCacheUsageToString(string_ int64) objectivec.IObject
	IsEqualToModelConfiguration(configuration objectivec.IObject) bool
	NeuralEngineCompilerOptions() objectivec.IObject
	ParentModelName() string
	SetParentModelName(value string)
	PredictionConcurrencyHint() int64
	SetPredictionConcurrencyHint(value int64)
	PreferredMTLDevice() objectivec.IObject
	SetPreferredMTLDevice(value objectivec.IObject)
	PreparesLazily() bool
	SetPreparesLazily(value bool)
	ProfilingOptions() int64
	SetProfilingOptions(value int64)
	RootModelURL() foundation.INSURL
	SetRootModelURL(value foundation.INSURL)
	SerializesMILTextForDebugging() bool
	SetSerializesMILTextForDebugging(value bool)
	SpecializationUsesMPSGraphExecutable() bool
	SetSpecializationUsesMPSGraphExecutable(value bool)
	TrainWithMLCompute() bool
	SetTrainWithMLCompute(value bool)
	UsePrecompiledE5Bundle() bool
	SetUsePrecompiledE5Bundle(value bool)
	UsePreloadedKey() bool
	SetUsePreloadedKey(value bool)
	UseWatchSPIForScribble() bool
	SetUseWatchSPIForScribble(value bool)
	UsesCompileTimeMPSGraphTypeInferenceForModelVersion(version objectivec.IObject) bool
	InitWithCoder(coder foundation.INSCoder) MLModelConfiguration
	InitWithComputeUnits(units int64) MLModelConfiguration
}

// Init initializes the instance.
func (m MLModelConfiguration) Init() MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelConfiguration) Autorelease() MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelConfiguration creates a new MLModelConfiguration instance.
func NewMLModelConfiguration() MLModelConfiguration {
	class := getMLModelConfigurationClass()
	rv := objc.Send[MLModelConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/initWithCoder:
func NewModelConfigurationWithCoder(coder objectivec.IObject) MLModelConfiguration {
	instance := getMLModelConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLModelConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/initWithComputeUnits:
func NewModelConfigurationWithComputeUnits(units int64) MLModelConfiguration {
	instance := getMLModelConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComputeUnits:"), units)
	return MLModelConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/bnnsGraphBackendUsageToString:
func (m MLModelConfiguration) BnnsGraphBackendUsageToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("bnnsGraphBackendUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/computeUnitsToString:
func (m MLModelConfiguration) ComputeUnitsToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("computeUnitsToString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLE5EngineUsageToString:
func (m MLModelConfiguration) ExperimentalMLE5EngineUsageToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("experimentalMLE5EngineUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLProgramEncryptedCacheUsageToString:
func (m MLModelConfiguration) ExperimentalMLProgramEncryptedCacheUsageToString(string_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("experimentalMLProgramEncryptedCacheUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/isEqualToModelConfiguration:
func (m MLModelConfiguration) IsEqualToModelConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEqualToModelConfiguration:"), configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/neuralEngineCompilerOptions
func (m MLModelConfiguration) NeuralEngineCompilerOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("neuralEngineCompilerOptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/usesCompileTimeMPSGraphTypeInferenceForModelVersion:
func (m MLModelConfiguration) UsesCompileTimeMPSGraphTypeInferenceForModelVersion(version objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usesCompileTimeMPSGraphTypeInferenceForModelVersion:"), version)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/initWithCoder:
func (m MLModelConfiguration) InitWithCoder(coder foundation.INSCoder) MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/initWithComputeUnits:
func (m MLModelConfiguration) InitWithComputeUnits(units int64) MLModelConfiguration {
	rv := objc.Send[MLModelConfiguration](m.ID, objc.Sel("initWithComputeUnits:"), units)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/defaultConfiguration
func (_MLModelConfigurationClass MLModelConfigurationClass) DefaultConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelConfigurationClass.class), objc.Sel("defaultConfiguration"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/supportsSecureCoding
func (_MLModelConfigurationClass MLModelConfigurationClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLModelConfigurationClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowBackgroundGPUCompute
func (m MLModelConfiguration) AllowBackgroundGPUCompute() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowBackgroundGPUCompute"))
	return rv
}
func (m MLModelConfiguration) SetAllowBackgroundGPUCompute(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowBackgroundGPUCompute:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowBackgroundGPUComputeSetting
func (m MLModelConfiguration) AllowBackgroundGPUComputeSetting() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowBackgroundGPUComputeSetting"))
	return rv
}
func (m MLModelConfiguration) SetAllowBackgroundGPUComputeSetting(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowBackgroundGPUComputeSetting:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowFloat16AccumulationOnGPU
func (m MLModelConfiguration) AllowFloat16AccumulationOnGPU() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowFloat16AccumulationOnGPU"))
	return rv
}
func (m MLModelConfiguration) SetAllowFloat16AccumulationOnGPU(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowFloat16AccumulationOnGPU:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowsInstrumentation
func (m MLModelConfiguration) AllowsInstrumentation() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsInstrumentation"))
	return rv
}
func (m MLModelConfiguration) SetAllowsInstrumentation(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsInstrumentation:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/e5rtComputeDeviceTypeMask
func (m MLModelConfiguration) E5rtComputeDeviceTypeMask() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("e5rtComputeDeviceTypeMask"))
	return rv
}
func (m MLModelConfiguration) SetE5rtComputeDeviceTypeMask(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setE5rtComputeDeviceTypeMask:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/e5rtCustomANECompilerOptions
func (m MLModelConfiguration) E5rtCustomANECompilerOptions() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("e5rtCustomANECompilerOptions"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetE5rtCustomANECompilerOptions(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setE5rtCustomANECompilerOptions:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/e5rtDynamicCallableFunctions
func (m MLModelConfiguration) E5rtDynamicCallableFunctions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("e5rtDynamicCallableFunctions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetE5rtDynamicCallableFunctions(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setE5rtDynamicCallableFunctions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/e5rtMutableMILWeightURLs
func (m MLModelConfiguration) E5rtMutableMILWeightURLs() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("e5rtMutableMILWeightURLs"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetE5rtMutableMILWeightURLs(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setE5rtMutableMILWeightURLs:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/enableTestVectorMode
func (m MLModelConfiguration) EnableTestVectorMode() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("enableTestVectorMode"))
	return rv
}
func (m MLModelConfiguration) SetEnableTestVectorMode(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnableTestVectorMode:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLE5BNNSGraphBackendUsage
func (m MLModelConfiguration) ExperimentalMLE5BNNSGraphBackendUsage() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("experimentalMLE5BNNSGraphBackendUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5BNNSGraphBackendUsage(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setExperimentalMLE5BNNSGraphBackendUsage:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLE5BNNSGraphBackendUsageMultiSegment
func (m MLModelConfiguration) ExperimentalMLE5BNNSGraphBackendUsageMultiSegment() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("experimentalMLE5BNNSGraphBackendUsageMultiSegment"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5BNNSGraphBackendUsageMultiSegment(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setExperimentalMLE5BNNSGraphBackendUsageMultiSegment:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLE5EngineUsage
func (m MLModelConfiguration) ExperimentalMLE5EngineUsage() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("experimentalMLE5EngineUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5EngineUsage(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setExperimentalMLE5EngineUsage:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/experimentalMLProgramEncryptedCacheUsage
func (m MLModelConfiguration) ExperimentalMLProgramEncryptedCacheUsage() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("experimentalMLProgramEncryptedCacheUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLProgramEncryptedCacheUsage(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setExperimentalMLProgramEncryptedCacheUsage:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/parentModelName
func (m MLModelConfiguration) ParentModelName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parentModelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetParentModelName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setParentModelName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/predictionConcurrencyHint
func (m MLModelConfiguration) PredictionConcurrencyHint() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("predictionConcurrencyHint"))
	return rv
}
func (m MLModelConfiguration) SetPredictionConcurrencyHint(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setPredictionConcurrencyHint:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/preferredMTLDevice
func (m MLModelConfiguration) PreferredMTLDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("preferredMTLDevice"))
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) SetPreferredMTLDevice(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMTLDevice:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/preparesLazily
func (m MLModelConfiguration) PreparesLazily() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("preparesLazily"))
	return rv
}
func (m MLModelConfiguration) SetPreparesLazily(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreparesLazily:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/profilingOptions
func (m MLModelConfiguration) ProfilingOptions() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("profilingOptions"))
	return rv
}
func (m MLModelConfiguration) SetProfilingOptions(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setProfilingOptions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/rootModelURL
func (m MLModelConfiguration) RootModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("rootModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetRootModelURL(value foundation.INSURL) {
	objc.Send[struct{}](m.ID, objc.Sel("setRootModelURL:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/serializesMILTextForDebugging
func (m MLModelConfiguration) SerializesMILTextForDebugging() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("serializesMILTextForDebugging"))
	return rv
}
func (m MLModelConfiguration) SetSerializesMILTextForDebugging(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSerializesMILTextForDebugging:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/specializationUsesMPSGraphExecutable
func (m MLModelConfiguration) SpecializationUsesMPSGraphExecutable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("specializationUsesMPSGraphExecutable"))
	return rv
}
func (m MLModelConfiguration) SetSpecializationUsesMPSGraphExecutable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSpecializationUsesMPSGraphExecutable:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/trainWithMLCompute
func (m MLModelConfiguration) TrainWithMLCompute() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("trainWithMLCompute"))
	return rv
}
func (m MLModelConfiguration) SetTrainWithMLCompute(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setTrainWithMLCompute:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/usePrecompiledE5Bundle
func (m MLModelConfiguration) UsePrecompiledE5Bundle() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usePrecompiledE5Bundle"))
	return rv
}
func (m MLModelConfiguration) SetUsePrecompiledE5Bundle(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsePrecompiledE5Bundle:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/usePreloadedKey
func (m MLModelConfiguration) UsePreloadedKey() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usePreloadedKey"))
	return rv
}
func (m MLModelConfiguration) SetUsePreloadedKey(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsePreloadedKey:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/useWatchSPIForScribble
func (m MLModelConfiguration) UseWatchSPIForScribble() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useWatchSPIForScribble"))
	return rv
}
func (m MLModelConfiguration) SetUseWatchSPIForScribble(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUseWatchSPIForScribble:"), value)
}

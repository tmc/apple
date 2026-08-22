// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
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
	rv := objc.SendIfResponds[MLModelConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
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
//   - [MLModelConfiguration.InitWithComputeUnits]
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
//   - [IMLModelConfiguration.InitWithComputeUnits]
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
	PreferredMTLDevice() metal.MTLDeviceObject
	SetPreferredMTLDevice(value metal.MTLDeviceObject)
	PreparesLazily() bool
	SetPreparesLazily(value bool)
	ProfilingOptions() int64
	SetProfilingOptions(value int64)
	RootModelURL() foundation.NSURL
	SetRootModelURL(value foundation.NSURL)
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
	InitWithComputeUnits(units int64) MLModelConfiguration
}

// Init initializes the instance.
func (m MLModelConfiguration) Init() MLModelConfiguration {
	rv := objc.SendIfResponds[MLModelConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelConfiguration) Autorelease() MLModelConfiguration {
	rv := objc.SendIfResponds[MLModelConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelConfiguration creates a new MLModelConfiguration instance.
func NewMLModelConfiguration() MLModelConfiguration {
	class := getMLModelConfigurationClass()
	rv := objc.SendIfResponds[MLModelConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelConfigurationWithComputeUnits(units int64) MLModelConfiguration {
	instance := getMLModelConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithComputeUnits:"), units)
	return MLModelConfigurationFromID(rv)
}

func (m MLModelConfiguration) BnnsGraphBackendUsageToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("bnnsGraphBackendUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) ComputeUnitsToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeUnitsToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) ExperimentalMLE5EngineUsageToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("experimentalMLE5EngineUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) ExperimentalMLProgramEncryptedCacheUsageToString(string_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("experimentalMLProgramEncryptedCacheUsageToString:"), string_)
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) IsEqualToModelConfiguration(configuration objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isEqualToModelConfiguration:"), configuration)
	return rv
}
func (m MLModelConfiguration) NeuralEngineCompilerOptions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("neuralEngineCompilerOptions"))
	return objectivec.Object{ID: rv}
}
func (m MLModelConfiguration) UsesCompileTimeMPSGraphTypeInferenceForModelVersion(version objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("usesCompileTimeMPSGraphTypeInferenceForModelVersion:"), version)
	return rv
}
func (m MLModelConfiguration) InitWithComputeUnits(units int64) MLModelConfiguration {
	rv := objc.SendIfResponds[MLModelConfiguration](m.ID, objc.Sel("initWithComputeUnits:"), units)
	return rv
}

func (_MLModelConfigurationClass MLModelConfigurationClass) DefaultConfiguration() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelConfigurationClass.class), objc.Sel("defaultConfiguration"))
	return objectivec.Object{ID: rv}
}
func (_MLModelConfigurationClass MLModelConfigurationClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLModelConfigurationClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLModelConfiguration) AllowBackgroundGPUCompute() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowBackgroundGPUCompute"))
	return rv
}
func (m MLModelConfiguration) SetAllowBackgroundGPUCompute(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowBackgroundGPUCompute:"), value)
}
func (m MLModelConfiguration) AllowBackgroundGPUComputeSetting() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowBackgroundGPUComputeSetting"))
	return rv
}
func (m MLModelConfiguration) SetAllowBackgroundGPUComputeSetting(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowBackgroundGPUComputeSetting:"), value)
}
func (m MLModelConfiguration) AllowFloat16AccumulationOnGPU() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowFloat16AccumulationOnGPU"))
	return rv
}
func (m MLModelConfiguration) SetAllowFloat16AccumulationOnGPU(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowFloat16AccumulationOnGPU:"), value)
}
func (m MLModelConfiguration) AllowsInstrumentation() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("allowsInstrumentation"))
	return rv
}
func (m MLModelConfiguration) SetAllowsInstrumentation(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setAllowsInstrumentation:"), value)
}
func (m MLModelConfiguration) E5rtComputeDeviceTypeMask() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("e5rtComputeDeviceTypeMask"))
	return rv
}
func (m MLModelConfiguration) SetE5rtComputeDeviceTypeMask(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setE5rtComputeDeviceTypeMask:"), value)
}
func (m MLModelConfiguration) E5rtCustomANECompilerOptions() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("e5rtCustomANECompilerOptions"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetE5rtCustomANECompilerOptions(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setE5rtCustomANECompilerOptions:"), objc.String(value))
}
func (m MLModelConfiguration) E5rtDynamicCallableFunctions() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("e5rtDynamicCallableFunctions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetE5rtDynamicCallableFunctions(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setE5rtDynamicCallableFunctions:"), value)
}
func (m MLModelConfiguration) E5rtMutableMILWeightURLs() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("e5rtMutableMILWeightURLs"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetE5rtMutableMILWeightURLs(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setE5rtMutableMILWeightURLs:"), value)
}
func (m MLModelConfiguration) EnableTestVectorMode() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("enableTestVectorMode"))
	return rv
}
func (m MLModelConfiguration) SetEnableTestVectorMode(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEnableTestVectorMode:"), value)
}
func (m MLModelConfiguration) ExperimentalMLE5BNNSGraphBackendUsage() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("experimentalMLE5BNNSGraphBackendUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5BNNSGraphBackendUsage(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setExperimentalMLE5BNNSGraphBackendUsage:"), value)
}
func (m MLModelConfiguration) ExperimentalMLE5BNNSGraphBackendUsageMultiSegment() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("experimentalMLE5BNNSGraphBackendUsageMultiSegment"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5BNNSGraphBackendUsageMultiSegment(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setExperimentalMLE5BNNSGraphBackendUsageMultiSegment:"), value)
}
func (m MLModelConfiguration) ExperimentalMLE5EngineUsage() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("experimentalMLE5EngineUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLE5EngineUsage(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setExperimentalMLE5EngineUsage:"), value)
}
func (m MLModelConfiguration) ExperimentalMLProgramEncryptedCacheUsage() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("experimentalMLProgramEncryptedCacheUsage"))
	return rv
}
func (m MLModelConfiguration) SetExperimentalMLProgramEncryptedCacheUsage(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setExperimentalMLProgramEncryptedCacheUsage:"), value)
}
func (m MLModelConfiguration) ParentModelName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parentModelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetParentModelName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setParentModelName:"), objc.String(value))
}
func (m MLModelConfiguration) PredictionConcurrencyHint() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("predictionConcurrencyHint"))
	return rv
}
func (m MLModelConfiguration) SetPredictionConcurrencyHint(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPredictionConcurrencyHint:"), value)
}
func (m MLModelConfiguration) PreferredMTLDevice() metal.MTLDeviceObject {
	rv := objc.SendIfResponds[metal.MTLDeviceObject](m.ID, objc.Sel("preferredMTLDevice"))
	return metal.MTLDeviceObject(rv)
}
func (m MLModelConfiguration) SetPreferredMTLDevice(value metal.MTLDeviceObject) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPreferredMTLDevice:"), value)
}
func (m MLModelConfiguration) PreparesLazily() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("preparesLazily"))
	return rv
}
func (m MLModelConfiguration) SetPreparesLazily(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPreparesLazily:"), value)
}
func (m MLModelConfiguration) ProfilingOptions() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("profilingOptions"))
	return rv
}
func (m MLModelConfiguration) SetProfilingOptions(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProfilingOptions:"), value)
}
func (m MLModelConfiguration) RootModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("rootModelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetRootModelURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setRootModelURL:"), value)
}
func (m MLModelConfiguration) SerializesMILTextForDebugging() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("serializesMILTextForDebugging"))
	return rv
}
func (m MLModelConfiguration) SetSerializesMILTextForDebugging(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSerializesMILTextForDebugging:"), value)
}
func (m MLModelConfiguration) SpecializationUsesMPSGraphExecutable() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("specializationUsesMPSGraphExecutable"))
	return rv
}
func (m MLModelConfiguration) SetSpecializationUsesMPSGraphExecutable(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSpecializationUsesMPSGraphExecutable:"), value)
}
func (m MLModelConfiguration) TrainWithMLCompute() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("trainWithMLCompute"))
	return rv
}
func (m MLModelConfiguration) SetTrainWithMLCompute(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTrainWithMLCompute:"), value)
}
func (m MLModelConfiguration) UsePrecompiledE5Bundle() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("usePrecompiledE5Bundle"))
	return rv
}
func (m MLModelConfiguration) SetUsePrecompiledE5Bundle(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUsePrecompiledE5Bundle:"), value)
}
func (m MLModelConfiguration) UsePreloadedKey() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("usePreloadedKey"))
	return rv
}
func (m MLModelConfiguration) SetUsePreloadedKey(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUsePreloadedKey:"), value)
}
func (m MLModelConfiguration) UseWatchSPIForScribble() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("useWatchSPIForScribble"))
	return rv
}
func (m MLModelConfiguration) SetUseWatchSPIForScribble(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUseWatchSPIForScribble:"), value)
}

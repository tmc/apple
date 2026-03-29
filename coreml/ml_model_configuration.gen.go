// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
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

// The settings for creating or updating a machine learning model.
//
// # Overview
// 
// Use a model configuration to:
// 
// - Set or override model parameters. - Designate which device the model uses
// to make predictions, such as a GPU. - Restrict the model to use a specific
// computational device category, such as a CPU.
// 
// You typically use a model configuration instance to configure an [MLModel]
// instance as you create it with [ModelWithContentsOfURLConfigurationError]
// or create an [MLUpdateTask]. See [Personalizing a Model with On-Device
// Updates].
// 
// Configure your model parameters by setting values for each relevant
// [MLParameterKey] in the [MLModelConfiguration.Parameters] property.
//
// [Personalizing a Model with On-Device Updates]: https://developer.apple.com/documentation/CoreML/personalizing-a-model-with-on-device-updates
//
// # Configuring model parameters
//
//   - [MLModelConfiguration.FunctionName]: Function name that [MLModel] will use.
//   - [MLModelConfiguration.SetFunctionName]
//   - [MLModelConfiguration.ModelDisplayName]: A human readable name of a model for display purposes.
//   - [MLModelConfiguration.SetModelDisplayName]
//   - [MLModelConfiguration.Parameters]: A dictionary of configuration settings your app can override when loading a model.
//   - [MLModelConfiguration.SetParameters]
//
// # Configuring GPU usage
//
//   - [MLModelConfiguration.PreferredMetalDevice]: The metal device you prefer this model use to make predictions (inference) and update the model.
//   - [MLModelConfiguration.SetPreferredMetalDevice]
//   - [MLModelConfiguration.AllowLowPrecisionAccumulationOnGPU]: A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//   - [MLModelConfiguration.SetAllowLowPrecisionAccumulationOnGPU]
//
// # Allowing access to processing units
//
//   - [MLModelConfiguration.ComputeUnits]: The processing unit or units the model uses to make predictions.
//   - [MLModelConfiguration.SetComputeUnits]
//
// # Getting optimization hints
//
//   - [MLModelConfiguration.OptimizationHints]: A group of hints for CoreML to optimize
//   - [MLModelConfiguration.SetOptimizationHints]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration
type MLModelConfiguration struct {
	objectivec.Object
}

// MLModelConfigurationFromID constructs a [MLModelConfiguration] from an objc.ID.
//
// The settings for creating or updating a machine learning model.
func MLModelConfigurationFromID(id objc.ID) MLModelConfiguration {
	return MLModelConfiguration{objectivec.Object{ID: id}}
}
// NOTE: MLModelConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLModelConfiguration] class.
//
// # Configuring model parameters
//
//   - [IMLModelConfiguration.FunctionName]: Function name that [MLModel] will use.
//   - [IMLModelConfiguration.SetFunctionName]
//   - [IMLModelConfiguration.ModelDisplayName]: A human readable name of a model for display purposes.
//   - [IMLModelConfiguration.SetModelDisplayName]
//   - [IMLModelConfiguration.Parameters]: A dictionary of configuration settings your app can override when loading a model.
//   - [IMLModelConfiguration.SetParameters]
//
// # Configuring GPU usage
//
//   - [IMLModelConfiguration.PreferredMetalDevice]: The metal device you prefer this model use to make predictions (inference) and update the model.
//   - [IMLModelConfiguration.SetPreferredMetalDevice]
//   - [IMLModelConfiguration.AllowLowPrecisionAccumulationOnGPU]: A Boolean value that determines whether to allow low-precision accumulation on a GPU.
//   - [IMLModelConfiguration.SetAllowLowPrecisionAccumulationOnGPU]
//
// # Allowing access to processing units
//
//   - [IMLModelConfiguration.ComputeUnits]: The processing unit or units the model uses to make predictions.
//   - [IMLModelConfiguration.SetComputeUnits]
//
// # Getting optimization hints
//
//   - [IMLModelConfiguration.OptimizationHints]: A group of hints for CoreML to optimize
//   - [IMLModelConfiguration.SetOptimizationHints]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration
type IMLModelConfiguration interface {
	objectivec.IObject

	// Topic: Configuring model parameters

	// Function name that [MLModel] will use.
	FunctionName() string
	SetFunctionName(value string)
	// A human readable name of a model for display purposes.
	ModelDisplayName() string
	SetModelDisplayName(value string)
	// A dictionary of configuration settings your app can override when loading a model.
	Parameters() foundation.INSDictionary
	SetParameters(value foundation.INSDictionary)

	// Topic: Configuring GPU usage

	// The metal device you prefer this model use to make predictions (inference) and update the model.
	PreferredMetalDevice() metal.MTLDevice
	SetPreferredMetalDevice(value metal.MTLDevice)
	// A Boolean value that determines whether to allow low-precision accumulation on a GPU.
	AllowLowPrecisionAccumulationOnGPU() bool
	SetAllowLowPrecisionAccumulationOnGPU(value bool)

	// Topic: Allowing access to processing units

	// The processing unit or units the model uses to make predictions.
	ComputeUnits() MLComputeUnits
	SetComputeUnits(value MLComputeUnits)

	// Topic: Getting optimization hints

	// A group of hints for CoreML to optimize
	OptimizationHints() IMLOptimizationHints
	SetOptimizationHints(value IMLOptimizationHints)

	EncodeWithCoder(coder foundation.INSCoder)
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

func (m MLModelConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Function name that [MLModel] will use.
//
// # Discussion
// 
// Some model types (e.g. ML Program) supports multiple functions in a model
// asset, where each [MLModel] instance is associated with a particular
// function.
// 
// Use [MLModelAsset] to get the list of available functions. Use `nil` to use
// a default function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/functionName
func (m MLModelConfiguration) FunctionName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetFunctionName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionName:"), objc.String(value))
}
// A human readable name of a model for display purposes.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/modelDisplayName
func (m MLModelConfiguration) ModelDisplayName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDisplayName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelConfiguration) SetModelDisplayName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDisplayName:"), objc.String(value))
}
// A dictionary of configuration settings your app can override when loading a
// model.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/parameters
func (m MLModelConfiguration) Parameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameters:"), value)
}
// The metal device you prefer this model use to make predictions (inference)
// and update the model.
//
// # Discussion
// 
// If [PreferredMetalDevice] is `nil`, the default value, Core ML chooses a
// metal device for you.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/preferredMetalDevice
func (m MLModelConfiguration) PreferredMetalDevice() metal.MTLDevice {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("preferredMetalDevice"))
	return metal.MTLDeviceObjectFromID(rv)
}
func (m MLModelConfiguration) SetPreferredMetalDevice(value metal.MTLDevice) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMetalDevice:"), value)
}
// A Boolean value that determines whether to allow low-precision accumulation
// on a GPU.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/allowLowPrecisionAccumulationOnGPU
func (m MLModelConfiguration) AllowLowPrecisionAccumulationOnGPU() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowLowPrecisionAccumulationOnGPU"))
	return rv
}
func (m MLModelConfiguration) SetAllowLowPrecisionAccumulationOnGPU(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowLowPrecisionAccumulationOnGPU:"), value)
}
// The processing unit or units the model uses to make predictions.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelConfiguration/computeUnits
func (m MLModelConfiguration) ComputeUnits() MLComputeUnits {
	rv := objc.Send[MLComputeUnits](m.ID, objc.Sel("computeUnits"))
	return MLComputeUnits(rv)
}
func (m MLModelConfiguration) SetComputeUnits(value MLComputeUnits) {
	objc.Send[struct{}](m.ID, objc.Sel("setComputeUnits:"), value)
}
// A group of hints for CoreML to optimize
//
// See: https://developer.apple.com/documentation/coreml/mlmodelconfiguration/optimizationhints-1oq0g
func (m MLModelConfiguration) OptimizationHints() IMLOptimizationHints {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("optimizationHints"))
	return MLOptimizationHintsFromID(objc.ID(rv))
}
func (m MLModelConfiguration) SetOptimizationHints(value IMLOptimizationHints) {
	objc.Send[struct{}](m.ID, objc.Sel("setOptimizationHints:"), value)
}


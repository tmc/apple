// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MERAWProcessingFloatParameter] class.
var (
	_MERAWProcessingFloatParameterClass     MERAWProcessingFloatParameterClass
	_MERAWProcessingFloatParameterClassOnce sync.Once
)

func getMERAWProcessingFloatParameterClass() MERAWProcessingFloatParameterClass {
	_MERAWProcessingFloatParameterClassOnce.Do(func() {
		_MERAWProcessingFloatParameterClass = MERAWProcessingFloatParameterClass{class: objc.GetClass("MERAWProcessingFloatParameter")}
	})
	return _MERAWProcessingFloatParameterClass
}

// GetMERAWProcessingFloatParameterClass returns the class object for MERAWProcessingFloatParameter.
func GetMERAWProcessingFloatParameterClass() MERAWProcessingFloatParameterClass {
	return getMERAWProcessingFloatParameterClass()
}

type MERAWProcessingFloatParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingFloatParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingFloatParameterClass) Alloc() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a floating-point parameter of a RAW processor.
//
// # Properties
//
//   - [MERAWProcessingFloatParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [MERAWProcessingFloatParameter.SetCurrentValue]
//   - [MERAWProcessingFloatParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [MERAWProcessingFloatParameter.MaximumValue]: The maximum value for this parameter.
//   - [MERAWProcessingFloatParameter.MinimumValue]: The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint
type MERAWProcessingFloatParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingFloatParameterFromID constructs a [MERAWProcessingFloatParameter] from an objc.ID.
//
// An object that describes a floating-point parameter of a RAW processor.
func MERAWProcessingFloatParameterFromID(id objc.ID) MERAWProcessingFloatParameter {
	return MERAWProcessingFloatParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingFloatParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingFloatParameter] class.
//
// # Properties
//
//   - [IMERAWProcessingFloatParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [IMERAWProcessingFloatParameter.SetCurrentValue]
//   - [IMERAWProcessingFloatParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [IMERAWProcessingFloatParameter.MaximumValue]: The maximum value for this parameter.
//   - [IMERAWProcessingFloatParameter.MinimumValue]: The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint
type IMERAWProcessingFloatParameter interface {
	IMERAWProcessingParameter

	// Topic: Properties

	// Get or set the current value for this parameter.
	CurrentValue() float32
	SetCurrentValue(value float32)
	// The initial value for this parameter as defined in the sequence metadata.
	InitialValue() float32
	// The maximum value for this parameter.
	MaximumValue() float32
	// The minimum value for this parameter.
	MinimumValue() float32
}

// Init initializes the instance.
func (m MERAWProcessingFloatParameter) Init() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingFloatParameter) Autorelease() MERAWProcessingFloatParameter {
	rv := objc.Send[MERAWProcessingFloatParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingFloatParameter creates a new MERAWProcessingFloatParameter instance.
func NewMERAWProcessingFloatParameter() MERAWProcessingFloatParameter {
	class := getMERAWProcessingFloatParameterClass()
	rv := objc.Send[MERAWProcessingFloatParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Get or set the current value for this parameter.
//
// # Discussion
//
// This property can be observed if appropriate in order to monitor changes to
// the set of [MERAWProcessingParameters] vended by the extension.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/currentValue
func (m MERAWProcessingFloatParameter) CurrentValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("currentValue"))
	return rv
}
func (m MERAWProcessingFloatParameter) SetCurrentValue(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentValue:"), value)
}

// The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/initialValue
func (m MERAWProcessingFloatParameter) InitialValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("initialValue"))
	return rv
}

// The maximum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/maximumValue
func (m MERAWProcessingFloatParameter) MaximumValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("maximumValue"))
	return rv
}

// The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/FloatingPoint/minimumValue
func (m MERAWProcessingFloatParameter) MinimumValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("minimumValue"))
	return rv
}

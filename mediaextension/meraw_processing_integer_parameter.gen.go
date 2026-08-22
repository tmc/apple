// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MERAWProcessingIntegerParameter] class.
var (
	_MERAWProcessingIntegerParameterClass     MERAWProcessingIntegerParameterClass
	_MERAWProcessingIntegerParameterClassOnce sync.Once
)

func getMERAWProcessingIntegerParameterClass() MERAWProcessingIntegerParameterClass {
	_MERAWProcessingIntegerParameterClassOnce.Do(func() {
		_MERAWProcessingIntegerParameterClass = MERAWProcessingIntegerParameterClass{class: objc.GetClass("MERAWProcessingIntegerParameter")}
	})
	return _MERAWProcessingIntegerParameterClass
}

// GetMERAWProcessingIntegerParameterClass returns the class object for MERAWProcessingIntegerParameter.
func GetMERAWProcessingIntegerParameterClass() MERAWProcessingIntegerParameterClass {
	return getMERAWProcessingIntegerParameterClass()
}

type MERAWProcessingIntegerParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingIntegerParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingIntegerParameterClass) Alloc() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an integer parameter of a RAW processor.
//
// # Properties
//
//   - [MERAWProcessingIntegerParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [MERAWProcessingIntegerParameter.SetCurrentValue]
//   - [MERAWProcessingIntegerParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [MERAWProcessingIntegerParameter.MaximumValue]: The maximum value for this parameter.
//   - [MERAWProcessingIntegerParameter.MinimumValue]: The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer
type MERAWProcessingIntegerParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingIntegerParameterFromID constructs a [MERAWProcessingIntegerParameter] from an objc.ID.
//
// An object that describes an integer parameter of a RAW processor.
func MERAWProcessingIntegerParameterFromID(id objc.ID) MERAWProcessingIntegerParameter {
	return MERAWProcessingIntegerParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingIntegerParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingIntegerParameter] class.
//
// # Properties
//
//   - [IMERAWProcessingIntegerParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [IMERAWProcessingIntegerParameter.SetCurrentValue]
//   - [IMERAWProcessingIntegerParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [IMERAWProcessingIntegerParameter.MaximumValue]: The maximum value for this parameter.
//   - [IMERAWProcessingIntegerParameter.MinimumValue]: The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer
type IMERAWProcessingIntegerParameter interface {
	IMERAWProcessingParameter

	// Topic: Properties

	// Get or set the current value for this parameter.
	CurrentValue() int
	SetCurrentValue(value int)
	// The initial value for this parameter as defined in the sequence metadata.
	InitialValue() int
	// The maximum value for this parameter.
	MaximumValue() int
	// The minimum value for this parameter.
	MinimumValue() int
}

// Init initializes the instance.
func (m MERAWProcessingIntegerParameter) Init() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingIntegerParameter) Autorelease() MERAWProcessingIntegerParameter {
	rv := objc.Send[MERAWProcessingIntegerParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingIntegerParameter creates a new MERAWProcessingIntegerParameter instance.
func NewMERAWProcessingIntegerParameter() MERAWProcessingIntegerParameter {
	class := getMERAWProcessingIntegerParameterClass()
	rv := objc.Send[MERAWProcessingIntegerParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Get or set the current value for this parameter.
//
// # Discussion
//
// This property can be observed if appropriate in order to monitor changes to
// the set of [MERAWProcessingParameters] vended by the extension.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/currentValue
func (m MERAWProcessingIntegerParameter) CurrentValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("currentValue"))
	return rv
}
func (m MERAWProcessingIntegerParameter) SetCurrentValue(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentValue:"), value)
}

// The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/initialValue
func (m MERAWProcessingIntegerParameter) InitialValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("initialValue"))
	return rv
}

// The maximum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/maximumValue
func (m MERAWProcessingIntegerParameter) MaximumValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("maximumValue"))
	return rv
}

// The minimum value for this parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Integer/minimumValue
func (m MERAWProcessingIntegerParameter) MinimumValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("minimumValue"))
	return rv
}

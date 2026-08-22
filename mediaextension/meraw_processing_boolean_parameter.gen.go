// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MERAWProcessingBooleanParameter] class.
var (
	_MERAWProcessingBooleanParameterClass     MERAWProcessingBooleanParameterClass
	_MERAWProcessingBooleanParameterClassOnce sync.Once
)

func getMERAWProcessingBooleanParameterClass() MERAWProcessingBooleanParameterClass {
	_MERAWProcessingBooleanParameterClassOnce.Do(func() {
		_MERAWProcessingBooleanParameterClass = MERAWProcessingBooleanParameterClass{class: objc.GetClass("MERAWProcessingBooleanParameter")}
	})
	return _MERAWProcessingBooleanParameterClass
}

// GetMERAWProcessingBooleanParameterClass returns the class object for MERAWProcessingBooleanParameter.
func GetMERAWProcessingBooleanParameterClass() MERAWProcessingBooleanParameterClass {
	return getMERAWProcessingBooleanParameterClass()
}

type MERAWProcessingBooleanParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingBooleanParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingBooleanParameterClass) Alloc() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a Boolean parameter of a RAW processor.
//
// # Properties
//
//   - [MERAWProcessingBooleanParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [MERAWProcessingBooleanParameter.SetCurrentValue]
//   - [MERAWProcessingBooleanParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean
type MERAWProcessingBooleanParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingBooleanParameterFromID constructs a [MERAWProcessingBooleanParameter] from an objc.ID.
//
// An object that describes a Boolean parameter of a RAW processor.
func MERAWProcessingBooleanParameterFromID(id objc.ID) MERAWProcessingBooleanParameter {
	return MERAWProcessingBooleanParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingBooleanParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingBooleanParameter] class.
//
// # Properties
//
//   - [IMERAWProcessingBooleanParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [IMERAWProcessingBooleanParameter.SetCurrentValue]
//   - [IMERAWProcessingBooleanParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean
type IMERAWProcessingBooleanParameter interface {
	IMERAWProcessingParameter

	// Topic: Properties

	// Get or set the current value for this parameter.
	CurrentValue() bool
	SetCurrentValue(value bool)
	// The initial value for this parameter as defined in the sequence metadata.
	InitialValue() bool
}

// Init initializes the instance.
func (m MERAWProcessingBooleanParameter) Init() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingBooleanParameter) Autorelease() MERAWProcessingBooleanParameter {
	rv := objc.Send[MERAWProcessingBooleanParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingBooleanParameter creates a new MERAWProcessingBooleanParameter instance.
func NewMERAWProcessingBooleanParameter() MERAWProcessingBooleanParameter {
	class := getMERAWProcessingBooleanParameterClass()
	rv := objc.Send[MERAWProcessingBooleanParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Get or set the current value for this parameter.
//
// # Discussion
//
// This property can be observed if appropriate in order to monitor changes to
// the set of [MERAWProcessingParameters] vended by the extension.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean/currentValue
func (m MERAWProcessingBooleanParameter) CurrentValue() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("currentValue"))
	return rv
}
func (m MERAWProcessingBooleanParameter) SetCurrentValue(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentValue:"), value)
}

// The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/Boolean/initialValue
func (m MERAWProcessingBooleanParameter) InitialValue() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("initialValue"))
	return rv
}

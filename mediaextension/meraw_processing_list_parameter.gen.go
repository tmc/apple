// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MERAWProcessingListParameter] class.
var (
	_MERAWProcessingListParameterClass     MERAWProcessingListParameterClass
	_MERAWProcessingListParameterClassOnce sync.Once
)

func getMERAWProcessingListParameterClass() MERAWProcessingListParameterClass {
	_MERAWProcessingListParameterClassOnce.Do(func() {
		_MERAWProcessingListParameterClass = MERAWProcessingListParameterClass{class: objc.GetClass("MERAWProcessingListParameter")}
	})
	return _MERAWProcessingListParameterClass
}

// GetMERAWProcessingListParameterClass returns the class object for MERAWProcessingListParameter.
func GetMERAWProcessingListParameterClass() MERAWProcessingListParameterClass {
	return getMERAWProcessingListParameterClass()
}

type MERAWProcessingListParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingListParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingListParameterClass) Alloc() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a list parameter of a RAW processor.
//
// # Properties
//
//   - [MERAWProcessingListParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [MERAWProcessingListParameter.SetCurrentValue]
//   - [MERAWProcessingListParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [MERAWProcessingListParameter.ListElements]: The ordered array of [MERAWProcessingListElementParameter] which make up this list.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List
type MERAWProcessingListParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingListParameterFromID constructs a [MERAWProcessingListParameter] from an objc.ID.
//
// An object that describes a list parameter of a RAW processor.
func MERAWProcessingListParameterFromID(id objc.ID) MERAWProcessingListParameter {
	return MERAWProcessingListParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingListParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingListParameter] class.
//
// # Properties
//
//   - [IMERAWProcessingListParameter.CurrentValue]: Get or set the current value for this parameter.
//   - [IMERAWProcessingListParameter.SetCurrentValue]
//   - [IMERAWProcessingListParameter.InitialValue]: The initial value for this parameter as defined in the sequence metadata.
//   - [IMERAWProcessingListParameter.ListElements]: The ordered array of [MERAWProcessingListElementParameter] which make up this list.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List
type IMERAWProcessingListParameter interface {
	IMERAWProcessingParameter

	// Topic: Properties

	// Get or set the current value for this parameter.
	CurrentValue() int
	SetCurrentValue(value int)
	// The initial value for this parameter as defined in the sequence metadata.
	InitialValue() int
	// The ordered array of [MERAWProcessingListElementParameter] which make up this list.
	ListElements() []MERAWProcessingListElementParameter
}

// Init initializes the instance.
func (m MERAWProcessingListParameter) Init() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingListParameter) Autorelease() MERAWProcessingListParameter {
	rv := objc.Send[MERAWProcessingListParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingListParameter creates a new MERAWProcessingListParameter instance.
func NewMERAWProcessingListParameter() MERAWProcessingListParameter {
	class := getMERAWProcessingListParameterClass()
	rv := objc.Send[MERAWProcessingListParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Get or set the current value for this parameter.
//
// # Discussion
//
// This property can be observed if appropriate in order to monitor changes to
// the set of [MERAWProcessingParameters] vended by the extension.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/currentValue
func (m MERAWProcessingListParameter) CurrentValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("currentValue"))
	return rv
}
func (m MERAWProcessingListParameter) SetCurrentValue(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentValue:"), value)
}

// The initial value for this parameter as defined in the sequence metadata.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/initialValue
func (m MERAWProcessingListParameter) InitialValue() int {
	rv := objc.Send[int](m.ID, objc.Sel("initialValue"))
	return rv
}

// The ordered array of [MERAWProcessingListElementParameter] which make up
// this list.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/List/listElements
func (m MERAWProcessingListParameter) ListElements() []MERAWProcessingListElementParameter {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("listElements"))
	return objc.ConvertSlice(rv, func(id objc.ID) MERAWProcessingListElementParameter {
		return MERAWProcessingListElementParameterFromID(id)
	})
}

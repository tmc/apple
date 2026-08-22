// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MERAWProcessingSubGroupParameter] class.
var (
	_MERAWProcessingSubGroupParameterClass     MERAWProcessingSubGroupParameterClass
	_MERAWProcessingSubGroupParameterClassOnce sync.Once
)

func getMERAWProcessingSubGroupParameterClass() MERAWProcessingSubGroupParameterClass {
	_MERAWProcessingSubGroupParameterClassOnce.Do(func() {
		_MERAWProcessingSubGroupParameterClass = MERAWProcessingSubGroupParameterClass{class: objc.GetClass("MERAWProcessingSubGroupParameter")}
	})
	return _MERAWProcessingSubGroupParameterClass
}

// GetMERAWProcessingSubGroupParameterClass returns the class object for MERAWProcessingSubGroupParameter.
func GetMERAWProcessingSubGroupParameterClass() MERAWProcessingSubGroupParameterClass {
	return getMERAWProcessingSubGroupParameterClass()
}

type MERAWProcessingSubGroupParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingSubGroupParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingSubGroupParameterClass) Alloc() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a sub group parameter of a RAW processor.
//
// # Overview
//
// Sub groups are logical groupings of [MERAWProcessingParameter] objects that
// should be displayed together in an application user interface.
//
// # Creating a sub group parameter object
//
//   - [MERAWProcessingSubGroupParameter.InitWithNameDescriptionParameters]: Creates a sub group parameter object with the parameters value.
//
// # Properties
//
//   - [MERAWProcessingSubGroupParameter.SubGroupParameters]: The array of [MERAWProcessingParameter](<https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter>) objects in the sub group.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup
type MERAWProcessingSubGroupParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingSubGroupParameterFromID constructs a [MERAWProcessingSubGroupParameter] from an objc.ID.
//
// An object that describes a sub group parameter of a RAW processor.
func MERAWProcessingSubGroupParameterFromID(id objc.ID) MERAWProcessingSubGroupParameter {
	return MERAWProcessingSubGroupParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingSubGroupParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingSubGroupParameter] class.
//
// # Creating a sub group parameter object
//
//   - [IMERAWProcessingSubGroupParameter.InitWithNameDescriptionParameters]: Creates a sub group parameter object with the parameters value.
//
// # Properties
//
//   - [IMERAWProcessingSubGroupParameter.SubGroupParameters]: The array of [MERAWProcessingParameter](<https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter>) objects in the sub group.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup
type IMERAWProcessingSubGroupParameter interface {
	IMERAWProcessingParameter

	// Topic: Creating a sub group parameter object

	// Creates a sub group parameter object with the parameters value.
	InitWithNameDescriptionParameters(name string, description string, parameters []MERAWProcessingParameter) MERAWProcessingSubGroupParameter

	// Topic: Properties

	// The array of [MERAWProcessingParameter](<https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter>) objects in the sub group.
	SubGroupParameters() []MERAWProcessingParameter
}

// Init initializes the instance.
func (m MERAWProcessingSubGroupParameter) Init() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingSubGroupParameter) Autorelease() MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingSubGroupParameter creates a new MERAWProcessingSubGroupParameter instance.
func NewMERAWProcessingSubGroupParameter() MERAWProcessingSubGroupParameter {
	class := getMERAWProcessingSubGroupParameterClass()
	rv := objc.Send[MERAWProcessingSubGroupParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a sub group parameter object with the parameters value.
//
// name: A localized human-readable name for the parameter, suitable for displaying
// in application UI.
//
// description: A localized description of the parameter, suitable for displaying in a tool
// tip or similar explanatory UI.
//
// parameters: The array of [MERAWProcessingParameter] objects in the sub group.
//
// # Return Value
//
// An instance of [MERAWProcessingSubGroupParameter].
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup/init(name:description:parameters:)
func NewMERAWProcessingSubGroupParameterWithNameDescriptionParameters(name string, description string, parameters []MERAWProcessingParameter) MERAWProcessingSubGroupParameter {
	instance := getMERAWProcessingSubGroupParameterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:description:parameters:"), objc.String(name), objc.String(description), objectivec.IObjectSliceToNSArray(parameters))
	return MERAWProcessingSubGroupParameterFromID(rv)
}

// Creates a sub group parameter object with the parameters value.
//
// name: A localized human-readable name for the parameter, suitable for displaying
// in application UI.
//
// description: A localized description of the parameter, suitable for displaying in a tool
// tip or similar explanatory UI.
//
// parameters: The array of [MERAWProcessingParameter] objects in the sub group.
//
// # Return Value
//
// An instance of [MERAWProcessingSubGroupParameter].
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup/init(name:description:parameters:)
func (m MERAWProcessingSubGroupParameter) InitWithNameDescriptionParameters(name string, description string, parameters []MERAWProcessingParameter) MERAWProcessingSubGroupParameter {
	rv := objc.Send[MERAWProcessingSubGroupParameter](m.ID, objc.Sel("initWithName:description:parameters:"), objc.String(name), objc.String(description), objectivec.IObjectSliceToNSArray(parameters))
	return rv
}

// The array of [MERAWProcessingParameter] objects in the sub group.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/SubGroup/subGroupParameters
func (m MERAWProcessingSubGroupParameter) SubGroupParameters() []MERAWProcessingParameter {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("subGroupParameters"))
	return objc.ConvertSlice(rv, func(id objc.ID) MERAWProcessingParameter {
		return MERAWProcessingParameterFromID(id)
	})
}

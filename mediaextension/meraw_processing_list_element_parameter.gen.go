// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MERAWProcessingListElementParameter] class.
var (
	_MERAWProcessingListElementParameterClass     MERAWProcessingListElementParameterClass
	_MERAWProcessingListElementParameterClassOnce sync.Once
)

func getMERAWProcessingListElementParameterClass() MERAWProcessingListElementParameterClass {
	_MERAWProcessingListElementParameterClassOnce.Do(func() {
		_MERAWProcessingListElementParameterClass = MERAWProcessingListElementParameterClass{class: objc.GetClass("MERAWProcessingListElementParameter")}
	})
	return _MERAWProcessingListElementParameterClass
}

// GetMERAWProcessingListElementParameterClass returns the class object for MERAWProcessingListElementParameter.
func GetMERAWProcessingListElementParameterClass() MERAWProcessingListElementParameterClass {
	return getMERAWProcessingListElementParameterClass()
}

type MERAWProcessingListElementParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingListElementParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingListElementParameterClass) Alloc() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a list element parameter of a RAW processor.
//
// # Overview
//
// The [MERAWProcessingParameterListElement] protocol provides an interface
// for [VideoToolbox] to query descriptions of the different elements in a
// parameter list for a list element in a [MERAWProcessingParameter]. A
// distinct [MERAWProcessingParameterListElement] is created for each list
// element.
//
// # Creating a list element parameter object
//
//   - [MERAWProcessingListElementParameter.InitWithNameDescriptionElementID]: Creates a list element parameter object with the element id value.
//
// # Properties
//
//   - [MERAWProcessingListElementParameter.ListElementID]: A unique number in the list which represents this list option.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement
type MERAWProcessingListElementParameter struct {
	MERAWProcessingParameter
}

// MERAWProcessingListElementParameterFromID constructs a [MERAWProcessingListElementParameter] from an objc.ID.
//
// An object that describes a list element parameter of a RAW processor.
func MERAWProcessingListElementParameterFromID(id objc.ID) MERAWProcessingListElementParameter {
	return MERAWProcessingListElementParameter{MERAWProcessingParameter: MERAWProcessingParameterFromID(id)}
}

// NOTE: MERAWProcessingListElementParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingListElementParameter] class.
//
// # Creating a list element parameter object
//
//   - [IMERAWProcessingListElementParameter.InitWithNameDescriptionElementID]: Creates a list element parameter object with the element id value.
//
// # Properties
//
//   - [IMERAWProcessingListElementParameter.ListElementID]: A unique number in the list which represents this list option.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement
type IMERAWProcessingListElementParameter interface {
	IMERAWProcessingParameter

	// Topic: Creating a list element parameter object

	// Creates a list element parameter object with the element id value.
	InitWithNameDescriptionElementID(name string, description string, elementID int) MERAWProcessingListElementParameter

	// Topic: Properties

	// A unique number in the list which represents this list option.
	ListElementID() int
}

// Init initializes the instance.
func (m MERAWProcessingListElementParameter) Init() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingListElementParameter) Autorelease() MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingListElementParameter creates a new MERAWProcessingListElementParameter instance.
func NewMERAWProcessingListElementParameter() MERAWProcessingListElementParameter {
	class := getMERAWProcessingListElementParameterClass()
	rv := objc.Send[MERAWProcessingListElementParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a list element parameter object with the element id value.
//
// name: A localized human-readable name for the parameter, suitable for displaying
// in application UI.
//
// description: A localized description of the parameter, suitable for displaying in a tool
// tip or similar explanatory UI.
//
// elementID: A unique number in the list which represents this list option.
//
// # Return Value
//
// An instance of [MERAWProcessingListElementParameter].
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement/init(name:description:elementID:)
func NewMERAWProcessingListElementParameterWithNameDescriptionElementID(name string, description string, elementID int) MERAWProcessingListElementParameter {
	instance := getMERAWProcessingListElementParameterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:description:elementID:"), objc.String(name), objc.String(description), elementID)
	return MERAWProcessingListElementParameterFromID(rv)
}

// Creates a list element parameter object with the element id value.
//
// name: A localized human-readable name for the parameter, suitable for displaying
// in application UI.
//
// description: A localized description of the parameter, suitable for displaying in a tool
// tip or similar explanatory UI.
//
// elementID: A unique number in the list which represents this list option.
//
// # Return Value
//
// An instance of [MERAWProcessingListElementParameter].
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement/init(name:description:elementID:)
func (m MERAWProcessingListElementParameter) InitWithNameDescriptionElementID(name string, description string, elementID int) MERAWProcessingListElementParameter {
	rv := objc.Send[MERAWProcessingListElementParameter](m.ID, objc.Sel("initWithName:description:elementID:"), objc.String(name), objc.String(description), elementID)
	return rv
}

// A unique number in the list which represents this list option.
//
// # Discussion
//
// The set of elements in the list may change depending on other configuration
// parameters, so while the index of an element in this list may change, this
// ID never changes and is used to report list element selection.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/ListElement/listElementID
func (m MERAWProcessingListElementParameter) ListElementID() int {
	rv := objc.Send[int](m.ID, objc.Sel("listElementID"))
	return rv
}

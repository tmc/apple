// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MERAWProcessingParameter] class.
var (
	_MERAWProcessingParameterClass     MERAWProcessingParameterClass
	_MERAWProcessingParameterClassOnce sync.Once
)

func getMERAWProcessingParameterClass() MERAWProcessingParameterClass {
	_MERAWProcessingParameterClassOnce.Do(func() {
		_MERAWProcessingParameterClass = MERAWProcessingParameterClass{class: objc.GetClass("MERAWProcessingParameter")}
	})
	return _MERAWProcessingParameterClass
}

// GetMERAWProcessingParameterClass returns the class object for MERAWProcessingParameter.
func GetMERAWProcessingParameterClass() MERAWProcessingParameterClass {
	return getMERAWProcessingParameterClass()
}

type MERAWProcessingParameterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessingParameterClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessingParameterClass) Alloc() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object for the RAW processor to describe each processing parameter the
// processor exposes.
//
// # Discussion
//
// This protocol provides an interface for Video Toolbox to query descriptions
// of the different parameters that can be used to influence RAW processor
// operation. A distinct [MERAWProcessingParameter] is created for each
// parameter supported by the RAW processor, and the set of supported
// parameters is returned by the [ProcessingParameters] interface.
//
// # Inspecting a processing parameter
//
//   - [MERAWProcessingParameter.Enabled]: A Boolean value that indicates whether the extension enables the parameter.
//   - [MERAWProcessingParameter.SetEnabled]
//   - [MERAWProcessingParameter.Key]: A unique key string identifying the parameter.
//   - [MERAWProcessingParameter.LongDescription]: A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//   - [MERAWProcessingParameter.Name]: A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter
type MERAWProcessingParameter struct {
	objectivec.Object
}

// MERAWProcessingParameterFromID constructs a [MERAWProcessingParameter] from an objc.ID.
//
// An object for the RAW processor to describe each processing parameter the
// processor exposes.
func MERAWProcessingParameterFromID(id objc.ID) MERAWProcessingParameter {
	return MERAWProcessingParameter{objectivec.Object{ID: id}}
}

// NOTE: MERAWProcessingParameter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessingParameter] class.
//
// # Inspecting a processing parameter
//
//   - [IMERAWProcessingParameter.Enabled]: A Boolean value that indicates whether the extension enables the parameter.
//   - [IMERAWProcessingParameter.SetEnabled]
//   - [IMERAWProcessingParameter.Key]: A unique key string identifying the parameter.
//   - [IMERAWProcessingParameter.LongDescription]: A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
//   - [IMERAWProcessingParameter.Name]: A localized human-readable name for the parameter, suitable for displaying in application UI.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter
type IMERAWProcessingParameter interface {
	objectivec.IObject

	// Topic: Inspecting a processing parameter

	// A Boolean value that indicates whether the extension enables the parameter.
	Enabled() bool
	SetEnabled(value bool)
	// A unique key string identifying the parameter.
	Key() string
	// A localized description of the parameter, suitable for displaying in a tool tip or similar explanatory UI.
	LongDescription() string
	// A localized human-readable name for the parameter, suitable for displaying in application UI.
	Name() string
}

// Init initializes the instance.
func (m MERAWProcessingParameter) Init() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessingParameter) Autorelease() MERAWProcessingParameter {
	rv := objc.Send[MERAWProcessingParameter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessingParameter creates a new MERAWProcessingParameter instance.
func NewMERAWProcessingParameter() MERAWProcessingParameter {
	class := getMERAWProcessingParameterClass()
	rv := objc.Send[MERAWProcessingParameter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the extension enables the parameter.
//
// # Discussion
//
// This parameter can only be modified by the extension. From the
// application-facing interface, [VTRAWProcessingSession], this is a read-only
// value which indicates whether the parameter should be grayed out and
// disabled in any UI being generated.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/enabled
func (m MERAWProcessingParameter) Enabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("enabled"))
	return rv
}
func (m MERAWProcessingParameter) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}

// A unique key string identifying the parameter.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/key
func (m MERAWProcessingParameter) Key() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}

// A localized description of the parameter, suitable for displaying in a tool
// tip or similar explanatory UI.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/longDescription
func (m MERAWProcessingParameter) LongDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("longDescription"))
	return foundation.NSStringFromID(rv).String()
}

// A localized human-readable name for the parameter, suitable for displaying
// in application UI.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessingParameter/name
func (m MERAWProcessingParameter) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEFormatReaderInstantiationOptions] class.
var (
	_MEFormatReaderInstantiationOptionsClass     MEFormatReaderInstantiationOptionsClass
	_MEFormatReaderInstantiationOptionsClassOnce sync.Once
)

func getMEFormatReaderInstantiationOptionsClass() MEFormatReaderInstantiationOptionsClass {
	_MEFormatReaderInstantiationOptionsClassOnce.Do(func() {
		_MEFormatReaderInstantiationOptionsClass = MEFormatReaderInstantiationOptionsClass{class: objc.GetClass("MEFormatReaderInstantiationOptions")}
	})
	return _MEFormatReaderInstantiationOptionsClass
}

// GetMEFormatReaderInstantiationOptionsClass returns the class object for MEFormatReaderInstantiationOptions.
func GetMEFormatReaderInstantiationOptionsClass() MEFormatReaderInstantiationOptionsClass {
	return getMEFormatReaderInstantiationOptionsClass()
}

type MEFormatReaderInstantiationOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEFormatReaderInstantiationOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEFormatReaderInstantiationOptionsClass) Alloc() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains options to pass to a format reader extension.
//
// # Overview
//
// This object is mutable with options set through instance properties.
//
// # Inspecting format reader extension options
//
//   - [MEFormatReaderInstantiationOptions.AllowIncrementalFragmentParsing]: Enables support for parsing additional fragments.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions
type MEFormatReaderInstantiationOptions struct {
	objectivec.Object
}

// MEFormatReaderInstantiationOptionsFromID constructs a [MEFormatReaderInstantiationOptions] from an objc.ID.
//
// An object that contains options to pass to a format reader extension.
func MEFormatReaderInstantiationOptionsFromID(id objc.ID) MEFormatReaderInstantiationOptions {
	return MEFormatReaderInstantiationOptions{objectivec.Object{ID: id}}
}

// NOTE: MEFormatReaderInstantiationOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEFormatReaderInstantiationOptions] class.
//
// # Inspecting format reader extension options
//
//   - [IMEFormatReaderInstantiationOptions.AllowIncrementalFragmentParsing]: Enables support for parsing additional fragments.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions
type IMEFormatReaderInstantiationOptions interface {
	objectivec.IObject

	// Topic: Inspecting format reader extension options

	// Enables support for parsing additional fragments.
	AllowIncrementalFragmentParsing() bool
}

// Init initializes the instance.
func (m MEFormatReaderInstantiationOptions) Init() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEFormatReaderInstantiationOptions) Autorelease() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEFormatReaderInstantiationOptions creates a new MEFormatReaderInstantiationOptions instance.
func NewMEFormatReaderInstantiationOptions() MEFormatReaderInstantiationOptions {
	class := getMEFormatReaderInstantiationOptionsClass()
	rv := objc.Send[MEFormatReaderInstantiationOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Enables support for parsing additional fragments.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions/allowIncrementalFragmentParsing
func (m MEFormatReaderInstantiationOptions) AllowIncrementalFragmentParsing() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowIncrementalFragmentParsing"))
	return rv
}

// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A general interface for objects that produce filters.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterConstructor
type CIFilterConstructor interface {
	objectivec.IObject

	// Returns a filter object specified by name.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFilterConstructor/filter(withName:)
	FilterWithName(name string) CIFilter
}

// CIFilterConstructorObject wraps an existing Objective-C object that conforms to the CIFilterConstructor protocol.
type CIFilterConstructorObject struct {
	objectivec.Object
}
func (o CIFilterConstructorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFilterConstructorObjectFromID constructs a [CIFilterConstructorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFilterConstructorObjectFromID(id objc.ID) CIFilterConstructorObject {
	return CIFilterConstructorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a filter object specified by name.
//
// name: The name of the requested custom filter.
//
// # Return Value
// 
// A [CIFilter] object implementing the custom filter.
//
// # Discussion
// 
// Core Image calls this method when a filter is requested by name using the
// [CIFilter] class method [FilterWithName] method (or related methods). Your
// implementation of this method should provide a new instance of the
// [CIFilter] subclass for your custom filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterConstructor/filter(withName:)
func (o CIFilterConstructorObject) FilterWithName(name string) CIFilter {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filterWithName:"), objc.String(name))
	return CIFilterFromID(rv)
	}


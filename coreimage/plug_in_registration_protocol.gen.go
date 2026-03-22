// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface for loading Core Image image units.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugInRegistration
type CIPlugInRegistration interface {
	objectivec.IObject

	// Loads and initializes an image unit, performing custom tasks as needed.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPlugInRegistration/load(_:)
	Load(host unsafe.Pointer) bool
}

// CIPlugInRegistrationObject wraps an existing Objective-C object that conforms to the CIPlugInRegistration protocol.
type CIPlugInRegistrationObject struct {
	objectivec.Object
}
func (o CIPlugInRegistrationObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPlugInRegistrationObjectFromID constructs a [CIPlugInRegistrationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPlugInRegistrationObjectFromID(id objc.ID) CIPlugInRegistrationObject {
	return CIPlugInRegistrationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Loads and initializes an image unit, performing custom tasks as needed.
//
// host: Reserved for future use.
//
// # Return Value
// 
// Returns `true` if the image unit is successfully initialized
//
// # Discussion
// 
// The `load` method is called once by the host to initialize the image unit
// when the first filter in the image unit is instantiated. The method
// provides the image unit with an opportunity to perform custom
// initialization, such as a registration check.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPlugInRegistration/load(_:)
func (o CIPlugInRegistrationObject) Load(host unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("load:"), host)
	return rv
	}


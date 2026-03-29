// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioConsolePortArray] class.
var (
	_VZVirtioConsolePortArrayClass     VZVirtioConsolePortArrayClass
	_VZVirtioConsolePortArrayClassOnce sync.Once
)

func getVZVirtioConsolePortArrayClass() VZVirtioConsolePortArrayClass {
	_VZVirtioConsolePortArrayClassOnce.Do(func() {
		_VZVirtioConsolePortArrayClass = VZVirtioConsolePortArrayClass{class: objc.GetClass("VZVirtioConsolePortArray")}
	})
	return _VZVirtioConsolePortArrayClass
}

// GetVZVirtioConsolePortArrayClass returns the class object for VZVirtioConsolePortArray.
func GetVZVirtioConsolePortArrayClass() VZVirtioConsolePortArrayClass {
	return getVZVirtioConsolePortArrayClass()
}

type VZVirtioConsolePortArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsolePortArrayClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsolePortArrayClass) Alloc() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a collection of Virtio console ports.
//
// # Determining the number of ports
//
//   - [VZVirtioConsolePortArray.MaximumPortCount]: An unsigned integer that represents the maximum number of ports allocated by this device.
//
// # Accessing a specific port
//
//   - [VZVirtioConsolePortArray.ObjectAtIndexedSubscript]: Returns the Virtio console port at the specified index.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray
type VZVirtioConsolePortArray struct {
	objectivec.Object
}

// VZVirtioConsolePortArrayFromID constructs a [VZVirtioConsolePortArray] from an objc.ID.
//
// A class that represents a collection of Virtio console ports.
func VZVirtioConsolePortArrayFromID(id objc.ID) VZVirtioConsolePortArray {
	return VZVirtioConsolePortArray{objectivec.Object{ID: id}}
}
// NOTE: VZVirtioConsolePortArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioConsolePortArray] class.
//
// # Determining the number of ports
//
//   - [IVZVirtioConsolePortArray.MaximumPortCount]: An unsigned integer that represents the maximum number of ports allocated by this device.
//
// # Accessing a specific port
//
//   - [IVZVirtioConsolePortArray.ObjectAtIndexedSubscript]: Returns the Virtio console port at the specified index.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray
type IVZVirtioConsolePortArray interface {
	objectivec.IObject

	// Topic: Determining the number of ports

	// An unsigned integer that represents the maximum number of ports allocated by this device.
	MaximumPortCount() uint32

	// Topic: Accessing a specific port

	// Returns the Virtio console port at the specified index.
	ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePort
}

// Init initializes the instance.
func (v VZVirtioConsolePortArray) Init() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsolePortArray) Autorelease() VZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortArray creates a new VZVirtioConsolePortArray instance.
func NewVZVirtioConsolePortArray() VZVirtioConsolePortArray {
	class := getVZVirtioConsolePortArrayClass()
	rv := objc.Send[VZVirtioConsolePortArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the Virtio console port at the specified index.
//
// portIndex: The index of the port to return, if present.
//
// # Return Value
// 
// A [VZVirtioConsolePort] port, or `nil` if the index is outside the bounds
// of the array.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/subscript(_:)
func (v VZVirtioConsolePortArray) ObjectAtIndexedSubscript(portIndex uint) IVZVirtioConsolePort {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("objectAtIndexedSubscript:"), portIndex)
	return VZVirtioConsolePortFromID(rv)
}

// An unsigned integer that represents the maximum number of ports allocated
// by this device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortArray/maximumPortCount
func (v VZVirtioConsolePortArray) MaximumPortCount() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("maximumPortCount"))
	return rv
}


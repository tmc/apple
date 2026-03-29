// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDevice] class.
var (
	_VZStorageDeviceClass     VZStorageDeviceClass
	_VZStorageDeviceClassOnce sync.Once
)

func getVZStorageDeviceClass() VZStorageDeviceClass {
	_VZStorageDeviceClassOnce.Do(func() {
		_VZStorageDeviceClass = VZStorageDeviceClass{class: objc.GetClass("VZStorageDevice")}
	})
	return _VZStorageDeviceClass
}

// GetVZStorageDeviceClass returns the class object for VZStorageDevice.
func GetVZStorageDeviceClass() VZStorageDeviceClass {
	return getVZStorageDeviceClass()
}

type VZStorageDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZStorageDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceClass) Alloc() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents a storage device in a VM.
//
// # Overview
// 
// Don’t create a [VZStorageDevice] directly. Use one of its subclasses,
// such as [VZUSBMassStorageDevice], instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type VZStorageDevice struct {
	objectivec.Object
}

// VZStorageDeviceFromID constructs a [VZStorageDevice] from an objc.ID.
//
// A class that represents a storage device in a VM.
func VZStorageDeviceFromID(id objc.ID) VZStorageDevice {
	return VZStorageDevice{objectivec.Object{ID: id}}
}
// NOTE: VZStorageDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZStorageDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDevice
type IVZStorageDevice interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s VZStorageDevice) Init() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDevice) Autorelease() VZStorageDevice {
	rv := objc.Send[VZStorageDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDevice creates a new VZStorageDevice instance.
func NewVZStorageDevice() VZStorageDevice {
	class := getVZStorageDeviceClass()
	rv := objc.Send[VZStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAudioDevice] class.
var (
	_VZAudioDeviceClass     VZAudioDeviceClass
	_VZAudioDeviceClassOnce sync.Once
)

func getVZAudioDeviceClass() VZAudioDeviceClass {
	_VZAudioDeviceClassOnce.Do(func() {
		_VZAudioDeviceClass = VZAudioDeviceClass{class: objc.GetClass("_VZAudioDevice")}
	})
	return _VZAudioDeviceClass
}

// GetVZAudioDeviceClass returns the class object for _VZAudioDevice.
func GetVZAudioDeviceClass() VZAudioDeviceClass {
	return getVZAudioDeviceClass()
}

type VZAudioDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAudioDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAudioDeviceClass) Alloc() VZAudioDevice {
	rv := objc.Send[VZAudioDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZAudioDevice._initWithVirtualMachineAudioDeviceIndex]
// See: https://developer.apple.com/documentation/Virtualization/_VZAudioDevice
type VZAudioDevice struct {
	objectivec.Object
}

// VZAudioDeviceFromID constructs a [VZAudioDevice] from an objc.ID.
func VZAudioDeviceFromID(id objc.ID) VZAudioDevice {
	return VZAudioDevice{objectivec.Object{ID: id}}
}
// Ensure VZAudioDevice implements IVZAudioDevice.
var _ IVZAudioDevice = VZAudioDevice{}

// An interface definition for the [VZAudioDevice] class.
//
// # Methods
//
//   - [IVZAudioDevice._initWithVirtualMachineAudioDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZAudioDevice
type IVZAudioDevice interface {
	objectivec.IObject

	// Topic: Methods

	_initWithVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
}

// Init initializes the instance.
func (v VZAudioDevice) Init() VZAudioDevice {
	rv := objc.Send[VZAudioDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAudioDevice) Autorelease() VZAudioDevice {
	rv := objc.Send[VZAudioDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDevice creates a new VZAudioDevice instance.
func NewVZAudioDevice() VZAudioDevice {
	class := getVZAudioDeviceClass()
	rv := objc.Send[VZAudioDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZAudioDevice/_initWithVirtualMachine:audioDeviceIndex:
func (v VZAudioDevice) _initWithVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithVirtualMachine:audioDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineAudioDeviceIndex is an exported wrapper for the private method _initWithVirtualMachineAudioDeviceIndex.
func (v VZAudioDevice) InitWithVirtualMachineAudioDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	return v._initWithVirtualMachineAudioDeviceIndex(machine, index)
}


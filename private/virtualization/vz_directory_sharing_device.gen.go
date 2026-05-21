// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectorySharingDevice] class.
var (
	_VZDirectorySharingDeviceClass     VZDirectorySharingDeviceClass
	_VZDirectorySharingDeviceClassOnce sync.Once
)

func getVZDirectorySharingDeviceClass() VZDirectorySharingDeviceClass {
	_VZDirectorySharingDeviceClassOnce.Do(func() {
		_VZDirectorySharingDeviceClass = VZDirectorySharingDeviceClass{class: objc.GetClass("VZDirectorySharingDevice")}
	})
	return _VZDirectorySharingDeviceClass
}

// GetVZDirectorySharingDeviceClass returns the class object for VZDirectorySharingDevice.
func GetVZDirectorySharingDeviceClass() VZDirectorySharingDeviceClass {
	return getVZDirectorySharingDeviceClass()
}

type VZDirectorySharingDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectorySharingDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectorySharingDeviceClass) Alloc() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDirectorySharingDevice._initWithVirtualMachineDirectorySharingDeviceIndex]
type VZDirectorySharingDevice struct {
	objectivec.Object
}

// VZDirectorySharingDeviceFromID constructs a [VZDirectorySharingDevice] from an objc.ID.
func VZDirectorySharingDeviceFromID(id objc.ID) VZDirectorySharingDevice {
	return VZDirectorySharingDevice{objectivec.Object{ID: id}}
}

// Ensure VZDirectorySharingDevice implements IVZDirectorySharingDevice.
var _ IVZDirectorySharingDevice = VZDirectorySharingDevice{}

// An interface definition for the [VZDirectorySharingDevice] class.
//
// # Methods
//
//   - [IVZDirectorySharingDevice._initWithVirtualMachineDirectorySharingDeviceIndex]
type IVZDirectorySharingDevice interface {
	objectivec.IObject

	// Topic: Methods

	_initWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
}

// Init initializes the instance.
func (v VZDirectorySharingDevice) Init() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDirectorySharingDevice) Autorelease() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDevice creates a new VZDirectorySharingDevice instance.
func NewVZDirectorySharingDevice() VZDirectorySharingDevice {
	class := getVZDirectorySharingDeviceClass()
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZDirectorySharingDevice) _initWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithVirtualMachine:directorySharingDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineDirectorySharingDeviceIndex is an exported wrapper for the private method _initWithVirtualMachineDirectorySharingDeviceIndex.
func (v VZDirectorySharingDevice) InitWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:directorySharingDeviceIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithVirtualMachine:directorySharingDeviceIndex:"}
		return nil, err
	}
	return v._initWithVirtualMachineDirectorySharingDeviceIndex(machine, index), nil
}

// CanInitWithVirtualMachineDirectorySharingDeviceIndex reports whether the receiver responds to the private selector _initWithVirtualMachine:directorySharingDeviceIndex:.
func (v VZDirectorySharingDevice) CanInitWithVirtualMachineDirectorySharingDeviceIndex() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:directorySharingDeviceIndex:"))
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioTraditionalMemoryBalloonDevice] class.
var (
	_VZVirtioTraditionalMemoryBalloonDeviceClass     VZVirtioTraditionalMemoryBalloonDeviceClass
	_VZVirtioTraditionalMemoryBalloonDeviceClassOnce sync.Once
)

func getVZVirtioTraditionalMemoryBalloonDeviceClass() VZVirtioTraditionalMemoryBalloonDeviceClass {
	_VZVirtioTraditionalMemoryBalloonDeviceClassOnce.Do(func() {
		_VZVirtioTraditionalMemoryBalloonDeviceClass = VZVirtioTraditionalMemoryBalloonDeviceClass{class: objc.GetClass("VZVirtioTraditionalMemoryBalloonDevice")}
	})
	return _VZVirtioTraditionalMemoryBalloonDeviceClass
}

// GetVZVirtioTraditionalMemoryBalloonDeviceClass returns the class object for VZVirtioTraditionalMemoryBalloonDevice.
func GetVZVirtioTraditionalMemoryBalloonDeviceClass() VZVirtioTraditionalMemoryBalloonDeviceClass {
	return getVZVirtioTraditionalMemoryBalloonDeviceClass()
}

type VZVirtioTraditionalMemoryBalloonDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioTraditionalMemoryBalloonDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioTraditionalMemoryBalloonDeviceClass) Alloc() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioTraditionalMemoryBalloonDevice._maxTargetVirtualMachineMemorySize]
type VZVirtioTraditionalMemoryBalloonDevice struct {
	VZMemoryBalloonDevice
}

// VZVirtioTraditionalMemoryBalloonDeviceFromID constructs a [VZVirtioTraditionalMemoryBalloonDevice] from an objc.ID.
func VZVirtioTraditionalMemoryBalloonDeviceFromID(id objc.ID) VZVirtioTraditionalMemoryBalloonDevice {
	return VZVirtioTraditionalMemoryBalloonDevice{VZMemoryBalloonDevice: VZMemoryBalloonDeviceFromID(id)}
}

// Ensure VZVirtioTraditionalMemoryBalloonDevice implements IVZVirtioTraditionalMemoryBalloonDevice.
var _ IVZVirtioTraditionalMemoryBalloonDevice = VZVirtioTraditionalMemoryBalloonDevice{}

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDevice] class.
//
// # Methods
//
//   - [IVZVirtioTraditionalMemoryBalloonDevice._maxTargetVirtualMachineMemorySize]
type IVZVirtioTraditionalMemoryBalloonDevice interface {
	IVZMemoryBalloonDevice

	// Topic: Methods

	_maxTargetVirtualMachineMemorySize() uint64
}

// Init initializes the instance.
func (v VZVirtioTraditionalMemoryBalloonDevice) Init() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioTraditionalMemoryBalloonDevice) Autorelease() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioTraditionalMemoryBalloonDevice creates a new VZVirtioTraditionalMemoryBalloonDevice instance.
func NewVZVirtioTraditionalMemoryBalloonDevice() VZVirtioTraditionalMemoryBalloonDevice {
	class := getVZVirtioTraditionalMemoryBalloonDeviceClass()
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioTraditionalMemoryBalloonDevice) _maxTargetVirtualMachineMemorySize() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("_maxTargetVirtualMachineMemorySize"))
	return rv
}

// CanMaxTargetVirtualMachineMemorySize reports whether the receiver responds to the private selector _maxTargetVirtualMachineMemorySize.
func (v VZVirtioTraditionalMemoryBalloonDevice) CanMaxTargetVirtualMachineMemorySize() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_maxTargetVirtualMachineMemorySize"))
}

// MaxTargetVirtualMachineMemorySize is an exported wrapper for the private property _maxTargetVirtualMachineMemorySize.
func (v VZVirtioTraditionalMemoryBalloonDevice) MaxTargetVirtualMachineMemorySize() (uint64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_maxTargetVirtualMachineMemorySize")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_maxTargetVirtualMachineMemorySize"}
	}
	return v._maxTargetVirtualMachineMemorySize(), nil
}

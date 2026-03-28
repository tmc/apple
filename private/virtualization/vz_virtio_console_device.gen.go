// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioConsoleDevice] class.
var (
	_VZVirtioConsoleDeviceClass     VZVirtioConsoleDeviceClass
	_VZVirtioConsoleDeviceClassOnce sync.Once
)

func getVZVirtioConsoleDeviceClass() VZVirtioConsoleDeviceClass {
	_VZVirtioConsoleDeviceClassOnce.Do(func() {
		_VZVirtioConsoleDeviceClass = VZVirtioConsoleDeviceClass{class: objc.GetClass("VZVirtioConsoleDevice")}
	})
	return _VZVirtioConsoleDeviceClass
}

// GetVZVirtioConsoleDeviceClass returns the class object for VZVirtioConsoleDevice.
func GetVZVirtioConsoleDeviceClass() VZVirtioConsoleDeviceClass {
	return getVZVirtioConsoleDeviceClass()
}

type VZVirtioConsoleDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsoleDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsoleDeviceClass) Alloc() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice
type VZVirtioConsoleDevice struct {
	VZConsoleDevice
}

// VZVirtioConsoleDeviceFromID constructs a [VZVirtioConsoleDevice] from an objc.ID.
func VZVirtioConsoleDeviceFromID(id objc.ID) VZVirtioConsoleDevice {
	return VZVirtioConsoleDevice{VZConsoleDevice: VZConsoleDeviceFromID(id)}
}
// Ensure VZVirtioConsoleDevice implements IVZVirtioConsoleDevice.
var _ IVZVirtioConsoleDevice = VZVirtioConsoleDevice{}

// An interface definition for the [VZVirtioConsoleDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice
type IVZVirtioConsoleDevice interface {
	IVZConsoleDevice
}

// Init initializes the instance.
func (v VZVirtioConsoleDevice) Init() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsoleDevice) Autorelease() VZVirtioConsoleDevice {
	rv := objc.Send[VZVirtioConsoleDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDevice creates a new VZVirtioConsoleDevice instance.
func NewVZVirtioConsoleDevice() VZVirtioConsoleDevice {
	class := getVZVirtioConsoleDeviceClass()
	rv := objc.Send[VZVirtioConsoleDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDevice/initWithVirtualMachine:consoleDeviceIndex:configuration:
func NewVirtioConsoleDeviceWithVirtualMachineConsoleDeviceIndexConfiguration(machine objectivec.IObject, index uint64, configuration objectivec.IObject) VZVirtioConsoleDevice {
	instance := getVZVirtioConsoleDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:consoleDeviceIndex:configuration:"), machine, index, configuration)
	return VZVirtioConsoleDeviceFromID(rv)
}


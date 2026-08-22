// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZConsoleDevice] class.
var (
	_VZConsoleDeviceClass     VZConsoleDeviceClass
	_VZConsoleDeviceClassOnce sync.Once
)

func getVZConsoleDeviceClass() VZConsoleDeviceClass {
	_VZConsoleDeviceClassOnce.Do(func() {
		_VZConsoleDeviceClass = VZConsoleDeviceClass{class: objc.GetClass("VZConsoleDevice")}
	})
	return _VZConsoleDeviceClass
}

// GetVZConsoleDeviceClass returns the class object for VZConsoleDevice.
func GetVZConsoleDeviceClass() VZConsoleDeviceClass {
	return getVZConsoleDeviceClass()
}

type VZConsoleDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZConsoleDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZConsoleDeviceClass) Alloc() VZConsoleDevice {
	rv := objc.SendIfResponds[VZConsoleDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZConsoleDevice.InitWithVirtualMachineConsoleDeviceIndexConfiguration]
type VZConsoleDevice struct {
	objectivec.Object
}

// VZConsoleDeviceFromID constructs a [VZConsoleDevice] from an objc.ID.
func VZConsoleDeviceFromID(id objc.ID) VZConsoleDevice {
	return VZConsoleDevice{objectivec.Object{ID: id}}
}

// Ensure VZConsoleDevice implements IVZConsoleDevice.
var _ IVZConsoleDevice = VZConsoleDevice{}

// An interface definition for the [VZConsoleDevice] class.
//
// # Methods
//
//   - [IVZConsoleDevice.InitWithVirtualMachineConsoleDeviceIndexConfiguration]
type IVZConsoleDevice interface {
	objectivec.IObject

	// Topic: Methods

	InitWithVirtualMachineConsoleDeviceIndexConfiguration(machine objectivec.IObject, index uint64, configuration objectivec.IObject) VZConsoleDevice
}

// Init initializes the instance.
func (v VZConsoleDevice) Init() VZConsoleDevice {
	rv := objc.SendIfResponds[VZConsoleDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZConsoleDevice) Autorelease() VZConsoleDevice {
	rv := objc.SendIfResponds[VZConsoleDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDevice creates a new VZConsoleDevice instance.
func NewVZConsoleDevice() VZConsoleDevice {
	class := getVZConsoleDeviceClass()
	rv := objc.SendIfResponds[VZConsoleDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewConsoleDeviceWithVirtualMachineConsoleDeviceIndexConfiguration(machine objectivec.IObject, index uint64, configuration objectivec.IObject) VZConsoleDevice {
	instance := getVZConsoleDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:consoleDeviceIndex:configuration:"), machine, index, configuration)
	return VZConsoleDeviceFromID(rv)
}

func (v VZConsoleDevice) InitWithVirtualMachineConsoleDeviceIndexConfiguration(machine objectivec.IObject, index uint64, configuration objectivec.IObject) VZConsoleDevice {
	rv := objc.SendIfResponds[VZConsoleDevice](v.ID, objc.Sel("initWithVirtualMachine:consoleDeviceIndex:configuration:"), machine, index, configuration)
	return rv
}

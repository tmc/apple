// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomMMIODevice] class.
var (
	_VZCustomMMIODeviceClass     VZCustomMMIODeviceClass
	_VZCustomMMIODeviceClassOnce sync.Once
)

func getVZCustomMMIODeviceClass() VZCustomMMIODeviceClass {
	_VZCustomMMIODeviceClassOnce.Do(func() {
		_VZCustomMMIODeviceClass = VZCustomMMIODeviceClass{class: objc.GetClass("_VZCustomMMIODevice")}
	})
	return _VZCustomMMIODeviceClass
}

// GetVZCustomMMIODeviceClass returns the class object for _VZCustomMMIODevice.
func GetVZCustomMMIODeviceClass() VZCustomMMIODeviceClass {
	return getVZCustomMMIODeviceClass()
}

type VZCustomMMIODeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomMMIODeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomMMIODeviceClass) Alloc() VZCustomMMIODevice {
	rv := objc.SendIfResponds[VZCustomMMIODevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomMMIODevice.Delegate]
//   - [VZCustomMMIODevice.SetDelegate]
//   - [VZCustomMMIODevice.DeviceQueue]
//   - [VZCustomMMIODevice.GuestMemoryAtPhysicalAddressLength]
//   - [VZCustomMMIODevice.GuestRAMRegions]
//   - [VZCustomMMIODevice.PulseIRQ]
//   - [VZCustomMMIODevice.SetIRQValue]
//   - [VZCustomMMIODevice.SharedInitializationWithDeviceQueueFromConfiguration]
type VZCustomMMIODevice struct {
	objectivec.Object
}

// VZCustomMMIODeviceFromID constructs a [VZCustomMMIODevice] from an objc.ID.
func VZCustomMMIODeviceFromID(id objc.ID) VZCustomMMIODevice {
	return VZCustomMMIODevice{objectivec.Object{ID: id}}
}

// Ensure VZCustomMMIODevice implements IVZCustomMMIODevice.
var _ IVZCustomMMIODevice = VZCustomMMIODevice{}

// An interface definition for the [VZCustomMMIODevice] class.
//
// # Methods
//
//   - [IVZCustomMMIODevice.Delegate]
//   - [IVZCustomMMIODevice.SetDelegate]
//   - [IVZCustomMMIODevice.DeviceQueue]
//   - [IVZCustomMMIODevice.GuestMemoryAtPhysicalAddressLength]
//   - [IVZCustomMMIODevice.GuestRAMRegions]
//   - [IVZCustomMMIODevice.PulseIRQ]
//   - [IVZCustomMMIODevice.SetIRQValue]
//   - [IVZCustomMMIODevice.SharedInitializationWithDeviceQueueFromConfiguration]
type IVZCustomMMIODevice interface {
	objectivec.IObject

	// Topic: Methods

	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DeviceQueue() objectivec.Object
	GuestMemoryAtPhysicalAddressLength(address uint64, length uint64) objectivec.IObject
	GuestRAMRegions() objectivec.IObject
	PulseIRQ(irq uint64)
	SetIRQValue(irq uint64, value bool)
	SharedInitializationWithDeviceQueueFromConfiguration(queue objectivec.IObject, configuration objectivec.IObject)
}

// Init initializes the instance.
func (v VZCustomMMIODevice) Init() VZCustomMMIODevice {
	rv := objc.SendIfResponds[VZCustomMMIODevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomMMIODevice) Autorelease() VZCustomMMIODevice {
	rv := objc.SendIfResponds[VZCustomMMIODevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomMMIODevice creates a new VZCustomMMIODevice instance.
func NewVZCustomMMIODevice() VZCustomMMIODevice {
	class := getVZCustomMMIODeviceClass()
	rv := objc.SendIfResponds[VZCustomMMIODevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCustomMMIODevice) GuestMemoryAtPhysicalAddressLength(address uint64, length uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("guestMemoryAtPhysicalAddress:length:"), address, length)
	return objectivec.Object{ID: rv}
}
func (v VZCustomMMIODevice) GuestRAMRegions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("guestRAMRegions"))
	return objectivec.Object{ID: rv}
}
func (v VZCustomMMIODevice) PulseIRQ(irq uint64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("pulseIRQ:"), irq)
}
func (v VZCustomMMIODevice) SetIRQValue(irq uint64, value bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setIRQ:value:"), irq, value)
}
func (v VZCustomMMIODevice) SharedInitializationWithDeviceQueueFromConfiguration(queue objectivec.IObject, configuration objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sharedInitializationWithDeviceQueue:fromConfiguration:"), queue, configuration)
}

func (v VZCustomMMIODevice) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("delegate"))
	return rv
}
func (v VZCustomMMIODevice) SetDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
func (v VZCustomMMIODevice) DeviceQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("deviceQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

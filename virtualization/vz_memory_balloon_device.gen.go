// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMemoryBalloonDevice] class.
var (
	_VZMemoryBalloonDeviceClass     VZMemoryBalloonDeviceClass
	_VZMemoryBalloonDeviceClassOnce sync.Once
)

func getVZMemoryBalloonDeviceClass() VZMemoryBalloonDeviceClass {
	_VZMemoryBalloonDeviceClassOnce.Do(func() {
		_VZMemoryBalloonDeviceClass = VZMemoryBalloonDeviceClass{class: objc.GetClass("VZMemoryBalloonDevice")}
	})
	return _VZMemoryBalloonDeviceClass
}

// GetVZMemoryBalloonDeviceClass returns the class object for VZMemoryBalloonDevice.
func GetVZMemoryBalloonDeviceClass() VZMemoryBalloonDeviceClass {
	return getVZMemoryBalloonDeviceClass()
}

type VZMemoryBalloonDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryBalloonDeviceClass) Alloc() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The common behavior for memory devices.
//
// # Overview
// 
// Don’t instantiate this class directly. To request a memory ballon device,
// add an appropriate configuration object to the [VZMemoryBalloonDevice.MemoryBalloonDevices]
// property of the [VZVirtualMachineConfiguration] object that you use to
// configure the virtual machine. In response, the system instantiates the
// subclass of [VZMemoryBalloonDevice] that matches your request. For example,
// if you supply a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration]
// object in your configuration, the system creates a
// [VZVirtioTraditionalMemoryBalloonDevice] object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type VZMemoryBalloonDevice struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceFromID constructs a [VZMemoryBalloonDevice] from an objc.ID.
//
// The common behavior for memory devices.
func VZMemoryBalloonDeviceFromID(id objc.ID) VZMemoryBalloonDevice {
	return VZMemoryBalloonDevice{objectivec.Object{id}}
}
// NOTE: VZMemoryBalloonDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZMemoryBalloonDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type IVZMemoryBalloonDevice interface {
	objectivec.IObject

	// An array that you configure with a memory balloon device, used to update the memory in the VM.
	MemoryBalloonDevices() IVZMemoryBalloonDeviceConfiguration
	SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration)
}





// Init initializes the instance.
func (m VZMemoryBalloonDevice) Init() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMemoryBalloonDevice) Autorelease() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDevice creates a new VZMemoryBalloonDevice instance.
func NewVZMemoryBalloonDevice() VZMemoryBalloonDevice {
	class := getVZMemoryBalloonDeviceClass()
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// An array that you configure with a memory balloon device, used to update
// the memory in the VM.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/memoryballoondevices
func (m VZMemoryBalloonDevice) MemoryBalloonDevices() IVZMemoryBalloonDeviceConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("memoryBalloonDevices"))
	return VZMemoryBalloonDeviceConfigurationFromID(objc.ID(rv))
}
func (m VZMemoryBalloonDevice) SetMemoryBalloonDevices(value IVZMemoryBalloonDeviceConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}
























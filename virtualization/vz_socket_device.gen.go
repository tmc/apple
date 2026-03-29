// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSocketDevice] class.
var (
	_VZSocketDeviceClass     VZSocketDeviceClass
	_VZSocketDeviceClassOnce sync.Once
)

func getVZSocketDeviceClass() VZSocketDeviceClass {
	_VZSocketDeviceClassOnce.Do(func() {
		_VZSocketDeviceClass = VZSocketDeviceClass{class: objc.GetClass("VZSocketDevice")}
	})
	return _VZSocketDeviceClass
}

// GetVZSocketDeviceClass returns the class object for VZSocketDevice.
func GetVZSocketDeviceClass() VZSocketDeviceClass {
	return getVZSocketDeviceClass()
}

type VZSocketDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSocketDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSocketDeviceClass) Alloc() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common behavior of socket devices.
//
// # Overview
// 
// Don’t create or use a [VZSocketDevice] object directly. If your virtual
// machine’s configuration includes a [VZVirtioSocketDeviceConfiguration]
// object, the virtual machine returns a [VZVirtioSocketDevice] object in its
// [VZSocketDevice.SocketDevices] property. Use that object to configure the port-based
// communications for your virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDevice
type VZSocketDevice struct {
	objectivec.Object
}

// VZSocketDeviceFromID constructs a [VZSocketDevice] from an objc.ID.
//
// The common behavior of socket devices.
func VZSocketDeviceFromID(id objc.ID) VZSocketDevice {
	return VZSocketDevice{objectivec.Object{ID: id}}
}
// NOTE: VZSocketDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZSocketDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDevice
type IVZSocketDevice interface {
	objectivec.IObject

	// The array of socket devices that the VM configures for use ports in the guest VM.
	SocketDevices() IVZSocketDevice
	SetSocketDevices(value IVZSocketDevice)
}

// Init initializes the instance.
func (s VZSocketDevice) Init() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSocketDevice) Autorelease() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDevice creates a new VZSocketDevice instance.
func NewVZSocketDevice() VZSocketDevice {
	class := getVZSocketDeviceClass()
	rv := objc.Send[VZSocketDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of socket devices that the VM configures for use ports in the
// guest VM.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/socketdevices
func (s VZSocketDevice) SocketDevices() IVZSocketDevice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("socketDevices"))
	return VZSocketDeviceFromID(objc.ID(rv))
}
func (s VZSocketDevice) SetSocketDevices(value IVZSocketDevice) {
	objc.Send[struct{}](s.ID, objc.Sel("setSocketDevices:"), value)
}


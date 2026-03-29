// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCustomVirtioDevice] class.
var (
	_VZCustomVirtioDeviceClass     VZCustomVirtioDeviceClass
	_VZCustomVirtioDeviceClassOnce sync.Once
)

func getVZCustomVirtioDeviceClass() VZCustomVirtioDeviceClass {
	_VZCustomVirtioDeviceClassOnce.Do(func() {
		_VZCustomVirtioDeviceClass = VZCustomVirtioDeviceClass{class: objc.GetClass("_VZCustomVirtioDevice")}
	})
	return _VZCustomVirtioDeviceClass
}

// GetVZCustomVirtioDeviceClass returns the class object for _VZCustomVirtioDevice.
func GetVZCustomVirtioDeviceClass() VZCustomVirtioDeviceClass {
	return getVZCustomVirtioDeviceClass()
}

type VZCustomVirtioDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomVirtioDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomVirtioDeviceClass) Alloc() VZCustomVirtioDevice {
	rv := objc.Send[VZCustomVirtioDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCustomVirtioDevice.Delegate]
//   - [VZCustomVirtioDevice.SetDelegate]
//   - [VZCustomVirtioDevice.DeviceQueue]
//   - [VZCustomVirtioDevice.DriverFeaturesAtError]
//   - [VZCustomVirtioDevice.GuestMemoryAtPhysicalAddressLength]
//   - [VZCustomVirtioDevice.QueueAtIndex]
//   - [VZCustomVirtioDevice.RequestDeviceReset]
//   - [VZCustomVirtioDevice.UpdateDeviceSpecificConfigurationCompletionHandler]
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice
type VZCustomVirtioDevice struct {
	objectivec.Object
}

// VZCustomVirtioDeviceFromID constructs a [VZCustomVirtioDevice] from an objc.ID.
func VZCustomVirtioDeviceFromID(id objc.ID) VZCustomVirtioDevice {
	return VZCustomVirtioDevice{objectivec.Object{ID: id}}
}
// Ensure VZCustomVirtioDevice implements IVZCustomVirtioDevice.
var _ IVZCustomVirtioDevice = VZCustomVirtioDevice{}

// An interface definition for the [VZCustomVirtioDevice] class.
//
// # Methods
//
//   - [IVZCustomVirtioDevice.Delegate]
//   - [IVZCustomVirtioDevice.SetDelegate]
//   - [IVZCustomVirtioDevice.DeviceQueue]
//   - [IVZCustomVirtioDevice.DriverFeaturesAtError]
//   - [IVZCustomVirtioDevice.GuestMemoryAtPhysicalAddressLength]
//   - [IVZCustomVirtioDevice.QueueAtIndex]
//   - [IVZCustomVirtioDevice.RequestDeviceReset]
//   - [IVZCustomVirtioDevice.UpdateDeviceSpecificConfigurationCompletionHandler]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice
type IVZCustomVirtioDevice interface {
	objectivec.IObject

	// Topic: Methods

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DeviceQueue() objectivec.Object
	DriverFeaturesAtError(at uint32) (uint32, error)
	GuestMemoryAtPhysicalAddressLength(address uint64, length uint64) objectivec.IObject
	QueueAtIndex(index uint16) objectivec.IObject
	RequestDeviceReset()
	UpdateDeviceSpecificConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler)
}

// Init initializes the instance.
func (v VZCustomVirtioDevice) Init() VZCustomVirtioDevice {
	rv := objc.Send[VZCustomVirtioDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomVirtioDevice) Autorelease() VZCustomVirtioDevice {
	rv := objc.Send[VZCustomVirtioDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomVirtioDevice creates a new VZCustomVirtioDevice instance.
func NewVZCustomVirtioDevice() VZCustomVirtioDevice {
	class := getVZCustomVirtioDeviceClass()
	rv := objc.Send[VZCustomVirtioDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/driverFeaturesAt:error:
func (v VZCustomVirtioDevice) DriverFeaturesAtError(at uint32) (uint32, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint32](v.ID, objc.Sel("driverFeaturesAt:error:"), at, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/guestMemoryAtPhysicalAddress:length:
func (v VZCustomVirtioDevice) GuestMemoryAtPhysicalAddressLength(address uint64, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("guestMemoryAtPhysicalAddress:length:"), address, length)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/queueAtIndex:
func (v VZCustomVirtioDevice) QueueAtIndex(index uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("queueAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/requestDeviceReset
func (v VZCustomVirtioDevice) RequestDeviceReset() {
	objc.Send[objc.ID](v.ID, objc.Sel("requestDeviceReset"))
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/updateDeviceSpecificConfiguration:completionHandler:
func (v VZCustomVirtioDevice) UpdateDeviceSpecificConfigurationCompletionHandler(configuration objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("updateDeviceSpecificConfiguration:completionHandler:"), configuration, _block1)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/delegate
func (v VZCustomVirtioDevice) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (v VZCustomVirtioDevice) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomVirtioDevice/deviceQueue
func (v VZCustomVirtioDevice) DeviceQueue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("deviceQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// UpdateDeviceSpecificConfiguration is a synchronous wrapper around [VZCustomVirtioDevice.UpdateDeviceSpecificConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v VZCustomVirtioDevice) UpdateDeviceSpecificConfiguration(ctx context.Context, configuration objectivec.IObject) error {
	done := make(chan error, 1)
	v.UpdateDeviceSpecificConfigurationCompletionHandler(configuration, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}


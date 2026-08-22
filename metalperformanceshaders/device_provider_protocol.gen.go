// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that enables the setting of a Metal device for unarchived objects.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceProvider
type MPSDeviceProvider interface {
	objectivec.IObject

	// MpsMTLDevice protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceProvider/mpsMTLDevice()
	MpsMTLDevice() metal.MTLDevice
}

// MPSDeviceProviderObject wraps an existing Objective-C object that conforms to the MPSDeviceProvider protocol.
type MPSDeviceProviderObject struct {
	objectivec.Object
}

func (o MPSDeviceProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MPSDeviceProviderObjectFromID constructs a [MPSDeviceProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSDeviceProviderObjectFromID(id objc.ID) MPSDeviceProviderObject {
	return MPSDeviceProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDeviceProvider/mpsMTLDevice()
func (o MPSDeviceProviderObject) MpsMTLDevice() metal.MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mpsMTLDevice"))
	return metal.MTLDeviceObjectFromID(rv)
}

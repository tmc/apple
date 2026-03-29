// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVhostUserNetworkDeviceAttachment] class.
var (
	_VZVhostUserNetworkDeviceAttachmentClass     VZVhostUserNetworkDeviceAttachmentClass
	_VZVhostUserNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZVhostUserNetworkDeviceAttachmentClass() VZVhostUserNetworkDeviceAttachmentClass {
	_VZVhostUserNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZVhostUserNetworkDeviceAttachmentClass = VZVhostUserNetworkDeviceAttachmentClass{class: objc.GetClass("_VZVhostUserNetworkDeviceAttachment")}
	})
	return _VZVhostUserNetworkDeviceAttachmentClass
}

// GetVZVhostUserNetworkDeviceAttachmentClass returns the class object for _VZVhostUserNetworkDeviceAttachment.
func GetVZVhostUserNetworkDeviceAttachmentClass() VZVhostUserNetworkDeviceAttachmentClass {
	return getVZVhostUserNetworkDeviceAttachmentClass()
}

type VZVhostUserNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVhostUserNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVhostUserNetworkDeviceAttachmentClass) Alloc() VZVhostUserNetworkDeviceAttachment {
	rv := objc.Send[VZVhostUserNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVhostUserNetworkDeviceAttachment.EncodeWithEncoder]
//   - [VZVhostUserNetworkDeviceAttachment.GuestChecksumOffload]
//   - [VZVhostUserNetworkDeviceAttachment.GuestTcpSegmentationOffload]
//   - [VZVhostUserNetworkDeviceAttachment.HostChecksumOffload]
//   - [VZVhostUserNetworkDeviceAttachment.HostTcpSegmentationOffload]
//   - [VZVhostUserNetworkDeviceAttachment.Interface]
//   - [VZVhostUserNetworkDeviceAttachment.MaximumTransmissionUnit]
//   - [VZVhostUserNetworkDeviceAttachment.XpcService]
//   - [VZVhostUserNetworkDeviceAttachment.InitWithInterfaceError]
//   - [VZVhostUserNetworkDeviceAttachment.InitWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError]
//   - [VZVhostUserNetworkDeviceAttachment.InitWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError]
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment
type VZVhostUserNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZVhostUserNetworkDeviceAttachmentFromID constructs a [VZVhostUserNetworkDeviceAttachment] from an objc.ID.
func VZVhostUserNetworkDeviceAttachmentFromID(id objc.ID) VZVhostUserNetworkDeviceAttachment {
	return VZVhostUserNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// Ensure VZVhostUserNetworkDeviceAttachment implements IVZVhostUserNetworkDeviceAttachment.
var _ IVZVhostUserNetworkDeviceAttachment = VZVhostUserNetworkDeviceAttachment{}

// An interface definition for the [VZVhostUserNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZVhostUserNetworkDeviceAttachment.EncodeWithEncoder]
//   - [IVZVhostUserNetworkDeviceAttachment.GuestChecksumOffload]
//   - [IVZVhostUserNetworkDeviceAttachment.GuestTcpSegmentationOffload]
//   - [IVZVhostUserNetworkDeviceAttachment.HostChecksumOffload]
//   - [IVZVhostUserNetworkDeviceAttachment.HostTcpSegmentationOffload]
//   - [IVZVhostUserNetworkDeviceAttachment.Interface]
//   - [IVZVhostUserNetworkDeviceAttachment.MaximumTransmissionUnit]
//   - [IVZVhostUserNetworkDeviceAttachment.XpcService]
//   - [IVZVhostUserNetworkDeviceAttachment.InitWithInterfaceError]
//   - [IVZVhostUserNetworkDeviceAttachment.InitWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError]
//   - [IVZVhostUserNetworkDeviceAttachment.InitWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment
type IVZVhostUserNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Methods

	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	GuestChecksumOffload() int64
	GuestTcpSegmentationOffload() int64
	HostChecksumOffload() int64
	HostTcpSegmentationOffload() int64
	Interface() string
	MaximumTransmissionUnit() int64
	XpcService() int64
	InitWithInterfaceError(interface_ objectivec.IObject) (VZVhostUserNetworkDeviceAttachment, error)
	InitWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error)
	InitWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, service int64, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error)
}

// Init initializes the instance.
func (v VZVhostUserNetworkDeviceAttachment) Init() VZVhostUserNetworkDeviceAttachment {
	rv := objc.Send[VZVhostUserNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVhostUserNetworkDeviceAttachment) Autorelease() VZVhostUserNetworkDeviceAttachment {
	rv := objc.Send[VZVhostUserNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVhostUserNetworkDeviceAttachment creates a new VZVhostUserNetworkDeviceAttachment instance.
func NewVZVhostUserNetworkDeviceAttachment() VZVhostUserNetworkDeviceAttachment {
	class := getVZVhostUserNetworkDeviceAttachmentClass()
	rv := objc.Send[VZVhostUserNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:error:
func NewVZVhostUserNetworkDeviceAttachmentWithInterfaceError(interface_ objectivec.IObject) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZVhostUserNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterface:error:"), interface_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:
func NewVZVhostUserNetworkDeviceAttachmentWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZVhostUserNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterface:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:"), interface_, unit, offload, offload2, offload3, offload4, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:xpcService:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:
func NewVZVhostUserNetworkDeviceAttachmentWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, service int64, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	instance := getVZVhostUserNetworkDeviceAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInterface:xpcService:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:"), interface_, service, unit, offload, offload2, offload3, offload4, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/encodeWithEncoder:
func (v VZVhostUserNetworkDeviceAttachment) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:error:
func (v VZVhostUserNetworkDeviceAttachment) InitWithInterfaceError(interface_ objectivec.IObject) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithInterface:error:"), interface_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:
func (v VZVhostUserNetworkDeviceAttachment) InitWithInterfaceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithInterface:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:"), interface_, unit, offload, offload2, offload3, offload4, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/initWithInterface:xpcService:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:
func (v VZVhostUserNetworkDeviceAttachment) InitWithInterfaceXpcServiceMaximumTransmissionUnitHostChecksumOffloadGuestChecksumOffloadHostTcpSegmentationOffloadGuestTcpSegmentationOffloadError(interface_ objectivec.IObject, service int64, unit int64, offload int64, offload2 int64, offload3 int64, offload4 int64) (VZVhostUserNetworkDeviceAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithInterface:xpcService:maximumTransmissionUnit:hostChecksumOffload:guestChecksumOffload:hostTcpSegmentationOffload:guestTcpSegmentationOffload:error:"), interface_, service, unit, offload, offload2, offload3, offload4, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZVhostUserNetworkDeviceAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZVhostUserNetworkDeviceAttachmentFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/_defaultMaximumTransmissionUnit
func (_VZVhostUserNetworkDeviceAttachmentClass VZVhostUserNetworkDeviceAttachmentClass) _defaultMaximumTransmissionUnit() int64 {
	rv := objc.Send[int64](objc.ID(_VZVhostUserNetworkDeviceAttachmentClass.class), objc.Sel("_defaultMaximumTransmissionUnit"))
	return rv
}

// DefaultMaximumTransmissionUnit is an exported wrapper for the private method _defaultMaximumTransmissionUnit.
func (_VZVhostUserNetworkDeviceAttachmentClass VZVhostUserNetworkDeviceAttachmentClass) DefaultMaximumTransmissionUnit() int64 {
	return _VZVhostUserNetworkDeviceAttachmentClass._defaultMaximumTransmissionUnit()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/_defaultOffloadMode
func (_VZVhostUserNetworkDeviceAttachmentClass VZVhostUserNetworkDeviceAttachmentClass) _defaultOffloadMode() int64 {
	rv := objc.Send[int64](objc.ID(_VZVhostUserNetworkDeviceAttachmentClass.class), objc.Sel("_defaultOffloadMode"))
	return rv
}

// DefaultOffloadMode is an exported wrapper for the private method _defaultOffloadMode.
func (_VZVhostUserNetworkDeviceAttachmentClass VZVhostUserNetworkDeviceAttachmentClass) DefaultOffloadMode() int64 {
	return _VZVhostUserNetworkDeviceAttachmentClass._defaultOffloadMode()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/guestChecksumOffload
func (v VZVhostUserNetworkDeviceAttachment) GuestChecksumOffload() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("guestChecksumOffload"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/guestTcpSegmentationOffload
func (v VZVhostUserNetworkDeviceAttachment) GuestTcpSegmentationOffload() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("guestTcpSegmentationOffload"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/hostChecksumOffload
func (v VZVhostUserNetworkDeviceAttachment) HostChecksumOffload() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("hostChecksumOffload"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/hostTcpSegmentationOffload
func (v VZVhostUserNetworkDeviceAttachment) HostTcpSegmentationOffload() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("hostTcpSegmentationOffload"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/interface
func (v VZVhostUserNetworkDeviceAttachment) Interface() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("interface"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/maximumTransmissionUnit
func (v VZVhostUserNetworkDeviceAttachment) MaximumTransmissionUnit() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("maximumTransmissionUnit"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZVhostUserNetworkDeviceAttachment/xpcService
func (v VZVhostUserNetworkDeviceAttachment) XpcService() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("xpcService"))
	return rv
}


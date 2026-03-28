// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZBridgedNetworkDeviceAttachment] class.
var (
	_VZBridgedNetworkDeviceAttachmentClass     VZBridgedNetworkDeviceAttachmentClass
	_VZBridgedNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	_VZBridgedNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZBridgedNetworkDeviceAttachmentClass = VZBridgedNetworkDeviceAttachmentClass{class: objc.GetClass("VZBridgedNetworkDeviceAttachment")}
	})
	return _VZBridgedNetworkDeviceAttachmentClass
}

// GetVZBridgedNetworkDeviceAttachmentClass returns the class object for VZBridgedNetworkDeviceAttachment.
func GetVZBridgedNetworkDeviceAttachmentClass() VZBridgedNetworkDeviceAttachmentClass {
	return getVZBridgedNetworkDeviceAttachmentClass()
}

type VZBridgedNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBridgedNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBridgedNetworkDeviceAttachmentClass) Alloc() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZBridgedNetworkDeviceAttachment._macNatEnabled]
//   - [VZBridgedNetworkDeviceAttachment.Set_macNatEnabled]
//   - [VZBridgedNetworkDeviceAttachment._setMacNatEnabled]
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment
type VZBridgedNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZBridgedNetworkDeviceAttachmentFromID constructs a [VZBridgedNetworkDeviceAttachment] from an objc.ID.
func VZBridgedNetworkDeviceAttachmentFromID(id objc.ID) VZBridgedNetworkDeviceAttachment {
	return VZBridgedNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}
// Ensure VZBridgedNetworkDeviceAttachment implements IVZBridgedNetworkDeviceAttachment.
var _ IVZBridgedNetworkDeviceAttachment = VZBridgedNetworkDeviceAttachment{}

// An interface definition for the [VZBridgedNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZBridgedNetworkDeviceAttachment._macNatEnabled]
//   - [IVZBridgedNetworkDeviceAttachment.Set_macNatEnabled]
//   - [IVZBridgedNetworkDeviceAttachment._setMacNatEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment
type IVZBridgedNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Methods

	_macNatEnabled() bool
	Set_macNatEnabled(value bool)
	_setMacNatEnabled(enabled bool)
}

// Init initializes the instance.
func (b VZBridgedNetworkDeviceAttachment) Init() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b VZBridgedNetworkDeviceAttachment) Autorelease() VZBridgedNetworkDeviceAttachment {
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkDeviceAttachment creates a new VZBridgedNetworkDeviceAttachment instance.
func NewVZBridgedNetworkDeviceAttachment() VZBridgedNetworkDeviceAttachment {
	class := getVZBridgedNetworkDeviceAttachmentClass()
	rv := objc.Send[VZBridgedNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/_setMacNatEnabled:
func (b VZBridgedNetworkDeviceAttachment) _setMacNatEnabled(enabled bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("_setMacNatEnabled:"), enabled)
}

// SetMacNatEnabled is an exported wrapper for the private method _setMacNatEnabled.
func (b VZBridgedNetworkDeviceAttachment) SetMacNatEnabled(enabled bool) {
	b._setMacNatEnabled(enabled)
}

// See: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkDeviceAttachment/_macNatEnabled
func (b VZBridgedNetworkDeviceAttachment) _macNatEnabled() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("_macNatEnabled"))
	return rv
}
func (b VZBridgedNetworkDeviceAttachment) Set_macNatEnabled(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("set_macNatEnabled:"), value)
}


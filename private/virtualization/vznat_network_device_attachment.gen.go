// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZNATNetworkDeviceAttachment] class.
var (
	_VZNATNetworkDeviceAttachmentClass     VZNATNetworkDeviceAttachmentClass
	_VZNATNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNATNetworkDeviceAttachmentClass() VZNATNetworkDeviceAttachmentClass {
	_VZNATNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZNATNetworkDeviceAttachmentClass = VZNATNetworkDeviceAttachmentClass{class: objc.GetClass("VZNATNetworkDeviceAttachment")}
	})
	return _VZNATNetworkDeviceAttachmentClass
}

// GetVZNATNetworkDeviceAttachmentClass returns the class object for VZNATNetworkDeviceAttachment.
func GetVZNATNetworkDeviceAttachmentClass() VZNATNetworkDeviceAttachmentClass {
	return getVZNATNetworkDeviceAttachmentClass()
}

type VZNATNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNATNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNATNetworkDeviceAttachmentClass) Alloc() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNATNetworkDeviceAttachment._interfaceIsolationEnabled]
//   - [VZNATNetworkDeviceAttachment.Set_interfaceIsolationEnabled]
//   - [VZNATNetworkDeviceAttachment._setInterfaceIsolationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment
type VZNATNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZNATNetworkDeviceAttachmentFromID constructs a [VZNATNetworkDeviceAttachment] from an objc.ID.
func VZNATNetworkDeviceAttachmentFromID(id objc.ID) VZNATNetworkDeviceAttachment {
	return VZNATNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZNATNetworkDeviceAttachment implements IVZNATNetworkDeviceAttachment.
var _ IVZNATNetworkDeviceAttachment = VZNATNetworkDeviceAttachment{}

// An interface definition for the [VZNATNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZNATNetworkDeviceAttachment._interfaceIsolationEnabled]
//   - [IVZNATNetworkDeviceAttachment.Set_interfaceIsolationEnabled]
//   - [IVZNATNetworkDeviceAttachment._setInterfaceIsolationEnabled]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment
type IVZNATNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment

	// Topic: Methods

	_interfaceIsolationEnabled() bool
	Set_interfaceIsolationEnabled(value bool)
	_setInterfaceIsolationEnabled(enabled bool)
}

// Init initializes the instance.
func (n VZNATNetworkDeviceAttachment) Init() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNATNetworkDeviceAttachment) Autorelease() VZNATNetworkDeviceAttachment {
	rv := objc.Send[VZNATNetworkDeviceAttachment](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNATNetworkDeviceAttachment creates a new VZNATNetworkDeviceAttachment instance.
func NewVZNATNetworkDeviceAttachment() VZNATNetworkDeviceAttachment {
	class := getVZNATNetworkDeviceAttachmentClass()
	rv := objc.Send[VZNATNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment/_setInterfaceIsolationEnabled:
func (n VZNATNetworkDeviceAttachment) _setInterfaceIsolationEnabled(enabled bool) {
	objc.Send[objc.ID](n.ID, objc.Sel("_setInterfaceIsolationEnabled:"), enabled)
}

// SetInterfaceIsolationEnabled is an exported wrapper for the private method _setInterfaceIsolationEnabled.
func (n VZNATNetworkDeviceAttachment) SetInterfaceIsolationEnabled(enabled bool) {
	n._setInterfaceIsolationEnabled(enabled)
}

// See: https://developer.apple.com/documentation/Virtualization/VZNATNetworkDeviceAttachment/_interfaceIsolationEnabled
func (n VZNATNetworkDeviceAttachment) _interfaceIsolationEnabled() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("_interfaceIsolationEnabled"))
	return rv
}
func (n VZNATNetworkDeviceAttachment) Set_interfaceIsolationEnabled(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("set_interfaceIsolationEnabled:"), value)
}

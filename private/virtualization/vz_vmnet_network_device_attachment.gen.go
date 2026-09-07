// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVmnetNetworkDeviceAttachment] class.
var (
	_VZVmnetNetworkDeviceAttachmentClass     VZVmnetNetworkDeviceAttachmentClass
	_VZVmnetNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	_VZVmnetNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZVmnetNetworkDeviceAttachmentClass = VZVmnetNetworkDeviceAttachmentClass{class: objc.GetClass("VZVmnetNetworkDeviceAttachment")}
	})
	return _VZVmnetNetworkDeviceAttachmentClass
}

// GetVZVmnetNetworkDeviceAttachmentClass returns the class object for VZVmnetNetworkDeviceAttachment.
func GetVZVmnetNetworkDeviceAttachmentClass() VZVmnetNetworkDeviceAttachmentClass {
	return getVZVmnetNetworkDeviceAttachmentClass()
}

type VZVmnetNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVmnetNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVmnetNetworkDeviceAttachmentClass) Alloc() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZVmnetNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZVmnetNetworkDeviceAttachmentFromID constructs a [VZVmnetNetworkDeviceAttachment] from an objc.ID.
func VZVmnetNetworkDeviceAttachmentFromID(id objc.ID) VZVmnetNetworkDeviceAttachment {
	return VZVmnetNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZVmnetNetworkDeviceAttachment implements IVZVmnetNetworkDeviceAttachment.
var _ IVZVmnetNetworkDeviceAttachment = VZVmnetNetworkDeviceAttachment{}

// An interface definition for the [VZVmnetNetworkDeviceAttachment] class.
type IVZVmnetNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
}

// Init initializes the instance.
func (v VZVmnetNetworkDeviceAttachment) Init() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVmnetNetworkDeviceAttachment) Autorelease() VZVmnetNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVmnetNetworkDeviceAttachment creates a new VZVmnetNetworkDeviceAttachment instance.
func NewVZVmnetNetworkDeviceAttachment() VZVmnetNetworkDeviceAttachment {
	class := getVZVmnetNetworkDeviceAttachmentClass()
	rv := objc.SendIfResponds[VZVmnetNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

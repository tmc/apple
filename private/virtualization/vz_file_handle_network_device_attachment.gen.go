// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZFileHandleNetworkDeviceAttachment] class.
var (
	_VZFileHandleNetworkDeviceAttachmentClass     VZFileHandleNetworkDeviceAttachmentClass
	_VZFileHandleNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZFileHandleNetworkDeviceAttachmentClass() VZFileHandleNetworkDeviceAttachmentClass {
	_VZFileHandleNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZFileHandleNetworkDeviceAttachmentClass = VZFileHandleNetworkDeviceAttachmentClass{class: objc.GetClass("VZFileHandleNetworkDeviceAttachment")}
	})
	return _VZFileHandleNetworkDeviceAttachmentClass
}

// GetVZFileHandleNetworkDeviceAttachmentClass returns the class object for VZFileHandleNetworkDeviceAttachment.
func GetVZFileHandleNetworkDeviceAttachmentClass() VZFileHandleNetworkDeviceAttachmentClass {
	return getVZFileHandleNetworkDeviceAttachmentClass()
}

type VZFileHandleNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFileHandleNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileHandleNetworkDeviceAttachmentClass) Alloc() VZFileHandleNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZFileHandleNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZFileHandleNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZFileHandleNetworkDeviceAttachmentFromID constructs a [VZFileHandleNetworkDeviceAttachment] from an objc.ID.
func VZFileHandleNetworkDeviceAttachmentFromID(id objc.ID) VZFileHandleNetworkDeviceAttachment {
	return VZFileHandleNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZFileHandleNetworkDeviceAttachment implements IVZFileHandleNetworkDeviceAttachment.
var _ IVZFileHandleNetworkDeviceAttachment = VZFileHandleNetworkDeviceAttachment{}

// An interface definition for the [VZFileHandleNetworkDeviceAttachment] class.
type IVZFileHandleNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
}

// Init initializes the instance.
func (v VZFileHandleNetworkDeviceAttachment) Init() VZFileHandleNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZFileHandleNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFileHandleNetworkDeviceAttachment) Autorelease() VZFileHandleNetworkDeviceAttachment {
	rv := objc.SendIfResponds[VZFileHandleNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleNetworkDeviceAttachment creates a new VZFileHandleNetworkDeviceAttachment instance.
func NewVZFileHandleNetworkDeviceAttachment() VZFileHandleNetworkDeviceAttachment {
	class := getVZFileHandleNetworkDeviceAttachmentClass()
	rv := objc.SendIfResponds[VZFileHandleNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

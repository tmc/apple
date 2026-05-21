// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZHostOnlyNetworkDeviceAttachment] class.
var (
	_VZHostOnlyNetworkDeviceAttachmentClass     VZHostOnlyNetworkDeviceAttachmentClass
	_VZHostOnlyNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZHostOnlyNetworkDeviceAttachmentClass() VZHostOnlyNetworkDeviceAttachmentClass {
	_VZHostOnlyNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZHostOnlyNetworkDeviceAttachmentClass = VZHostOnlyNetworkDeviceAttachmentClass{class: objc.GetClass("_VZHostOnlyNetworkDeviceAttachment")}
	})
	return _VZHostOnlyNetworkDeviceAttachmentClass
}

// GetVZHostOnlyNetworkDeviceAttachmentClass returns the class object for _VZHostOnlyNetworkDeviceAttachment.
func GetVZHostOnlyNetworkDeviceAttachmentClass() VZHostOnlyNetworkDeviceAttachmentClass {
	return getVZHostOnlyNetworkDeviceAttachmentClass()
}

type VZHostOnlyNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHostOnlyNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHostOnlyNetworkDeviceAttachmentClass) Alloc() VZHostOnlyNetworkDeviceAttachment {
	rv := objc.Send[VZHostOnlyNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZHostOnlyNetworkDeviceAttachment struct {
	VZNetworkDeviceAttachment
}

// VZHostOnlyNetworkDeviceAttachmentFromID constructs a [VZHostOnlyNetworkDeviceAttachment] from an objc.ID.
func VZHostOnlyNetworkDeviceAttachmentFromID(id objc.ID) VZHostOnlyNetworkDeviceAttachment {
	return VZHostOnlyNetworkDeviceAttachment{VZNetworkDeviceAttachment: VZNetworkDeviceAttachmentFromID(id)}
}

// Ensure VZHostOnlyNetworkDeviceAttachment implements IVZHostOnlyNetworkDeviceAttachment.
var _ IVZHostOnlyNetworkDeviceAttachment = VZHostOnlyNetworkDeviceAttachment{}

// An interface definition for the [VZHostOnlyNetworkDeviceAttachment] class.
type IVZHostOnlyNetworkDeviceAttachment interface {
	IVZNetworkDeviceAttachment
}

// Init initializes the instance.
func (v VZHostOnlyNetworkDeviceAttachment) Init() VZHostOnlyNetworkDeviceAttachment {
	rv := objc.Send[VZHostOnlyNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHostOnlyNetworkDeviceAttachment) Autorelease() VZHostOnlyNetworkDeviceAttachment {
	rv := objc.Send[VZHostOnlyNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostOnlyNetworkDeviceAttachment creates a new VZHostOnlyNetworkDeviceAttachment instance.
func NewVZHostOnlyNetworkDeviceAttachment() VZHostOnlyNetworkDeviceAttachment {
	class := getVZHostOnlyNetworkDeviceAttachmentClass()
	rv := objc.Send[VZHostOnlyNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

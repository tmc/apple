// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDeviceAttachment] class.
var (
	_VZStorageDeviceAttachmentClass     VZStorageDeviceAttachmentClass
	_VZStorageDeviceAttachmentClassOnce sync.Once
)

func getVZStorageDeviceAttachmentClass() VZStorageDeviceAttachmentClass {
	_VZStorageDeviceAttachmentClassOnce.Do(func() {
		_VZStorageDeviceAttachmentClass = VZStorageDeviceAttachmentClass{class: objc.GetClass("VZStorageDeviceAttachment")}
	})
	return _VZStorageDeviceAttachmentClass
}

// GetVZStorageDeviceAttachmentClass returns the class object for VZStorageDeviceAttachment.
func GetVZStorageDeviceAttachmentClass() VZStorageDeviceAttachmentClass {
	return getVZStorageDeviceAttachmentClass()
}

type VZStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceAttachmentClass) Alloc() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The common behaviors for storage devices in the guest system.
//
// # Overview
// 
// A [VZStorageDeviceAttachment] object defines the implementation of a
// storage interface in a virtual machine. You use the attachment object to
// specify the source of the storage on the host computer.
// 
// Don’t create [VZStorageDeviceAttachment] objects directly. Instead,
// instantiate an appropriate subclass such as
// [VZDiskImageStorageDeviceAttachment], which provides storage using a disk
// image.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment
type VZStorageDeviceAttachment struct {
	objectivec.Object
}

// VZStorageDeviceAttachmentFromID constructs a [VZStorageDeviceAttachment] from an objc.ID.
//
// The common behaviors for storage devices in the guest system.
func VZStorageDeviceAttachmentFromID(id objc.ID) VZStorageDeviceAttachment {
	return VZStorageDeviceAttachment{objectivec.Object{id}}
}
// NOTE: VZStorageDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZStorageDeviceAttachment] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment
type IVZStorageDeviceAttachment interface {
	objectivec.IObject
}





// Init initializes the instance.
func (s VZStorageDeviceAttachment) Init() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDeviceAttachment) Autorelease() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceAttachment creates a new VZStorageDeviceAttachment instance.
func NewVZStorageDeviceAttachment() VZStorageDeviceAttachment {
	class := getVZStorageDeviceAttachmentClass()
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































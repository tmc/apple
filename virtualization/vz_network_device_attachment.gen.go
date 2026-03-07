// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDeviceAttachment] class.
var (
	_VZNetworkDeviceAttachmentClass     VZNetworkDeviceAttachmentClass
	_VZNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkDeviceAttachmentClass() VZNetworkDeviceAttachmentClass {
	_VZNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZNetworkDeviceAttachmentClass = VZNetworkDeviceAttachmentClass{class: objc.GetClass("VZNetworkDeviceAttachment")}
	})
	return _VZNetworkDeviceAttachmentClass
}

// GetVZNetworkDeviceAttachmentClass returns the class object for VZNetworkDeviceAttachment.
func GetVZNetworkDeviceAttachmentClass() VZNetworkDeviceAttachmentClass {
	return getVZNetworkDeviceAttachmentClass()
}

type VZNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceAttachmentClass) Alloc() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The common behaviors for the network attachment points of your virtual
// machine.
//
// # Overview
// 
// Don’t create a [VZNetworkDeviceAttachment] object directly. Instead,
// instantiate one of its concrete subclasses and use that object to configure
// your network devices. Each concrete subclass represents a specific type of
// network interface on the host computer.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment
type VZNetworkDeviceAttachment struct {
	objectivec.Object
}

// VZNetworkDeviceAttachmentFromID constructs a [VZNetworkDeviceAttachment] from an objc.ID.
//
// The common behaviors for the network attachment points of your virtual
// machine.
func VZNetworkDeviceAttachmentFromID(id objc.ID) VZNetworkDeviceAttachment {
	return VZNetworkDeviceAttachment{objectivec.Object{id}}
}
// NOTE: VZNetworkDeviceAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZNetworkDeviceAttachment] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment
type IVZNetworkDeviceAttachment interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n VZNetworkDeviceAttachment) Init() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkDeviceAttachment) Autorelease() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceAttachment creates a new VZNetworkDeviceAttachment instance.
func NewVZNetworkDeviceAttachment() VZNetworkDeviceAttachment {
	class := getVZNetworkDeviceAttachmentClass()
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































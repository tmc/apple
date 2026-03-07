// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSerialPortAttachment] class.
var (
	_VZSerialPortAttachmentClass     VZSerialPortAttachmentClass
	_VZSerialPortAttachmentClassOnce sync.Once
)

func getVZSerialPortAttachmentClass() VZSerialPortAttachmentClass {
	_VZSerialPortAttachmentClassOnce.Do(func() {
		_VZSerialPortAttachmentClass = VZSerialPortAttachmentClass{class: objc.GetClass("VZSerialPortAttachment")}
	})
	return _VZSerialPortAttachmentClass
}

// GetVZSerialPortAttachmentClass returns the class object for VZSerialPortAttachment.
func GetVZSerialPortAttachmentClass() VZSerialPortAttachmentClass {
	return getVZSerialPortAttachmentClass()
}

type VZSerialPortAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSerialPortAttachmentClass) Alloc() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The common behaviors for the serial attachment points of your virtual
// machine.
//
// # Overview
// 
// Don’t create a [VZSerialPortAttachment] object directly. Instead,
// instantiate a concrete subclass such as [VZFileHandleSerialPortAttachment]
// to configure how the virtual machine’s serial port connects with the host
// computer.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type VZSerialPortAttachment struct {
	objectivec.Object
}

// VZSerialPortAttachmentFromID constructs a [VZSerialPortAttachment] from an objc.ID.
//
// The common behaviors for the serial attachment points of your virtual
// machine.
func VZSerialPortAttachmentFromID(id objc.ID) VZSerialPortAttachment {
	return VZSerialPortAttachment{objectivec.Object{id}}
}
// NOTE: VZSerialPortAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZSerialPortAttachment] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type IVZSerialPortAttachment interface {
	objectivec.IObject
}





// Init initializes the instance.
func (s VZSerialPortAttachment) Init() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSerialPortAttachment) Autorelease() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortAttachment creates a new VZSerialPortAttachment instance.
func NewVZSerialPortAttachment() VZSerialPortAttachment {
	class := getVZSerialPortAttachmentClass()
	rv := objc.Send[VZSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































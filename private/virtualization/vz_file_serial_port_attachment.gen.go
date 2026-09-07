// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZFileSerialPortAttachment] class.
var (
	_VZFileSerialPortAttachmentClass     VZFileSerialPortAttachmentClass
	_VZFileSerialPortAttachmentClassOnce sync.Once
)

func getVZFileSerialPortAttachmentClass() VZFileSerialPortAttachmentClass {
	_VZFileSerialPortAttachmentClassOnce.Do(func() {
		_VZFileSerialPortAttachmentClass = VZFileSerialPortAttachmentClass{class: objc.GetClass("VZFileSerialPortAttachment")}
	})
	return _VZFileSerialPortAttachmentClass
}

// GetVZFileSerialPortAttachmentClass returns the class object for VZFileSerialPortAttachment.
func GetVZFileSerialPortAttachmentClass() VZFileSerialPortAttachmentClass {
	return getVZFileSerialPortAttachmentClass()
}

type VZFileSerialPortAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFileSerialPortAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileSerialPortAttachmentClass) Alloc() VZFileSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZFileSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileSerialPortAttachmentFromID constructs a [VZFileSerialPortAttachment] from an objc.ID.
func VZFileSerialPortAttachmentFromID(id objc.ID) VZFileSerialPortAttachment {
	return VZFileSerialPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}

// Ensure VZFileSerialPortAttachment implements IVZFileSerialPortAttachment.
var _ IVZFileSerialPortAttachment = VZFileSerialPortAttachment{}

// An interface definition for the [VZFileSerialPortAttachment] class.
type IVZFileSerialPortAttachment interface {
	IVZSerialPortAttachment
}

// Init initializes the instance.
func (v VZFileSerialPortAttachment) Init() VZFileSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileSerialPortAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFileSerialPortAttachment) Autorelease() VZFileSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileSerialPortAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileSerialPortAttachment creates a new VZFileSerialPortAttachment instance.
func NewVZFileSerialPortAttachment() VZFileSerialPortAttachment {
	class := getVZFileSerialPortAttachmentClass()
	rv := objc.SendIfResponds[VZFileSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

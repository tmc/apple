// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZFileHandleSerialPortAttachment] class.
var (
	_VZFileHandleSerialPortAttachmentClass     VZFileHandleSerialPortAttachmentClass
	_VZFileHandleSerialPortAttachmentClassOnce sync.Once
)

func getVZFileHandleSerialPortAttachmentClass() VZFileHandleSerialPortAttachmentClass {
	_VZFileHandleSerialPortAttachmentClassOnce.Do(func() {
		_VZFileHandleSerialPortAttachmentClass = VZFileHandleSerialPortAttachmentClass{class: objc.GetClass("VZFileHandleSerialPortAttachment")}
	})
	return _VZFileHandleSerialPortAttachmentClass
}

// GetVZFileHandleSerialPortAttachmentClass returns the class object for VZFileHandleSerialPortAttachment.
func GetVZFileHandleSerialPortAttachmentClass() VZFileHandleSerialPortAttachmentClass {
	return getVZFileHandleSerialPortAttachmentClass()
}

type VZFileHandleSerialPortAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFileHandleSerialPortAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileHandleSerialPortAttachmentClass) Alloc() VZFileHandleSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileHandleSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZFileHandleSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileHandleSerialPortAttachmentFromID constructs a [VZFileHandleSerialPortAttachment] from an objc.ID.
func VZFileHandleSerialPortAttachmentFromID(id objc.ID) VZFileHandleSerialPortAttachment {
	return VZFileHandleSerialPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}

// Ensure VZFileHandleSerialPortAttachment implements IVZFileHandleSerialPortAttachment.
var _ IVZFileHandleSerialPortAttachment = VZFileHandleSerialPortAttachment{}

// An interface definition for the [VZFileHandleSerialPortAttachment] class.
type IVZFileHandleSerialPortAttachment interface {
	IVZSerialPortAttachment
}

// Init initializes the instance.
func (v VZFileHandleSerialPortAttachment) Init() VZFileHandleSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileHandleSerialPortAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFileHandleSerialPortAttachment) Autorelease() VZFileHandleSerialPortAttachment {
	rv := objc.SendIfResponds[VZFileHandleSerialPortAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleSerialPortAttachment creates a new VZFileHandleSerialPortAttachment instance.
func NewVZFileHandleSerialPortAttachment() VZFileHandleSerialPortAttachment {
	class := getVZFileHandleSerialPortAttachmentClass()
	rv := objc.SendIfResponds[VZFileHandleSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

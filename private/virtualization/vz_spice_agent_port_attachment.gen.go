// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZSpiceAgentPortAttachment] class.
var (
	_VZSpiceAgentPortAttachmentClass     VZSpiceAgentPortAttachmentClass
	_VZSpiceAgentPortAttachmentClassOnce sync.Once
)

func getVZSpiceAgentPortAttachmentClass() VZSpiceAgentPortAttachmentClass {
	_VZSpiceAgentPortAttachmentClassOnce.Do(func() {
		_VZSpiceAgentPortAttachmentClass = VZSpiceAgentPortAttachmentClass{class: objc.GetClass("VZSpiceAgentPortAttachment")}
	})
	return _VZSpiceAgentPortAttachmentClass
}

// GetVZSpiceAgentPortAttachmentClass returns the class object for VZSpiceAgentPortAttachment.
func GetVZSpiceAgentPortAttachmentClass() VZSpiceAgentPortAttachmentClass {
	return getVZSpiceAgentPortAttachmentClass()
}

type VZSpiceAgentPortAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSpiceAgentPortAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSpiceAgentPortAttachmentClass) Alloc() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment
type VZSpiceAgentPortAttachment struct {
	VZSerialPortAttachment
}

// VZSpiceAgentPortAttachmentFromID constructs a [VZSpiceAgentPortAttachment] from an objc.ID.
func VZSpiceAgentPortAttachmentFromID(id objc.ID) VZSpiceAgentPortAttachment {
	return VZSpiceAgentPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}

// Ensure VZSpiceAgentPortAttachment implements IVZSpiceAgentPortAttachment.
var _ IVZSpiceAgentPortAttachment = VZSpiceAgentPortAttachment{}

// An interface definition for the [VZSpiceAgentPortAttachment] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment
type IVZSpiceAgentPortAttachment interface {
	IVZSerialPortAttachment
}

// Init initializes the instance.
func (s VZSpiceAgentPortAttachment) Init() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSpiceAgentPortAttachment) Autorelease() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentPortAttachment creates a new VZSpiceAgentPortAttachment instance.
func NewVZSpiceAgentPortAttachment() VZSpiceAgentPortAttachment {
	class := getVZSpiceAgentPortAttachmentClass()
	rv := objc.Send[VZSpiceAgentPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

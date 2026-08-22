// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZXPCBifrostAttachment] class.
var (
	_VZXPCBifrostAttachmentClass     VZXPCBifrostAttachmentClass
	_VZXPCBifrostAttachmentClassOnce sync.Once
)

func getVZXPCBifrostAttachmentClass() VZXPCBifrostAttachmentClass {
	_VZXPCBifrostAttachmentClassOnce.Do(func() {
		_VZXPCBifrostAttachmentClass = VZXPCBifrostAttachmentClass{class: objc.GetClass("_VZXPCBifrostAttachment")}
	})
	return _VZXPCBifrostAttachmentClass
}

// GetVZXPCBifrostAttachmentClass returns the class object for _VZXPCBifrostAttachment.
func GetVZXPCBifrostAttachmentClass() VZXPCBifrostAttachmentClass {
	return getVZXPCBifrostAttachmentClass()
}

type VZXPCBifrostAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZXPCBifrostAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZXPCBifrostAttachmentClass) Alloc() VZXPCBifrostAttachment {
	rv := objc.SendIfResponds[VZXPCBifrostAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZXPCBifrostAttachment struct {
	VZBifrostAttachment
}

// VZXPCBifrostAttachmentFromID constructs a [VZXPCBifrostAttachment] from an objc.ID.
func VZXPCBifrostAttachmentFromID(id objc.ID) VZXPCBifrostAttachment {
	return VZXPCBifrostAttachment{VZBifrostAttachment: VZBifrostAttachmentFromID(id)}
}

// Ensure VZXPCBifrostAttachment implements IVZXPCBifrostAttachment.
var _ IVZXPCBifrostAttachment = VZXPCBifrostAttachment{}

// An interface definition for the [VZXPCBifrostAttachment] class.
type IVZXPCBifrostAttachment interface {
	IVZBifrostAttachment
}

// Init initializes the instance.
func (v VZXPCBifrostAttachment) Init() VZXPCBifrostAttachment {
	rv := objc.SendIfResponds[VZXPCBifrostAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZXPCBifrostAttachment) Autorelease() VZXPCBifrostAttachment {
	rv := objc.SendIfResponds[VZXPCBifrostAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXPCBifrostAttachment creates a new VZXPCBifrostAttachment instance.
func NewVZXPCBifrostAttachment() VZXPCBifrostAttachment {
	class := getVZXPCBifrostAttachmentClass()
	rv := objc.SendIfResponds[VZXPCBifrostAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

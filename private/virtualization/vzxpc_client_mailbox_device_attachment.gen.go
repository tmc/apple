// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZXPCClientMailboxDeviceAttachment] class.
var (
	_VZXPCClientMailboxDeviceAttachmentClass     VZXPCClientMailboxDeviceAttachmentClass
	_VZXPCClientMailboxDeviceAttachmentClassOnce sync.Once
)

func getVZXPCClientMailboxDeviceAttachmentClass() VZXPCClientMailboxDeviceAttachmentClass {
	_VZXPCClientMailboxDeviceAttachmentClassOnce.Do(func() {
		_VZXPCClientMailboxDeviceAttachmentClass = VZXPCClientMailboxDeviceAttachmentClass{class: objc.GetClass("_VZXPCClientMailboxDeviceAttachment")}
	})
	return _VZXPCClientMailboxDeviceAttachmentClass
}

// GetVZXPCClientMailboxDeviceAttachmentClass returns the class object for _VZXPCClientMailboxDeviceAttachment.
func GetVZXPCClientMailboxDeviceAttachmentClass() VZXPCClientMailboxDeviceAttachmentClass {
	return getVZXPCClientMailboxDeviceAttachmentClass()
}

type VZXPCClientMailboxDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZXPCClientMailboxDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZXPCClientMailboxDeviceAttachmentClass) Alloc() VZXPCClientMailboxDeviceAttachment {
	rv := objc.SendIfResponds[VZXPCClientMailboxDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZXPCClientMailboxDeviceAttachment._initWithMailboxHandle]
type VZXPCClientMailboxDeviceAttachment struct {
	VZMailboxDeviceAttachment
}

// VZXPCClientMailboxDeviceAttachmentFromID constructs a [VZXPCClientMailboxDeviceAttachment] from an objc.ID.
func VZXPCClientMailboxDeviceAttachmentFromID(id objc.ID) VZXPCClientMailboxDeviceAttachment {
	return VZXPCClientMailboxDeviceAttachment{VZMailboxDeviceAttachment: VZMailboxDeviceAttachmentFromID(id)}
}

// Ensure VZXPCClientMailboxDeviceAttachment implements IVZXPCClientMailboxDeviceAttachment.
var _ IVZXPCClientMailboxDeviceAttachment = VZXPCClientMailboxDeviceAttachment{}

// An interface definition for the [VZXPCClientMailboxDeviceAttachment] class.
//
// # Methods
//
//   - [IVZXPCClientMailboxDeviceAttachment._initWithMailboxHandle]
type IVZXPCClientMailboxDeviceAttachment interface {
	IVZMailboxDeviceAttachment

	// Topic: Methods

	_initWithMailboxHandle(handle MailboxHandle) objectivec.IObject
}

// Init initializes the instance.
func (v VZXPCClientMailboxDeviceAttachment) Init() VZXPCClientMailboxDeviceAttachment {
	rv := objc.SendIfResponds[VZXPCClientMailboxDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZXPCClientMailboxDeviceAttachment) Autorelease() VZXPCClientMailboxDeviceAttachment {
	rv := objc.SendIfResponds[VZXPCClientMailboxDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXPCClientMailboxDeviceAttachment creates a new VZXPCClientMailboxDeviceAttachment instance.
func NewVZXPCClientMailboxDeviceAttachment() VZXPCClientMailboxDeviceAttachment {
	class := getVZXPCClientMailboxDeviceAttachmentClass()
	rv := objc.SendIfResponds[VZXPCClientMailboxDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZXPCClientMailboxDeviceAttachment) _initWithMailboxHandle(handle MailboxHandle) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_initWithMailboxHandle:"), handle)
	return objectivec.Object{ID: rv}
}

// InitWithMailboxHandle is an exported wrapper for the private method _initWithMailboxHandle.
func (v VZXPCClientMailboxDeviceAttachment) InitWithMailboxHandle(handle MailboxHandle) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithMailboxHandle:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithMailboxHandle:"}
		return nil, err
	}
	return v._initWithMailboxHandle(handle), nil
}

// CanInitWithMailboxHandle reports whether the receiver responds to the private selector _initWithMailboxHandle:.
func (v VZXPCClientMailboxDeviceAttachment) CanInitWithMailboxHandle() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithMailboxHandle:"))
}

// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// An attachment point that enables the Spice clipboard sharing capability.
//
// # Enabling clipboard sharing between the host and the VM
//
//   - [VZSpiceAgentPortAttachment.SharesClipboard]: A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
//   - [VZSpiceAgentPortAttachment.SetSharesClipboard]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment
type VZSpiceAgentPortAttachment struct {
	VZSerialPortAttachment
}

// VZSpiceAgentPortAttachmentFromID constructs a [VZSpiceAgentPortAttachment] from an objc.ID.
//
// An attachment point that enables the Spice clipboard sharing capability.
func VZSpiceAgentPortAttachmentFromID(id objc.ID) VZSpiceAgentPortAttachment {
	return VZSpiceAgentPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}
// NOTE: VZSpiceAgentPortAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZSpiceAgentPortAttachment] class.
//
// # Enabling clipboard sharing between the host and the VM
//
//   - [IVZSpiceAgentPortAttachment.SharesClipboard]: A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
//   - [IVZSpiceAgentPortAttachment.SetSharesClipboard]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment
type IVZSpiceAgentPortAttachment interface {
	IVZSerialPortAttachment

	// Topic: Enabling clipboard sharing between the host and the VM

	// A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
	SharesClipboard() bool
	SetSharesClipboard(value bool)
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

// A Boolean value that indicates whether the framework needs to share the
// clipboard between the host and the VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/sharesClipboard
func (s VZSpiceAgentPortAttachment) SharesClipboard() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("sharesClipboard"))
	return rv
}
func (s VZSpiceAgentPortAttachment) SetSharesClipboard(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSharesClipboard:"), value)
}

// The name of the Virtio console port for the Spice guest agent.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/spiceAgentPortName
func (_VZSpiceAgentPortAttachmentClass VZSpiceAgentPortAttachmentClass) SpiceAgentPortName() string {
	rv := objc.Send[objc.ID](objc.ID(_VZSpiceAgentPortAttachmentClass.class), objc.Sel("spiceAgentPortName"))
	return foundation.NSStringFromID(rv).String()
}


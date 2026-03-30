// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMailboxDeviceAttachment] class.
var (
	_VZMailboxDeviceAttachmentClass     VZMailboxDeviceAttachmentClass
	_VZMailboxDeviceAttachmentClassOnce sync.Once
)

func getVZMailboxDeviceAttachmentClass() VZMailboxDeviceAttachmentClass {
	_VZMailboxDeviceAttachmentClassOnce.Do(func() {
		_VZMailboxDeviceAttachmentClass = VZMailboxDeviceAttachmentClass{class: objc.GetClass("_VZMailboxDeviceAttachment")}
	})
	return _VZMailboxDeviceAttachmentClass
}

// GetVZMailboxDeviceAttachmentClass returns the class object for _VZMailboxDeviceAttachment.
func GetVZMailboxDeviceAttachmentClass() VZMailboxDeviceAttachmentClass {
	return getVZMailboxDeviceAttachmentClass()
}

type VZMailboxDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMailboxDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMailboxDeviceAttachmentClass) Alloc() VZMailboxDeviceAttachment {
	rv := objc.Send[VZMailboxDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMailboxDeviceAttachment._attachment]
//   - [VZMailboxDeviceAttachment._init]
//   - [VZMailboxDeviceAttachment.EncodeWithEncoder]
//   - [VZMailboxDeviceAttachment.DebugDescription]
//   - [VZMailboxDeviceAttachment.Description]
//   - [VZMailboxDeviceAttachment.Hash]
//   - [VZMailboxDeviceAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment
type VZMailboxDeviceAttachment struct {
	objectivec.Object
}

// VZMailboxDeviceAttachmentFromID constructs a [VZMailboxDeviceAttachment] from an objc.ID.
func VZMailboxDeviceAttachmentFromID(id objc.ID) VZMailboxDeviceAttachment {
	return VZMailboxDeviceAttachment{objectivec.Object{ID: id}}
}

// Ensure VZMailboxDeviceAttachment implements IVZMailboxDeviceAttachment.
var _ IVZMailboxDeviceAttachment = VZMailboxDeviceAttachment{}

// An interface definition for the [VZMailboxDeviceAttachment] class.
//
// # Methods
//
//   - [IVZMailboxDeviceAttachment._attachment]
//   - [IVZMailboxDeviceAttachment._init]
//   - [IVZMailboxDeviceAttachment.EncodeWithEncoder]
//   - [IVZMailboxDeviceAttachment.DebugDescription]
//   - [IVZMailboxDeviceAttachment.Description]
//   - [IVZMailboxDeviceAttachment.Hash]
//   - [IVZMailboxDeviceAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment
type IVZMailboxDeviceAttachment interface {
	objectivec.IObject

	// Topic: Methods

	_attachment() objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZMailboxDeviceAttachment) Init() VZMailboxDeviceAttachment {
	rv := objc.Send[VZMailboxDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMailboxDeviceAttachment) Autorelease() VZMailboxDeviceAttachment {
	rv := objc.Send[VZMailboxDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMailboxDeviceAttachment creates a new VZMailboxDeviceAttachment instance.
func NewVZMailboxDeviceAttachment() VZMailboxDeviceAttachment {
	class := getVZMailboxDeviceAttachmentClass()
	rv := objc.Send[VZMailboxDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/_init
func (v VZMailboxDeviceAttachment) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/encodeWithEncoder:
func (v VZMailboxDeviceAttachment) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/_attachment
func (v VZMailboxDeviceAttachment) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/debugDescription
func (v VZMailboxDeviceAttachment) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/description
func (v VZMailboxDeviceAttachment) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/hash
func (v VZMailboxDeviceAttachment) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxDeviceAttachment/superclass
func (v VZMailboxDeviceAttachment) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

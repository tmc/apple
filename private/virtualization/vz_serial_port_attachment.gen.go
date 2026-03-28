// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Class returns the underlying Objective-C class pointer.
func (vc VZSerialPortAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSerialPortAttachmentClass) Alloc() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZSerialPortAttachment._init]
//   - [VZSerialPortAttachment._attachment]
//   - [VZSerialPortAttachment.DebugDescription]
//   - [VZSerialPortAttachment.Description]
//   - [VZSerialPortAttachment.Hash]
//   - [VZSerialPortAttachment.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type VZSerialPortAttachment struct {
	objectivec.Object
}

// VZSerialPortAttachmentFromID constructs a [VZSerialPortAttachment] from an objc.ID.
func VZSerialPortAttachmentFromID(id objc.ID) VZSerialPortAttachment {
	return VZSerialPortAttachment{objectivec.Object{ID: id}}
}
// Ensure VZSerialPortAttachment implements IVZSerialPortAttachment.
var _ IVZSerialPortAttachment = VZSerialPortAttachment{}

// An interface definition for the [VZSerialPortAttachment] class.
//
// # Methods
//
//   - [IVZSerialPortAttachment._init]
//   - [IVZSerialPortAttachment._attachment]
//   - [IVZSerialPortAttachment.DebugDescription]
//   - [IVZSerialPortAttachment.Description]
//   - [IVZSerialPortAttachment.Hash]
//   - [IVZSerialPortAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type IVZSerialPortAttachment interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_attachment() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
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

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/_init
func (s VZSerialPortAttachment) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/_attachment
func (s VZSerialPortAttachment) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/debugDescription
func (s VZSerialPortAttachment) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/description
func (s VZSerialPortAttachment) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/hash
func (s VZSerialPortAttachment) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment/superclass
func (s VZSerialPortAttachment) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}


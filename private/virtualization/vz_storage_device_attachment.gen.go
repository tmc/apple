// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZStorageDeviceAttachment] class.
var (
	_VZStorageDeviceAttachmentClass     VZStorageDeviceAttachmentClass
	_VZStorageDeviceAttachmentClassOnce sync.Once
)

func getVZStorageDeviceAttachmentClass() VZStorageDeviceAttachmentClass {
	_VZStorageDeviceAttachmentClassOnce.Do(func() {
		_VZStorageDeviceAttachmentClass = VZStorageDeviceAttachmentClass{class: objc.GetClass("VZStorageDeviceAttachment")}
	})
	return _VZStorageDeviceAttachmentClass
}

// GetVZStorageDeviceAttachmentClass returns the class object for VZStorageDeviceAttachment.
func GetVZStorageDeviceAttachmentClass() VZStorageDeviceAttachmentClass {
	return getVZStorageDeviceAttachmentClass()
}

type VZStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZStorageDeviceAttachmentClass) Alloc() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZStorageDeviceAttachment._init]
//   - [VZStorageDeviceAttachment.DebugDescription]
//   - [VZStorageDeviceAttachment.Description]
//   - [VZStorageDeviceAttachment.Hash]
//   - [VZStorageDeviceAttachment.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment
type VZStorageDeviceAttachment struct {
	objectivec.Object
}

// VZStorageDeviceAttachmentFromID constructs a [VZStorageDeviceAttachment] from an objc.ID.
func VZStorageDeviceAttachmentFromID(id objc.ID) VZStorageDeviceAttachment {
	return VZStorageDeviceAttachment{objectivec.Object{ID: id}}
}
// Ensure VZStorageDeviceAttachment implements IVZStorageDeviceAttachment.
var _ IVZStorageDeviceAttachment = VZStorageDeviceAttachment{}

// An interface definition for the [VZStorageDeviceAttachment] class.
//
// # Methods
//
//   - [IVZStorageDeviceAttachment._init]
//   - [IVZStorageDeviceAttachment.DebugDescription]
//   - [IVZStorageDeviceAttachment.Description]
//   - [IVZStorageDeviceAttachment.Hash]
//   - [IVZStorageDeviceAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment
type IVZStorageDeviceAttachment interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s VZStorageDeviceAttachment) Init() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZStorageDeviceAttachment) Autorelease() VZStorageDeviceAttachment {
	rv := objc.Send[VZStorageDeviceAttachment](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZStorageDeviceAttachment creates a new VZStorageDeviceAttachment instance.
func NewVZStorageDeviceAttachment() VZStorageDeviceAttachment {
	class := getVZStorageDeviceAttachmentClass()
	rv := objc.Send[VZStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment/_init
func (s VZStorageDeviceAttachment) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment/debugDescription
func (s VZStorageDeviceAttachment) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment/description
func (s VZStorageDeviceAttachment) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment/hash
func (s VZStorageDeviceAttachment) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZStorageDeviceAttachment/superclass
func (s VZStorageDeviceAttachment) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}


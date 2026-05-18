// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDeviceAttachment] class.
var (
	_VZNetworkDeviceAttachmentClass     VZNetworkDeviceAttachmentClass
	_VZNetworkDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkDeviceAttachmentClass() VZNetworkDeviceAttachmentClass {
	_VZNetworkDeviceAttachmentClassOnce.Do(func() {
		_VZNetworkDeviceAttachmentClass = VZNetworkDeviceAttachmentClass{class: objc.GetClass("VZNetworkDeviceAttachment")}
	})
	return _VZNetworkDeviceAttachmentClass
}

// GetVZNetworkDeviceAttachmentClass returns the class object for VZNetworkDeviceAttachment.
func GetVZNetworkDeviceAttachmentClass() VZNetworkDeviceAttachmentClass {
	return getVZNetworkDeviceAttachmentClass()
}

type VZNetworkDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNetworkDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceAttachmentClass) Alloc() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNetworkDeviceAttachment._init]
//   - [VZNetworkDeviceAttachment._attachment]
//   - [VZNetworkDeviceAttachment.DebugDescription]
//   - [VZNetworkDeviceAttachment.Description]
//   - [VZNetworkDeviceAttachment.Hash]
//   - [VZNetworkDeviceAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment
type VZNetworkDeviceAttachment struct {
	objectivec.Object
}

// VZNetworkDeviceAttachmentFromID constructs a [VZNetworkDeviceAttachment] from an objc.ID.
func VZNetworkDeviceAttachmentFromID(id objc.ID) VZNetworkDeviceAttachment {
	return VZNetworkDeviceAttachment{objectivec.Object{ID: id}}
}

// Ensure VZNetworkDeviceAttachment implements IVZNetworkDeviceAttachment.
var _ IVZNetworkDeviceAttachment = VZNetworkDeviceAttachment{}

// An interface definition for the [VZNetworkDeviceAttachment] class.
//
// # Methods
//
//   - [IVZNetworkDeviceAttachment._init]
//   - [IVZNetworkDeviceAttachment._attachment]
//   - [IVZNetworkDeviceAttachment.DebugDescription]
//   - [IVZNetworkDeviceAttachment.Description]
//   - [IVZNetworkDeviceAttachment.Hash]
//   - [IVZNetworkDeviceAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment
type IVZNetworkDeviceAttachment interface {
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
func (v VZNetworkDeviceAttachment) Init() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNetworkDeviceAttachment) Autorelease() VZNetworkDeviceAttachment {
	rv := objc.Send[VZNetworkDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceAttachment creates a new VZNetworkDeviceAttachment instance.
func NewVZNetworkDeviceAttachment() VZNetworkDeviceAttachment {
	class := getVZNetworkDeviceAttachmentClass()
	rv := objc.Send[VZNetworkDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/_init
func (v VZNetworkDeviceAttachment) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/_attachment
func (v VZNetworkDeviceAttachment) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}

// CanAttachment reports whether the receiver responds to the private selector _attachment.
func (v VZNetworkDeviceAttachment) CanAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_attachment"))
}

// Attachment is an exported wrapper for the private property _attachment.
func (v VZNetworkDeviceAttachment) Attachment() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_attachment")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_attachment"}
	}
	return v._attachment(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/debugDescription
func (v VZNetworkDeviceAttachment) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/description
func (v VZNetworkDeviceAttachment) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/hash
func (v VZNetworkDeviceAttachment) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceAttachment/superclass
func (v VZNetworkDeviceAttachment) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

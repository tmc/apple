// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZBifrostAttachment] class.
var (
	_VZBifrostAttachmentClass     VZBifrostAttachmentClass
	_VZBifrostAttachmentClassOnce sync.Once
)

func getVZBifrostAttachmentClass() VZBifrostAttachmentClass {
	_VZBifrostAttachmentClassOnce.Do(func() {
		_VZBifrostAttachmentClass = VZBifrostAttachmentClass{class: objc.GetClass("_VZBifrostAttachment")}
	})
	return _VZBifrostAttachmentClass
}

// GetVZBifrostAttachmentClass returns the class object for _VZBifrostAttachment.
func GetVZBifrostAttachmentClass() VZBifrostAttachmentClass {
	return getVZBifrostAttachmentClass()
}

type VZBifrostAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZBifrostAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZBifrostAttachmentClass) Alloc() VZBifrostAttachment {
	rv := objc.Send[VZBifrostAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZBifrostAttachment._attachment]
//   - [VZBifrostAttachment._init]
//   - [VZBifrostAttachment.EncodeWithEncoder]
//   - [VZBifrostAttachment.DebugDescription]
//   - [VZBifrostAttachment.Description]
//   - [VZBifrostAttachment.Hash]
//   - [VZBifrostAttachment.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment
type VZBifrostAttachment struct {
	objectivec.Object
}

// VZBifrostAttachmentFromID constructs a [VZBifrostAttachment] from an objc.ID.
func VZBifrostAttachmentFromID(id objc.ID) VZBifrostAttachment {
	return VZBifrostAttachment{objectivec.Object{ID: id}}
}
// Ensure VZBifrostAttachment implements IVZBifrostAttachment.
var _ IVZBifrostAttachment = VZBifrostAttachment{}

// An interface definition for the [VZBifrostAttachment] class.
//
// # Methods
//
//   - [IVZBifrostAttachment._attachment]
//   - [IVZBifrostAttachment._init]
//   - [IVZBifrostAttachment.EncodeWithEncoder]
//   - [IVZBifrostAttachment.DebugDescription]
//   - [IVZBifrostAttachment.Description]
//   - [IVZBifrostAttachment.Hash]
//   - [IVZBifrostAttachment.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment
type IVZBifrostAttachment interface {
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
func (v VZBifrostAttachment) Init() VZBifrostAttachment {
	rv := objc.Send[VZBifrostAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZBifrostAttachment) Autorelease() VZBifrostAttachment {
	rv := objc.Send[VZBifrostAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBifrostAttachment creates a new VZBifrostAttachment instance.
func NewVZBifrostAttachment() VZBifrostAttachment {
	class := getVZBifrostAttachmentClass()
	rv := objc.Send[VZBifrostAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/_init
func (v VZBifrostAttachment) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/encodeWithEncoder:
func (v VZBifrostAttachment) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/_attachment
func (v VZBifrostAttachment) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/debugDescription
func (v VZBifrostAttachment) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/description
func (v VZBifrostAttachment) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/hash
func (v VZBifrostAttachment) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZBifrostAttachment/superclass
func (v VZBifrostAttachment) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}


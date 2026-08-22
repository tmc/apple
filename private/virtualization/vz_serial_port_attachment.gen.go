// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSerialPortAttachment._attachment]
//   - [VZSerialPortAttachment._init]
//   - [VZSerialPortAttachment.DebugDescription]
//   - [VZSerialPortAttachment.Description]
//   - [VZSerialPortAttachment.Hash]
//   - [VZSerialPortAttachment.Superclass]
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
//   - [IVZSerialPortAttachment._attachment]
//   - [IVZSerialPortAttachment._init]
//   - [IVZSerialPortAttachment.DebugDescription]
//   - [IVZSerialPortAttachment.Description]
//   - [IVZSerialPortAttachment.Hash]
//   - [IVZSerialPortAttachment.Superclass]
type IVZSerialPortAttachment interface {
	objectivec.IObject

	// Topic: Methods

	_attachment() unsafe.Pointer
	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZSerialPortAttachment) Init() VZSerialPortAttachment {
	rv := objc.SendIfResponds[VZSerialPortAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSerialPortAttachment) Autorelease() VZSerialPortAttachment {
	rv := objc.SendIfResponds[VZSerialPortAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortAttachment creates a new VZSerialPortAttachment instance.
func NewVZSerialPortAttachment() VZSerialPortAttachment {
	class := getVZSerialPortAttachmentClass()
	rv := objc.SendIfResponds[VZSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZSerialPortAttachment) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZSerialPortAttachment) _attachment() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_attachment"))
	return rv
}

// CanAttachment reports whether the receiver responds to the private selector _attachment.
func (v VZSerialPortAttachment) CanAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_attachment"))
}

// Attachment is an exported wrapper for the private property _attachment.
func (v VZSerialPortAttachment) Attachment() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_attachment")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_attachment"}
	}
	return v._attachment(), nil
}
func (v VZSerialPortAttachment) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSerialPortAttachment) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZSerialPortAttachment) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZSerialPortAttachment) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

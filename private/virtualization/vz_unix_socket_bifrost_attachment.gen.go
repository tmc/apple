// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUnixSocketBifrostAttachment] class.
var (
	_VZUnixSocketBifrostAttachmentClass     VZUnixSocketBifrostAttachmentClass
	_VZUnixSocketBifrostAttachmentClassOnce sync.Once
)

func getVZUnixSocketBifrostAttachmentClass() VZUnixSocketBifrostAttachmentClass {
	_VZUnixSocketBifrostAttachmentClassOnce.Do(func() {
		_VZUnixSocketBifrostAttachmentClass = VZUnixSocketBifrostAttachmentClass{class: objc.GetClass("_VZUnixSocketBifrostAttachment")}
	})
	return _VZUnixSocketBifrostAttachmentClass
}

// GetVZUnixSocketBifrostAttachmentClass returns the class object for _VZUnixSocketBifrostAttachment.
func GetVZUnixSocketBifrostAttachmentClass() VZUnixSocketBifrostAttachmentClass {
	return getVZUnixSocketBifrostAttachmentClass()
}

type VZUnixSocketBifrostAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUnixSocketBifrostAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUnixSocketBifrostAttachmentClass) Alloc() VZUnixSocketBifrostAttachment {
	rv := objc.Send[VZUnixSocketBifrostAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUnixSocketBifrostAttachment.Path]
//   - [VZUnixSocketBifrostAttachment.InitWithPathError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment
type VZUnixSocketBifrostAttachment struct {
	VZBifrostAttachment
}

// VZUnixSocketBifrostAttachmentFromID constructs a [VZUnixSocketBifrostAttachment] from an objc.ID.
func VZUnixSocketBifrostAttachmentFromID(id objc.ID) VZUnixSocketBifrostAttachment {
	return VZUnixSocketBifrostAttachment{VZBifrostAttachment: VZBifrostAttachmentFromID(id)}
}

// Ensure VZUnixSocketBifrostAttachment implements IVZUnixSocketBifrostAttachment.
var _ IVZUnixSocketBifrostAttachment = VZUnixSocketBifrostAttachment{}

// An interface definition for the [VZUnixSocketBifrostAttachment] class.
//
// # Methods
//
//   - [IVZUnixSocketBifrostAttachment.Path]
//   - [IVZUnixSocketBifrostAttachment.InitWithPathError]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment
type IVZUnixSocketBifrostAttachment interface {
	IVZBifrostAttachment

	// Topic: Methods

	Path() string
	InitWithPathError(path objectivec.IObject) (VZUnixSocketBifrostAttachment, error)
}

// Init initializes the instance.
func (v VZUnixSocketBifrostAttachment) Init() VZUnixSocketBifrostAttachment {
	rv := objc.Send[VZUnixSocketBifrostAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUnixSocketBifrostAttachment) Autorelease() VZUnixSocketBifrostAttachment {
	rv := objc.Send[VZUnixSocketBifrostAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUnixSocketBifrostAttachment creates a new VZUnixSocketBifrostAttachment instance.
func NewVZUnixSocketBifrostAttachment() VZUnixSocketBifrostAttachment {
	class := getVZUnixSocketBifrostAttachmentClass()
	rv := objc.Send[VZUnixSocketBifrostAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment/initWithPath:error:
func NewVZUnixSocketBifrostAttachmentWithPathError(path objectivec.IObject) (VZUnixSocketBifrostAttachment, error) {
	var errorPtr objc.ID
	instance := getVZUnixSocketBifrostAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZUnixSocketBifrostAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZUnixSocketBifrostAttachmentFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment/initWithPath:error:
func (v VZUnixSocketBifrostAttachment) InitWithPathError(path objectivec.IObject) (VZUnixSocketBifrostAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZUnixSocketBifrostAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZUnixSocketBifrostAttachmentFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment/maximumPathLength
func (_VZUnixSocketBifrostAttachmentClass VZUnixSocketBifrostAttachmentClass) MaximumPathLength() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZUnixSocketBifrostAttachmentClass.class), objc.Sel("maximumPathLength"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUnixSocketBifrostAttachment/path
func (v VZUnixSocketBifrostAttachment) Path() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}

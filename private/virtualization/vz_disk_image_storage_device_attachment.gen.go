// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDiskImageStorageDeviceAttachment] class.
var (
	_VZDiskImageStorageDeviceAttachmentClass     VZDiskImageStorageDeviceAttachmentClass
	_VZDiskImageStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskImageStorageDeviceAttachmentClass() VZDiskImageStorageDeviceAttachmentClass {
	_VZDiskImageStorageDeviceAttachmentClassOnce.Do(func() {
		_VZDiskImageStorageDeviceAttachmentClass = VZDiskImageStorageDeviceAttachmentClass{class: objc.GetClass("VZDiskImageStorageDeviceAttachment")}
	})
	return _VZDiskImageStorageDeviceAttachmentClass
}

// GetVZDiskImageStorageDeviceAttachmentClass returns the class object for VZDiskImageStorageDeviceAttachment.
func GetVZDiskImageStorageDeviceAttachmentClass() VZDiskImageStorageDeviceAttachmentClass {
	return getVZDiskImageStorageDeviceAttachmentClass()
}

type VZDiskImageStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskImageStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskImageStorageDeviceAttachmentClass) Alloc() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDiskImageStorageDeviceAttachment._updateDiskSize]
//   - [VZDiskImageStorageDeviceAttachment.ReadOnly]
type VZDiskImageStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskImageStorageDeviceAttachmentFromID constructs a [VZDiskImageStorageDeviceAttachment] from an objc.ID.
func VZDiskImageStorageDeviceAttachmentFromID(id objc.ID) VZDiskImageStorageDeviceAttachment {
	return VZDiskImageStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}

// Ensure VZDiskImageStorageDeviceAttachment implements IVZDiskImageStorageDeviceAttachment.
var _ IVZDiskImageStorageDeviceAttachment = VZDiskImageStorageDeviceAttachment{}

// An interface definition for the [VZDiskImageStorageDeviceAttachment] class.
//
// # Methods
//
//   - [IVZDiskImageStorageDeviceAttachment._updateDiskSize]
//   - [IVZDiskImageStorageDeviceAttachment.ReadOnly]
type IVZDiskImageStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Methods

	_updateDiskSize(size uint64)
	ReadOnly() bool
}

// Init initializes the instance.
func (v VZDiskImageStorageDeviceAttachment) Init() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDiskImageStorageDeviceAttachment) Autorelease() VZDiskImageStorageDeviceAttachment {
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImageStorageDeviceAttachment creates a new VZDiskImageStorageDeviceAttachment instance.
func NewVZDiskImageStorageDeviceAttachment() VZDiskImageStorageDeviceAttachment {
	class := getVZDiskImageStorageDeviceAttachmentClass()
	rv := objc.Send[VZDiskImageStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZDiskImageStorageDeviceAttachment) _updateDiskSize(size uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_updateDiskSize:"), size)
}

// UpdateDiskSize is an exported wrapper for the private method _updateDiskSize.
func (v VZDiskImageStorageDeviceAttachment) UpdateDiskSize(size uint64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_updateDiskSize:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_updateDiskSize:"}
		return err
	}
	v._updateDiskSize(size)
	return nil
}

// CanUpdateDiskSize reports whether the receiver responds to the private selector _updateDiskSize:.
func (v VZDiskImageStorageDeviceAttachment) CanUpdateDiskSize() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_updateDiskSize:"))
}

func (_VZDiskImageStorageDeviceAttachmentClass VZDiskImageStorageDeviceAttachmentClass) _diskImageStorageDeviceAttachmentWithDiskImage(image objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VZDiskImageStorageDeviceAttachmentClass.class), objc.Sel("_diskImageStorageDeviceAttachmentWithDiskImage:"), image)
	return objectivec.Object{ID: rv}
}

// DiskImageStorageDeviceAttachmentWithDiskImage is an exported wrapper for the private method _diskImageStorageDeviceAttachmentWithDiskImage.
func (_VZDiskImageStorageDeviceAttachmentClass VZDiskImageStorageDeviceAttachmentClass) DiskImageStorageDeviceAttachmentWithDiskImage(image objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_VZDiskImageStorageDeviceAttachmentClass.class), objc.Sel("_diskImageStorageDeviceAttachmentWithDiskImage:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_diskImageStorageDeviceAttachmentWithDiskImage:"}
		return nil, err
	}
	return _VZDiskImageStorageDeviceAttachmentClass._diskImageStorageDeviceAttachmentWithDiskImage(image), nil
}

// CanDiskImageStorageDeviceAttachmentWithDiskImage reports whether the receiver responds to the private selector _diskImageStorageDeviceAttachmentWithDiskImage:.
func (_VZDiskImageStorageDeviceAttachmentClass VZDiskImageStorageDeviceAttachmentClass) CanDiskImageStorageDeviceAttachmentWithDiskImage() bool {
	return objc.RespondsToSelector(objc.ID(_VZDiskImageStorageDeviceAttachmentClass.class), objc.Sel("_diskImageStorageDeviceAttachmentWithDiskImage:"))
}

func (v VZDiskImageStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("readOnly"))
	return rv
}

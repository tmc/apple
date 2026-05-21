// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
var (
	_VZNetworkBlockDeviceStorageDeviceAttachmentClass     VZNetworkBlockDeviceStorageDeviceAttachmentClass
	_VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZNetworkBlockDeviceStorageDeviceAttachmentClass() VZNetworkBlockDeviceStorageDeviceAttachmentClass {
	_VZNetworkBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		_VZNetworkBlockDeviceStorageDeviceAttachmentClass = VZNetworkBlockDeviceStorageDeviceAttachmentClass{class: objc.GetClass("VZNetworkBlockDeviceStorageDeviceAttachment")}
	})
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass
}

// GetVZNetworkBlockDeviceStorageDeviceAttachmentClass returns the class object for VZNetworkBlockDeviceStorageDeviceAttachment.
func GetVZNetworkBlockDeviceStorageDeviceAttachmentClass() VZNetworkBlockDeviceStorageDeviceAttachmentClass {
	return getVZNetworkBlockDeviceStorageDeviceAttachmentClass()
}

type VZNetworkBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNetworkBlockDeviceStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkBlockDeviceStorageDeviceAttachmentClass) Alloc() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.ForcedReadOnly]
type VZNetworkBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZNetworkBlockDeviceStorageDeviceAttachmentFromID constructs a [VZNetworkBlockDeviceStorageDeviceAttachment] from an objc.ID.
func VZNetworkBlockDeviceStorageDeviceAttachmentFromID(id objc.ID) VZNetworkBlockDeviceStorageDeviceAttachment {
	return VZNetworkBlockDeviceStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}

// Ensure VZNetworkBlockDeviceStorageDeviceAttachment implements IVZNetworkBlockDeviceStorageDeviceAttachment.
var _ IVZNetworkBlockDeviceStorageDeviceAttachment = VZNetworkBlockDeviceStorageDeviceAttachment{}

// An interface definition for the [VZNetworkBlockDeviceStorageDeviceAttachment] class.
//
// # Methods
//
//   - [IVZNetworkBlockDeviceStorageDeviceAttachment.ForcedReadOnly]
type IVZNetworkBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Methods

	ForcedReadOnly() bool
}

// Init initializes the instance.
func (v VZNetworkBlockDeviceStorageDeviceAttachment) Init() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNetworkBlockDeviceStorageDeviceAttachment) Autorelease() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkBlockDeviceStorageDeviceAttachment creates a new VZNetworkBlockDeviceStorageDeviceAttachment instance.
func NewVZNetworkBlockDeviceStorageDeviceAttachment() VZNetworkBlockDeviceStorageDeviceAttachment {
	class := getVZNetworkBlockDeviceStorageDeviceAttachmentClass()
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultReadOnly() bool {
	rv := objc.Send[bool](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultReadOnly"))
	return rv
}

// DefaultReadOnly is an exported wrapper for the private method _defaultReadOnly.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultReadOnly() (bool, error) {
	if !objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultReadOnly")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_defaultReadOnly"}
		return false, err
	}
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultReadOnly(), nil
}

// CanDefaultReadOnly reports whether the receiver responds to the private selector _defaultReadOnly.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) CanDefaultReadOnly() bool {
	return objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultReadOnly"))
}
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultSynchronizationMode() int64 {
	rv := objc.Send[int64](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultSynchronizationMode"))
	return rv
}

// DefaultSynchronizationMode is an exported wrapper for the private method _defaultSynchronizationMode.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultSynchronizationMode() (int64, error) {
	if !objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultSynchronizationMode")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_defaultSynchronizationMode"}
		return 0, err
	}
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultSynchronizationMode(), nil
}

// CanDefaultSynchronizationMode reports whether the receiver responds to the private selector _defaultSynchronizationMode.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) CanDefaultSynchronizationMode() bool {
	return objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultSynchronizationMode"))
}
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultTimeout() float64 {
	rv := objc.Send[float64](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultTimeout"))
	return rv
}

// DefaultTimeout is an exported wrapper for the private method _defaultTimeout.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultTimeout() (float64, error) {
	if !objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultTimeout")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_defaultTimeout"}
		return 0.0, err
	}
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultTimeout(), nil
}

// CanDefaultTimeout reports whether the receiver responds to the private selector _defaultTimeout.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) CanDefaultTimeout() bool {
	return objc.RespondsToSelector(objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultTimeout"))
}

func (v VZNetworkBlockDeviceStorageDeviceAttachment) ForcedReadOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("forcedReadOnly"))
	return rv
}

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

//
// # Methods
//
//   - [VZNetworkBlockDeviceStorageDeviceAttachment.ForcedReadOnly]
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment
type IVZNetworkBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Methods

	ForcedReadOnly() bool
}

// Init initializes the instance.
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Init() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkBlockDeviceStorageDeviceAttachment) Autorelease() VZNetworkBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkBlockDeviceStorageDeviceAttachment creates a new VZNetworkBlockDeviceStorageDeviceAttachment instance.
func NewVZNetworkBlockDeviceStorageDeviceAttachment() VZNetworkBlockDeviceStorageDeviceAttachment {
	class := getVZNetworkBlockDeviceStorageDeviceAttachmentClass()
	rv := objc.Send[VZNetworkBlockDeviceStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/_defaultReadOnly
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultReadOnly() bool {
	rv := objc.Send[bool](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultReadOnly"))
	return rv
}

// DefaultReadOnly is an exported wrapper for the private method _defaultReadOnly.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultReadOnly() bool {
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultReadOnly()
}
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/_defaultSynchronizationMode
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultSynchronizationMode() int64 {
	rv := objc.Send[int64](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultSynchronizationMode"))
	return rv
}

// DefaultSynchronizationMode is an exported wrapper for the private method _defaultSynchronizationMode.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultSynchronizationMode() int64 {
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultSynchronizationMode()
}
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/_defaultTimeout
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) _defaultTimeout() float64 {
	rv := objc.Send[float64](objc.ID(_VZNetworkBlockDeviceStorageDeviceAttachmentClass.class), objc.Sel("_defaultTimeout"))
	return rv
}

// DefaultTimeout is an exported wrapper for the private method _defaultTimeout.
func (_VZNetworkBlockDeviceStorageDeviceAttachmentClass VZNetworkBlockDeviceStorageDeviceAttachmentClass) DefaultTimeout() float64 {
	return _VZNetworkBlockDeviceStorageDeviceAttachmentClass._defaultTimeout()
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkBlockDeviceStorageDeviceAttachment/forcedReadOnly
func (n VZNetworkBlockDeviceStorageDeviceAttachment) ForcedReadOnly() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("forcedReadOnly"))
	return rv
}


// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
var (
	_VZDiskBlockDeviceStorageDeviceAttachmentClass     VZDiskBlockDeviceStorageDeviceAttachmentClass
	_VZDiskBlockDeviceStorageDeviceAttachmentClassOnce sync.Once
)

func getVZDiskBlockDeviceStorageDeviceAttachmentClass() VZDiskBlockDeviceStorageDeviceAttachmentClass {
	_VZDiskBlockDeviceStorageDeviceAttachmentClassOnce.Do(func() {
		_VZDiskBlockDeviceStorageDeviceAttachmentClass = VZDiskBlockDeviceStorageDeviceAttachmentClass{class: objc.GetClass("VZDiskBlockDeviceStorageDeviceAttachment")}
	})
	return _VZDiskBlockDeviceStorageDeviceAttachmentClass
}

// GetVZDiskBlockDeviceStorageDeviceAttachmentClass returns the class object for VZDiskBlockDeviceStorageDeviceAttachment.
func GetVZDiskBlockDeviceStorageDeviceAttachmentClass() VZDiskBlockDeviceStorageDeviceAttachmentClass {
	return getVZDiskBlockDeviceStorageDeviceAttachmentClass()
}

type VZDiskBlockDeviceStorageDeviceAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskBlockDeviceStorageDeviceAttachmentClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskBlockDeviceStorageDeviceAttachmentClass) Alloc() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDiskBlockDeviceStorageDeviceAttachment._initWithURLReadOnlySynchronizationModeError]
//   - [VZDiskBlockDeviceStorageDeviceAttachment._url]
//   - [VZDiskBlockDeviceStorageDeviceAttachment.ReadOnly]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment
type VZDiskBlockDeviceStorageDeviceAttachment struct {
	VZStorageDeviceAttachment
}

// VZDiskBlockDeviceStorageDeviceAttachmentFromID constructs a [VZDiskBlockDeviceStorageDeviceAttachment] from an objc.ID.
func VZDiskBlockDeviceStorageDeviceAttachmentFromID(id objc.ID) VZDiskBlockDeviceStorageDeviceAttachment {
	return VZDiskBlockDeviceStorageDeviceAttachment{VZStorageDeviceAttachment: VZStorageDeviceAttachmentFromID(id)}
}

// Ensure VZDiskBlockDeviceStorageDeviceAttachment implements IVZDiskBlockDeviceStorageDeviceAttachment.
var _ IVZDiskBlockDeviceStorageDeviceAttachment = VZDiskBlockDeviceStorageDeviceAttachment{}

// An interface definition for the [VZDiskBlockDeviceStorageDeviceAttachment] class.
//
// # Methods
//
//   - [IVZDiskBlockDeviceStorageDeviceAttachment._initWithURLReadOnlySynchronizationModeError]
//   - [IVZDiskBlockDeviceStorageDeviceAttachment._url]
//   - [IVZDiskBlockDeviceStorageDeviceAttachment.ReadOnly]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment
type IVZDiskBlockDeviceStorageDeviceAttachment interface {
	IVZStorageDeviceAttachment

	// Topic: Methods

	_initWithURLReadOnlySynchronizationModeError(url foundation.INSURL, only bool, mode int64) (objectivec.IObject, error)
	_url() foundation.INSURL
	ReadOnly() bool
}

// Init initializes the instance.
func (d VZDiskBlockDeviceStorageDeviceAttachment) Init() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDiskBlockDeviceStorageDeviceAttachment) Autorelease() VZDiskBlockDeviceStorageDeviceAttachment {
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskBlockDeviceStorageDeviceAttachment creates a new VZDiskBlockDeviceStorageDeviceAttachment instance.
func NewVZDiskBlockDeviceStorageDeviceAttachment() VZDiskBlockDeviceStorageDeviceAttachment {
	class := getVZDiskBlockDeviceStorageDeviceAttachmentClass()
	rv := objc.Send[VZDiskBlockDeviceStorageDeviceAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/_initWithURL:readOnly:synchronizationMode:error:
func (d VZDiskBlockDeviceStorageDeviceAttachment) _initWithURLReadOnlySynchronizationModeError(url foundation.INSURL, only bool, mode int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_initWithURL:readOnly:synchronizationMode:error:"), url, only, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// InitWithURLReadOnlySynchronizationModeError is an exported wrapper for the private method _initWithURLReadOnlySynchronizationModeError.
func (d VZDiskBlockDeviceStorageDeviceAttachment) InitWithURLReadOnlySynchronizationModeError(url foundation.INSURL, only bool, mode int64) (objectivec.IObject, error) {
	return d._initWithURLReadOnlySynchronizationModeError(url, only, mode)
}

// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/_url
func (d VZDiskBlockDeviceStorageDeviceAttachment) _url() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_url"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/VZDiskBlockDeviceStorageDeviceAttachment/readOnly
func (d VZDiskBlockDeviceStorageDeviceAttachment) ReadOnly() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("readOnly"))
	return rv
}

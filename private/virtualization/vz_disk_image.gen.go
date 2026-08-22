// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDiskImage] class.
var (
	_VZDiskImageClass     VZDiskImageClass
	_VZDiskImageClassOnce sync.Once
)

func getVZDiskImageClass() VZDiskImageClass {
	_VZDiskImageClassOnce.Do(func() {
		_VZDiskImageClass = VZDiskImageClass{class: objc.GetClass("_VZDiskImage")}
	})
	return _VZDiskImageClass
}

// GetVZDiskImageClass returns the class object for _VZDiskImage.
func GetVZDiskImageClass() VZDiskImageClass {
	return getVZDiskImageClass()
}

type VZDiskImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDiskImageClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDiskImageClass) Alloc() VZDiskImage {
	rv := objc.SendIfResponds[VZDiskImage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDiskImage.URL]
//   - [VZDiskImage.CachingMode]
//   - [VZDiskImage.Identifier]
//   - [VZDiskImage.IsReadOnly]
//   - [VZDiskImage.SynchronizationMode]
//   - [VZDiskImage.UpdateDiskSize]
//   - [VZDiskImage.SetUpdateDiskSize]
//   - [VZDiskImage.InitWithDescriptorError]
//   - [VZDiskImage.InitWithURLReadOnlyError]
//   - [VZDiskImage.ReadOnly]
type VZDiskImage struct {
	objectivec.Object
}

// VZDiskImageFromID constructs a [VZDiskImage] from an objc.ID.
func VZDiskImageFromID(id objc.ID) VZDiskImage {
	return VZDiskImage{objectivec.Object{ID: id}}
}

// Ensure VZDiskImage implements IVZDiskImage.
var _ IVZDiskImage = VZDiskImage{}

// An interface definition for the [VZDiskImage] class.
//
// # Methods
//
//   - [IVZDiskImage.URL]
//   - [IVZDiskImage.CachingMode]
//   - [IVZDiskImage.Identifier]
//   - [IVZDiskImage.IsReadOnly]
//   - [IVZDiskImage.SynchronizationMode]
//   - [IVZDiskImage.UpdateDiskSize]
//   - [IVZDiskImage.SetUpdateDiskSize]
//   - [IVZDiskImage.InitWithDescriptorError]
//   - [IVZDiskImage.InitWithURLReadOnlyError]
//   - [IVZDiskImage.ReadOnly]
type IVZDiskImage interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.NSURL
	CachingMode() int64
	Identifier() string
	IsReadOnly() bool
	SynchronizationMode() int64
	UpdateDiskSize() foundation.NSNumber
	SetUpdateDiskSize(value foundation.NSNumber)
	InitWithDescriptorError(descriptor objectivec.IObject) (VZDiskImage, error)
	InitWithURLReadOnlyError(url foundation.NSURL, only bool) (VZDiskImage, error)
	ReadOnly() bool
}

// Init initializes the instance.
func (v VZDiskImage) Init() VZDiskImage {
	rv := objc.SendIfResponds[VZDiskImage](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDiskImage) Autorelease() VZDiskImage {
	rv := objc.SendIfResponds[VZDiskImage](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDiskImage creates a new VZDiskImage instance.
func NewVZDiskImage() VZDiskImage {
	class := getVZDiskImageClass()
	rv := objc.SendIfResponds[VZDiskImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZDiskImageWithDescriptorError(descriptor objectivec.IObject) (VZDiskImage, error) {
	var errorPtr objc.ID
	instance := getVZDiskImageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImage{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZDiskImage{}, objc.ErrInitFailed
	}
	return VZDiskImageFromID(rv), nil
}

func NewVZDiskImageWithURLReadOnlyError(url foundation.NSURL, only bool) (VZDiskImage, error) {
	var errorPtr objc.ID
	instance := getVZDiskImageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:readOnly:error:"), url, only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZDiskImage{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZDiskImage{}, objc.ErrInitFailed
	}
	return VZDiskImageFromID(rv), nil
}

func (v VZDiskImage) IsReadOnly() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isReadOnly"))
	return rv
}
func (v VZDiskImage) InitWithDescriptorError(descriptor objectivec.IObject) (VZDiskImage, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZDiskImage), foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageFromID(rv), nil

}
func (v VZDiskImage) InitWithURLReadOnlyError(url foundation.NSURL, only bool) (VZDiskImage, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithURL:readOnly:error:"), url, only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZDiskImage), foundation.NSErrorFrom(errorPtr)
	}
	return VZDiskImageFromID(rv), nil

}

func (v VZDiskImage) URL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (v VZDiskImage) CachingMode() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("cachingMode"))
	return rv
}
func (v VZDiskImage) Identifier() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZDiskImage) ReadOnly() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("readOnly"))
	return rv
}
func (v VZDiskImage) SynchronizationMode() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("synchronizationMode"))
	return rv
}
func (v VZDiskImage) UpdateDiskSize() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("updateDiskSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v VZDiskImage) SetUpdateDiskSize(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUpdateDiskSize:"), value)
}

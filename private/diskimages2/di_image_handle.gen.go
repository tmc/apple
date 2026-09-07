// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIImageHandle] class.
var (
	_DIImageHandleClass     DIImageHandleClass
	_DIImageHandleClassOnce sync.Once
)

func getDIImageHandleClass() DIImageHandleClass {
	_DIImageHandleClassOnce.Do(func() {
		_DIImageHandleClass = DIImageHandleClass{class: objc.GetClass("DIImageHandle")}
	})
	return _DIImageHandleClass
}

// GetDIImageHandleClass returns the class object for DIImageHandle.
func GetDIImageHandleClass() DIImageHandleClass {
	return getDIImageHandleClass()
}

type DIImageHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIImageHandleClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIImageHandleClass) Alloc() DIImageHandle {
	rv := objc.SendIfResponds[DIImageHandle](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIImageHandle.BlockCount]
//   - [DIImageHandle.BlockSize]
//   - [DIImageHandle.CacheImage]
//   - [DIImageHandle.FileDescriptor]
//   - [DIImageHandle.ImageFormat]
//   - [DIImageHandle.OpenMode]
//   - [DIImageHandle.ParentUUID]
//   - [DIImageHandle.ResizeToSizeError]
//   - [DIImageHandle.StackableUUID]
//   - [DIImageHandle.InitWithDiskImageLockableResources]
//   - [DIImageHandle.InitWithOpenParamsLockableResources]
type DIImageHandle struct {
	objectivec.Object
}

// DIImageHandleFromID constructs a [DIImageHandle] from an objc.ID.
func DIImageHandleFromID(id objc.ID) DIImageHandle {
	return DIImageHandle{objectivec.Object{ID: id}}
}

// Ensure DIImageHandle implements IDIImageHandle.
var _ IDIImageHandle = DIImageHandle{}

// An interface definition for the [DIImageHandle] class.
//
// # Methods
//
//   - [IDIImageHandle.BlockCount]
//   - [IDIImageHandle.BlockSize]
//   - [IDIImageHandle.CacheImage]
//   - [IDIImageHandle.FileDescriptor]
//   - [IDIImageHandle.ImageFormat]
//   - [IDIImageHandle.OpenMode]
//   - [IDIImageHandle.ParentUUID]
//   - [IDIImageHandle.ResizeToSizeError]
//   - [IDIImageHandle.StackableUUID]
//   - [IDIImageHandle.InitWithDiskImageLockableResources]
//   - [IDIImageHandle.InitWithOpenParamsLockableResources]
type IDIImageHandle interface {
	objectivec.IObject

	// Topic: Methods

	BlockCount() uint64
	BlockSize() uint64
	CacheImage() bool
	FileDescriptor() int32
	ImageFormat() int64
	OpenMode() int64
	ParentUUID() foundation.NSUUID
	ResizeToSizeError(size uint64) (bool, error)
	StackableUUID() foundation.NSUUID
	InitWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle
	InitWithOpenParamsLockableResources(params *DiskimageOpenParams, resources unsafe.Pointer) DIImageHandle
}

// Init initializes the instance.
func (d DIImageHandle) Init() DIImageHandle {
	rv := objc.SendIfResponds[DIImageHandle](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIImageHandle) Autorelease() DIImageHandle {
	rv := objc.SendIfResponds[DIImageHandle](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIImageHandle creates a new DIImageHandle instance.
func NewDIImageHandle() DIImageHandle {
	class := getDIImageHandleClass()
	rv := objc.SendIfResponds[DIImageHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIImageHandleWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle {
	instance := getDIImageHandleClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDiskImage:lockableResources:"), image, resources)
	return DIImageHandleFromID(rv)
}

func NewDIImageHandleWithOpenParamsLockableResources(params *DiskimageOpenParams, resources unsafe.Pointer) DIImageHandle {
	instance := getDIImageHandleClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOpenParams:lockableResources:"), params, resources)
	return DIImageHandleFromID(rv)
}

func (d DIImageHandle) ResizeToSizeError(size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeToSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeToSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIImageHandle) InitWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle {
	rv := objc.SendIfResponds[DIImageHandle](d.ID, objc.Sel("initWithDiskImage:lockableResources:"), image, resources)
	return rv
}
func (d DIImageHandle) InitWithOpenParamsLockableResources(params *DiskimageOpenParams, resources unsafe.Pointer) DIImageHandle {
	rv := objc.SendIfResponds[DIImageHandle](d.ID, objc.Sel("initWithOpenParams:lockableResources:"), params, resources)
	return rv
}

func (d DIImageHandle) BlockCount() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("blockCount"))
	return rv
}
func (d DIImageHandle) BlockSize() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("blockSize"))
	return rv
}
func (d DIImageHandle) CacheImage() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("cacheImage"))
	return rv
}
func (d DIImageHandle) FileDescriptor() int32 {
	rv := objc.SendIfResponds[int32](d.ID, objc.Sel("fileDescriptor"))
	return rv
}
func (d DIImageHandle) ImageFormat() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("imageFormat"))
	return rv
}
func (d DIImageHandle) OpenMode() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("openMode"))
	return rv
}
func (d DIImageHandle) ParentUUID() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](d.ID, objc.Sel("parentUUID"))
	return foundation.NSUUID(rv)
}
func (d DIImageHandle) StackableUUID() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](d.ID, objc.Sel("stackableUUID"))
	return foundation.NSUUID(rv)
}

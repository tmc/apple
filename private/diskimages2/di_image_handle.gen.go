// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
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
	rv := objc.Send[DIImageHandle](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIImageHandle.MoveDiskImage]
//   - [DIImageHandle.InitWithDiskImageLockableResources]
// See: https://developer.apple.com/documentation/DiskImages2/DIImageHandle
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
//   - [IDIImageHandle.MoveDiskImage]
//   - [IDIImageHandle.InitWithDiskImageLockableResources]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageHandle
type IDIImageHandle interface {
	objectivec.IObject

	// Topic: Methods

	MoveDiskImage() objectivec.IObject
	InitWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle
}

// Init initializes the instance.
func (d DIImageHandle) Init() DIImageHandle {
	rv := objc.Send[DIImageHandle](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIImageHandle) Autorelease() DIImageHandle {
	rv := objc.Send[DIImageHandle](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIImageHandle creates a new DIImageHandle instance.
func NewDIImageHandle() DIImageHandle {
	class := getDIImageHandleClass()
	rv := objc.Send[DIImageHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageHandle/initWithDiskImage:lockableResources:
func NewDIImageHandleWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle {
	instance := getDIImageHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDiskImage:lockableResources:"), image, resources)
	return DIImageHandleFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIImageHandle/moveDiskImage
func (d DIImageHandle) MoveDiskImage() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("moveDiskImage"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIImageHandle/initWithDiskImage:lockableResources:
func (d DIImageHandle) InitWithDiskImageLockableResources(image unsafe.Pointer, resources unsafe.Pointer) DIImageHandle {
	rv := objc.Send[DIImageHandle](d.ID, objc.Sel("initWithDiskImage:lockableResources:"), image, resources)
	return rv
}


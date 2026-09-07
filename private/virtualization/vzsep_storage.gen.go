// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZSEPStorage] class.
var (
	_VZSEPStorageClass     VZSEPStorageClass
	_VZSEPStorageClassOnce sync.Once
)

func getVZSEPStorageClass() VZSEPStorageClass {
	_VZSEPStorageClassOnce.Do(func() {
		_VZSEPStorageClass = VZSEPStorageClass{class: objc.GetClass("_VZSEPStorage")}
	})
	return _VZSEPStorageClass
}

// GetVZSEPStorageClass returns the class object for _VZSEPStorage.
func GetVZSEPStorageClass() VZSEPStorageClass {
	return getVZSEPStorageClass()
}

type VZSEPStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSEPStorageClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSEPStorageClass) Alloc() VZSEPStorage {
	rv := objc.SendIfResponds[VZSEPStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZSEPStorage struct {
	VZCoprocessorStorage
}

// VZSEPStorageFromID constructs a [VZSEPStorage] from an objc.ID.
func VZSEPStorageFromID(id objc.ID) VZSEPStorage {
	return VZSEPStorage{VZCoprocessorStorage: VZCoprocessorStorageFromID(id)}
}

// Ensure VZSEPStorage implements IVZSEPStorage.
var _ IVZSEPStorage = VZSEPStorage{}

// An interface definition for the [VZSEPStorage] class.
type IVZSEPStorage interface {
	IVZCoprocessorStorage
}

// Init initializes the instance.
func (v VZSEPStorage) Init() VZSEPStorage {
	rv := objc.SendIfResponds[VZSEPStorage](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSEPStorage) Autorelease() VZSEPStorage {
	rv := objc.SendIfResponds[VZSEPStorage](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSEPStorage creates a new VZSEPStorage instance.
func NewVZSEPStorage() VZSEPStorage {
	class := getVZSEPStorageClass()
	rv := objc.SendIfResponds[VZSEPStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZSEPStorageCreatingStorageAtURLError(url foundation.NSURL) (VZSEPStorage, error) {
	var errorPtr objc.ID
	instance := getVZSEPStorageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initCreatingStorageAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZSEPStorage{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZSEPStorage{}, objc.ErrInitFailed
	}
	return VZSEPStorageFromID(rv), nil
}

func NewVZSEPStorageWithURL(url foundation.NSURL) VZSEPStorage {
	instance := getVZSEPStorageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return VZSEPStorageFromID(rv)
}

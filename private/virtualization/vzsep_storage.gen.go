// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	rv := objc.Send[VZSEPStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSEPStorage.URL]
//   - [VZSEPStorage.InitCreatingStorageAtURLError]
//   - [VZSEPStorage.InitWithURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage
type VZSEPStorage struct {
	objectivec.Object
}

// VZSEPStorageFromID constructs a [VZSEPStorage] from an objc.ID.
func VZSEPStorageFromID(id objc.ID) VZSEPStorage {
	return VZSEPStorage{objectivec.Object{ID: id}}
}

// Ensure VZSEPStorage implements IVZSEPStorage.
var _ IVZSEPStorage = VZSEPStorage{}

// An interface definition for the [VZSEPStorage] class.
//
// # Methods
//
//   - [IVZSEPStorage.URL]
//   - [IVZSEPStorage.InitCreatingStorageAtURLError]
//   - [IVZSEPStorage.InitWithURL]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage
type IVZSEPStorage interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.INSURL
	InitCreatingStorageAtURLError(url foundation.INSURL) (VZSEPStorage, error)
	InitWithURL(url foundation.INSURL) VZSEPStorage
}

// Init initializes the instance.
func (v VZSEPStorage) Init() VZSEPStorage {
	rv := objc.Send[VZSEPStorage](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSEPStorage) Autorelease() VZSEPStorage {
	rv := objc.Send[VZSEPStorage](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSEPStorage creates a new VZSEPStorage instance.
func NewVZSEPStorage() VZSEPStorage {
	class := getVZSEPStorageClass()
	rv := objc.Send[VZSEPStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage/initCreatingStorageAtURL:error:
func NewVZSEPStorageCreatingStorageAtURLError(url foundation.INSURL) (VZSEPStorage, error) {
	var errorPtr objc.ID
	instance := getVZSEPStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initCreatingStorageAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZSEPStorage{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZSEPStorageFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage/initWithURL:
func NewVZSEPStorageWithURL(url foundation.INSURL) VZSEPStorage {
	instance := getVZSEPStorageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return VZSEPStorageFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage/initCreatingStorageAtURL:error:
func (v VZSEPStorage) InitCreatingStorageAtURLError(url foundation.INSURL) (VZSEPStorage, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initCreatingStorageAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZSEPStorage), foundation.NSErrorFrom(errorPtr)
	}
	return VZSEPStorageFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage/initWithURL:
func (v VZSEPStorage) InitWithURL(url foundation.INSURL) VZSEPStorage {
	rv := objc.Send[VZSEPStorage](v.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPStorage/URL
func (v VZSEPStorage) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCoprocessorStorage] class.
var (
	_VZCoprocessorStorageClass     VZCoprocessorStorageClass
	_VZCoprocessorStorageClassOnce sync.Once
)

func getVZCoprocessorStorageClass() VZCoprocessorStorageClass {
	_VZCoprocessorStorageClassOnce.Do(func() {
		_VZCoprocessorStorageClass = VZCoprocessorStorageClass{class: objc.GetClass("_VZCoprocessorStorage")}
	})
	return _VZCoprocessorStorageClass
}

// GetVZCoprocessorStorageClass returns the class object for _VZCoprocessorStorage.
func GetVZCoprocessorStorageClass() VZCoprocessorStorageClass {
	return getVZCoprocessorStorageClass()
}

type VZCoprocessorStorageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCoprocessorStorageClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCoprocessorStorageClass) Alloc() VZCoprocessorStorage {
	rv := objc.SendIfResponds[VZCoprocessorStorage](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCoprocessorStorage.URL]
//   - [VZCoprocessorStorage._storageSize]
//   - [VZCoprocessorStorage.InitCreatingStorageAtURLError]
//   - [VZCoprocessorStorage.InitWithURL]
type VZCoprocessorStorage struct {
	objectivec.Object
}

// VZCoprocessorStorageFromID constructs a [VZCoprocessorStorage] from an objc.ID.
func VZCoprocessorStorageFromID(id objc.ID) VZCoprocessorStorage {
	return VZCoprocessorStorage{objectivec.Object{ID: id}}
}

// Ensure VZCoprocessorStorage implements IVZCoprocessorStorage.
var _ IVZCoprocessorStorage = VZCoprocessorStorage{}

// An interface definition for the [VZCoprocessorStorage] class.
//
// # Methods
//
//   - [IVZCoprocessorStorage.URL]
//   - [IVZCoprocessorStorage._storageSize]
//   - [IVZCoprocessorStorage.InitCreatingStorageAtURLError]
//   - [IVZCoprocessorStorage.InitWithURL]
type IVZCoprocessorStorage interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.NSURL
	_storageSize() uint64
	InitCreatingStorageAtURLError(url foundation.NSURL) (VZCoprocessorStorage, error)
	InitWithURL(url foundation.NSURL) VZCoprocessorStorage
}

// Init initializes the instance.
func (v VZCoprocessorStorage) Init() VZCoprocessorStorage {
	rv := objc.SendIfResponds[VZCoprocessorStorage](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCoprocessorStorage) Autorelease() VZCoprocessorStorage {
	rv := objc.SendIfResponds[VZCoprocessorStorage](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCoprocessorStorage creates a new VZCoprocessorStorage instance.
func NewVZCoprocessorStorage() VZCoprocessorStorage {
	class := getVZCoprocessorStorageClass()
	rv := objc.SendIfResponds[VZCoprocessorStorage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZCoprocessorStorageCreatingStorageAtURLError(url foundation.NSURL) (VZCoprocessorStorage, error) {
	var errorPtr objc.ID
	instance := getVZCoprocessorStorageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initCreatingStorageAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZCoprocessorStorage{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZCoprocessorStorage{}, objc.ErrInitFailed
	}
	return VZCoprocessorStorageFromID(rv), nil
}

func NewVZCoprocessorStorageWithURL(url foundation.NSURL) VZCoprocessorStorage {
	instance := getVZCoprocessorStorageClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return VZCoprocessorStorageFromID(rv)
}

func (v VZCoprocessorStorage) InitCreatingStorageAtURLError(url foundation.NSURL) (VZCoprocessorStorage, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initCreatingStorageAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZCoprocessorStorage), foundation.NSErrorFrom(errorPtr)
	}
	return VZCoprocessorStorageFromID(rv), nil

}
func (v VZCoprocessorStorage) InitWithURL(url foundation.NSURL) VZCoprocessorStorage {
	rv := objc.SendIfResponds[VZCoprocessorStorage](v.ID, objc.Sel("initWithURL:"), url)
	return rv
}

func (v VZCoprocessorStorage) URL() foundation.NSURL {
	rv := objc.SendIfResponds[foundation.NSURL](v.ID, objc.Sel("URL"))
	return foundation.NSURL(rv)
}
func (v VZCoprocessorStorage) _storageSize() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("_storageSize"))
	return rv
}

// CanStorageSize reports whether the receiver responds to the private selector _storageSize.
func (v VZCoprocessorStorage) CanStorageSize() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_storageSize"))
}

// StorageSize is an exported wrapper for the private property _storageSize.
func (v VZCoprocessorStorage) StorageSize() (uint64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_storageSize")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_storageSize"}
	}
	return v._storageSize(), nil
}

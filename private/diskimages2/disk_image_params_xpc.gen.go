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

// The class instance for the [DiskImageParamsXPC] class.
var (
	_DiskImageParamsXPCClass     DiskImageParamsXPCClass
	_DiskImageParamsXPCClassOnce sync.Once
)

func getDiskImageParamsXPCClass() DiskImageParamsXPCClass {
	_DiskImageParamsXPCClassOnce.Do(func() {
		_DiskImageParamsXPCClass = DiskImageParamsXPCClass{class: objc.GetClass("DiskImageParamsXPC")}
	})
	return _DiskImageParamsXPCClass
}

// GetDiskImageParamsXPCClass returns the class object for DiskImageParamsXPC.
func GetDiskImageParamsXPCClass() DiskImageParamsXPCClass {
	return getDiskImageParamsXPCClass()
}

type DiskImageParamsXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsXPCClass) Alloc() DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DiskImageParamsXPC.AllowOnDiskCacheWithSinkDiskImage]
//   - [DiskImageParamsXPC.BackendXPC]
//   - [DiskImageParamsXPC.SetBackendXPC]
//   - [DiskImageParamsXPC.BlockSize]
//   - [DiskImageParamsXPC.SetBlockSize]
//   - [DiskImageParamsXPC.CacheBackendXPC]
//   - [DiskImageParamsXPC.SetCacheBackendXPC]
//   - [DiskImageParamsXPC.CacheURL]
//   - [DiskImageParamsXPC.SetCacheURL]
//   - [DiskImageParamsXPC.CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size]
//   - [DiskImageParamsXPC.EncodeWithCoder]
//   - [DiskImageParamsXPC.InstanceID]
//   - [DiskImageParamsXPC.SetInstanceID]
//   - [DiskImageParamsXPC.IsSparseFormat]
//   - [DiskImageParamsXPC.IsWritableFormat]
//   - [DiskImageParamsXPC.LockBackendsWithError]
//   - [DiskImageParamsXPC.LockBackendsWithWritableOnlyError]
//   - [DiskImageParamsXPC.LockWritableBackendsWithError]
//   - [DiskImageParamsXPC.LockableResources]
//   - [DiskImageParamsXPC.MountedOnAPFS]
//   - [DiskImageParamsXPC.NumBlocks]
//   - [DiskImageParamsXPC.SetBlockSizeError]
//   - [DiskImageParamsXPC.SetSizeWithDiskImageNewSize]
//   - [DiskImageParamsXPC.ShadowChain]
//   - [DiskImageParamsXPC.SetShadowChain]
//   - [DiskImageParamsXPC.InitWithBackendXPC]
//   - [DiskImageParamsXPC.InitWithBackendXPCBlockSize]
//   - [DiskImageParamsXPC.InitWithCoder]
type DiskImageParamsXPC struct {
	objectivec.Object
}

// DiskImageParamsXPCFromID constructs a [DiskImageParamsXPC] from an objc.ID.
func DiskImageParamsXPCFromID(id objc.ID) DiskImageParamsXPC {
	return DiskImageParamsXPC{objectivec.Object{ID: id}}
}

// Ensure DiskImageParamsXPC implements IDiskImageParamsXPC.
var _ IDiskImageParamsXPC = DiskImageParamsXPC{}

// An interface definition for the [DiskImageParamsXPC] class.
//
// # Methods
//
//   - [IDiskImageParamsXPC.AllowOnDiskCacheWithSinkDiskImage]
//   - [IDiskImageParamsXPC.BackendXPC]
//   - [IDiskImageParamsXPC.SetBackendXPC]
//   - [IDiskImageParamsXPC.BlockSize]
//   - [IDiskImageParamsXPC.SetBlockSize]
//   - [IDiskImageParamsXPC.CacheBackendXPC]
//   - [IDiskImageParamsXPC.SetCacheBackendXPC]
//   - [IDiskImageParamsXPC.CacheURL]
//   - [IDiskImageParamsXPC.SetCacheURL]
//   - [IDiskImageParamsXPC.CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size]
//   - [IDiskImageParamsXPC.EncodeWithCoder]
//   - [IDiskImageParamsXPC.InstanceID]
//   - [IDiskImageParamsXPC.SetInstanceID]
//   - [IDiskImageParamsXPC.IsSparseFormat]
//   - [IDiskImageParamsXPC.IsWritableFormat]
//   - [IDiskImageParamsXPC.LockBackendsWithError]
//   - [IDiskImageParamsXPC.LockBackendsWithWritableOnlyError]
//   - [IDiskImageParamsXPC.LockWritableBackendsWithError]
//   - [IDiskImageParamsXPC.LockableResources]
//   - [IDiskImageParamsXPC.MountedOnAPFS]
//   - [IDiskImageParamsXPC.NumBlocks]
//   - [IDiskImageParamsXPC.SetBlockSizeError]
//   - [IDiskImageParamsXPC.SetSizeWithDiskImageNewSize]
//   - [IDiskImageParamsXPC.ShadowChain]
//   - [IDiskImageParamsXPC.SetShadowChain]
//   - [IDiskImageParamsXPC.InitWithBackendXPC]
//   - [IDiskImageParamsXPC.InitWithBackendXPCBlockSize]
//   - [IDiskImageParamsXPC.InitWithCoder]
type IDiskImageParamsXPC interface {
	objectivec.IObject

	// Topic: Methods

	AllowOnDiskCacheWithSinkDiskImage(image unsafe.Pointer) bool
	BackendXPC() IBackendXPC
	SetBackendXPC(value IBackendXPC)
	BlockSize() uint64
	SetBlockSize(value uint64)
	CacheBackendXPC() IBackendXPC
	SetCacheBackendXPC(value IBackendXPC)
	CacheURL() foundation.NSURL
	SetCacheURL(value foundation.NSURL)
	CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size(backend unsafe.Pointer, blocks uint64, image unsafe.Pointer, cache_only bool, stack_size uint64) unsafe.Pointer
	EncodeWithCoder(coder foundation.INSCoder)
	InstanceID() foundation.NSUUID
	SetInstanceID(value foundation.NSUUID)
	IsSparseFormat() bool
	IsWritableFormat() bool
	LockBackendsWithError() (bool, error)
	LockBackendsWithWritableOnlyError(only bool) (bool, error)
	LockWritableBackendsWithError() (bool, error)
	LockableResources() unsafe.Pointer
	MountedOnAPFS() bool
	NumBlocks() uint64
	SetBlockSizeError(size uint64) (bool, error)
	SetSizeWithDiskImageNewSize(image unsafe.Pointer, size uint64) int
	ShadowChain() IDIShadowChain
	SetShadowChain(value IDIShadowChain)
	InitWithBackendXPC(xpc objectivec.IObject) DiskImageParamsXPC
	InitWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsXPC
	InitWithCoder(coder foundation.INSCoder) DiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsXPC) Init() DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsXPC) Autorelease() DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsXPC creates a new DiskImageParamsXPC instance.
func NewDiskImageParamsXPC() DiskImageParamsXPC {
	class := getDiskImageParamsXPCClass()
	rv := objc.Send[DiskImageParamsXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsXPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsXPCFromID(rv)
}

func NewDiskImageParamsXPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsXPCFromID(rv)
}

func NewDiskImageParamsXPCWithCoder(coder objectivec.IObject) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsXPCFromID(rv)
}

func (d DiskImageParamsXPC) AllowOnDiskCacheWithSinkDiskImage(image unsafe.Pointer) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowOnDiskCacheWithSinkDiskImage:"), image)
	return rv
}
func (d DiskImageParamsXPC) CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size(backend unsafe.Pointer, blocks uint64, image unsafe.Pointer, cache_only bool, stack_size uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("createShadowDiskImageWithBackend:numBlocks:sinkDiskImage:cache_only:stack_size:"), backend, blocks, image, cache_only, stack_size)
	return rv
}
func (d DiskImageParamsXPC) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (d DiskImageParamsXPC) LockBackendsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("lockBackendsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lockBackendsWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DiskImageParamsXPC) LockBackendsWithWritableOnlyError(only bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("lockBackendsWithWritableOnly:error:"), only, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lockBackendsWithWritableOnly:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DiskImageParamsXPC) LockWritableBackendsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("lockWritableBackendsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lockWritableBackendsWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DiskImageParamsXPC) MountedOnAPFS() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("mountedOnAPFS"))
	return rv
}
func (d DiskImageParamsXPC) SetBlockSizeError(size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setBlockSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setBlockSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DiskImageParamsXPC) SetSizeWithDiskImageNewSize(image unsafe.Pointer, size uint64) int {
	rv := objc.Send[int](d.ID, objc.Sel("setSizeWithDiskImage:newSize:"), image, size)
	return rv
}
func (d DiskImageParamsXPC) InitWithBackendXPC(xpc objectivec.IObject) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return rv
}
func (d DiskImageParamsXPC) InitWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return rv
}
func (d DiskImageParamsXPC) InitWithCoder(coder foundation.INSCoder) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) GetAEAKeyFromSAKSWithMetadataKeyError(metadata objectivec.IObject, key string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("getAEAKeyFromSAKSWithMetadata:key:error:"), metadata, unsafe.Pointer(unsafe.StringData(key+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getAEAKeyFromSAKSWithMetadata:key:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) GetAEAKeyWithHelperKeyBufferBufferSizeError(helper unsafe.Pointer, buffer string, size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("getAEAKeyWithHelper:keyBuffer:bufferSize:error:"), helper, unsafe.Pointer(unsafe.StringData(buffer+"\x00")), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getAEAKeyWithHelper:keyBuffer:bufferSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewAEABackendThrowsWithBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newAEABackendThrowsWithBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithURLFileOpenFlagsError(url foundation.NSURL, flags int) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithURL:fileOpenFlags:error:"), url, flags, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithUnlockedBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithUnlockedBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) ValidateSupportedFormatWithBackendXPCError(xpc objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("validateSupportedFormatWithBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateSupportedFormatWithBackendXPC:error: returned NO with nil NSError")
	}
	return rv, nil

}

func (d DiskImageParamsXPC) BackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetBackendXPC(value IBackendXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setBackendXPC:"), value)
}
func (d DiskImageParamsXPC) BlockSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("blockSize"))
	return rv
}
func (d DiskImageParamsXPC) SetBlockSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setBlockSize:"), value)
}
func (d DiskImageParamsXPC) CacheBackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cacheBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetCacheBackendXPC(value IBackendXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setCacheBackendXPC:"), value)
}
func (d DiskImageParamsXPC) CacheURL() foundation.NSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cacheURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetCacheURL(value foundation.NSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setCacheURL:"), value)
}
func (d DiskImageParamsXPC) InstanceID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("instanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetInstanceID(value foundation.NSUUID) {
	objc.Send[struct{}](d.ID, objc.Sel("setInstanceID:"), value)
}
func (d DiskImageParamsXPC) IsSparseFormat() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isSparseFormat"))
	return rv
}
func (d DiskImageParamsXPC) IsWritableFormat() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isWritableFormat"))
	return rv
}
func (d DiskImageParamsXPC) LockableResources() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("lockableResources"))
	return rv
}
func (d DiskImageParamsXPC) NumBlocks() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("numBlocks"))
	return rv
}
func (d DiskImageParamsXPC) ShadowChain() IDIShadowChain {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("shadowChain"))
	return DIShadowChainFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetShadowChain(value IDIShadowChain) {
	objc.Send[struct{}](d.ID, objc.Sel("setShadowChain:"), value)
}

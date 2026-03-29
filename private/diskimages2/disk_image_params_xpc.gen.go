// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
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

//
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
//   - [DiskImageParamsXPC.CreateDiskImageWithCacheShadowValidation]
//   - [DiskImageParamsXPC.CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size]
//   - [DiskImageParamsXPC.CreateSinkDiskImage]
//   - [DiskImageParamsXPC.EncodeWithCoder]
//   - [DiskImageParamsXPC.GetImageInfoWithExtraError]
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
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC
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
//   - [IDiskImageParamsXPC.CreateDiskImageWithCacheShadowValidation]
//   - [IDiskImageParamsXPC.CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size]
//   - [IDiskImageParamsXPC.CreateSinkDiskImage]
//   - [IDiskImageParamsXPC.EncodeWithCoder]
//   - [IDiskImageParamsXPC.GetImageInfoWithExtraError]
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
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC
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
	CacheURL() foundation.INSURL
	SetCacheURL(value foundation.INSURL)
	CreateDiskImageWithCacheShadowValidation(cache bool, validation bool) objectivec.IObject
	CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size(backend objectivec.IObject, blocks uint64, image unsafe.Pointer, cache_only bool, stack_size uint64) objectivec.IObject
	CreateSinkDiskImage() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	GetImageInfoWithExtraError(extra bool) (objectivec.IObject, error)
	InstanceID() foundation.NSUUID
	SetInstanceID(value foundation.NSUUID)
	IsSparseFormat() bool
	IsWritableFormat() bool
	LockBackendsWithError() (bool, error)
	LockBackendsWithWritableOnlyError(only bool) (bool, error)
	LockWritableBackendsWithError() (bool, error)
	LockableResources() objectivec.IObject
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

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsXPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsXPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsXPCWithCoder(coder objectivec.IObject) DiskImageParamsXPC {
	instance := getDiskImageParamsXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/allowOnDiskCacheWithSinkDiskImage:
func (d DiskImageParamsXPC) AllowOnDiskCacheWithSinkDiskImage(image unsafe.Pointer) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowOnDiskCacheWithSinkDiskImage:"), image)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/createDiskImageWithCache:shadowValidation:
func (d DiskImageParamsXPC) CreateDiskImageWithCacheShadowValidation(cache bool, validation bool) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createDiskImageWithCache:shadowValidation:"), cache, validation)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/createShadowDiskImageWithBackend:numBlocks:sinkDiskImage:cache_only:stack_size:
func (d DiskImageParamsXPC) CreateShadowDiskImageWithBackendNumBlocksSinkDiskImageCache_onlyStack_size(backend objectivec.IObject, blocks uint64, image unsafe.Pointer, cache_only bool, stack_size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createShadowDiskImageWithBackend:numBlocks:sinkDiskImage:cache_only:stack_size:"), backend, blocks, image, cache_only, stack_size)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/createSinkDiskImage
func (d DiskImageParamsXPC) CreateSinkDiskImage() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createSinkDiskImage"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/encodeWithCoder:
func (d DiskImageParamsXPC) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/getImageInfoWithExtra:error:
func (d DiskImageParamsXPC) GetImageInfoWithExtraError(extra bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getImageInfoWithExtra:error:"), extra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/lockBackendsWithError:
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
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/lockBackendsWithWritableOnly:error:
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
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/lockWritableBackendsWithError:
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
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/mountedOnAPFS
func (d DiskImageParamsXPC) MountedOnAPFS() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("mountedOnAPFS"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/setBlockSize:error:
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
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/setSizeWithDiskImage:newSize:
func (d DiskImageParamsXPC) SetSizeWithDiskImageNewSize(image unsafe.Pointer, size uint64) int {
	rv := objc.Send[int](d.ID, objc.Sel("setSizeWithDiskImage:newSize:"), image, size)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func (d DiskImageParamsXPC) InitWithBackendXPC(xpc objectivec.IObject) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func (d DiskImageParamsXPC) InitWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func (d DiskImageParamsXPC) InitWithCoder(coder foundation.INSCoder) DiskImageParamsXPC {
	rv := objc.Send[DiskImageParamsXPC](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/getAEAKeyFromSAKSWithMetadata:key:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) GetAEAKeyFromSAKSWithMetadataKeyError(metadata objectivec.IObject, key string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("getAEAKeyFromSAKSWithMetadata:key:error:"), metadata, unsafe.Pointer(unsafe.StringData(key + "\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getAEAKeyFromSAKSWithMetadata:key:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/getAEAKeyWithHelper:keyBuffer:bufferSize:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) GetAEAKeyWithHelperKeyBufferBufferSizeError(helper unsafe.Pointer, buffer string, size uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("getAEAKeyWithHelper:keyBuffer:bufferSize:error:"), helper, unsafe.Pointer(unsafe.StringData(buffer + "\x00")), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getAEAKeyWithHelper:keyBuffer:bufferSize:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/newAEABackendThrowsWithBackendXPC:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewAEABackendThrowsWithBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newAEABackendThrowsWithBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/newWithBackendXPC:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/newWithURL:fileOpenFlags:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithURLFileOpenFlagsError(url foundation.INSURL, flags int) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithURL:fileOpenFlags:error:"), url, flags, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/newWithUnlockedBackendXPC:error:
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) NewWithUnlockedBackendXPCError(xpc objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("newWithUnlockedBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/supportsSecureCoding
func (_DiskImageParamsXPCClass DiskImageParamsXPCClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DiskImageParamsXPCClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/validateSupportedFormatWithBackendXPC:error:
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

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/backendXPC
func (d DiskImageParamsXPC) BackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetBackendXPC(value IBackendXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setBackendXPC:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/blockSize
func (d DiskImageParamsXPC) BlockSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("blockSize"))
	return rv
}
func (d DiskImageParamsXPC) SetBlockSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setBlockSize:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/cacheBackendXPC
func (d DiskImageParamsXPC) CacheBackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cacheBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetCacheBackendXPC(value IBackendXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setCacheBackendXPC:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/cacheURL
func (d DiskImageParamsXPC) CacheURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cacheURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetCacheURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setCacheURL:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/instanceID
func (d DiskImageParamsXPC) InstanceID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("instanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetInstanceID(value foundation.NSUUID) {
	objc.Send[struct{}](d.ID, objc.Sel("setInstanceID:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/isSparseFormat
func (d DiskImageParamsXPC) IsSparseFormat() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isSparseFormat"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/isWritableFormat
func (d DiskImageParamsXPC) IsWritableFormat() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isWritableFormat"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/lockableResources
func (d DiskImageParamsXPC) LockableResources() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("lockableResources"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/numBlocks
func (d DiskImageParamsXPC) NumBlocks() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("numBlocks"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/shadowChain
func (d DiskImageParamsXPC) ShadowChain() IDIShadowChain {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("shadowChain"))
	return DIShadowChainFromID(objc.ID(rv))
}
func (d DiskImageParamsXPC) SetShadowChain(value IDIShadowChain) {
	objc.Send[struct{}](d.ID, objc.Sel("setShadowChain:"), value)
}


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

// The class instance for the [DIAttachParams] class.
var (
	_DIAttachParamsClass     DIAttachParamsClass
	_DIAttachParamsClassOnce sync.Once
)

func getDIAttachParamsClass() DIAttachParamsClass {
	_DIAttachParamsClassOnce.Do(func() {
		_DIAttachParamsClass = DIAttachParamsClass{class: objc.GetClass("DIAttachParams")}
	})
	return _DIAttachParamsClass
}

// GetDIAttachParamsClass returns the class object for DIAttachParams.
func GetDIAttachParamsClass() DIAttachParamsClass {
	return getDIAttachParamsClass()
}

type DIAttachParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIAttachParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIAttachParamsClass) Alloc() DIAttachParams {
	rv := objc.Send[DIAttachParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIAttachParams.AutoMount]
//   - [DIAttachParams.SetAutoMount]
//   - [DIAttachParams.FileMode]
//   - [DIAttachParams.SetFileMode]
//   - [DIAttachParams.EmulateExternalDisk]
//   - [DIAttachParams.SetEmulateExternalDisk]
//   - [DIAttachParams.OnDiskCache]
//   - [DIAttachParams.SetOnDiskCache]
//   - [DIAttachParams.UniqueDevice]
//   - [DIAttachParams.SetUniqueDevice]
//   - [DIAttachParams.HandleRefCount]
//   - [DIAttachParams.SetHandleRefCount]
//   - [DIAttachParams.SingleInstanceDaemon]
//   - [DIAttachParams.SetSingleInstanceDaemon]
//   - [DIAttachParams.SuppressSsdFlags]
//   - [DIAttachParams.SetSuppressSsdFlags]
//   - [DIAttachParams.ShouldValidateShadows]
//   - [DIAttachParams.SetShouldValidateShadows]
//   - [DIAttachParams.CommandSize]
//   - [DIAttachParams.SetCommandSize]
//   - [DIAttachParams.RegEntryID]
//   - [DIAttachParams.SetRegEntryID]
//   - [DIAttachParams.CustomCacheURL]
//   - [DIAttachParams.SetCustomCacheURL]
//   - [DIAttachParams.InputMountedFrom]
//   - [DIAttachParams.SetInputMountedFrom]
//   - [DIAttachParams.SetPassphraseError]
//   - [DIAttachParams.NewAttachWithError]
//   - [DIAttachParams.SetupDefaults]
//   - [DIAttachParams.InputStatFS]
//   - [DIAttachParams.SetInputStatFS]
//   - [DIAttachParams.IsDeviceHighThroughputWithRegistryEntryID]
//   - [DIAttachParams.IsDeviceSolidStateWithRegistryEntryID]
//   - [DIAttachParams.IsDeviceWithPropertyRegistryEntryIDPredicate]
//   - [DIAttachParams.ReOpenIfWritableWithError]
//   - [DIAttachParams.ToDI1ParamsWithError]
//   - [DIAttachParams.UpdateStatFSWithError]
//   - [DIAttachParams.InitWithURLShadowURLsError]
//   - [DIAttachParams.InitWithExistingParamsError]
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams
type DIAttachParams struct {
	DIBaseParams
}

// DIAttachParamsFromID constructs a [DIAttachParams] from an objc.ID.
func DIAttachParamsFromID(id objc.ID) DIAttachParams {
	return DIAttachParams{DIBaseParams: DIBaseParamsFromID(id)}
}
// Ensure DIAttachParams implements IDIAttachParams.
var _ IDIAttachParams = DIAttachParams{}

// An interface definition for the [DIAttachParams] class.
//
// # Methods
//
//   - [IDIAttachParams.AutoMount]
//   - [IDIAttachParams.SetAutoMount]
//   - [IDIAttachParams.FileMode]
//   - [IDIAttachParams.SetFileMode]
//   - [IDIAttachParams.EmulateExternalDisk]
//   - [IDIAttachParams.SetEmulateExternalDisk]
//   - [IDIAttachParams.OnDiskCache]
//   - [IDIAttachParams.SetOnDiskCache]
//   - [IDIAttachParams.UniqueDevice]
//   - [IDIAttachParams.SetUniqueDevice]
//   - [IDIAttachParams.HandleRefCount]
//   - [IDIAttachParams.SetHandleRefCount]
//   - [IDIAttachParams.SingleInstanceDaemon]
//   - [IDIAttachParams.SetSingleInstanceDaemon]
//   - [IDIAttachParams.SuppressSsdFlags]
//   - [IDIAttachParams.SetSuppressSsdFlags]
//   - [IDIAttachParams.ShouldValidateShadows]
//   - [IDIAttachParams.SetShouldValidateShadows]
//   - [IDIAttachParams.CommandSize]
//   - [IDIAttachParams.SetCommandSize]
//   - [IDIAttachParams.RegEntryID]
//   - [IDIAttachParams.SetRegEntryID]
//   - [IDIAttachParams.CustomCacheURL]
//   - [IDIAttachParams.SetCustomCacheURL]
//   - [IDIAttachParams.InputMountedFrom]
//   - [IDIAttachParams.SetInputMountedFrom]
//   - [IDIAttachParams.SetPassphraseError]
//   - [IDIAttachParams.NewAttachWithError]
//   - [IDIAttachParams.SetupDefaults]
//   - [IDIAttachParams.InputStatFS]
//   - [IDIAttachParams.SetInputStatFS]
//   - [IDIAttachParams.IsDeviceHighThroughputWithRegistryEntryID]
//   - [IDIAttachParams.IsDeviceSolidStateWithRegistryEntryID]
//   - [IDIAttachParams.IsDeviceWithPropertyRegistryEntryIDPredicate]
//   - [IDIAttachParams.ReOpenIfWritableWithError]
//   - [IDIAttachParams.ToDI1ParamsWithError]
//   - [IDIAttachParams.UpdateStatFSWithError]
//   - [IDIAttachParams.InitWithURLShadowURLsError]
//   - [IDIAttachParams.InitWithExistingParamsError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams
type IDIAttachParams interface {
	IDIBaseParams

	// Topic: Methods

	AutoMount() bool
	SetAutoMount(value bool)
	FileMode() int64
	SetFileMode(value int64)
	EmulateExternalDisk() bool
	SetEmulateExternalDisk(value bool)
	OnDiskCache() bool
	SetOnDiskCache(value bool)
	UniqueDevice() bool
	SetUniqueDevice(value bool)
	HandleRefCount() bool
	SetHandleRefCount(value bool)
	SingleInstanceDaemon() bool
	SetSingleInstanceDaemon(value bool)
	SuppressSsdFlags() bool
	SetSuppressSsdFlags(value bool)
	ShouldValidateShadows() bool
	SetShouldValidateShadows(value bool)
	CommandSize() uint64
	SetCommandSize(value uint64)
	RegEntryID() uint64
	SetRegEntryID(value uint64)
	CustomCacheURL() foundation.INSURL
	SetCustomCacheURL(value foundation.INSURL)
	InputMountedFrom() string
	SetInputMountedFrom(value string)
	SetPassphraseError(passphrase string) (bool, error)
	NewAttachWithError() (IDIDeviceHandle, error)
	SetupDefaults()
	InputStatFS() IDIStatFS
	SetInputStatFS(value IDIStatFS)
	IsDeviceHighThroughputWithRegistryEntryID(id uint64) bool
	IsDeviceSolidStateWithRegistryEntryID(id uint64) bool
	IsDeviceWithPropertyRegistryEntryIDPredicate(property string, id uint64, predicate objectivec.IObject) bool
	ReOpenIfWritableWithError() (bool, error)
	ToDI1ParamsWithError() (objectivec.IObject, error)
	UpdateStatFSWithError() (bool, error)
	InitWithURLShadowURLsError(url foundation.INSURL, shadowURLs foundation.INSArray) (DIAttachParams, error)
	InitWithExistingParamsError(params IDIAttachParams) (DIAttachParams, error)
}

// Init initializes the instance.
func (d DIAttachParams) Init() DIAttachParams {
	rv := objc.Send[DIAttachParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIAttachParams) Autorelease() DIAttachParams {
	rv := objc.Send[DIAttachParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIAttachParams creates a new DIAttachParams instance.
func NewDIAttachParams() DIAttachParams {
	class := getDIAttachParamsClass()
	rv := objc.Send[DIAttachParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithCoder:
func NewDIAttachParamsWithCoder(coder objectivec.IObject) DIAttachParams {
	instance := getDIAttachParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIAttachParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithExistingParams:error:
func NewDIAttachParamsWithExistingParamsError(params IDIAttachParams) (DIAttachParams, error) {
	var errorPtr objc.ID
	instance := getDIAttachParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithURL:error:
func NewDIAttachParamsWithURLError(url foundation.INSURL) (DIAttachParams, error) {
	var errorPtr objc.ID
	instance := getDIAttachParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithURL:shadowURLs:error:
func NewDIAttachParamsWithURLShadowURLsError(url foundation.INSURL, shadowURLs foundation.INSArray) (DIAttachParams, error) {
	var errorPtr objc.ID
	instance := getDIAttachParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:shadowURLs:error:"), url, shadowURLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/setPassphrase:error:
func (d DIAttachParams) SetPassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setPassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase + "\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/newAttachWithError:
func (d DIAttachParams) NewAttachWithError() (IDIDeviceHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("newAttachWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDeviceHandle{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIDeviceHandleFromID(rv), nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/setupDefaults
func (d DIAttachParams) SetupDefaults() {
	objc.Send[objc.ID](d.ID, objc.Sel("setupDefaults"))
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/isDeviceHighThroughputWithRegistryEntryID:
func (d DIAttachParams) IsDeviceHighThroughputWithRegistryEntryID(id uint64) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDeviceHighThroughputWithRegistryEntryID:"), id)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/isDeviceSolidStateWithRegistryEntryID:
func (d DIAttachParams) IsDeviceSolidStateWithRegistryEntryID(id uint64) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDeviceSolidStateWithRegistryEntryID:"), id)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/isDeviceWithProperty:registryEntryID:predicate:
func (d DIAttachParams) IsDeviceWithPropertyRegistryEntryIDPredicate(property string, id uint64, predicate objectivec.IObject) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDeviceWithProperty:registryEntryID:predicate:"), unsafe.Pointer(unsafe.StringData(property + "\x00")), id, predicate)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/reOpenIfWritableWithError:
func (d DIAttachParams) ReOpenIfWritableWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("reOpenIfWritableWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reOpenIfWritableWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/toDI1ParamsWithError:
func (d DIAttachParams) ToDI1ParamsWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("toDI1ParamsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/updateStatFSWithError:
func (d DIAttachParams) UpdateStatFSWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("updateStatFSWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateStatFSWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithURL:shadowURLs:error:
func (d DIAttachParams) InitWithURLShadowURLsError(url foundation.INSURL, shadowURLs foundation.INSArray) (DIAttachParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:shadowURLs:error:"), url, shadowURLs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachParamsFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/initWithExistingParams:error:
func (d DIAttachParams) InitWithExistingParamsError(params IDIAttachParams) (DIAttachParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithExistingParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachParamsFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/copyWithURL:outURLStr:maxLen:error:
func (_DIAttachParamsClass DIAttachParamsClass) CopyWithURLOutURLStrMaxLenError(url foundation.INSURL, uRLStr string, len_ uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIAttachParamsClass.class), objc.Sel("copyWithURL:outURLStr:maxLen:error:"), url, unsafe.Pointer(unsafe.StringData(uRLStr + "\x00")), len_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyWithURL:outURLStr:maxLen:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/autoMount
func (d DIAttachParams) AutoMount() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("autoMount"))
	return rv
}
func (d DIAttachParams) SetAutoMount(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setAutoMount:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/commandSize
func (d DIAttachParams) CommandSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("commandSize"))
	return rv
}
func (d DIAttachParams) SetCommandSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setCommandSize:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/customCacheURL
func (d DIAttachParams) CustomCacheURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("customCacheURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d DIAttachParams) SetCustomCacheURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setCustomCacheURL:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/emulateExternalDisk
func (d DIAttachParams) EmulateExternalDisk() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("emulateExternalDisk"))
	return rv
}
func (d DIAttachParams) SetEmulateExternalDisk(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setEmulateExternalDisk:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/fileMode
func (d DIAttachParams) FileMode() int64 {
	rv := objc.Send[int64](d.ID, objc.Sel("fileMode"))
	return rv
}
func (d DIAttachParams) SetFileMode(value int64) {
	objc.Send[struct{}](d.ID, objc.Sel("setFileMode:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/handleRefCount
func (d DIAttachParams) HandleRefCount() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("handleRefCount"))
	return rv
}
func (d DIAttachParams) SetHandleRefCount(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setHandleRefCount:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/inputMountedFrom
func (d DIAttachParams) InputMountedFrom() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputMountedFrom"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIAttachParams) SetInputMountedFrom(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setInputMountedFrom:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/onDiskCache
func (d DIAttachParams) OnDiskCache() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("onDiskCache"))
	return rv
}
func (d DIAttachParams) SetOnDiskCache(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setOnDiskCache:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/regEntryID
func (d DIAttachParams) RegEntryID() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("regEntryID"))
	return rv
}
func (d DIAttachParams) SetRegEntryID(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setRegEntryID:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/shouldValidateShadows
func (d DIAttachParams) ShouldValidateShadows() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldValidateShadows"))
	return rv
}
func (d DIAttachParams) SetShouldValidateShadows(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setShouldValidateShadows:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/singleInstanceDaemon
func (d DIAttachParams) SingleInstanceDaemon() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("singleInstanceDaemon"))
	return rv
}
func (d DIAttachParams) SetSingleInstanceDaemon(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setSingleInstanceDaemon:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/suppressSsdFlags
func (d DIAttachParams) SuppressSsdFlags() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("suppressSsdFlags"))
	return rv
}
func (d DIAttachParams) SetSuppressSsdFlags(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setSuppressSsdFlags:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/uniqueDevice
func (d DIAttachParams) UniqueDevice() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("uniqueDevice"))
	return rv
}
func (d DIAttachParams) SetUniqueDevice(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setUniqueDevice:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachParams/inputStatFS
func (d DIAttachParams) InputStatFS() IDIStatFS {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputStatFS"))
	return DIStatFSFromID(objc.ID(rv))
}
func (d DIAttachParams) SetInputStatFS(value IDIStatFS) {
	objc.Send[struct{}](d.ID, objc.Sel("setInputStatFS:"), value)
}


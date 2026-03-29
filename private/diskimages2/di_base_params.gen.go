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

// The class instance for the [DIBaseParams] class.
var (
	_DIBaseParamsClass     DIBaseParamsClass
	_DIBaseParamsClassOnce sync.Once
)

func getDIBaseParamsClass() DIBaseParamsClass {
	_DIBaseParamsClassOnce.Do(func() {
		_DIBaseParamsClass = DIBaseParamsClass{class: objc.GetClass("DIBaseParams")}
	})
	return _DIBaseParamsClass
}

// GetDIBaseParamsClass returns the class object for DIBaseParams.
func GetDIBaseParamsClass() DIBaseParamsClass {
	return getDIBaseParamsClass()
}

type DIBaseParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIBaseParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIBaseParamsClass) Alloc() DIBaseParams {
	rv := objc.Send[DIBaseParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIBaseParams.RAMdisk]
//   - [DIBaseParams.Backend]
//   - [DIBaseParams.BlockSize]
//   - [DIBaseParams.SetBlockSize]
//   - [DIBaseParams.CryptoHeader]
//   - [DIBaseParams.DeserializationError]
//   - [DIBaseParams.SetDeserializationError]
//   - [DIBaseParams.DiskImageParamsXPC]
//   - [DIBaseParams.SetDiskImageParamsXPC]
//   - [DIBaseParams.EncodeWithCoder]
//   - [DIBaseParams.EncryptionUUID]
//   - [DIBaseParams.HasUnlockedBackend]
//   - [DIBaseParams.InputURL]
//   - [DIBaseParams.SetInputURL]
//   - [DIBaseParams.InstanceID]
//   - [DIBaseParams.Invalidate]
//   - [DIBaseParams.IsPstack]
//   - [DIBaseParams.MutableSymmetricKey]
//   - [DIBaseParams.OpenExistingImageWithError]
//   - [DIBaseParams.OpenExistingImageWithFlagsError]
//   - [DIBaseParams.OverrideBlockSize]
//   - [DIBaseParams.PrepareImageWithXpcHandlerFileModeError]
//   - [DIBaseParams.ReadPassphraseFlags]
//   - [DIBaseParams.SetReadPassphraseFlags]
//   - [DIBaseParams.ReplaceDiskImageWithUnlockedBackendXPCError]
//   - [DIBaseParams.RequiresRootDaemon]
//   - [DIBaseParams.ShadowChain]
//   - [DIBaseParams.SupportsPstack]
//   - [DIBaseParams.SymmetricKey]
//   - [DIBaseParams.SetSymmetricKey]
//   - [DIBaseParams.TryResolvePstackChain]
//   - [DIBaseParams.UnlockWithPassphraseError]
//   - [DIBaseParams.ValidateDeserializationWithError]
//   - [DIBaseParams.InitWithCoder]
//   - [DIBaseParams.InitWithURLError]
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams
type DIBaseParams struct {
	objectivec.Object
}

// DIBaseParamsFromID constructs a [DIBaseParams] from an objc.ID.
func DIBaseParamsFromID(id objc.ID) DIBaseParams {
	return DIBaseParams{objectivec.Object{ID: id}}
}
// Ensure DIBaseParams implements IDIBaseParams.
var _ IDIBaseParams = DIBaseParams{}

// An interface definition for the [DIBaseParams] class.
//
// # Methods
//
//   - [IDIBaseParams.RAMdisk]
//   - [IDIBaseParams.Backend]
//   - [IDIBaseParams.BlockSize]
//   - [IDIBaseParams.SetBlockSize]
//   - [IDIBaseParams.CryptoHeader]
//   - [IDIBaseParams.DeserializationError]
//   - [IDIBaseParams.SetDeserializationError]
//   - [IDIBaseParams.DiskImageParamsXPC]
//   - [IDIBaseParams.SetDiskImageParamsXPC]
//   - [IDIBaseParams.EncodeWithCoder]
//   - [IDIBaseParams.EncryptionUUID]
//   - [IDIBaseParams.HasUnlockedBackend]
//   - [IDIBaseParams.InputURL]
//   - [IDIBaseParams.SetInputURL]
//   - [IDIBaseParams.InstanceID]
//   - [IDIBaseParams.Invalidate]
//   - [IDIBaseParams.IsPstack]
//   - [IDIBaseParams.MutableSymmetricKey]
//   - [IDIBaseParams.OpenExistingImageWithError]
//   - [IDIBaseParams.OpenExistingImageWithFlagsError]
//   - [IDIBaseParams.OverrideBlockSize]
//   - [IDIBaseParams.PrepareImageWithXpcHandlerFileModeError]
//   - [IDIBaseParams.ReadPassphraseFlags]
//   - [IDIBaseParams.SetReadPassphraseFlags]
//   - [IDIBaseParams.ReplaceDiskImageWithUnlockedBackendXPCError]
//   - [IDIBaseParams.RequiresRootDaemon]
//   - [IDIBaseParams.ShadowChain]
//   - [IDIBaseParams.SupportsPstack]
//   - [IDIBaseParams.SymmetricKey]
//   - [IDIBaseParams.SetSymmetricKey]
//   - [IDIBaseParams.TryResolvePstackChain]
//   - [IDIBaseParams.UnlockWithPassphraseError]
//   - [IDIBaseParams.ValidateDeserializationWithError]
//   - [IDIBaseParams.InitWithCoder]
//   - [IDIBaseParams.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams
type IDIBaseParams interface {
	objectivec.IObject

	// Topic: Methods

	RAMdisk() bool
	Backend() objectivec.IObject
	BlockSize() uint32
	SetBlockSize(value uint32)
	CryptoHeader() unsafe.Pointer
	DeserializationError() foundation.INSError
	SetDeserializationError(value foundation.INSError)
	DiskImageParamsXPC() IDiskImageParamsXPC
	SetDiskImageParamsXPC(value IDiskImageParamsXPC)
	EncodeWithCoder(coder foundation.INSCoder)
	EncryptionUUID() foundation.NSUUID
	HasUnlockedBackend() bool
	InputURL() IDIURL
	SetInputURL(value IDIURL)
	InstanceID() foundation.NSUUID
	Invalidate()
	IsPstack() bool
	MutableSymmetricKey() foundation.NSMutableData
	OpenExistingImageWithError() (bool, error)
	OpenExistingImageWithFlagsError(flags int) (bool, error)
	OverrideBlockSize() foundation.NSNumber
	PrepareImageWithXpcHandlerFileModeError(handler objectivec.IObject, mode int64) (bool, error)
	ReadPassphraseFlags() uint64
	SetReadPassphraseFlags(value uint64)
	ReplaceDiskImageWithUnlockedBackendXPCError(xpc objectivec.IObject) (bool, error)
	RequiresRootDaemon() bool
	ShadowChain() IDIShadowChain
	SupportsPstack() bool
	SymmetricKey() foundation.INSData
	SetSymmetricKey(value foundation.INSData)
	TryResolvePstackChain(chain []objectivec.IObject) bool
	UnlockWithPassphraseError(passphrase string) (bool, error)
	ValidateDeserializationWithError() (bool, error)
	InitWithCoder(coder foundation.INSCoder) DIBaseParams
	InitWithURLError(url foundation.INSURL) (DIBaseParams, error)
}

// Init initializes the instance.
func (d DIBaseParams) Init() DIBaseParams {
	rv := objc.Send[DIBaseParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIBaseParams) Autorelease() DIBaseParams {
	rv := objc.Send[DIBaseParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIBaseParams creates a new DIBaseParams instance.
func NewDIBaseParams() DIBaseParams {
	class := getDIBaseParamsClass()
	rv := objc.Send[DIBaseParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithCoder:
func NewDIBaseParamsWithCoder(coder objectivec.IObject) DIBaseParams {
	instance := getDIBaseParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIBaseParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithURL:error:
func NewDIBaseParamsWithURLError(url foundation.INSURL) (DIBaseParams, error) {
	var errorPtr objc.ID
	instance := getDIBaseParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIBaseParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIBaseParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/encodeWithCoder:
func (d DIBaseParams) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/invalidate
func (d DIBaseParams) Invalidate() {
	objc.Send[objc.ID](d.ID, objc.Sel("invalidate"))
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/openExistingImageWithError:
func (d DIBaseParams) OpenExistingImageWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("openExistingImageWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("openExistingImageWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/openExistingImageWithFlags:error:
func (d DIBaseParams) OpenExistingImageWithFlagsError(flags int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("openExistingImageWithFlags:error:"), flags, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("openExistingImageWithFlags:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/prepareImageWithXpcHandler:fileMode:error:
func (d DIBaseParams) PrepareImageWithXpcHandlerFileModeError(handler objectivec.IObject, mode int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("prepareImageWithXpcHandler:fileMode:error:"), handler, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareImageWithXpcHandler:fileMode:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/replaceDiskImageWithUnlockedBackendXPC:error:
func (d DIBaseParams) ReplaceDiskImageWithUnlockedBackendXPCError(xpc objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("replaceDiskImageWithUnlockedBackendXPC:error:"), xpc, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("replaceDiskImageWithUnlockedBackendXPC:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/supportsPstack
func (d DIBaseParams) SupportsPstack() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("supportsPstack"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/tryResolvePstackChain:
func (d DIBaseParams) TryResolvePstackChain(chain []objectivec.IObject) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("tryResolvePstackChain:"), objectivec.IObjectSliceToNSArray(chain))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/unlockWithPassphrase:error:
func (d DIBaseParams) UnlockWithPassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockWithPassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase + "\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockWithPassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/validateDeserializationWithError:
func (d DIBaseParams) ValidateDeserializationWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("validateDeserializationWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateDeserializationWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithCoder:
func (d DIBaseParams) InitWithCoder(coder foundation.INSCoder) DIBaseParams {
	rv := objc.Send[DIBaseParams](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/initWithURL:error:
func (d DIBaseParams) InitWithURLError(url foundation.INSURL) (DIBaseParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIBaseParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIBaseParamsFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/supportsSecureCoding
func (_DIBaseParamsClass DIBaseParamsClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIBaseParamsClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/RAMdisk
func (d DIBaseParams) RAMdisk() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("RAMdisk"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/backend
func (d DIBaseParams) Backend() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backend"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/blockSize
func (d DIBaseParams) BlockSize() uint32 {
	rv := objc.Send[uint32](d.ID, objc.Sel("blockSize"))
	return rv
}
func (d DIBaseParams) SetBlockSize(value uint32) {
	objc.Send[struct{}](d.ID, objc.Sel("setBlockSize:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/cryptoHeader
func (d DIBaseParams) CryptoHeader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("cryptoHeader"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/deserializationError
func (d DIBaseParams) DeserializationError() foundation.INSError {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("deserializationError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
func (d DIBaseParams) SetDeserializationError(value foundation.INSError) {
	objc.Send[struct{}](d.ID, objc.Sel("setDeserializationError:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/diskImageParamsXPC
func (d DIBaseParams) DiskImageParamsXPC() IDiskImageParamsXPC {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("diskImageParamsXPC"))
	return DiskImageParamsXPCFromID(objc.ID(rv))
}
func (d DIBaseParams) SetDiskImageParamsXPC(value IDiskImageParamsXPC) {
	objc.Send[struct{}](d.ID, objc.Sel("setDiskImageParamsXPC:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/encryptionUUID
func (d DIBaseParams) EncryptionUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("encryptionUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/hasUnlockedBackend
func (d DIBaseParams) HasUnlockedBackend() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasUnlockedBackend"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/inputURL
func (d DIBaseParams) InputURL() IDIURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputURL"))
	return DIURLFromID(objc.ID(rv))
}
func (d DIBaseParams) SetInputURL(value IDIURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setInputURL:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/instanceID
func (d DIBaseParams) InstanceID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("instanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/isPstack
func (d DIBaseParams) IsPstack() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isPstack"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/mutableSymmetricKey
func (d DIBaseParams) MutableSymmetricKey() foundation.NSMutableData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mutableSymmetricKey"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/overrideBlockSize
func (d DIBaseParams) OverrideBlockSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("overrideBlockSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/readPassphraseFlags
func (d DIBaseParams) ReadPassphraseFlags() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("readPassphraseFlags"))
	return rv
}
func (d DIBaseParams) SetReadPassphraseFlags(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setReadPassphraseFlags:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/requiresRootDaemon
func (d DIBaseParams) RequiresRootDaemon() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("requiresRootDaemon"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/shadowChain
func (d DIBaseParams) ShadowChain() IDIShadowChain {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("shadowChain"))
	return DIShadowChainFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIBaseParams/symmetricKey
func (d DIBaseParams) SymmetricKey() foundation.INSData {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("symmetricKey"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (d DIBaseParams) SetSymmetricKey(value foundation.INSData) {
	objc.Send[struct{}](d.ID, objc.Sel("setSymmetricKey:"), value)
}


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

// The class instance for the [DICreateParams] class.
var (
	_DICreateParamsClass     DICreateParamsClass
	_DICreateParamsClassOnce sync.Once
)

func getDICreateParamsClass() DICreateParamsClass {
	_DICreateParamsClassOnce.Do(func() {
		_DICreateParamsClass = DICreateParamsClass{class: objc.GetClass("DICreateParams")}
	})
	return _DICreateParamsClass
}

// GetDICreateParamsClass returns the class object for DICreateParams.
func GetDICreateParamsClass() DICreateParamsClass {
	return getDICreateParamsClass()
}

type DICreateParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DICreateParamsClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DICreateParamsClass) Alloc() DICreateParams {
	rv := objc.Send[DICreateParams](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DICreateParams.Certificate]
//   - [DICreateParams.SetCertificate]
//   - [DICreateParams.CreateDiskImageParamsWithError]
//   - [DICreateParams.CreateDiskImageParamsXPC]
//   - [DICreateParams.CreateEncryptionWithXPCHandlerError]
//   - [DICreateParams.CreateFromAuthRef]
//   - [DICreateParams.SetCreateFromAuthRef]
//   - [DICreateParams.CreateInternalWithError]
//   - [DICreateParams.CreateWithError]
//   - [DICreateParams.EncryptionMethod]
//   - [DICreateParams.SetEncryptionMethod]
//   - [DICreateParams.FolderCopyXPCHandler]
//   - [DICreateParams.SetFolderCopyXPCHandler]
//   - [DICreateParams.NumBlocks]
//   - [DICreateParams.SetNumBlocks]
//   - [DICreateParams.OnErrorCleanup]
//   - [DICreateParams.Passphrase]
//   - [DICreateParams.SetPassphrase]
//   - [DICreateParams.PublicKey]
//   - [DICreateParams.SetPublicKey]
//   - [DICreateParams.ResizeWithDiskImageNumberOfBlocksError]
//   - [DICreateParams.ResizeWithNumBlocksError]
//   - [DICreateParams.RootCopierWithDstFolderURLSrcFolderURLProgressError]
//   - [DICreateParams.SetPassphraseEncryptionMethodError]
//   - [DICreateParams.SystemKeychainAccount]
//   - [DICreateParams.SetSystemKeychainAccount]
//   - [DICreateParams.TemporaryPassphrase]
//   - [DICreateParams.TraverseSrcFolderAsRootWithURLParallelModeProgressFolderSizeNumFilesError]
//   - [DICreateParams.ValidateBlockSizeSupport]
//   - [DICreateParams.InitWithURLNumBlocksError]
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams
type DICreateParams struct {
	DIBaseParams
}

// DICreateParamsFromID constructs a [DICreateParams] from an objc.ID.
func DICreateParamsFromID(id objc.ID) DICreateParams {
	return DICreateParams{DIBaseParams: DIBaseParamsFromID(id)}
}
// Ensure DICreateParams implements IDICreateParams.
var _ IDICreateParams = DICreateParams{}

// An interface definition for the [DICreateParams] class.
//
// # Methods
//
//   - [IDICreateParams.Certificate]
//   - [IDICreateParams.SetCertificate]
//   - [IDICreateParams.CreateDiskImageParamsWithError]
//   - [IDICreateParams.CreateDiskImageParamsXPC]
//   - [IDICreateParams.CreateEncryptionWithXPCHandlerError]
//   - [IDICreateParams.CreateFromAuthRef]
//   - [IDICreateParams.SetCreateFromAuthRef]
//   - [IDICreateParams.CreateInternalWithError]
//   - [IDICreateParams.CreateWithError]
//   - [IDICreateParams.EncryptionMethod]
//   - [IDICreateParams.SetEncryptionMethod]
//   - [IDICreateParams.FolderCopyXPCHandler]
//   - [IDICreateParams.SetFolderCopyXPCHandler]
//   - [IDICreateParams.NumBlocks]
//   - [IDICreateParams.SetNumBlocks]
//   - [IDICreateParams.OnErrorCleanup]
//   - [IDICreateParams.Passphrase]
//   - [IDICreateParams.SetPassphrase]
//   - [IDICreateParams.PublicKey]
//   - [IDICreateParams.SetPublicKey]
//   - [IDICreateParams.ResizeWithDiskImageNumberOfBlocksError]
//   - [IDICreateParams.ResizeWithNumBlocksError]
//   - [IDICreateParams.RootCopierWithDstFolderURLSrcFolderURLProgressError]
//   - [IDICreateParams.SetPassphraseEncryptionMethodError]
//   - [IDICreateParams.SystemKeychainAccount]
//   - [IDICreateParams.SetSystemKeychainAccount]
//   - [IDICreateParams.TemporaryPassphrase]
//   - [IDICreateParams.TraverseSrcFolderAsRootWithURLParallelModeProgressFolderSizeNumFilesError]
//   - [IDICreateParams.ValidateBlockSizeSupport]
//   - [IDICreateParams.InitWithURLNumBlocksError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams
type IDICreateParams interface {
	IDIBaseParams

	// Topic: Methods

	Certificate() string
	SetCertificate(value string)
	CreateDiskImageParamsWithError() (bool, error)
	CreateDiskImageParamsXPC()
	CreateEncryptionWithXPCHandlerError(xPCHandler objectivec.IObject) (bool, error)
	CreateFromAuthRef() unsafe.Pointer
	SetCreateFromAuthRef(value unsafe.Pointer)
	CreateInternalWithError() (objectivec.IObject, error)
	CreateWithError() (objectivec.IObject, error)
	EncryptionMethod() uint64
	SetEncryptionMethod(value uint64)
	FolderCopyXPCHandler() IDIClient2Controller_XPCHandler
	SetFolderCopyXPCHandler(value IDIClient2Controller_XPCHandler)
	NumBlocks() uint64
	SetNumBlocks(value uint64)
	OnErrorCleanup() bool
	Passphrase() bool
	SetPassphrase(value bool)
	PublicKey() string
	SetPublicKey(value string)
	ResizeWithDiskImageNumberOfBlocksError(image unsafe.Pointer, blocks uint64) (bool, error)
	ResizeWithNumBlocksError(blocks uint64) (bool, error)
	RootCopierWithDstFolderURLSrcFolderURLProgressError(url foundation.INSURL, url2 foundation.INSURL, progress objectivec.IObject) (bool, error)
	SetPassphraseEncryptionMethodError(passphrase string, method uint64) (bool, error)
	SystemKeychainAccount() string
	SetSystemKeychainAccount(value string)
	TemporaryPassphrase() IDITemporaryPassphrase
	TraverseSrcFolderAsRootWithURLParallelModeProgressFolderSizeNumFilesError(url foundation.NSURL, mode bool, progress objectivec.IObject) (uint64, uint64, error)
	ValidateBlockSizeSupport() bool
	InitWithURLNumBlocksError(url foundation.INSURL, blocks uint64) (DICreateParams, error)
}

// Init initializes the instance.
func (d DICreateParams) Init() DICreateParams {
	rv := objc.Send[DICreateParams](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DICreateParams) Autorelease() DICreateParams {
	rv := objc.Send[DICreateParams](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDICreateParams creates a new DICreateParams instance.
func NewDICreateParams() DICreateParams {
	class := getDICreateParamsClass()
	rv := objc.Send[DICreateParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithCoder:
func NewDICreateParamsWithCoder(coder objectivec.IObject) DICreateParams {
	instance := getDICreateParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DICreateParamsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:error:
func NewDICreateParamsWithURLError(url foundation.INSURL) (DICreateParams, error) {
	var errorPtr objc.ID
	instance := getDICreateParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:numBlocks:error:
func NewDICreateParamsWithURLNumBlocksError(url foundation.INSURL, blocks uint64) (DICreateParams, error) {
	var errorPtr objc.ID
	instance := getDICreateParamsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:numBlocks:error:"), url, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateParamsFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createDiskImageParamsWithError:
func (d DICreateParams) CreateDiskImageParamsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createDiskImageParamsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createDiskImageParamsWithError: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createDiskImageParamsXPC
func (d DICreateParams) CreateDiskImageParamsXPC() {
	objc.Send[objc.ID](d.ID, objc.Sel("createDiskImageParamsXPC"))
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createEncryptionWithXPCHandler:error:
func (d DICreateParams) CreateEncryptionWithXPCHandlerError(xPCHandler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createEncryptionWithXPCHandler:error:"), xPCHandler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createEncryptionWithXPCHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createInternalWithError:
func (d DICreateParams) CreateInternalWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createInternalWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createWithError:
func (d DICreateParams) CreateWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/onErrorCleanup
func (d DICreateParams) OnErrorCleanup() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("onErrorCleanup"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/resizeWithDiskImage:numberOfBlocks:error:
func (d DICreateParams) ResizeWithDiskImageNumberOfBlocksError(image unsafe.Pointer, blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeWithDiskImage:numberOfBlocks:error:"), image, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeWithDiskImage:numberOfBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/resizeWithNumBlocks:error:
func (d DICreateParams) ResizeWithNumBlocksError(blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeWithNumBlocks:error:"), blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeWithNumBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/rootCopierWithDstFolderURL:srcFolderURL:progress:error:
func (d DICreateParams) RootCopierWithDstFolderURLSrcFolderURLProgressError(url foundation.INSURL, url2 foundation.INSURL, progress objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("rootCopierWithDstFolderURL:srcFolderURL:progress:error:"), url, url2, progress, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("rootCopierWithDstFolderURL:srcFolderURL:progress:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/setPassphrase:encryptionMethod:error:
func (d DICreateParams) SetPassphraseEncryptionMethodError(passphrase string, method uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setPassphrase:encryptionMethod:error:"), unsafe.Pointer(unsafe.StringData(passphrase + "\x00")), method, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPassphrase:encryptionMethod:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/traverseSrcFolderAsRootWithURL:parallelMode:progress:folderSize:numFiles:error:
func (d DICreateParams) TraverseSrcFolderAsRootWithURLParallelModeProgressFolderSizeNumFilesError(url foundation.NSURL, mode bool, progress objectivec.IObject) (uint64, uint64, error) {
	var size uint64
	var files uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("traverseSrcFolderAsRootWithURL:parallelMode:progress:folderSize:numFiles:error:"), url, mode, progress, unsafe.Pointer(&size), unsafe.Pointer(&files), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, 0, errors.New("traverseSrcFolderAsRootWithURL:parallelMode:progress:folderSize:numFiles:error: returned NO with nil NSError")
	}
	return size, files, nil
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/validateBlockSizeSupport
func (d DICreateParams) ValidateBlockSizeSupport() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("validateBlockSizeSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/initWithURL:numBlocks:error:
func (d DICreateParams) InitWithURLNumBlocksError(url foundation.INSURL, blocks uint64) (DICreateParams, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:numBlocks:error:"), url, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DICreateParams{}, foundation.NSErrorFrom(errorPtr)
	}
	return DICreateParamsFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/checkExistingFileWithURL:isDirectory:error:
func (_DICreateParamsClass DICreateParamsClass) CheckExistingFileWithURLIsDirectoryError(url foundation.INSURL, directory bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICreateParamsClass.class), objc.Sel("checkExistingFileWithURL:isDirectory:error:"), url, directory, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkExistingFileWithURL:isDirectory:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/eraseIfExistingWithURL:error:
func (_DICreateParamsClass DICreateParamsClass) EraseIfExistingWithURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICreateParamsClass.class), objc.Sel("eraseIfExistingWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("eraseIfExistingWithURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/toHeaderEncryptionMode:headerEncMode:error:
func (_DICreateParamsClass DICreateParamsClass) ToHeaderEncryptionModeHeaderEncModeError(mode uint64, mode2 unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DICreateParamsClass.class), objc.Sel("toHeaderEncryptionMode:headerEncMode:error:"), mode, mode2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("toHeaderEncryptionMode:headerEncMode:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/certificate
func (d DICreateParams) Certificate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("certificate"))
	return foundation.NSStringFromID(rv).String()
}
func (d DICreateParams) SetCertificate(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setCertificate:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/createFromAuthRef
func (d DICreateParams) CreateFromAuthRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("createFromAuthRef"))
	return rv
}
func (d DICreateParams) SetCreateFromAuthRef(value unsafe.Pointer) {
	objc.Send[struct{}](d.ID, objc.Sel("setCreateFromAuthRef:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/encryptionMethod
func (d DICreateParams) EncryptionMethod() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("encryptionMethod"))
	return rv
}
func (d DICreateParams) SetEncryptionMethod(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setEncryptionMethod:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/folderCopyXPCHandler
func (d DICreateParams) FolderCopyXPCHandler() IDIClient2Controller_XPCHandler {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("folderCopyXPCHandler"))
	return DIClient2Controller_XPCHandlerFromID(objc.ID(rv))
}
func (d DICreateParams) SetFolderCopyXPCHandler(value IDIClient2Controller_XPCHandler) {
	objc.Send[struct{}](d.ID, objc.Sel("setFolderCopyXPCHandler:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/numBlocks
func (d DICreateParams) NumBlocks() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("numBlocks"))
	return rv
}
func (d DICreateParams) SetNumBlocks(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setNumBlocks:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/passphrase
func (d DICreateParams) Passphrase() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("passphrase"))
	return rv
}
func (d DICreateParams) SetPassphrase(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setPassphrase:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/publicKey
func (d DICreateParams) PublicKey() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("publicKey"))
	return foundation.NSStringFromID(rv).String()
}
func (d DICreateParams) SetPublicKey(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setPublicKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/systemKeychainAccount
func (d DICreateParams) SystemKeychainAccount() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("systemKeychainAccount"))
	return foundation.NSStringFromID(rv).String()
}
func (d DICreateParams) SetSystemKeychainAccount(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setSystemKeychainAccount:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/DICreateParams/temporaryPassphrase
func (d DICreateParams) TemporaryPassphrase() IDITemporaryPassphrase {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("temporaryPassphrase"))
	return DITemporaryPassphraseFromID(objc.ID(rv))
}


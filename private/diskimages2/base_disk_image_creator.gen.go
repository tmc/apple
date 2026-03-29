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

// The class instance for the [BaseDiskImageCreator] class.
var (
	_BaseDiskImageCreatorClass     BaseDiskImageCreatorClass
	_BaseDiskImageCreatorClassOnce sync.Once
)

func getBaseDiskImageCreatorClass() BaseDiskImageCreatorClass {
	_BaseDiskImageCreatorClassOnce.Do(func() {
		_BaseDiskImageCreatorClass = BaseDiskImageCreatorClass{class: objc.GetClass("BaseDiskImageCreator")}
	})
	return _BaseDiskImageCreatorClass
}

// GetBaseDiskImageCreatorClass returns the class object for BaseDiskImageCreator.
func GetBaseDiskImageCreatorClass() BaseDiskImageCreatorClass {
	return getBaseDiskImageCreatorClass()
}

type BaseDiskImageCreatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BaseDiskImageCreatorClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BaseDiskImageCreatorClass) Alloc() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [BaseDiskImageCreator.URL]
//   - [BaseDiskImageCreator.BlockSize]
//   - [BaseDiskImageCreator.SetBlockSize]
//   - [BaseDiskImageCreator.Certificate]
//   - [BaseDiskImageCreator.SetCertificate]
//   - [BaseDiskImageCreator.CreateEmptyImageWithError]
//   - [BaseDiskImageCreator.DataPartition]
//   - [BaseDiskImageCreator.SetDataPartition]
//   - [BaseDiskImageCreator.DevBSDName]
//   - [BaseDiskImageCreator.SetDevBSDName]
//   - [BaseDiskImageCreator.EjectWithError]
//   - [BaseDiskImageCreator.EncryptionMethod]
//   - [BaseDiskImageCreator.SetEncryptionMethod]
//   - [BaseDiskImageCreator.FileSystem]
//   - [BaseDiskImageCreator.SetFileSystem]
//   - [BaseDiskImageCreator.FormatImageWithCreateParamsError]
//   - [BaseDiskImageCreator.ImageFormat]
//   - [BaseDiskImageCreator.SetImageFormat]
//   - [BaseDiskImageCreator.MutableSymmetricKey]
//   - [BaseDiskImageCreator.NewAttachWithCreateParamsError]
//   - [BaseDiskImageCreator.NewMKDIDeviceWithError]
//   - [BaseDiskImageCreator.NumBlocks]
//   - [BaseDiskImageCreator.SetNumBlocks]
//   - [BaseDiskImageCreator.PartitionDiskWithError]
//   - [BaseDiskImageCreator.Passphrase]
//   - [BaseDiskImageCreator.SetPassphrase]
//   - [BaseDiskImageCreator.PublicKey]
//   - [BaseDiskImageCreator.SetPublicKey]
//   - [BaseDiskImageCreator.ReadPassphraseFlags]
//   - [BaseDiskImageCreator.SetReadPassphraseFlags]
//   - [BaseDiskImageCreator.SetPassphraseEncryptionMethod]
//   - [BaseDiskImageCreator.SparseBundleBandSize]
//   - [BaseDiskImageCreator.SetSparseBundleBandSize]
//   - [BaseDiskImageCreator.SymmetricKey]
//   - [BaseDiskImageCreator.SetSymmetricKey]
//   - [BaseDiskImageCreator.TemporaryPassphrase]
//   - [BaseDiskImageCreator.SetTemporaryPassphrase]
//   - [BaseDiskImageCreator.VolumeName]
//   - [BaseDiskImageCreator.SetVolumeName]
//   - [BaseDiskImageCreator.InitWithURLDefaultFormatError]
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator
type BaseDiskImageCreator struct {
	objectivec.Object
}

// BaseDiskImageCreatorFromID constructs a [BaseDiskImageCreator] from an objc.ID.
func BaseDiskImageCreatorFromID(id objc.ID) BaseDiskImageCreator {
	return BaseDiskImageCreator{objectivec.Object{ID: id}}
}
// Ensure BaseDiskImageCreator implements IBaseDiskImageCreator.
var _ IBaseDiskImageCreator = BaseDiskImageCreator{}

// An interface definition for the [BaseDiskImageCreator] class.
//
// # Methods
//
//   - [IBaseDiskImageCreator.URL]
//   - [IBaseDiskImageCreator.BlockSize]
//   - [IBaseDiskImageCreator.SetBlockSize]
//   - [IBaseDiskImageCreator.Certificate]
//   - [IBaseDiskImageCreator.SetCertificate]
//   - [IBaseDiskImageCreator.CreateEmptyImageWithError]
//   - [IBaseDiskImageCreator.DataPartition]
//   - [IBaseDiskImageCreator.SetDataPartition]
//   - [IBaseDiskImageCreator.DevBSDName]
//   - [IBaseDiskImageCreator.SetDevBSDName]
//   - [IBaseDiskImageCreator.EjectWithError]
//   - [IBaseDiskImageCreator.EncryptionMethod]
//   - [IBaseDiskImageCreator.SetEncryptionMethod]
//   - [IBaseDiskImageCreator.FileSystem]
//   - [IBaseDiskImageCreator.SetFileSystem]
//   - [IBaseDiskImageCreator.FormatImageWithCreateParamsError]
//   - [IBaseDiskImageCreator.ImageFormat]
//   - [IBaseDiskImageCreator.SetImageFormat]
//   - [IBaseDiskImageCreator.MutableSymmetricKey]
//   - [IBaseDiskImageCreator.NewAttachWithCreateParamsError]
//   - [IBaseDiskImageCreator.NewMKDIDeviceWithError]
//   - [IBaseDiskImageCreator.NumBlocks]
//   - [IBaseDiskImageCreator.SetNumBlocks]
//   - [IBaseDiskImageCreator.PartitionDiskWithError]
//   - [IBaseDiskImageCreator.Passphrase]
//   - [IBaseDiskImageCreator.SetPassphrase]
//   - [IBaseDiskImageCreator.PublicKey]
//   - [IBaseDiskImageCreator.SetPublicKey]
//   - [IBaseDiskImageCreator.ReadPassphraseFlags]
//   - [IBaseDiskImageCreator.SetReadPassphraseFlags]
//   - [IBaseDiskImageCreator.SetPassphraseEncryptionMethod]
//   - [IBaseDiskImageCreator.SparseBundleBandSize]
//   - [IBaseDiskImageCreator.SetSparseBundleBandSize]
//   - [IBaseDiskImageCreator.SymmetricKey]
//   - [IBaseDiskImageCreator.SetSymmetricKey]
//   - [IBaseDiskImageCreator.TemporaryPassphrase]
//   - [IBaseDiskImageCreator.SetTemporaryPassphrase]
//   - [IBaseDiskImageCreator.VolumeName]
//   - [IBaseDiskImageCreator.SetVolumeName]
//   - [IBaseDiskImageCreator.InitWithURLDefaultFormatError]
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator
type IBaseDiskImageCreator interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.INSURL
	BlockSize() uint32
	SetBlockSize(value uint32)
	Certificate() string
	SetCertificate(value string)
	CreateEmptyImageWithError() (objectivec.IObject, error)
	DataPartition() IDIDataPartition
	SetDataPartition(value IDIDataPartition)
	DevBSDName() string
	SetDevBSDName(value string)
	EjectWithError() (bool, error)
	EncryptionMethod() uint64
	SetEncryptionMethod(value uint64)
	FileSystem() uint64
	SetFileSystem(value uint64)
	FormatImageWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error)
	ImageFormat() int64
	SetImageFormat(value int64)
	MutableSymmetricKey() foundation.NSMutableData
	NewAttachWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error)
	NewMKDIDeviceWithError() (objectivec.IObject, error)
	NumBlocks() uint64
	SetNumBlocks(value uint64)
	PartitionDiskWithError() (bool, error)
	Passphrase() bool
	SetPassphrase(value bool)
	PublicKey() string
	SetPublicKey(value string)
	ReadPassphraseFlags() uint64
	SetReadPassphraseFlags(value uint64)
	SetPassphraseEncryptionMethod(passphrase string, method uint64)
	SparseBundleBandSize() uint64
	SetSparseBundleBandSize(value uint64)
	SymmetricKey() foundation.INSData
	SetSymmetricKey(value foundation.INSData)
	TemporaryPassphrase() IDITemporaryPassphrase
	SetTemporaryPassphrase(value IDITemporaryPassphrase)
	VolumeName() string
	SetVolumeName(value string)
	InitWithURLDefaultFormatError(url foundation.INSURL, format int64) (BaseDiskImageCreator, error)
}

// Init initializes the instance.
func (b BaseDiskImageCreator) Init() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BaseDiskImageCreator) Autorelease() BaseDiskImageCreator {
	rv := objc.Send[BaseDiskImageCreator](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBaseDiskImageCreator creates a new BaseDiskImageCreator instance.
func NewBaseDiskImageCreator() BaseDiskImageCreator {
	class := getBaseDiskImageCreatorClass()
	rv := objc.Send[BaseDiskImageCreator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/initWithURL:defaultFormat:error:
func NewBaseDiskImageCreatorWithURLDefaultFormatError(url foundation.INSURL, format int64) (BaseDiskImageCreator, error) {
	var errorPtr objc.ID
	instance := getBaseDiskImageCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BaseDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BaseDiskImageCreatorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/createEmptyImageWithError:
func (b BaseDiskImageCreator) CreateEmptyImageWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("createEmptyImageWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/ejectWithError:
func (b BaseDiskImageCreator) EjectWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("ejectWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("ejectWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/formatImageWithCreateParams:error:
func (b BaseDiskImageCreator) FormatImageWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("formatImageWithCreateParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/newAttachWithCreateParams:error:
func (b BaseDiskImageCreator) NewAttachWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newAttachWithCreateParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/newMKDIDeviceWithError:
func (b BaseDiskImageCreator) NewMKDIDeviceWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newMKDIDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/partitionDiskWithError:
func (b BaseDiskImageCreator) PartitionDiskWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("partitionDiskWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("partitionDiskWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/setPassphrase:encryptionMethod:
func (b BaseDiskImageCreator) SetPassphraseEncryptionMethod(passphrase string, method uint64) {
	objc.Send[objc.ID](b.ID, objc.Sel("setPassphrase:encryptionMethod:"), unsafe.Pointer(unsafe.StringData(passphrase + "\x00")), method)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/initWithURL:defaultFormat:error:
func (b BaseDiskImageCreator) InitWithURLDefaultFormatError(url foundation.INSURL, format int64) (BaseDiskImageCreator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BaseDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BaseDiskImageCreatorFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/debugLogsEnabled
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) DebugLogsEnabled() bool {
	rv := objc.Send[bool](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("debugLogsEnabled"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/forwardLogs
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) ForwardLogs() bool {
	rv := objc.Send[bool](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("forwardLogs"))
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/setDebugLogsEnabled:
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) SetDebugLogsEnabled(enabled bool) {
	objc.Send[objc.ID](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("setDebugLogsEnabled:"), enabled)
}
//
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/setForwardLogs:
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) SetForwardLogs(logs bool) {
	objc.Send[objc.ID](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("setForwardLogs:"), logs)
}

// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/URL
func (b BaseDiskImageCreator) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/blockSize
func (b BaseDiskImageCreator) BlockSize() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("blockSize"))
	return rv
}
func (b BaseDiskImageCreator) SetBlockSize(value uint32) {
	objc.Send[struct{}](b.ID, objc.Sel("setBlockSize:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/certificate
func (b BaseDiskImageCreator) Certificate() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("certificate"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetCertificate(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setCertificate:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/dataPartition
func (b BaseDiskImageCreator) DataPartition() IDIDataPartition {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("dataPartition"))
	return DIDataPartitionFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetDataPartition(value IDIDataPartition) {
	objc.Send[struct{}](b.ID, objc.Sel("setDataPartition:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/devBSDName
func (b BaseDiskImageCreator) DevBSDName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("devBSDName"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetDevBSDName(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setDevBSDName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/encryptionMethod
func (b BaseDiskImageCreator) EncryptionMethod() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("encryptionMethod"))
	return rv
}
func (b BaseDiskImageCreator) SetEncryptionMethod(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setEncryptionMethod:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/fileSystem
func (b BaseDiskImageCreator) FileSystem() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("fileSystem"))
	return rv
}
func (b BaseDiskImageCreator) SetFileSystem(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setFileSystem:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/imageFormat
func (b BaseDiskImageCreator) ImageFormat() int64 {
	rv := objc.Send[int64](b.ID, objc.Sel("imageFormat"))
	return rv
}
func (b BaseDiskImageCreator) SetImageFormat(value int64) {
	objc.Send[struct{}](b.ID, objc.Sel("setImageFormat:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/mutableSymmetricKey
func (b BaseDiskImageCreator) MutableSymmetricKey() foundation.NSMutableData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("mutableSymmetricKey"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/numBlocks
func (b BaseDiskImageCreator) NumBlocks() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("numBlocks"))
	return rv
}
func (b BaseDiskImageCreator) SetNumBlocks(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setNumBlocks:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/passphrase
func (b BaseDiskImageCreator) Passphrase() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("passphrase"))
	return rv
}
func (b BaseDiskImageCreator) SetPassphrase(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setPassphrase:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/publicKey
func (b BaseDiskImageCreator) PublicKey() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("publicKey"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetPublicKey(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setPublicKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/readPassphraseFlags
func (b BaseDiskImageCreator) ReadPassphraseFlags() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("readPassphraseFlags"))
	return rv
}
func (b BaseDiskImageCreator) SetReadPassphraseFlags(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setReadPassphraseFlags:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/sparseBundleBandSize
func (b BaseDiskImageCreator) SparseBundleBandSize() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("sparseBundleBandSize"))
	return rv
}
func (b BaseDiskImageCreator) SetSparseBundleBandSize(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setSparseBundleBandSize:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/symmetricKey
func (b BaseDiskImageCreator) SymmetricKey() foundation.INSData {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("symmetricKey"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetSymmetricKey(value foundation.INSData) {
	objc.Send[struct{}](b.ID, objc.Sel("setSymmetricKey:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/temporaryPassphrase
func (b BaseDiskImageCreator) TemporaryPassphrase() IDITemporaryPassphrase {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("temporaryPassphrase"))
	return DITemporaryPassphraseFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetTemporaryPassphrase(value IDITemporaryPassphrase) {
	objc.Send[struct{}](b.ID, objc.Sel("setTemporaryPassphrase:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/volumeName
func (b BaseDiskImageCreator) VolumeName() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("volumeName"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetVolumeName(value string) {
	objc.Send[struct{}](b.ID, objc.Sel("setVolumeName:"), objc.String(value))
}


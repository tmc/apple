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
	rv := objc.SendIfResponds[BaseDiskImageCreator](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

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
type IBaseDiskImageCreator interface {
	objectivec.IObject

	// Topic: Methods

	URL() foundation.NSURL
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
	SymmetricKey() foundation.NSData
	SetSymmetricKey(value foundation.NSData)
	TemporaryPassphrase() IDITemporaryPassphrase
	SetTemporaryPassphrase(value IDITemporaryPassphrase)
	VolumeName() string
	SetVolumeName(value string)
	InitWithURLDefaultFormatError(url foundation.NSURL, format int64) (BaseDiskImageCreator, error)
}

// Init initializes the instance.
func (b BaseDiskImageCreator) Init() BaseDiskImageCreator {
	rv := objc.SendIfResponds[BaseDiskImageCreator](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BaseDiskImageCreator) Autorelease() BaseDiskImageCreator {
	rv := objc.SendIfResponds[BaseDiskImageCreator](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBaseDiskImageCreator creates a new BaseDiskImageCreator instance.
func NewBaseDiskImageCreator() BaseDiskImageCreator {
	class := getBaseDiskImageCreatorClass()
	rv := objc.SendIfResponds[BaseDiskImageCreator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBaseDiskImageCreatorWithURLDefaultFormatError(url foundation.NSURL, format int64) (BaseDiskImageCreator, error) {
	var errorPtr objc.ID
	instance := getBaseDiskImageCreatorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BaseDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return BaseDiskImageCreator{}, objc.ErrInitFailed
	}
	return BaseDiskImageCreatorFromID(rv), nil
}

func (b BaseDiskImageCreator) CreateEmptyImageWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("createEmptyImageWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (b BaseDiskImageCreator) FormatImageWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("formatImageWithCreateParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (b BaseDiskImageCreator) NewAttachWithCreateParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newAttachWithCreateParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (b BaseDiskImageCreator) NewMKDIDeviceWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newMKDIDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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
func (b BaseDiskImageCreator) SetPassphraseEncryptionMethod(passphrase string, method uint64) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setPassphrase:encryptionMethod:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), method)
}
func (b BaseDiskImageCreator) InitWithURLDefaultFormatError(url foundation.NSURL, format int64) (BaseDiskImageCreator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BaseDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BaseDiskImageCreatorFromID(rv), nil

}

func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) DebugLogsEnabled() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("debugLogsEnabled"))
	return rv
}
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) ForwardLogs() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("forwardLogs"))
	return rv
}
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) SetDebugLogsEnabled(enabled bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("setDebugLogsEnabled:"), enabled)
}
func (_BaseDiskImageCreatorClass BaseDiskImageCreatorClass) SetForwardLogs(logs bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_BaseDiskImageCreatorClass.class), objc.Sel("setForwardLogs:"), logs)
}

func (b BaseDiskImageCreator) URL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) BlockSize() uint32 {
	rv := objc.SendIfResponds[uint32](b.ID, objc.Sel("blockSize"))
	return rv
}
func (b BaseDiskImageCreator) SetBlockSize(value uint32) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setBlockSize:"), value)
}
func (b BaseDiskImageCreator) Certificate() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("certificate"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetCertificate(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCertificate:"), objc.String(value))
}
func (b BaseDiskImageCreator) DataPartition() IDIDataPartition {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("dataPartition"))
	return DIDataPartitionFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetDataPartition(value IDIDataPartition) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setDataPartition:"), value)
}
func (b BaseDiskImageCreator) DevBSDName() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("devBSDName"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetDevBSDName(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setDevBSDName:"), objc.String(value))
}
func (b BaseDiskImageCreator) EncryptionMethod() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("encryptionMethod"))
	return rv
}
func (b BaseDiskImageCreator) SetEncryptionMethod(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setEncryptionMethod:"), value)
}
func (b BaseDiskImageCreator) FileSystem() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("fileSystem"))
	return rv
}
func (b BaseDiskImageCreator) SetFileSystem(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setFileSystem:"), value)
}
func (b BaseDiskImageCreator) ImageFormat() int64 {
	rv := objc.SendIfResponds[int64](b.ID, objc.Sel("imageFormat"))
	return rv
}
func (b BaseDiskImageCreator) SetImageFormat(value int64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setImageFormat:"), value)
}
func (b BaseDiskImageCreator) MutableSymmetricKey() foundation.NSMutableData {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("mutableSymmetricKey"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) NumBlocks() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("numBlocks"))
	return rv
}
func (b BaseDiskImageCreator) SetNumBlocks(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setNumBlocks:"), value)
}
func (b BaseDiskImageCreator) Passphrase() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("passphrase"))
	return rv
}
func (b BaseDiskImageCreator) SetPassphrase(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPassphrase:"), value)
}
func (b BaseDiskImageCreator) PublicKey() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("publicKey"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetPublicKey(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPublicKey:"), objc.String(value))
}
func (b BaseDiskImageCreator) ReadPassphraseFlags() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("readPassphraseFlags"))
	return rv
}
func (b BaseDiskImageCreator) SetReadPassphraseFlags(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setReadPassphraseFlags:"), value)
}
func (b BaseDiskImageCreator) SparseBundleBandSize() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("sparseBundleBandSize"))
	return rv
}
func (b BaseDiskImageCreator) SetSparseBundleBandSize(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSparseBundleBandSize:"), value)
}
func (b BaseDiskImageCreator) SymmetricKey() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("symmetricKey"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetSymmetricKey(value foundation.NSData) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSymmetricKey:"), value)
}
func (b BaseDiskImageCreator) TemporaryPassphrase() IDITemporaryPassphrase {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("temporaryPassphrase"))
	return DITemporaryPassphraseFromID(objc.ID(rv))
}
func (b BaseDiskImageCreator) SetTemporaryPassphrase(value IDITemporaryPassphrase) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setTemporaryPassphrase:"), value)
}
func (b BaseDiskImageCreator) VolumeName() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("volumeName"))
	return foundation.NSStringFromID(rv).String()
}
func (b BaseDiskImageCreator) SetVolumeName(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVolumeName:"), objc.String(value))
}

// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BackendXPC] class.
var (
	_BackendXPCClass     BackendXPCClass
	_BackendXPCClassOnce sync.Once
)

func getBackendXPCClass() BackendXPCClass {
	_BackendXPCClassOnce.Do(func() {
		_BackendXPCClass = BackendXPCClass{class: objc.GetClass("BackendXPC")}
	})
	return _BackendXPCClass
}

// GetBackendXPCClass returns the class object for BackendXPC.
func GetBackendXPCClass() BackendXPCClass {
	return getBackendXPCClass()
}

type BackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BackendXPCClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BackendXPCClass) Alloc() BackendXPC {
	rv := objc.Send[BackendXPC](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BackendXPC.Backend]
//   - [BackendXPC.SetBackend]
//   - [BackendXPC.CryptoHeader]
//   - [BackendXPC.EncodeWithCoder]
//   - [BackendXPC.GetCryptoHeaderBackend]
//   - [BackendXPC.InstanceID]
//   - [BackendXPC.IsUnlocked]
//   - [BackendXPC.Lock]
//   - [BackendXPC.NewWithCryptoFormatError]
//   - [BackendXPC.ReplaceWithBackendXPC]
//   - [BackendXPC.TryCreatingCryptoHeader]
//   - [BackendXPC.InitWithCoder]
//
// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC
type BackendXPC struct {
	objectivec.Object
}

// BackendXPCFromID constructs a [BackendXPC] from an objc.ID.
func BackendXPCFromID(id objc.ID) BackendXPC {
	return BackendXPC{objectivec.Object{ID: id}}
}

// Ensure BackendXPC implements IBackendXPC.
var _ IBackendXPC = BackendXPC{}

// An interface definition for the [BackendXPC] class.
//
// # Methods
//
//   - [IBackendXPC.Backend]
//   - [IBackendXPC.SetBackend]
//   - [IBackendXPC.CryptoHeader]
//   - [IBackendXPC.EncodeWithCoder]
//   - [IBackendXPC.GetCryptoHeaderBackend]
//   - [IBackendXPC.InstanceID]
//   - [IBackendXPC.IsUnlocked]
//   - [IBackendXPC.Lock]
//   - [IBackendXPC.NewWithCryptoFormatError]
//   - [IBackendXPC.ReplaceWithBackendXPC]
//   - [IBackendXPC.TryCreatingCryptoHeader]
//   - [IBackendXPC.InitWithCoder]
//
// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC
type IBackendXPC interface {
	objectivec.IObject

	// Topic: Methods

	Backend() objectivec.IObject
	SetBackend(value objectivec.IObject)
	CryptoHeader() unsafe.Pointer
	EncodeWithCoder(coder foundation.INSCoder)
	GetCryptoHeaderBackend() objectivec.IObject
	InstanceID() foundation.NSUUID
	IsUnlocked() bool
	Lock() int
	NewWithCryptoFormatError(format unsafe.Pointer) (objectivec.IObject, error)
	ReplaceWithBackendXPC(xpc objectivec.IObject)
	TryCreatingCryptoHeader() bool
	InitWithCoder(coder foundation.INSCoder) BackendXPC
}

// Init initializes the instance.
func (b BackendXPC) Init() BackendXPC {
	rv := objc.Send[BackendXPC](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BackendXPC) Autorelease() BackendXPC {
	rv := objc.Send[BackendXPC](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackendXPC creates a new BackendXPC instance.
func NewBackendXPC() BackendXPC {
	class := getBackendXPCClass()
	rv := objc.Send[BackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/initWithCoder:
func NewBackendXPCWithCoder(coder objectivec.IObject) BackendXPC {
	instance := getBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return BackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/encodeWithCoder:
func (b BackendXPC) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/getCryptoHeaderBackend
func (b BackendXPC) GetCryptoHeaderBackend() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getCryptoHeaderBackend"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/isUnlocked
func (b BackendXPC) IsUnlocked() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isUnlocked"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/lock
func (b BackendXPC) Lock() int {
	rv := objc.Send[int](b.ID, objc.Sel("lock"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/newWithCryptoFormat:error:
func (b BackendXPC) NewWithCryptoFormatError(format unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newWithCryptoFormat:error:"), format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/replaceWithBackendXPC:
func (b BackendXPC) ReplaceWithBackendXPC(xpc objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("replaceWithBackendXPC:"), xpc)
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/tryCreatingCryptoHeader
func (b BackendXPC) TryCreatingCryptoHeader() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("tryCreatingCryptoHeader"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/initWithCoder:
func (b BackendXPC) InitWithCoder(coder foundation.INSCoder) BackendXPC {
	rv := objc.Send[BackendXPC](b.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/newFileBackendWithURL:fileOpenFlags:error:
func (_BackendXPCClass BackendXPCClass) NewFileBackendWithURLFileOpenFlagsError(url foundation.INSURL, flags int) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_BackendXPCClass.class), objc.Sel("newFileBackendWithURL:fileOpenFlags:error:"), url, flags, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/supportsSecureCoding
func (_BackendXPCClass BackendXPCClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_BackendXPCClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/backend
func (b BackendXPC) Backend() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("backend"))
	return objectivec.Object{ID: rv}
}
func (b BackendXPC) SetBackend(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setBackend:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/cryptoHeader
func (b BackendXPC) CryptoHeader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b.ID, objc.Sel("cryptoHeader"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/BackendXPC/instanceID
func (b BackendXPC) InstanceID() foundation.NSUUID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("instanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

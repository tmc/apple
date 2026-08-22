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
	rv := objc.SendIfResponds[BackendXPC](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BackendXPC.Backend]
//   - [BackendXPC.SetBackend]
//   - [BackendXPC.CryptoHeader]
//   - [BackendXPC.EncodeWithCoder]
//   - [BackendXPC.InstanceID]
//   - [BackendXPC.IsUnlocked]
//   - [BackendXPC.Lock]
//   - [BackendXPC.NewWithCryptoFormatError]
//   - [BackendXPC.ReplaceWithBackendXPC]
//   - [BackendXPC.TryCreatingCryptoHeader]
//   - [BackendXPC.InitWithCoder]
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
//   - [IBackendXPC.InstanceID]
//   - [IBackendXPC.IsUnlocked]
//   - [IBackendXPC.Lock]
//   - [IBackendXPC.NewWithCryptoFormatError]
//   - [IBackendXPC.ReplaceWithBackendXPC]
//   - [IBackendXPC.TryCreatingCryptoHeader]
//   - [IBackendXPC.InitWithCoder]
type IBackendXPC interface {
	objectivec.IObject

	// Topic: Methods

	Backend() unsafe.Pointer
	SetBackend(value unsafe.Pointer)
	CryptoHeader() unsafe.Pointer
	EncodeWithCoder(coder foundation.INSCoder)
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
	rv := objc.SendIfResponds[BackendXPC](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BackendXPC) Autorelease() BackendXPC {
	rv := objc.SendIfResponds[BackendXPC](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackendXPC creates a new BackendXPC instance.
func NewBackendXPC() BackendXPC {
	class := getBackendXPCClass()
	rv := objc.SendIfResponds[BackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBackendXPCWithCoder(coder objectivec.IObject) BackendXPC {
	instance := getBackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return BackendXPCFromID(rv)
}

func (b BackendXPC) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (b BackendXPC) IsUnlocked() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("isUnlocked"))
	return rv
}
func (b BackendXPC) Lock() int {
	rv := objc.SendIfResponds[int](b.ID, objc.Sel("lock"))
	return rv
}
func (b BackendXPC) NewWithCryptoFormatError(format unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("newWithCryptoFormat:error:"), format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (b BackendXPC) ReplaceWithBackendXPC(xpc objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("replaceWithBackendXPC:"), xpc)
}
func (b BackendXPC) TryCreatingCryptoHeader() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("tryCreatingCryptoHeader"))
	return rv
}
func (b BackendXPC) InitWithCoder(coder foundation.INSCoder) BackendXPC {
	rv := objc.SendIfResponds[BackendXPC](b.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_BackendXPCClass BackendXPCClass) NewFileBackendWithURLFileOpenFlagsError(url foundation.NSURL, flags int) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_BackendXPCClass.class), objc.Sel("newFileBackendWithURL:fileOpenFlags:error:"), url, flags, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_BackendXPCClass BackendXPCClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_BackendXPCClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (b BackendXPC) Backend() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](b.ID, objc.Sel("backend"))
	return rv
}
func (b BackendXPC) SetBackend(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setBackend:"), value)
}
func (b BackendXPC) CryptoHeader() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](b.ID, objc.Sel("cryptoHeader"))
	return rv
}
func (b BackendXPC) InstanceID() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("instanceID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

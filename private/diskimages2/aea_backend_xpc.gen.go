// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AEABackendXPC] class.
var (
	_AEABackendXPCClass     AEABackendXPCClass
	_AEABackendXPCClassOnce sync.Once
)

func getAEABackendXPCClass() AEABackendXPCClass {
	_AEABackendXPCClassOnce.Do(func() {
		_AEABackendXPCClass = AEABackendXPCClass{class: objc.GetClass("AEABackendXPC")}
	})
	return _AEABackendXPCClass
}

// GetAEABackendXPCClass returns the class object for AEABackendXPC.
func GetAEABackendXPCClass() AEABackendXPCClass {
	return getAEABackendXPCClass()
}

type AEABackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AEABackendXPCClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AEABackendXPCClass) Alloc() AEABackendXPC {
	rv := objc.SendIfResponds[AEABackendXPC](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AEABackendXPC.BaseBackendXPC]
//   - [AEABackendXPC.Key]
//   - [AEABackendXPC.InitWithBackendKey]
type AEABackendXPC struct {
	BackendXPC
}

// AEABackendXPCFromID constructs a [AEABackendXPC] from an objc.ID.
func AEABackendXPCFromID(id objc.ID) AEABackendXPC {
	return AEABackendXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure AEABackendXPC implements IAEABackendXPC.
var _ IAEABackendXPC = AEABackendXPC{}

// An interface definition for the [AEABackendXPC] class.
//
// # Methods
//
//   - [IAEABackendXPC.BaseBackendXPC]
//   - [IAEABackendXPC.Key]
//   - [IAEABackendXPC.InitWithBackendKey]
type IAEABackendXPC interface {
	IBackendXPC

	// Topic: Methods

	BaseBackendXPC() IBackendXPC
	Key() unsafe.Pointer
	InitWithBackendKey(backend objectivec.IObject, key unsafe.Pointer) AEABackendXPC
}

// Init initializes the instance.
func (a AEABackendXPC) Init() AEABackendXPC {
	rv := objc.SendIfResponds[AEABackendXPC](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AEABackendXPC) Autorelease() AEABackendXPC {
	rv := objc.SendIfResponds[AEABackendXPC](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEABackendXPC creates a new AEABackendXPC instance.
func NewAEABackendXPC() AEABackendXPC {
	class := getAEABackendXPCClass()
	rv := objc.SendIfResponds[AEABackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAEABackendXPCWithBackendKey(backend objectivec.IObject, key unsafe.Pointer) AEABackendXPC {
	instance := getAEABackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackend:key:"), backend, key)
	return AEABackendXPCFromID(rv)
}

func NewAEABackendXPCWithCoder(coder objectivec.IObject) AEABackendXPC {
	instance := getAEABackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AEABackendXPCFromID(rv)
}

func (a AEABackendXPC) InitWithBackendKey(backend objectivec.IObject, key unsafe.Pointer) AEABackendXPC {
	rv := objc.SendIfResponds[AEABackendXPC](a.ID, objc.Sel("initWithBackend:key:"), backend, key)
	return rv
}

func (a AEABackendXPC) BaseBackendXPC() IBackendXPC {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("baseBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
func (a AEABackendXPC) Key() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("key"))
	return rv
}

// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
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
	rv := objc.Send[AEABackendXPC](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AEABackendXPC.BaseBackendXPC]
//   - [AEABackendXPC.Key]
//   - [AEABackendXPC.InitWithBackendKey]
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC
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
//
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC
type IAEABackendXPC interface {
	IBackendXPC

	// Topic: Methods

	BaseBackendXPC() IBackendXPC
	Key() objectivec.IObject
	InitWithBackendKey(backend objectivec.IObject, key objectivec.IObject) AEABackendXPC
}

// Init initializes the instance.
func (a AEABackendXPC) Init() AEABackendXPC {
	rv := objc.Send[AEABackendXPC](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AEABackendXPC) Autorelease() AEABackendXPC {
	rv := objc.Send[AEABackendXPC](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEABackendXPC creates a new AEABackendXPC instance.
func NewAEABackendXPC() AEABackendXPC {
	class := getAEABackendXPCClass()
	rv := objc.Send[AEABackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC/initWithBackend:key:
func NewAEABackendXPCWithBackendKey(backend objectivec.IObject, key objectivec.IObject) AEABackendXPC {
	instance := getAEABackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackend:key:"), backend, key)
	return AEABackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC/initWithCoder:
func NewAEABackendXPCWithCoder(coder objectivec.IObject) AEABackendXPC {
	instance := getAEABackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AEABackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC/initWithBackend:key:
func (a AEABackendXPC) InitWithBackendKey(backend objectivec.IObject, key objectivec.IObject) AEABackendXPC {
	rv := objc.Send[AEABackendXPC](a.ID, objc.Sel("initWithBackend:key:"), backend, key)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC/baseBackendXPC
func (a AEABackendXPC) BaseBackendXPC() IBackendXPC {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("baseBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/AEABackendXPC/key
func (a AEABackendXPC) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}


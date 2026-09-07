// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CryptoBackendXPC] class.
var (
	_CryptoBackendXPCClass     CryptoBackendXPCClass
	_CryptoBackendXPCClassOnce sync.Once
)

func getCryptoBackendXPCClass() CryptoBackendXPCClass {
	_CryptoBackendXPCClassOnce.Do(func() {
		_CryptoBackendXPCClass = CryptoBackendXPCClass{class: objc.GetClass("CryptoBackendXPC")}
	})
	return _CryptoBackendXPCClass
}

// GetCryptoBackendXPCClass returns the class object for CryptoBackendXPC.
func GetCryptoBackendXPCClass() CryptoBackendXPCClass {
	return getCryptoBackendXPCClass()
}

type CryptoBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CryptoBackendXPCClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CryptoBackendXPCClass) Alloc() CryptoBackendXPC {
	rv := objc.SendIfResponds[CryptoBackendXPC](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CryptoBackendXPC.BaseBackendXPC]
//   - [CryptoBackendXPC.InitWithFormatBaseBackendXPC]
type CryptoBackendXPC struct {
	BackendXPC
}

// CryptoBackendXPCFromID constructs a [CryptoBackendXPC] from an objc.ID.
func CryptoBackendXPCFromID(id objc.ID) CryptoBackendXPC {
	return CryptoBackendXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure CryptoBackendXPC implements ICryptoBackendXPC.
var _ ICryptoBackendXPC = CryptoBackendXPC{}

// An interface definition for the [CryptoBackendXPC] class.
//
// # Methods
//
//   - [ICryptoBackendXPC.BaseBackendXPC]
//   - [ICryptoBackendXPC.InitWithFormatBaseBackendXPC]
type ICryptoBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	BaseBackendXPC() IBackendXPC
	InitWithFormatBaseBackendXPC(format unsafe.Pointer, xpc objectivec.IObject) CryptoBackendXPC
}

// Init initializes the instance.
func (c CryptoBackendXPC) Init() CryptoBackendXPC {
	rv := objc.SendIfResponds[CryptoBackendXPC](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CryptoBackendXPC) Autorelease() CryptoBackendXPC {
	rv := objc.SendIfResponds[CryptoBackendXPC](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCryptoBackendXPC creates a new CryptoBackendXPC instance.
func NewCryptoBackendXPC() CryptoBackendXPC {
	class := getCryptoBackendXPCClass()
	rv := objc.SendIfResponds[CryptoBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCryptoBackendXPCWithCoder(coder objectivec.IObject) CryptoBackendXPC {
	instance := getCryptoBackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CryptoBackendXPCFromID(rv)
}

func NewCryptoBackendXPCWithFormatBaseBackendXPC(format unsafe.Pointer, xpc objectivec.IObject) CryptoBackendXPC {
	instance := getCryptoBackendXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFormat:baseBackendXPC:"), format, xpc)
	return CryptoBackendXPCFromID(rv)
}

func (c CryptoBackendXPC) InitWithFormatBaseBackendXPC(format unsafe.Pointer, xpc objectivec.IObject) CryptoBackendXPC {
	rv := objc.SendIfResponds[CryptoBackendXPC](c.ID, objc.Sel("initWithFormat:baseBackendXPC:"), format, xpc)
	return rv
}

func (c CryptoBackendXPC) BaseBackendXPC() IBackendXPC {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("baseBackendXPC"))
	return BackendXPCFromID(objc.ID(rv))
}

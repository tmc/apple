// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsRAM_XPC] class.
var (
	_DiskImageParamsRAM_XPCClass     DiskImageParamsRAM_XPCClass
	_DiskImageParamsRAM_XPCClassOnce sync.Once
)

func getDiskImageParamsRAM_XPCClass() DiskImageParamsRAM_XPCClass {
	_DiskImageParamsRAM_XPCClassOnce.Do(func() {
		_DiskImageParamsRAM_XPCClass = DiskImageParamsRAM_XPCClass{class: objc.GetClass("DiskImageParamsRAM_XPC")}
	})
	return _DiskImageParamsRAM_XPCClass
}

// GetDiskImageParamsRAM_XPCClass returns the class object for DiskImageParamsRAM_XPC.
func GetDiskImageParamsRAM_XPCClass() DiskImageParamsRAM_XPCClass {
	return getDiskImageParamsRAM_XPCClass()
}

type DiskImageParamsRAM_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsRAM_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsRAM_XPCClass) Alloc() DiskImageParamsRAM_XPC {
	rv := objc.Send[DiskImageParamsRAM_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DiskImageParamsRAM_XPC.RamSizeStr]
//   - [DiskImageParamsRAM_XPC.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC
type DiskImageParamsRAM_XPC struct {
	DiskImageParamsRaw_XPC
}

// DiskImageParamsRAM_XPCFromID constructs a [DiskImageParamsRAM_XPC] from an objc.ID.
func DiskImageParamsRAM_XPCFromID(id objc.ID) DiskImageParamsRAM_XPC {
	return DiskImageParamsRAM_XPC{DiskImageParamsRaw_XPC: DiskImageParamsRaw_XPCFromID(id)}
}

// Ensure DiskImageParamsRAM_XPC implements IDiskImageParamsRAM_XPC.
var _ IDiskImageParamsRAM_XPC = DiskImageParamsRAM_XPC{}

// An interface definition for the [DiskImageParamsRAM_XPC] class.
//
// # Methods
//
//   - [IDiskImageParamsRAM_XPC.RamSizeStr]
//   - [IDiskImageParamsRAM_XPC.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC
type IDiskImageParamsRAM_XPC interface {
	IDiskImageParamsRaw_XPC

	// Topic: Methods

	RamSizeStr() string
	InitWithURLError(url foundation.INSURL) (DiskImageParamsRAM_XPC, error)
}

// Init initializes the instance.
func (d DiskImageParamsRAM_XPC) Init() DiskImageParamsRAM_XPC {
	rv := objc.Send[DiskImageParamsRAM_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsRAM_XPC) Autorelease() DiskImageParamsRAM_XPC {
	rv := objc.Send[DiskImageParamsRAM_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsRAM_XPC creates a new DiskImageParamsRAM_XPC instance.
func NewDiskImageParamsRAM_XPC() DiskImageParamsRAM_XPC {
	class := getDiskImageParamsRAM_XPCClass()
	rv := objc.Send[DiskImageParamsRAM_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsRAM_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsRAM_XPC {
	instance := getDiskImageParamsRAM_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsRAM_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsRAM_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsRAM_XPC {
	instance := getDiskImageParamsRAM_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsRAM_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsRAM_XPCWithCoder(coder objectivec.IObject) DiskImageParamsRAM_XPC {
	instance := getDiskImageParamsRAM_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsRAM_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/initWithURL:error:
func NewDiskImageParamsRAM_XPCWithURLError(url foundation.INSURL) (DiskImageParamsRAM_XPC, error) {
	var errorPtr objc.ID
	instance := getDiskImageParamsRAM_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageParamsRAM_XPC{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageParamsRAM_XPCFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/initWithURL:error:
func (d DiskImageParamsRAM_XPC) InitWithURLError(url foundation.INSURL) (DiskImageParamsRAM_XPC, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageParamsRAM_XPC{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageParamsRAM_XPCFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/ramSizeStr
func (d DiskImageParamsRAM_XPC) RamSizeStr() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ramSizeStr"))
	return foundation.NSStringFromID(rv).String()
}

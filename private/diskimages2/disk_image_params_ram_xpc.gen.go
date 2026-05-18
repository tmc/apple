// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsRAMXPC] class.
var (
	_DiskImageParamsRAMXPCClass     DiskImageParamsRAMXPCClass
	_DiskImageParamsRAMXPCClassOnce sync.Once
)

func getDiskImageParamsRAMXPCClass() DiskImageParamsRAMXPCClass {
	_DiskImageParamsRAMXPCClassOnce.Do(func() {
		_DiskImageParamsRAMXPCClass = DiskImageParamsRAMXPCClass{class: objc.GetClass("DiskImageParamsRAM_XPC")}
	})
	return _DiskImageParamsRAMXPCClass
}

// GetDiskImageParamsRAMXPCClass returns the class object for DiskImageParamsRAM_XPC.
func GetDiskImageParamsRAMXPCClass() DiskImageParamsRAMXPCClass {
	return getDiskImageParamsRAMXPCClass()
}

type DiskImageParamsRAMXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsRAMXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsRAMXPCClass) Alloc() DiskImageParamsRAMXPC {
	rv := objc.Send[DiskImageParamsRAMXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DiskImageParamsRAMXPC.RamSizeStr]
//   - [DiskImageParamsRAMXPC.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC
type DiskImageParamsRAMXPC struct {
	DiskImageParamsRawXPC
}

// DiskImageParamsRAMXPCFromID constructs a [DiskImageParamsRAMXPC] from an objc.ID.
func DiskImageParamsRAMXPCFromID(id objc.ID) DiskImageParamsRAMXPC {
	return DiskImageParamsRAMXPC{DiskImageParamsRawXPC: DiskImageParamsRawXPCFromID(id)}
}

// DiskImageParamsRAM_XPCFromID is an alias for [DiskImageParamsRAMXPCFromID] for cross-framework compatibility.
func DiskImageParamsRAM_XPCFromID(id objc.ID) DiskImageParamsRAMXPC {
	return DiskImageParamsRAMXPCFromID(id)
}

// Ensure DiskImageParamsRAMXPC implements IDiskImageParamsRAMXPC.
var _ IDiskImageParamsRAMXPC = DiskImageParamsRAMXPC{}

// An interface definition for the [DiskImageParamsRAMXPC] class.
//
// # Methods
//
//   - [IDiskImageParamsRAMXPC.RamSizeStr]
//   - [IDiskImageParamsRAMXPC.InitWithURLError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC
type IDiskImageParamsRAMXPC interface {
	IDiskImageParamsRawXPC

	// Topic: Methods

	RamSizeStr() string
	InitWithURLError(url foundation.INSURL) (DiskImageParamsRAMXPC, error)
}

// Init initializes the instance.
func (d DiskImageParamsRAMXPC) Init() DiskImageParamsRAMXPC {
	rv := objc.Send[DiskImageParamsRAMXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsRAMXPC) Autorelease() DiskImageParamsRAMXPC {
	rv := objc.Send[DiskImageParamsRAMXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsRAMXPC creates a new DiskImageParamsRAMXPC instance.
func NewDiskImageParamsRAMXPC() DiskImageParamsRAMXPC {
	class := getDiskImageParamsRAMXPCClass()
	rv := objc.Send[DiskImageParamsRAMXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsRAM_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsRAMXPC {
	instance := getDiskImageParamsRAMXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsRAMXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsRAM_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsRAMXPC {
	instance := getDiskImageParamsRAMXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsRAMXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsRAM_XPCWithCoder(coder objectivec.IObject) DiskImageParamsRAMXPC {
	instance := getDiskImageParamsRAMXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsRAMXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/initWithURL:error:
func NewDiskImageParamsRAM_XPCWithURLError(url foundation.INSURL) (DiskImageParamsRAMXPC, error) {
	var errorPtr objc.ID
	instance := getDiskImageParamsRAMXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageParamsRAMXPC{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageParamsRAMXPCFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/initWithURL:error:
func (d DiskImageParamsRAMXPC) InitWithURLError(url foundation.INSURL) (DiskImageParamsRAMXPC, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DiskImageParamsRAMXPC{}, foundation.NSErrorFrom(errorPtr)
	}
	return DiskImageParamsRAMXPCFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRAM_XPC/ramSizeStr
func (d DiskImageParamsRAMXPC) RamSizeStr() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ramSizeStr"))
	return foundation.NSStringFromID(rv).String()
}

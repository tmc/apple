// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsUDIF_XPC] class.
var (
	_DiskImageParamsUDIF_XPCClass     DiskImageParamsUDIF_XPCClass
	_DiskImageParamsUDIF_XPCClassOnce sync.Once
)

func getDiskImageParamsUDIF_XPCClass() DiskImageParamsUDIF_XPCClass {
	_DiskImageParamsUDIF_XPCClassOnce.Do(func() {
		_DiskImageParamsUDIF_XPCClass = DiskImageParamsUDIF_XPCClass{class: objc.GetClass("DiskImageParamsUDIF_XPC")}
	})
	return _DiskImageParamsUDIF_XPCClass
}

// GetDiskImageParamsUDIF_XPCClass returns the class object for DiskImageParamsUDIF_XPC.
func GetDiskImageParamsUDIF_XPCClass() DiskImageParamsUDIF_XPCClass {
	return getDiskImageParamsUDIF_XPCClass()
}

type DiskImageParamsUDIF_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsUDIF_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsUDIF_XPCClass) Alloc() DiskImageParamsUDIF_XPC {
	rv := objc.Send[DiskImageParamsUDIF_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DiskImageParamsUDIF_XPC.InitWithBackendXPCHeader]
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsUDIF_XPC
type DiskImageParamsUDIF_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsUDIF_XPCFromID constructs a [DiskImageParamsUDIF_XPC] from an objc.ID.
func DiskImageParamsUDIF_XPCFromID(id objc.ID) DiskImageParamsUDIF_XPC {
	return DiskImageParamsUDIF_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}
// Ensure DiskImageParamsUDIF_XPC implements IDiskImageParamsUDIF_XPC.
var _ IDiskImageParamsUDIF_XPC = DiskImageParamsUDIF_XPC{}

// An interface definition for the [DiskImageParamsUDIF_XPC] class.
//
// # Methods
//
//   - [IDiskImageParamsUDIF_XPC.InitWithBackendXPCHeader]
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsUDIF_XPC
type IDiskImageParamsUDIF_XPC interface {
	IDiskImageParamsXPC

	// Topic: Methods

	InitWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsUDIF_XPC
}

// Init initializes the instance.
func (d DiskImageParamsUDIF_XPC) Init() DiskImageParamsUDIF_XPC {
	rv := objc.Send[DiskImageParamsUDIF_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsUDIF_XPC) Autorelease() DiskImageParamsUDIF_XPC {
	rv := objc.Send[DiskImageParamsUDIF_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsUDIF_XPC creates a new DiskImageParamsUDIF_XPC instance.
func NewDiskImageParamsUDIF_XPC() DiskImageParamsUDIF_XPC {
	class := getDiskImageParamsUDIF_XPCClass()
	rv := objc.Send[DiskImageParamsUDIF_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsUDIF_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsUDIF_XPC {
	instance := getDiskImageParamsUDIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsUDIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsUDIF_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsUDIF_XPC {
	instance := getDiskImageParamsUDIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsUDIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsUDIF_XPC/initWithBackendXPC:header:
func NewDiskImageParamsUDIF_XPCWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsUDIF_XPC {
	instance := getDiskImageParamsUDIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return DiskImageParamsUDIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsUDIF_XPC/initWithCoder:
func NewDiskImageParamsUDIF_XPCWithCoder(coder objectivec.IObject) DiskImageParamsUDIF_XPC {
	instance := getDiskImageParamsUDIF_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsUDIF_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsUDIF_XPC/initWithBackendXPC:header:
func (d DiskImageParamsUDIF_XPC) InitWithBackendXPCHeader(xpc objectivec.IObject, header objectivec.IObject) DiskImageParamsUDIF_XPC {
	rv := objc.Send[DiskImageParamsUDIF_XPC](d.ID, objc.Sel("initWithBackendXPC:header:"), xpc, header)
	return rv
}


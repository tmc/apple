// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsRaw_XPC] class.
var (
	_DiskImageParamsRaw_XPCClass     DiskImageParamsRaw_XPCClass
	_DiskImageParamsRaw_XPCClassOnce sync.Once
)

func getDiskImageParamsRaw_XPCClass() DiskImageParamsRaw_XPCClass {
	_DiskImageParamsRaw_XPCClassOnce.Do(func() {
		_DiskImageParamsRaw_XPCClass = DiskImageParamsRaw_XPCClass{class: objc.GetClass("DiskImageParamsRaw_XPC")}
	})
	return _DiskImageParamsRaw_XPCClass
}

// GetDiskImageParamsRaw_XPCClass returns the class object for DiskImageParamsRaw_XPC.
func GetDiskImageParamsRaw_XPCClass() DiskImageParamsRaw_XPCClass {
	return getDiskImageParamsRaw_XPCClass()
}

type DiskImageParamsRaw_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsRaw_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsRaw_XPCClass) Alloc() DiskImageParamsRaw_XPC {
	rv := objc.Send[DiskImageParamsRaw_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRaw_XPC
type DiskImageParamsRaw_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsRaw_XPCFromID constructs a [DiskImageParamsRaw_XPC] from an objc.ID.
func DiskImageParamsRaw_XPCFromID(id objc.ID) DiskImageParamsRaw_XPC {
	return DiskImageParamsRaw_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}
// Ensure DiskImageParamsRaw_XPC implements IDiskImageParamsRaw_XPC.
var _ IDiskImageParamsRaw_XPC = DiskImageParamsRaw_XPC{}

// An interface definition for the [DiskImageParamsRaw_XPC] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsRaw_XPC
type IDiskImageParamsRaw_XPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsRaw_XPC) Init() DiskImageParamsRaw_XPC {
	rv := objc.Send[DiskImageParamsRaw_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsRaw_XPC) Autorelease() DiskImageParamsRaw_XPC {
	rv := objc.Send[DiskImageParamsRaw_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsRaw_XPC creates a new DiskImageParamsRaw_XPC instance.
func NewDiskImageParamsRaw_XPC() DiskImageParamsRaw_XPC {
	class := getDiskImageParamsRaw_XPCClass()
	rv := objc.Send[DiskImageParamsRaw_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsRaw_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsRaw_XPC {
	instance := getDiskImageParamsRaw_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsRaw_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsRaw_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsRaw_XPC {
	instance := getDiskImageParamsRaw_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsRaw_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsRaw_XPCWithCoder(coder objectivec.IObject) DiskImageParamsRaw_XPC {
	instance := getDiskImageParamsRaw_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsRaw_XPCFromID(rv)
}


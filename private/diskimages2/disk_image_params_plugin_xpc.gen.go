// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsPlugin_XPC] class.
var (
	_DiskImageParamsPlugin_XPCClass     DiskImageParamsPlugin_XPCClass
	_DiskImageParamsPlugin_XPCClassOnce sync.Once
)

func getDiskImageParamsPlugin_XPCClass() DiskImageParamsPlugin_XPCClass {
	_DiskImageParamsPlugin_XPCClassOnce.Do(func() {
		_DiskImageParamsPlugin_XPCClass = DiskImageParamsPlugin_XPCClass{class: objc.GetClass("DiskImageParamsPlugin_XPC")}
	})
	return _DiskImageParamsPlugin_XPCClass
}

// GetDiskImageParamsPlugin_XPCClass returns the class object for DiskImageParamsPlugin_XPC.
func GetDiskImageParamsPlugin_XPCClass() DiskImageParamsPlugin_XPCClass {
	return getDiskImageParamsPlugin_XPCClass()
}

type DiskImageParamsPlugin_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsPlugin_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsPlugin_XPCClass) Alloc() DiskImageParamsPlugin_XPC {
	rv := objc.Send[DiskImageParamsPlugin_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsPlugin_XPC
type DiskImageParamsPlugin_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsPlugin_XPCFromID constructs a [DiskImageParamsPlugin_XPC] from an objc.ID.
func DiskImageParamsPlugin_XPCFromID(id objc.ID) DiskImageParamsPlugin_XPC {
	return DiskImageParamsPlugin_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}
// Ensure DiskImageParamsPlugin_XPC implements IDiskImageParamsPlugin_XPC.
var _ IDiskImageParamsPlugin_XPC = DiskImageParamsPlugin_XPC{}

// An interface definition for the [DiskImageParamsPlugin_XPC] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsPlugin_XPC
type IDiskImageParamsPlugin_XPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsPlugin_XPC) Init() DiskImageParamsPlugin_XPC {
	rv := objc.Send[DiskImageParamsPlugin_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsPlugin_XPC) Autorelease() DiskImageParamsPlugin_XPC {
	rv := objc.Send[DiskImageParamsPlugin_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsPlugin_XPC creates a new DiskImageParamsPlugin_XPC instance.
func NewDiskImageParamsPlugin_XPC() DiskImageParamsPlugin_XPC {
	class := getDiskImageParamsPlugin_XPCClass()
	rv := objc.Send[DiskImageParamsPlugin_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsPlugin_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsPlugin_XPC {
	instance := getDiskImageParamsPlugin_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsPlugin_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsPlugin_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsPlugin_XPC {
	instance := getDiskImageParamsPlugin_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsPlugin_XPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsPlugin_XPCWithCoder(coder objectivec.IObject) DiskImageParamsPlugin_XPC {
	instance := getDiskImageParamsPlugin_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsPlugin_XPCFromID(rv)
}


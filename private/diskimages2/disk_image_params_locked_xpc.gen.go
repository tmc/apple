// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsLocked_XPC] class.
var (
	_DiskImageParamsLocked_XPCClass     DiskImageParamsLocked_XPCClass
	_DiskImageParamsLocked_XPCClassOnce sync.Once
)

func getDiskImageParamsLocked_XPCClass() DiskImageParamsLocked_XPCClass {
	_DiskImageParamsLocked_XPCClassOnce.Do(func() {
		_DiskImageParamsLocked_XPCClass = DiskImageParamsLocked_XPCClass{class: objc.GetClass("DiskImageParamsLocked_XPC")}
	})
	return _DiskImageParamsLocked_XPCClass
}

// GetDiskImageParamsLocked_XPCClass returns the class object for DiskImageParamsLocked_XPC.
func GetDiskImageParamsLocked_XPCClass() DiskImageParamsLocked_XPCClass {
	return getDiskImageParamsLocked_XPCClass()
}

type DiskImageParamsLocked_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsLocked_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsLocked_XPCClass) Alloc() DiskImageParamsLocked_XPC {
	rv := objc.Send[DiskImageParamsLocked_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsLocked_XPC
type DiskImageParamsLocked_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsLocked_XPCFromID constructs a [DiskImageParamsLocked_XPC] from an objc.ID.
func DiskImageParamsLocked_XPCFromID(id objc.ID) DiskImageParamsLocked_XPC {
	return DiskImageParamsLocked_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// Ensure DiskImageParamsLocked_XPC implements IDiskImageParamsLocked_XPC.
var _ IDiskImageParamsLocked_XPC = DiskImageParamsLocked_XPC{}

// An interface definition for the [DiskImageParamsLocked_XPC] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsLocked_XPC
type IDiskImageParamsLocked_XPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsLocked_XPC) Init() DiskImageParamsLocked_XPC {
	rv := objc.Send[DiskImageParamsLocked_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsLocked_XPC) Autorelease() DiskImageParamsLocked_XPC {
	rv := objc.Send[DiskImageParamsLocked_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsLocked_XPC creates a new DiskImageParamsLocked_XPC instance.
func NewDiskImageParamsLocked_XPC() DiskImageParamsLocked_XPC {
	class := getDiskImageParamsLocked_XPCClass()
	rv := objc.Send[DiskImageParamsLocked_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsLocked_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsLocked_XPC {
	instance := getDiskImageParamsLocked_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsLocked_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsLocked_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsLocked_XPC {
	instance := getDiskImageParamsLocked_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsLocked_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsLocked_XPCWithCoder(coder objectivec.IObject) DiskImageParamsLocked_XPC {
	instance := getDiskImageParamsLocked_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsLocked_XPCFromID(rv)
}

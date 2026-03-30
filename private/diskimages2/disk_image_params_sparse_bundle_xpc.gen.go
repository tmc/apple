// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsSparseBundle_XPC] class.
var (
	_DiskImageParamsSparseBundle_XPCClass     DiskImageParamsSparseBundle_XPCClass
	_DiskImageParamsSparseBundle_XPCClassOnce sync.Once
)

func getDiskImageParamsSparseBundle_XPCClass() DiskImageParamsSparseBundle_XPCClass {
	_DiskImageParamsSparseBundle_XPCClassOnce.Do(func() {
		_DiskImageParamsSparseBundle_XPCClass = DiskImageParamsSparseBundle_XPCClass{class: objc.GetClass("DiskImageParamsSparseBundle_XPC")}
	})
	return _DiskImageParamsSparseBundle_XPCClass
}

// GetDiskImageParamsSparseBundle_XPCClass returns the class object for DiskImageParamsSparseBundle_XPC.
func GetDiskImageParamsSparseBundle_XPCClass() DiskImageParamsSparseBundle_XPCClass {
	return getDiskImageParamsSparseBundle_XPCClass()
}

type DiskImageParamsSparseBundle_XPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsSparseBundle_XPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsSparseBundle_XPCClass) Alloc() DiskImageParamsSparseBundle_XPC {
	rv := objc.Send[DiskImageParamsSparseBundle_XPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsSparseBundle_XPC
type DiskImageParamsSparseBundle_XPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsSparseBundle_XPCFromID constructs a [DiskImageParamsSparseBundle_XPC] from an objc.ID.
func DiskImageParamsSparseBundle_XPCFromID(id objc.ID) DiskImageParamsSparseBundle_XPC {
	return DiskImageParamsSparseBundle_XPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// Ensure DiskImageParamsSparseBundle_XPC implements IDiskImageParamsSparseBundle_XPC.
var _ IDiskImageParamsSparseBundle_XPC = DiskImageParamsSparseBundle_XPC{}

// An interface definition for the [DiskImageParamsSparseBundle_XPC] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsSparseBundle_XPC
type IDiskImageParamsSparseBundle_XPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsSparseBundle_XPC) Init() DiskImageParamsSparseBundle_XPC {
	rv := objc.Send[DiskImageParamsSparseBundle_XPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsSparseBundle_XPC) Autorelease() DiskImageParamsSparseBundle_XPC {
	rv := objc.Send[DiskImageParamsSparseBundle_XPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsSparseBundle_XPC creates a new DiskImageParamsSparseBundle_XPC instance.
func NewDiskImageParamsSparseBundle_XPC() DiskImageParamsSparseBundle_XPC {
	class := getDiskImageParamsSparseBundle_XPCClass()
	rv := objc.Send[DiskImageParamsSparseBundle_XPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:
func NewDiskImageParamsSparseBundle_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsSparseBundle_XPC {
	instance := getDiskImageParamsSparseBundle_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsSparseBundle_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithBackendXPC:blockSize:
func NewDiskImageParamsSparseBundle_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsSparseBundle_XPC {
	instance := getDiskImageParamsSparseBundle_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsSparseBundle_XPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DiskImageParamsXPC/initWithCoder:
func NewDiskImageParamsSparseBundle_XPCWithCoder(coder objectivec.IObject) DiskImageParamsSparseBundle_XPC {
	instance := getDiskImageParamsSparseBundle_XPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsSparseBundle_XPCFromID(rv)
}

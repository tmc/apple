// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DiskImageParamsSparseBundleXPC] class.
var (
	_DiskImageParamsSparseBundleXPCClass     DiskImageParamsSparseBundleXPCClass
	_DiskImageParamsSparseBundleXPCClassOnce sync.Once
)

func getDiskImageParamsSparseBundleXPCClass() DiskImageParamsSparseBundleXPCClass {
	_DiskImageParamsSparseBundleXPCClassOnce.Do(func() {
		_DiskImageParamsSparseBundleXPCClass = DiskImageParamsSparseBundleXPCClass{class: objc.GetClass("DiskImageParamsSparseBundle_XPC")}
	})
	return _DiskImageParamsSparseBundleXPCClass
}

// GetDiskImageParamsSparseBundleXPCClass returns the class object for DiskImageParamsSparseBundle_XPC.
func GetDiskImageParamsSparseBundleXPCClass() DiskImageParamsSparseBundleXPCClass {
	return getDiskImageParamsSparseBundleXPCClass()
}

type DiskImageParamsSparseBundleXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DiskImageParamsSparseBundleXPCClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DiskImageParamsSparseBundleXPCClass) Alloc() DiskImageParamsSparseBundleXPC {
	rv := objc.SendIfResponds[DiskImageParamsSparseBundleXPC](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DiskImageParamsSparseBundleXPC struct {
	DiskImageParamsXPC
}

// DiskImageParamsSparseBundleXPCFromID constructs a [DiskImageParamsSparseBundleXPC] from an objc.ID.
func DiskImageParamsSparseBundleXPCFromID(id objc.ID) DiskImageParamsSparseBundleXPC {
	return DiskImageParamsSparseBundleXPC{DiskImageParamsXPC: DiskImageParamsXPCFromID(id)}
}

// DiskImageParamsSparseBundle_XPCFromID is an alias for [DiskImageParamsSparseBundleXPCFromID] for cross-framework compatibility.
func DiskImageParamsSparseBundle_XPCFromID(id objc.ID) DiskImageParamsSparseBundleXPC {
	return DiskImageParamsSparseBundleXPCFromID(id)
}

// Ensure DiskImageParamsSparseBundleXPC implements IDiskImageParamsSparseBundleXPC.
var _ IDiskImageParamsSparseBundleXPC = DiskImageParamsSparseBundleXPC{}

// An interface definition for the [DiskImageParamsSparseBundleXPC] class.
type IDiskImageParamsSparseBundleXPC interface {
	IDiskImageParamsXPC
}

// Init initializes the instance.
func (d DiskImageParamsSparseBundleXPC) Init() DiskImageParamsSparseBundleXPC {
	rv := objc.SendIfResponds[DiskImageParamsSparseBundleXPC](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DiskImageParamsSparseBundleXPC) Autorelease() DiskImageParamsSparseBundleXPC {
	rv := objc.SendIfResponds[DiskImageParamsSparseBundleXPC](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiskImageParamsSparseBundleXPC creates a new DiskImageParamsSparseBundleXPC instance.
func NewDiskImageParamsSparseBundleXPC() DiskImageParamsSparseBundleXPC {
	class := getDiskImageParamsSparseBundleXPCClass()
	rv := objc.SendIfResponds[DiskImageParamsSparseBundleXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDiskImageParamsSparseBundle_XPCWithBackendXPC(xpc objectivec.IObject) DiskImageParamsSparseBundleXPC {
	instance := getDiskImageParamsSparseBundleXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:"), xpc)
	return DiskImageParamsSparseBundleXPCFromID(rv)
}

func NewDiskImageParamsSparseBundle_XPCWithBackendXPCBlockSize(xpc objectivec.IObject, size uint64) DiskImageParamsSparseBundleXPC {
	instance := getDiskImageParamsSparseBundleXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackendXPC:blockSize:"), xpc, size)
	return DiskImageParamsSparseBundleXPCFromID(rv)
}

func NewDiskImageParamsSparseBundle_XPCWithCoder(coder objectivec.IObject) DiskImageParamsSparseBundleXPC {
	instance := getDiskImageParamsSparseBundleXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DiskImageParamsSparseBundleXPCFromID(rv)
}
